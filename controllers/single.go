package controllers

import (
	"installer/models"
	"installer/utils"
)

type SingleController struct {
	BaseController
}

func (c *SingleController) DeployK8sDashboard() {
	info := models.DeployK8sDashboard()
	c.Data["json"] = info
	c.ServeJSON()
}
func (c *SingleController) KubeProxy() {
	info := models.KubeProxy()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) CheckAndUseConfig() {
	info := models.CheckAndUseConfig()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) CheckKubectlExisted() {
	info := models.CheckKubectlPath()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) CheckDockerExisted() {
	info := models.CheckDockerPath()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) GetImagesStatus() {
	info := models.GetImagesStatus()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) PullImages() {
	info := models.PullImages()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) TagImages() {
	info := models.TagImages()
	c.Data["json"] = info
	c.ServeJSON()
}
func (c *SingleController) RemoveImages() {
	info := models.RemoveImages()
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SingleController) Test() {
	utils.ExecSH()
	c.Data["json"] = "test"
	c.ServeJSON()
}
