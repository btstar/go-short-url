// 通过新浪接口地址和source(开发者账号申请) 来转换
package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

var (
	sina_url = "https://api.weibo.com/2/short_url/shorten.json"
)

type SinaUrls struct {
	Urls []SinaUrl `json:"urls"`
}

type SinaUrl struct {
	Result bool `json:"result"`
	UrlShort string `json:"url_short"`
	UrlLong string `json:"url_long"`
	ObjectType string `json:"object_type"`
	Types int `json:"type"`
	ObjectId string `json:"object_id"`
}

// GetShortUrl 通过新浪获取短链接
func GetShortUrl(longUrl string) *SinaUrls  {
	source := beego.AppConfig.String("sinaSource")
	req := httplib.Post(sina_url)
	req.Param("source", source)
	req.Param("url_long", longUrl)
	sina := &SinaUrls{}
	s, err := req.String()

	if err != nil {
		return nil
	}
	json.Unmarshal([]byte(s),sina)
	return sina
}

func GetMD5(url string) string {
	hash := md5.Sum([]byte(url))
	return fmt.Sprintf("%x", hash)
}