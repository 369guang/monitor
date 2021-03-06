package utils

import (
	"clients/config"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSystemInfo(t *testing.T) {
	data := SystemInfo()
	if data == nil {
		t.Error("SystemInfo failed!")
	}

	formatJson, _ := json.Marshal(data)
	fmt.Println(string(formatJson))
	t.Log("SystemInfo test pass")
}

func TestPushInfoData(t *testing.T) {
	err := config.Init("/Users/cengcanguang/work/clients/conf/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	data := SystemInfo()
	if data == nil {
		t.Error("TestSystemInfo failed!")
	}
	res := StructToMap(data)
	pushData(res, "/api/host/info/update")
	t.Log("PushData test pass")
}
