package utility

import (
	"strings"

	"golang.org/x/text/width"
)

// 半角全角変換
func HalfWidthConvertFullWidth(text string) string {
	return width.Narrow.String(text)
}

// 半角英小文字を与え、大文字小文字半角全角それぞれのパターンを返す（半角英小文字は返さない）
// @Param text String 半角英小文字
func BuildReplaceString(text string) []string {
	return StringArrayUnique([]string{HalfWidthConvertFullWidth(text), strings.ToUpper(text), HalfWidthConvertFullWidth(strings.ToUpper(text))})
}

// 文字配列をユニークにして返す
func StringArrayUnique(arr []string) []string {
	array := make(map[string]struct{})
	unique := []string{}

	for _, ele := range arr {
		array[ele] = struct{}{}
	}
	for arr := range array {
		unique = append(unique, arr)
	}

	return unique
}
