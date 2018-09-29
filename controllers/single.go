package controllers

import "installer/utils"

type SingleController struct {
	BaseController
}

func (c *SingleController) TestShell() {
	result := utils.TagImage()
	utils.RemoveImage(result)
	c.Data["json"] = "test"
	c.ServeJSON()
}
func (c *SingleController) TestSSH() {
	utils.SimpleSSH()
	c.Data["json"] = "test SSH"
	c.ServeJSON()
}
