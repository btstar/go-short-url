package controllers

import (
	"github.com/astaxie/beego"
	"go-short-url/vo"
	"fmt"
)

type ShortController struct {
	beego.Controller
}

// @Title 长链接转短链接
// @Description 长链接转短链接
// @Param	url		query    string  true        "长链接"
// @Success 200 {object} vo.ResponseData
// @Failure 400 Invalid url
// @router / [get]
func (c *ShortController) Get()  {
	response := vo.ResponseData{
		Code: vo.SUCCESS_CODE,
	}
	fmt.Println(response)
}
