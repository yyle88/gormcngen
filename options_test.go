package gormcngen

import (
	"testing"

	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewOptions(t *testing.T) {
	t.Log(neatjsons.S(NewOptions()))
}
