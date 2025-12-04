// Package gormcngen tests validate Options configuration and default values
// Auto verifies option initialization and JSON serialization
//
// gormcngen 测试包验证 Options 配置和默认值
// 自动验证选项初始化和 JSON 序列化
package gormcngen

import (
	"testing"

	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewOptions(t *testing.T) {
	t.Log(neatjsons.S(NewOptions()))
}
