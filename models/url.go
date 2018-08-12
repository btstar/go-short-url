package models

import (
	"go-short-url/utils"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris/core/errors"
)

// Url模型
type Url struct {
	Id int `orm:"pk;auto"`
	LongUrl string `orm:"size(256);column(long_url)"`
	ShortUrl string `orm:"size(100);column(short_url)"`
	Sign string `orm:"size(32)"`
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp)"`
}

func (url *Url) TableName() string  {
	return "urls"
}

func (url *Url) TableUnique() [][]string  {
	return [][]string{
		[]string{"Sign"},
	}
}

// GenerateShortUrl 根据长链接生成短链接
func GenerateShortUrl(longUrl string) (*Url, error)  {
	db := orm.NewOrm()
	url := &Url{}
	var sina_url *utils.SinaUrls
	sina_url = utils.GetShortUrl(longUrl)
	if sina_url == nil {
		return nil, errors.New("新浪接口请求错误")
	}
	url.LongUrl = sina_url.Urls[0].UrlLong
	url.ShortUrl = sina_url.Urls[0].UrlShort
	url.Sign = utils.GetMD5(url.LongUrl)
	_, err := db.Insert(url)
	if err != nil {
		return nil, errors.New("url新增错误")
	}
	return url, nil
}

// 验证url是否存在
func IsLongUrlExist(url string) *Url  {
	db := orm.NewOrm()
	sign := utils.GetMD5(url)

	urlObj := &Url{
		Sign: sign,
	}
	err := db.Read(urlObj, "Sign")
	if err != nil {
		return nil
	}
	return urlObj
}