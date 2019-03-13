package byteutil

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// EncodeUCS2ToStr ucs2编码
func EncodeUCS2ToStr(txt string) string {
	list := []rune(txt)
	result := ""
	for _, item := range list {
		result += fmt.Sprintf("%04X", int(item))
	}
	return result
}

// DecodeUCS2ToStr ucs2解码
func DecodeUCS2ToStr(txt string) string {
	if len(txt)%4 != 0 {
		return txt
	}

	r := []string{}
	for i := 0; i < len(txt); i += 4 {
		t, _ := strconv.ParseInt(txt[i:i+4], 16, 32)
		r = append(r, fmt.Sprintf("%c", t))
	}
	return strings.Join(r, "")
}

// Decode7bitToStr 7bit解码
// 31D98C06 => 1234
func Decode7bitToStr(str string) string {
	binStr := ""
	for i := len(str); i > 0; i -= 2 {
		ii, _ := strconv.ParseInt(str[i-2:i], 16, 32)
		binStr += int32toBinary(int(ii))
	}
	result := ""
	for i := len(binStr); i > 0; i -= 7 {
		if i > 7 {
			result += fmt.Sprintf("%c", binStrToInt(binStr[i-7:i]))
		} else {
			result += fmt.Sprintf("%c", binStrToInt(binStr[:i]))
		}
	}
	return result
}

func int32toBinary(orig int) string {
	r := ""
	i := 7
	for orig != 0 {
		t := orig % 2
		r = fmt.Sprintf("%d", t) + r
		i--
		orig /= 2
	}
	for i >= 0 {
		r = "0" + r
		i--
	}
	return r
}

// Encode7bitToStr 7bit编码
// 1234 => 31D98C06
func Encode7bitToStr(str string) string {
	tStr := []string{}
	for _, v := range str {
		tStr = append(tStr, int32toBinary(int(v))[1:])
	}

	binStr := ""
	for i := len(tStr) - 1; i >= 0; i-- {
		binStr += tStr[i]
	}
	r := ""
	for i := len(binStr); i > 0; i -= 8 {
		if i > 8 {
			r += fmt.Sprintf("%02X", binStrToInt(binStr[i-8:i]))
		} else {
			r += fmt.Sprintf("%02X", binStrToInt(binStr[:i]))
		}
	}
	return r
}

func binStrToInt(orig string) int {
	r := 0
	for _, i := range orig {
		r *= 2
		if string(i) == "1" {
			r++
		}
	}
	return r
}

// IsChiness 是否有中文
func IsChiness(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
