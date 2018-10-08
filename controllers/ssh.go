package controllers

import (
	"encoding/json"
	"fmt"
	"installer/models"
)

type SSHController struct {
	BaseController
}

func (c *SSHController) GetConnect() {
	var sshHost models.SSHHost
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &sshHost)
	info := models.Dossh(sshHost.Username, sshHost.Password, sshHost.Host, sshHost.Key, sshHost.CmdList, sshHost.Port, sshHost.CipherList)
	c.Data["json"] = info
	c.ServeJSON()
}
