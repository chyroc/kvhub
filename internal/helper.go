package internal

import (
	"crypto/md5"
	"fmt"
)

func trimPrefixPath(s string) string {
	ii := []int32(s)
	for i, s := range ii {
		if s == '.' || s == '/' {
			continue
		}
		return string(ii[i:])
	}
	return ""
}

func getMd5(s []byte) string {
	m := md5.New()
	m.Write(s)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func splitBySizeChat(s string, size int) (res []string) {
	tmp := []int32{}
	for _, v := range []int32(s) {
		if len(tmp) < size {
			tmp = append(tmp, v)
		} else {
			res = append(res, string(tmp))
			tmp = []int32{}
		}
	}
	if len(tmp) > 0 {
		res = append(res, string(tmp))
	}
	return res
}
