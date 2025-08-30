package models

import "github.com/yyle88/gormcnm"

func (c *Student) Columns() *StudentColumns {
	return &StudentColumns{
		V班级编码: gormcnm.Cnm(c.ClassCode, "class_code"),
		V学号:   gormcnm.Cnm(c.Num, "num"),
		V名字:   gormcnm.Cnm(c.Name, "name"),
		V性别:   gormcnm.Cnm(c.Sex, "sex"),
		V生日:   gormcnm.Cnm(c.BornDate, "age"),
	}
}

type StudentColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	V班级编码 gormcnm.ColumnName[string]
	V学号   gormcnm.ColumnName[int]
	V名字   gormcnm.ColumnName[string]
	V性别   gormcnm.ColumnName[SexType]
	V生日   gormcnm.ColumnName[string]
}

func (c *Class) Columns() *ClassColumns {
	return &ClassColumns{
		V班级编码: gormcnm.Cnm(c.V班级编码, "class_code"),
		V班级名称: gormcnm.Cnm(c.V班级名称, "class_name"),
		V班主任名: gormcnm.Cnm(c.V班主任名, "main_teacher_name"),
	}
}

type ClassColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	V班级编码 gormcnm.ColumnName[string]
	V班级名称 gormcnm.ColumnName[string]
	V班主任名 gormcnm.ColumnName[string]
}
