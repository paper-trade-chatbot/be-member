//go:build test
// +build test

package middleware

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/paper-trade-chatbot/be-member/config"
)

var appName string
var key string

func init() {
	appName = config.GetString("PROJECT_NAME")
	key = config.GetString("NEW_RELIC_LICENSE")
}

// NewRelicMiddleware ...
func NewRelicMiddleware() gin.HandlerFunc {

	if appName == "" || key == "" {
		return func(c *gin.Context) {}
	}

	config := newrelic.NewConfig(appName, key)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path, c.Writer, c.Request)
		defer txn.End()
		c.Next()
	}
}
