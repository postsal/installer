package models

const STATE_SUCCESS string = "success"
const STATE_ERROR string = "error"

type RetStruct struct { //all interface return struct
	Status  string      `json:status`  //success or error
	Message string      `json:message` //description message
	Data    interface{} `json:data`    //return data in this member
}

func GetJsonValue(state string, msg string, data interface{}) RetStruct {
	var ret RetStruct
	ret.Status = state
	ret.Message = msg
	ret.Data = data
	return ret
}
