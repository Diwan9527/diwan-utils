package main

import (
	"fmt"
	"testing"

	"github.com/Diwan9527/diwan-utils/utils"
)

func Test123(t *testing.T) {
	test_map := make(map[string]interface{})
	test_map["ss"] = "test"
	// data, err := yaml.Marshal(test_map)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	str, err := utils.CreateFile("/Users/wangui.ding/workspace/go_workspace/diwan-utils/test/test111.yaml", test_map, "yaml")
	if err != nil {
		fmt.Println(err)
		fmt.Println(str)
	}
}
