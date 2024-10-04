package example4

import "github.com/yyle88/gormcnm"

func (*Student) Columns() *StudentColumns {
	return &StudentColumns{
		V班级编码: "class_code",
		V学号:   "num",
		V名字:   "name",
		V性别:   "sex",
		V生日:   "age",
	}
}

type StudentColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	V班级编码 gormcnm.ColumnName[string]
	V学号   gormcnm.ColumnName[int]
	V名字   gormcnm.ColumnName[string]
	V性别   gormcnm.ColumnName[SexType]
	V生日   gormcnm.ColumnName[string]
}

func (*Class) Columns() *ClassColumns {
	return &ClassColumns{
		V班级编码: "class_code",
		V班级名称: "class_name",
		V班主任名: "main_teacher_name",
	}
}

type ClassColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	V班级编码 gormcnm.ColumnName[string]
	V班级名称 gormcnm.ColumnName[string]
	V班主任名 gormcnm.ColumnName[string]
}
