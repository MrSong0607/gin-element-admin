package test

import (
	"admin/service"
	"admin/utils"
	"testing"
)

func TestPassword(t *testing.T) {
	b, _ := service.EncryptPassword("admin")
	t.Log(utils.Hex(b))
}
