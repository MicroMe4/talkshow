package convert

import "unsafe"

//UnsafeBytesToString []byte转string 函数
func UnsafeBytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

//UnsafeStringToBytes string 转 []byte 函数
func UnsafeStringToBytes(str string) []byte {
	mid := (*[2]uintptr)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&([3]uintptr{mid[0], mid[1], mid[1]})))
}
