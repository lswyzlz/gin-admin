package entity

import "github.com/LyricTian/gin-admin/internal/app/schema"

//SchemaCustomerInfo 客户信息对象
type SchemaCustomerInfo schema.CustomerInfo

//CustomerInfo 菜单实体
type CustomerInfo struct {
	Model
	Mobile string `json:"mobile" binding:"required"` //电话
	Name   string `json:"name"`                      //姓名
	Age    uint8  `json:"age"`                       //年龄
	Gender int8   `json:"gender"`                    //性别
	Note   string `json:"note"`                      //备注
}
