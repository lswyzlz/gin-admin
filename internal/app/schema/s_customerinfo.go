package schema

//CustomerInfo 客户信息
type CustomerInfo struct {
	Mobile string `json:"mobile" binding:"required"` //电话
	Name   string `json:"name"`                      //姓名
	Age    uint8  `json:"age"`                       //年龄
	Gender int8   `json:"gender"`                    //性别
	Note   string `json:"note"`                      //备注
}
