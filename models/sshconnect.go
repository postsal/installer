package models

import (
	"bytes"
	"fmt"
	"installer/utils"
	"strings"
)

type SSHHost struct {
	Host       string
	Port       int
	Username   string
	Password   string
	CipherList []string
	Key        string
	CmdList    []string
	Result     SSHResult
}

type HostJson struct {
	SshHosts []SSHHost
}

type SSHResult struct {
	Host    string
	Success bool
	Result  string
}

func Dossh(username, password, host, key string, cmdlist []string, port int, cipherList []string) RetStruct {
	session, err := utils.Connect(username, password, host, key, port, cipherList)
	var sshResult SSHResult
	sshResult.Host = host

	if err != nil {
		sshResult.Success = false
		sshResult.Result = fmt.Sprintf("<%s>", err.Error())
		return GetJsonValue(STATE_ERROR, "", sshResult)
	}
	defer session.Close()

	cmdlist = append(cmdlist, "exit")
	newcmd := strings.Join(cmdlist, "&&")

	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt

	session.Stderr = &errbt
	err = session.Run(newcmd)
	if err != nil {
		sshResult.Success = false
		sshResult.Result = fmt.Sprintf("<%s>", err.Error())
		return GetJsonValue(STATE_ERROR, "", sshResult)
	}

	if errbt.String() != "" {
		sshResult.Success = false
		sshResult.Result = errbt.String()
	} else {
		sshResult.Success = true
		sshResult.Result = outbt.String()
	}
	return GetJsonValue(STATE_SUCCESS, "", sshResult)
}
