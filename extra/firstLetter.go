package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type colorGroup struct {
	ID    int      `json:"id"` // 如果希望，json化之后的属性名是小写字母的，可以使用struct tag
	name  string   // name小写，该字段无法被外部包访问和解析
	Color []string `json:"color"`
}

func main() {
	group := colorGroup{
		ID:    1,
		name:  "Reds",
		Color: []string{"Red", "Pink", "Ruby"},
	}
	res, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error marshalling", err)
	}
	os.Stdout.Write(res)
}

// output:
// {"id":1,"color":["Red","Pink","Ruby"]}
