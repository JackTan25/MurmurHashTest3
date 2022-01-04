package main

import (
	"fmt"
)

func main() {
	var input_data string = "Jack"
	var seed uint32 = 123
	fmt.Println("Murmurhash3(Jack,123) = ", Murmurhash3(input_data, seed))
}

func getBlkno(input_data string, idx int) uint32 {
	var s string = input_data[idx*4 : idx*4+4]
	var res uint32 = 0
	for i := 0; i < 4; i++ {
		res += uint32(s[i] << (24 - i*8))
	}
	return res
}
func _rotl(K uint32, r uint32) uint32 {
	return K<<r | K>>(32-r)
}
func Murmurhash3(input_data string, seed uint32) uint32 {
	var h uint32 = seed
	//声明常量
	const (
		c1 uint32 = 0xcc9e2d51 // 3,432,918,353
		c2 uint32 = 0x1b873593 // 461,845,907

		r1 uint32 = 15
		r2 uint32 = 13

		m uint32 = 5
		n uint32 = 0xe6546b64 //3,864,292,196
	)

	//分块处理
	var blkNums int = len(input_data) / 4
	//1.一个块一个块地处理,这是第一部分地工作
	for i := 0; i < blkNums; i++ {
		var K uint32 = getBlkno(input_data, i)
		K *= c1
		K = _rotl(K, r1)
		K *= c2
		K = _rotl(K, r2)
		h = h*m + n
	}
	//2.处理剩余量
	var remaining_bytes string = input_data[blkNums*4:]
	var k uint32 = 0
	switch len(remaining_bytes) {
	case 3:
		k ^= uint32(remaining_bytes[2] << 16)
	case 2:
		k ^= uint32(remaining_bytes[1] << 8)
	case 1:
		k ^= uint32(remaining_bytes[0])
	}
	k = k * c1
	k = _rotl(k, r1)
	k = k * c2
	h ^= k
	h ^= uint32(len(input_data))
	//3.加强雪崩测试
	h ^= h >> 16
	h *= 0x85ebca6b // 2,246,822,507
	h ^= h >> 13
	h *= 0xc2b2ae35 // 3,266,489,909
	h ^= h >> 16
	return h
}
/***************************************
 * Written By I am Jack(JackTan)
 * 
 * Murmurhash3 Test Programming
 **/