package models

import "installer/utils"

func DeployK8sDashboard() RetStruct {
	if utils.DeployK8sDashboard() {
		return GetJsonValue(STATE_SUCCESS, "K8S dashboard has deployed", "http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/")
	} else {
		return GetJsonValue(STATE_ERROR, "K8S dashboard deploy failed", "")
	}
}
func KubeProxy() RetStruct {
	if utils.KubeProxy() {
		return GetJsonValue(STATE_SUCCESS, "Starting to serve on 127.0.0.1:8001", "")
	} else {
		return GetJsonValue(STATE_ERROR, "start failed", "")
	}
}
func CheckAndUseConfig() RetStruct {
	if utils.CheckAndUseConfig() {
		return GetJsonValue(STATE_SUCCESS, "kubectl configuration has switched to docker-for-desktop", "")
	} else {
		return GetJsonValue(STATE_ERROR, "kubectl can not find docker-for-desktop config", "")
	}
}

func CheckKubectlPath() RetStruct {
	ok, path := utils.CheckKubectlExists()
	if ok {
		return GetJsonValue(STATE_SUCCESS, "kubectl existed", path)
	} else {
		return GetJsonValue(STATE_SUCCESS, "kubectl not existed", "")
	}
}

//检查是否安装docker
func CheckDockerPath() RetStruct {
	ok, path := utils.CheckDockerExists()
	if ok {
		return GetJsonValue(STATE_SUCCESS, "docker existed", path)
	} else {
		return GetJsonValue(STATE_SUCCESS, "docker not existed", "")
	}
}

// 查看基础镜像拉取的进度
func GetImagesStatus() RetStruct {
	result := utils.GetImagesStatus()
	return GetJsonValue(STATE_SUCCESS, "get images status list", result)
}

//拉取启动kubernetes 所需的镜像，并返回成功与否的结果
func PullImages() RetStruct {
	result := utils.PullImages()
	return GetJsonValue(STATE_SUCCESS, "temporary mirror pull state list", result)
}

//为镜像打tag
func TagImages() RetStruct {
	result := utils.TagImages()
	return GetJsonValue(STATE_SUCCESS, "mirror tag state list", result)
}

//删除临时镜像
func RemoveImages() RetStruct {
	result := utils.RemoveImages()
	return GetJsonValue(STATE_SUCCESS, "temporary mirror remove state list", result)
}
