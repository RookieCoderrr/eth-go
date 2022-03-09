package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

// usage: HexToBigInt("0x15b048467613093f818")
func HexToBigInt(str string) *big.Int {
	n := new(big.Int)
	if len(str) > 2 {
		str = str[2:]
	}
	n, _ = n.SetString(str, 16)
	return n
}

// 十六进制转十进制方法
func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

// 移除左侧所有的0
func TrimLeftZeroes(hex string) string {
	idx := 0
	for ; idx < len(hex); idx++ {
		if hex[idx] != '0' {
			break
		}
	}
	return hex[idx:]
}