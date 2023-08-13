package utils

import "regexp"

func Normalization(s string) string {
	// 控制序列(又名 ANSI 转义序列)的完整正则表达式是`/(\x9B|\x1B\[)[0-?]*[ -\/]*[@-~]/`
	reg := regexp.MustCompile(`/(\x9B|\x1B\[)[0-?]*[ -\/]*[@-~]/`)
	return reg.ReplaceAllString(s, "")
}
