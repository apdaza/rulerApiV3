// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/apdaza/rulerApiV3/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins: true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
			AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/dominio",
			beego.NSInclude(
				&controllers.DominioController{},
			),
		),

		beego.NSNamespace("/tipo_predicado",
			beego.NSInclude(
				&controllers.TipoPredicadoController{},
			),
		),

		beego.NSNamespace("/predicado",
			beego.NSInclude(
				&controllers.PredicadoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
