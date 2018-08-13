package controllers

import (
	"github.com/astaxie/beego"
	"go-short-url/models"
	"go-short-url/vo"
	"log"
	"strings"
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
	response := &vo.ResponseData{
		Code: vo.SUCCESS_CODE,
	}
	c.Data["json"] = response
	url := c.GetString("url")
	if url == "" || !strings.HasPrefix(url, "http") {
		response.Code = vo.ERROR_CODE
		response.Message = "url不合法"
		c.ServeJSON()
		c.StopRun()
	}
	log.Print(url)
	if shortUrl := models.IsLongUrlExist(url); shortUrl != nil {
		response.Data = make(map[string]interface{})
		response.Data["long_url"] = shortUrl.LongUrl
		response.Data["short_url"] = shortUrl.ShortUrl
	} else {
		shortUrl, err := models.GenerateShortUrl(url)
		if err != nil {
			response.Code = vo.ERROR_CODE
			response.Message = err.Error()
			c.ServeJSON()
			c.StopRun()
		}
		response.Data = make(map[string]interface{})
		response.Data["long_url"] = shortUrl.LongUrl
		response.Data["short_url"] = shortUrl.ShortUrl
	}

	c.ServeJSON()
}
