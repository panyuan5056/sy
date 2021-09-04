package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cast"
)

func BKDRHash(str string) string {
	seed := uint64(131) // 31 131 1313 13131 131313 etc..
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint64(str[i])
	}
	return cast.ToString(hash)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}

func StringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func Similarity(s, t string) float64 {
	s1 := []byte(s)
	t1 := []byte(t)
	max, min := 0, 0
	if len(s1) > len(t1) {
		max = len(s1)
		min = len(t1)
	} else {
		max = len(t1)
		min = len(s1)
	}
	total := 0
	for i := 0; i < min; i++ {
		if s1[i] == t1[i] {
			total = total + 1
		}
	}

	return cast.ToFloat64(total) / cast.ToFloat64(max)
}

func LevenshteinDistance(s, t string, ignoreCase bool) int {
	if ignoreCase {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}
