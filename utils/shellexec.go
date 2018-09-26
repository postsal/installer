package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func DeployK8sDashboard() {
	command := "./deploy_dashboard.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
func ClearImages() {
	command := "./clear_images.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func LoadImages() bool {
	command := "../deploy/curlBaidu.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return false
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	return true
}
func TestShellFile() string {
	command := "deploy/curlBaidu.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr + errStr
}

func CheckDockerExists() (bool, string) {
	path, err := exec.LookPath("docker")
	if err != nil {
		return false, ""
	} else {
		return true, path
	}
}
func CheckKubectlExists() (bool, string) {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		return false, ""
	} else {
		return true, path
	}
}
