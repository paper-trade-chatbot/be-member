package util

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// URL of the MOHW ASP.NET lookup webpage.
const mohwURL = "https://rao.mohw.gov.tw/LicenceLookup.aspx?type=reserve_p"

func init() {
}

// GetMedCertID fetches & returns the medical personnel certificate ID, given
// the R.O.C identification number of the individual.
func GetMedCertID(realUserID string) (string, error) {
	// Retrieve form first.
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: transport}
	response, err := client.Get(mohwURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read form data from body.
	root, err := html.Parse(response.Body)
	if err != nil {
		return "", err
	}

	// Get form input values.
	viewStateNode := getElementByID(root, "__VIEWSTATE")
	viewState, _ := getAttribute(viewStateNode, "value")
	viewStateGeneratorNode := getElementByID(root, "__VIEWSTATEGENERATOR")
	viewStateGenerator, _ := getAttribute(viewStateGeneratorNode, "value")
	eventValidationNode := getElementByID(root, "__EVENTVALIDATION")
	eventValidation, _ := getAttribute(eventValidationNode, "value")

	// Create HTTP client & request form data.
	client = &http.Client{Transport: transport}
	response, err = client.PostForm(
		mohwURL,
		url.Values{
			"__VIEWSTATE":          {viewState},
			"__VIEWSTATEGENERATOR": {viewStateGenerator},
			"__EVENTVALIDATION":    {eventValidation},
			"doc_id_tbx":           {realUserID},
			"lookup_btn":           {"查詢"},
		},
	)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read form data from body.
	root, err = html.Parse(response.Body)
	if err != nil {
		return "", err
	}

	// Get response node.
	certIDParentNode := getElementByID(root, "Adoc_id_lbl")
	certIDChildNode := certIDParentNode.FirstChild
	if certIDChildNode == nil || certIDChildNode.FirstChild == nil {
		return "N/A", nil
	}
	certID := certIDChildNode.FirstChild.Data

	return certID, nil
}

func getElementByID(node *html.Node, id string) *html.Node {
	return traverse(node, id)
}

func traverse(node *html.Node, id string) *html.Node {
	if checkID(node, id) {
		return node
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result := traverse(child, id)
		if result != nil {
			return result
		}
	}

	return nil
}

func checkID(node *html.Node, id string) bool {
	if node.Type == html.ElementNode {
		value, ok := getAttribute(node, "id")
		if ok && value == id {
			return true
		}
	}

	return false
}

func getAttribute(node *html.Node, key string) (string, bool) {
	for _, attribute := range node.Attr {
		if attribute.Key == key {
			return attribute.Val, true
		}
	}

	return "", false
}
