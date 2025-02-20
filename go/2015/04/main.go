package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Sovle1(instruction string) int {
	var formatted_res_str string
	var str_hash string
	for i := 0; i < 10000000; i++ {
		formatted_res_str = fmt.Sprintf("%s%d", instruction, i)
		str_hash = GetMD5Hash(formatted_res_str)
		if strings.HasPrefix(str_hash, "00000") {
			return i
		}
	}
	return -1
}

func Sovle2(instruction string) int {
	var formatted_res_str string
	var str_hash string
	for i := 0; i < 10000000; i++ {
		formatted_res_str = fmt.Sprintf("%s%d", instruction, i)
		str_hash = GetMD5Hash(formatted_res_str)
		if strings.HasPrefix(str_hash, "000000") {
			return i
		}
	}
	return -1
}

func main() {
	dat := "yzbqklnj"
	fmt.Println(Sovle1(dat))
	fmt.Println(Sovle2(dat))
}
