package strfns

import "strings"

// St1 替换所有非字符串
func St1(tr string) string {
	re := []string{` `, `_`, `-`, `,`, `.`, `/`, `\\`, `(`, `)`, `{`, `}`, `[`, `]`, `|`, `*`, `!`, `@`, `#`, `$`, `%`, `^`, `&`, `+`, `=`, `:`, `;`, `'`, `"`}
	for _, v := range re {
		tr = strings.Replace(tr, v, ``, -1)
	}
	return tr
}

func St2(tr string) string {
	tr = strings.Replace(tr, ` `, ``, -1)
	tr = strings.Replace(tr, `,`, `;`, -1)
	return tr
}

func St3(tr string) string {
	tr = strings.Replace(tr, ` `, ``, -1)
	return tr
}
