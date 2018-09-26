package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                                            //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")                     //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Requested-With") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}
func (c *BaseController) Options() {
	c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}
