package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

var images = map[string]string{
	"k8s.gcr.io/kube-proxy-amd64:v1.10.3":              "52.83.113.145:5000/k8s.gcr.io/kube-proxy-amd64:v1.10.3",
	"k8s.gcr.io/kube-controller-manager-amd64:v1.10.3": "52.83.113.145:5000/k8s.gcr.io/kube-controller-manager-amd64:v1.10.3",
	"k8s.gcr.io/kube-scheduler-amd64:v1.10.3":          "52.83.113.145:5000/k8s.gcr.io/kube-scheduler-amd64:v1.10.3",
	"k8s.gcr.io/kube-apiserver-amd64:v1.10.3":          "52.83.113.145:5000/k8s.gcr.io/kube-apiserver-amd64:v1.10.3",
	"k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.8":    "52.83.113.145:5000/k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.8",
	"k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.5":    "52.83.113.145:5000/k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.5",
	"k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.8":          "52.83.113.145:5000/k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.8",
	"k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.5":          "52.83.113.145:5000/k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.5",
	"k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.8":         "52.83.113.145:5000/k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.8",
	"k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.5":         "52.83.113.145:5000/k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.5",
	"k8s.gcr.io/pause-amd64:3.1":                       "52.83.113.145:5000/k8s.gcr.io/pause-amd64:3.1",
	"k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.0":    "52.83.113.145:5000/k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.0",
	"k8s.gcr.io/etcd-amd64:3.1.12":                     "52.83.113.145:5000/k8s.gcr.io/etcd-amd64:3.1.12",
}

func Pwd() string {
	cmd := exec.Command("pwd")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "pwd", err.Error())
		return ""
	} else {
		pwd := string(output)
		fmt.Printf("pwd : %s", pwd)
		return pwd
	}
}

func ExecSH() {
	command := "deploy/load_images.sh"
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

/**
deploy
**/

func DeployK8sDashboard() bool {
	command := "kubectl apply -f deploy/kubernetes-dashboard.yaml"
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

func GetImagesStatus() map[string]bool {
	result := make(map[string]bool)
	for k, _ := range images {
		if isImageExists(k) {
			result[k] = true
		} else {
			result[k] = false
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
