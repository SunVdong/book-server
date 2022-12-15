package book

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// book struct, 书籍结构体
type Book struct {
	global.GVA_MODEL
	FileName   string
	FileMd5    string
	FilePath   string
	ChunkTotal int
	IsFinish   bool
}
