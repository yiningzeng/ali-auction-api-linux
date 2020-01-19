// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/baymin/project/ali-auction-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/auction_item",
			beego.NSInclude(
				&controllers.AuctionItemController{},
			),
		),

		beego.NSNamespace("/auction_target",
			beego.NSInclude(
				&controllers.AuctionTargetController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
