package tools

import (
	"crypto/md5"
	"encoding/hex"
	"path"
	"strconv"
)

/* 字符串转md5 */
func StringToMd5(str string, length int64) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	data := hex.EncodeToString(md5Ctx.Sum(nil))
	if length == 0 {
		return data
	} else {
		return data[0:length]
	}
}

/* 字节转兆 */
func ByteToMegabyte(data string) string {
	base, _ := strconv.Atoi(data)
	result := float64(base) / (1024 * 1024)
	return strconv.FormatFloat(result, 'f', 2, 32)
}

/* 字符串转int */
func StringToInt(data string) int {
	num, _ := strconv.Atoi(data)
	return num
}

/* 字符串转int64 */
func StringToInt64(data string) int64 {
	num, _ := strconv.ParseInt(data, 10, 64)
	return num
}

/* 整形64转字符串 */
func Int64ToString(timeStamp int64) string {
	return strconv.FormatInt(timeStamp, 10)
}

/* 整形转字符串 */
func IntToString(timeStamp int) string {
	return strconv.Itoa(timeStamp)
}

/* 获得文件后缀 */
func GetFileSuffix(tempname string) string {
	return path.Ext(tempname)
}
