package book

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BookApi
}

var (
	excelService = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
)
