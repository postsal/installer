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

func TestShellFile() {
	command := "./deploy/load_images.sh"
	for key, value := range images {
		fmt.Println(key)
		cmd := exec.Command("/bin/bash", "-c", command, key, value)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	}
}

func RemoveImage(toRmImage map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for value, bl := range toRmImage {
		if !bl {
			continue
		}
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

func TagImage() map[string]bool {
	result := make(map[string]bool)
	for key, value := range images {
		command := "docker tag " + value + " " + key
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

func PullImage() {
	for key, value := range images {
		if isImageExists(key) {
			continue
		}
		command := "docker pull " + value
		str := strings.Split(command, " ")
		cmd := exec.Command(str[0], str[1:]...)
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		}
	}
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
func CheckKubectlExists() (bool, string) {
	path, err := exec.LookPath("kubectl")
	if err != nil {
		return false, ""
	} else {
		return true, path
	}
}
