# Golang首字母大小写问题
---
>  Go语言通过首字母的大小写来控制访问权限。无论是方法，变量，常量或是自定义的变量类型，如果首字母大写，则可以被外部包访问，反之则不可以。
>
> 而结构体中的字段名，如果首字母小写的话，则该字段无法被外部包访问和解析，比如，json解析。

```go
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
```

**输出：**

```
{"id":1,"color":["Red","Pink","Ruby"]}
```

## struct tag补充

- tag中含有`-`的字段不会输出；

- tag中含有`omitempty`，如果该字段的值为空值那么该字段就不会被输出，空值包括：false、0、nil指针或nil接口，任何长度为0的array, slice, map或者string。

