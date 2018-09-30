package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

var images = map[string]string{
	"k8s.gcr.io/kube-proxy-amd64:v1.10.3":              "52.83.113.145:5000/k8s.gcr.io/kube-proxy-amd64:v1.10.3",
	"k8s.gcr.io/kube-controller-manager-amd64:v1.10.3": "52.83.113.145:5000/k8s.gcr.io/kube-controller-manager-amd64:v1.10.3",
}

/**
deploy
**/

func DeployK8sDashboard() bool {
	command := "kubectl apply -f /Users/chainnova/work/go/src/installer/deploy/kubernetes-dashboard.yaml"
	str := strings.Split(command, " ")
	cmd := exec.Command(str[0], str[1:]...)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return false
	} else {
		return true
	}
}

func KubeProxy() bool {
	command := "kubectl proxy &"
	str := strings.Split(command, " ")
	cmd := exec.Command(str[0], str[1:]...)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return false
	} else {
		return true
	}

}

//切换到local config
func CheckAndUseConfig() bool {
	command := "kubectl config use-context docker-for-desktop"
	str := strings.Split(command, " ")
	cmd := exec.Command(str[0], str[1:]...)
	_, err := cmd.Output()
	if err != nil {
		return false
	} else {
		return true
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

/**
pull images
***/

func RemoveImages() map[string]bool {
	result := make(map[string]bool)
	for _, value := range images {
		command := "docker rmi " + value
		str := strings.Split(command, " ")
		cmd := exec.Command(str[0], str[1:]...)
		_, err := cmd.Output()
		if err != nil {
			result[value] = false
		} else {
			result[value] = true
		}
	}
	return result
}

func TagImages() map[string]bool {
	result := make(map[string]bool)
	for key, value := range images {
		command := "docker tag " + value + " " + key
		str := strings.Split(command, " ")
		cmd := exec.Command(str[0], str[1:]...)
		_, err := cmd.Output()
		if err != nil {
			result[key] = false
		} else {
			result[key] = true
		}
	}
	return result
}

func PullImages() map[string]bool {
	result := make(map[string]bool)
	for key, value := range images {
		if isImageExists(key) {
			result[key] = true
			continue
		}
		command := "docker pull " + value
		str := strings.Split(command, " ")
		cmd := exec.Command(str[0], str[1:]...)
		_, err := cmd.Output()
		if err != nil {
			result[value] = false
		} else {
			result[value] = true
		}
	}
	return result
}

func isImageExists(image string) bool {
	cmd := exec.Command("docker", "inspect", image)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}

func CheckDockerExists() (bool, string) {
	path, err := exec.LookPath("docker")
	if err != nil {
		return false, ""
	} else {
		return true, path
	}
}
