package model

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// get的参数会根据表单的值进行对应
type People struct {
	Name string `json:"name" form:"name"`
	Age  string `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

// xml数据进行解析绑定
type Article2 struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}
