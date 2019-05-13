package strfns

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Pagination ...
type Pagination struct {
	Now   int
	Num   int64
	Page  int64 // 查询总页数
	Begin int
	Pb    int64
	Pe    int64
	Ps    []bool
	URL   string
}

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

// IsMail ...
func IsMail(email string) (ok bool) {
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

// Md5 ... md5 加密
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

// Mail ...
// usr 	= "1679153844@qq.com"
// host	= "smtp.qq.com:25"
// pwd 	= "********"
func Mail(from, to, host, pwd, subject, body, mailtype string) (err error) {
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", from, pwd, hp[0])
	msg := []byte("To: " + to + "\r\n From:" + from + "<" + from + ">\r\n Subject:" + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	send := strings.Split(to, ";")
	err = smtp.SendMail(host, auth, from, send, msg)
	return
}

// NewPagination ...  分页函数
func NewPagination(now, num, row int64, url string) (p *Pagination) {
	var (
		begin, end int64
		col        int64 = 10 // 每页 展示页码数
	)
	p = &Pagination{
		Now: int(now),
		Num: num,
		URL: url,
	}
	if num == 0 {
		p.Page = 0
		p.Num = 0
	} else {
		p.Page = Ceil(num, row) // 查询总页数
		pages := make([]bool, p.Page)
		pages[now] = true

		// 如果总页数大于应展示页数
		if p.Page > col {
			// 当前的页码数 和 应展示页数 的关系
			if now > (col - 1) {
				if times := Floor(now, col); times < Floor(p.Page, col) {
					begin = times * col
					end = begin + col
					p.Ps = pages[begin:end]
				} else {
					begin = times * col
					p.Ps = pages[begin:]
				}
				p.Begin = int(begin)
			} else {
				p.Ps = pages[:col]
			}
		} else {
			p.Ps = pages
		}

		if num > col {
			if (now + 1) < p.Page {
				p.Pe = (now + 1) * row
			} else {
				p.Pe = num
			}

		} else {
			p.Pe = num
		}
		p.Pb = now*row + 1
	}
	return
}

// ReadCsv ... 读取 csv文件
func ReadCsv(file string) (tr [][]string) {
	r, _ := os.Open(file)
	read := csv.NewReader(r)
	tr, _ = read.ReadAll()
	return
}
