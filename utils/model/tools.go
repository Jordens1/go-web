package model

import "time"

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// 模板函数的实现
func UnixToTime(ts int) string {
	t := time.Unix(int64(ts), 0)
	return t.Format("2006-01-02 15:04:05")
}
