// @APIVersion 1.0.0
// @Title 短连接服务API
// @Description 使用Beego实现短连接服务
// @Contact jyx15221613915@gmail.com
// @TermsOfServiceUrl http://blog.jiangyixin.top
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-short-url/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/short",
			beego.NSInclude(
				&controllers.ShortController{},
			),
		),
		beego.NSNamespace("/expand",
			beego.NSInclude(
				&controllers.ExpandController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
