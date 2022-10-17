package specification

//PaginationStruct stores return info
type PaginationStruct struct {
	TableName string
	PageSize  int32
	Index     int32
}

// func NewPagination(info *PaginateInfo) (result map[string]interface{}) {
// 	// if info == nil {
// 	// 	return nil
// 	// }

// 	// result = map[string]interface{}{
// 	// 	"currentPage":  info.CurrentPage,
// 	// 	"nextPage":     info.NextPage,
// 	// 	"previousPage": info.PreviousPage,
// 	// 	"totalPages":   info.TotalPages,
// 	// 	"totalRows":    info.TotalRows,
// 	// }
// 	// return result
// 	return nil
// }
