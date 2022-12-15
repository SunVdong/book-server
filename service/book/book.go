package book

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/book"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type BookService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBook
//@description: 创建书籍
//@param: e model.Book
//@return: err error

func (exa *BookService) CreateBook(e book.Book) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除书籍
//@param: e model.Book
//@return: err error

func (exa *BookService) DeleteBook(e book.Book) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBook
//@description: 更新书籍
//@param: e *model.Book
//@return: err error

func (exa *BookService) UpdateBook(e *book.Book) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBook
//@description: 获取书籍信息
//@param: id uint
//@return: customer model.Book, err error

func (exa *BookService) GetBook(id uint) (customer book.Book, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取书籍列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *BookService) GetCustomerInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&book.Book{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []book.Book
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
