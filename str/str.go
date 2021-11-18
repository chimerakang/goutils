package str

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// 字符串转uint数组, 默认逗号分割
func Str2UintArr(str string) (ids []uint) {
	idArr := strings.Split(str, ",")
	for _, v := range idArr {
		ids = append(ids, Str2Uint(v))
	}
	return
}

// 字符串转uint
func Str2Uint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(num)
}

var (
	camelRe = regexp.MustCompile("(_)([a-zA-Z]+)")
	snakeRe = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// 字符串转为驼峰
func CamelCase(str string) string {
	camel := camelRe.ReplaceAllString(str, " $2")
	camel = strings.Title(camel)
	camel = strings.Replace(camel, " ", "", -1)
	return camel
}

// 字符串转为驼峰(首字母小写)
func CamelCaseLowerFirst(str string) string {
	camel := CamelCase(str)
	for i, v := range camel {
		return string(unicode.ToLower(v)) + camel[i+1:]
	}
	return camel
}

// 驼峰式写法转为下划线蛇形写法
func SnakeCase(str string) string {
	snake := snakeRe.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

// 加密base64字符串
func EncodeStr2Base64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解密base64字符串
func DecodeStrFromBase64(str string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(decodeBytes)
}

// Length returns number of characters
func Length(s string) int {
	return len([]rune(s))
}

// Before returns the string before the first occurrence of the substr string
func Before(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.Index(s, substr)
	if i != -1 {
		return s[:i]
	}
	return s
}

// BeforeLast returns the string before the last occurrence of the substr string
func BeforeLast(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.LastIndex(s, substr)
	if i != -1 {
		return s[:i]
	}
	return s
}

// After returns the string after the first occurrence of the substr string
func After(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.Index(s, substr)
	if i != -1 {
		i = i + len(substr)
		return s[i:]
	}
	return s
}

// AfterLast returns the string after the last occurrence of the substr string
func AfterLast(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.LastIndex(s, substr)
	if i != -1 {
		i = i + len(substr)
		return s[i:]
	}
	return s
}

func Index(s, substr string) int {
	return strings.Index(s, substr)
}

func RuneIndex(s, substr string) int {
	p := strings.Index(s, substr)
	if p == -1 || p == 0 {
		return p
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(s)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		pos++
		totalSize += size

		if totalSize == p {
			return pos
		}
	}
	return pos
}

func Contians(s, substr string) bool {
	return strings.Contains(s, substr)
}

func StartsWith(s, substr string) bool {
	if substr != "" && Substr(s, 0, len([]rune(substr))) == substr {
		return true
	}
	return false
}

func EndsWith(s, substr string) bool {
	if Substr(s, -len([]rune(substr)), len(s)) == substr {
		return true
	}
	return false
}

// Substr returns a string of length length from the start position
func Substr(s string, start int, strlength ...int) string {
	charlist := []rune(s)
	l := len(charlist)
	length := 0
	end := 0

	if len(strlength) == 0 {
		length = l
	} else {
		length = strlength[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}

	return string(charlist[start:end])
}

func SubByte(str string, length int) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

// Char returns a char slice
func Char(str string) []string {
	c := make([]string, 0)
	for _, v := range str {
		c = append(c, string(v))
	}
	return c
}

func Escape(s string) string {
	str := strconv.Quote(s)
	str = strings.Replace(str, "'", "\\'", -1)
	strlist := []rune(str)
	l := len(strlist)
	return Substr(str, 1, l-2)
}

func Ufirst(s string) string {
	r := []rune(s)
	if len(s) > 0 && unicode.IsLetter(r[0]) && unicode.IsLower(r[0]) {
		r[0] -= 32
		return string(r)
	}
	return s
}

// String returns a string of any type
func String(iface interface{}) string {
	switch val := iface.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(iface)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", iface)
}
