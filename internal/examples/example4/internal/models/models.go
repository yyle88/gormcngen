package models

type SexType string

const (
	Male   SexType = "Male"
	Female SexType = "Female"
)

/*
假设开发者英语水平不高，只会些塑料英语，甚至有字段名对齐的强迫症，则允许开发者使用别名
*/

type Student struct {
	ClassCode string  `gorm:"column:class_code" cnm:"V班级编码"`
	Num       int     `gorm:"column:num;primaryKey;autoIncrement:true" cnm:"V学号"` //塑料英语，正确的是“student ID”
	Name      string  `gorm:"column:name" cnm:"V名字"`
	Sex       SexType `gorm:"column:sex" cnm:"V性别"` //塑料英语，正确的是“Gender”
	BornDate  string  `gorm:"column:age" cnm:"V生日"` //塑料英语，正确的是“Date of Birth”
}

/*
彻底放弃治疗直接使用汉语编程，毕竟汉语天然的表意比较全面，而且非常精悍，而且字段名长度相同有利于代码对齐和整洁（这点对强迫症非常友善）
*/

type Class struct {
	V班级编码 string `gorm:"column:class_code;primaryKey;"`
	V班级名称 string `gorm:"column:class_name;unique"`
	V班主任名 string `gorm:"column:main_teacher_name;unique;"`
}
