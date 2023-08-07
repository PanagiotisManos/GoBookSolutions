/* In the code below, the iota starts at '0' and increments by '1' for each constant declaration.
The '<<' operator is used for bit-shifting to calculate the value of each constant in bytes
(1 KB = 1024 bytes, 1 MB = 1024 KB, etc.). However, any further use of iota would lead to an
overflow because iota can only represent up to 2^63-1 in signed integers, and we need to go
beyond that to represent ZB and YB. For that reason, we assigned the two constants as strings
giving them the exact value of a zettabyte and a yottabyte. */

package main

const (
	_         = iota
	KB uint64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB = "1180591620717411303424"
	YB = "1208925819614629174706176"
)

func main() {
	println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
