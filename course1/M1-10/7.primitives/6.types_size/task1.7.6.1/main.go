package main

import "unsafe"

func main() {

}
func sizeOfBool(b bool) int {
i := unsafe.Sizeof(b)
return int(i)
}
func sizeOfInt(n int) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
func sizeOfInt8(n int8) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
func sizeOfInt16(n int16) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
func sizeOfInt32(n int32) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
func sizeOfInt64(n int64) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
func sizeOfUint(n uint) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
func sizeOfUint8(n uint8) int {
	i := unsafe.Sizeof(n)
	return int(i)
}
