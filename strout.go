package strfns

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"regexp"
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

// TimeStr 获取时间戳 字符串格式
func TimeStr() (name string) {
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

// isMail ...
func isMail(email string) (ok bool) {
	if len(email) > 0 {
		ok, _ = regexp.MatchString("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$", email)
	}
	return
}

// RandStr ... 随机字符串
func RandStr() string {
	by := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, by); err != nil {
		return ``
	}
	return base64.URLEncoding.EncodeToString(by)
}

// RandPwd ... 系统自设密码时，字符转换， 可以是数字，字母，或者下划线 √√
func RandPwd() string {
	str := RandStr()
	re := []string{` `, `-`, `,`, `.`, `/`, `\\`, `(`, `)`, `{`, `}`, `[`, `]`, `|`, `*`, `!`, `@`, `#`, `$`, `%`, `^`, `&`, `+`, `=`, `:`, `;`, `'`, `"`}
	for _, v := range re {
		str = strings.Replace(str, v, ``, -1)
	}
	str = str[12:19]
	return str
}

// Md5 ...
func Md5(s string) string {
	w := md5.New()
	io.WriteString(w, s)
	return fmt.Sprintf("%x", w.Sum(nil))
}

// Ceil ... 向上取整
func Ceil(i, n int64) (j int64) {
	ii := float64(i)
	nn := float64(n)
	jj := math.Ceil(ii / nn)
	j = int64(jj)
	return
}

// Floor ... 向下取整
func Floor(i, n int64) (j int64) {
	ii := float64(i)
	nn := float64(n)
	jj := math.Floor(ii / nn)
	j = int64(jj)
	return
}
