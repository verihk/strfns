package strfns

import (
	"strconv"
	"strings"
	"time"
)

// St1 替换所有非字符串
func St1(tr string) string {
	re := []string{` `, `_`, `-`, `,`, `.`, `/`, `\\`, `(`, `)`, `{`, `}`, `[`, `]`, `|`, `*`, `!`, `@`, `#`, `$`, `%`, `^`, `&`, `+`, `=`, `:`, `;`, `'`, `"`}
	for _, v := range re {
		tr = strings.Replace(tr, v, ``, -1)
	}
	return tr
}

// St2 替换 空格 和 ','
func St2(tr string) string {
	// tr = strings.Replace(tr, ` `, ``, -1)
	tr = St3(tr)
	tr = strings.Replace(tr, `,`, `;`, -1)
	return tr
}

// St3 替换 空格
func St3(tr string) string {
	tr = strings.Replace(tr, ` `, ``, -1)
	return tr
}

// TimeString 获取时间戳 字符串格式
func TimeString() (name string) {
	name = strconv.FormatInt(time.Now().Unix(), 10)
	return
}

// Slice2Map ...
func Slice2Map(b []interface{}) (bb map[int]interface{}) {
	bb = make(map[int]interface{})
	for k, v := range b {
		bb[k] = v
	}
	return
}

// SliceUnpeat 切片去重 ...
func SliceUnpeat(s1 []interface{}) (s2 []interface{}, l int) {
	m := make(map[interface{}]interface{})
	for k, v := range s1 {
		m[v] = k
	}
	l = len(m)
	s2 = make([]interface{}, l)
	// for k, _ := range m {
	for k := range m {
		s2 = append(s2, k)
	}
	return
}
