package ciphers

import (
	"testing"
	"fmt"
)

func TestRc4_md5_NewDecrypter(t *testing.T) {
	key := []byte{90, 17, 251, 106, 32, 39, 86, 23, 215, 124, 214, 28, 176, 4, 226, 110}
	iv := []byte{72, 182, 149, 38, 168, 65, 205, 146, 117, 148, 100, 50, 57, 145, 11, 235}
	rc4 := &rc4_md5{16, 16}
	d, err := rc4.NewDecrypter(key, iv)
	if err != nil {
		panic(d)
	}
	src := []byte{220, 38, 125, 109, 79, 90, 167, 48, 71, 50, 249, 242, 166, 170, 58, 64, 119, 21, 118}
	dst := make([]byte, len(src))
	d.XORKeyStream(dst, src)
	fmt.Println(dst)
	src = []byte{68, 206, 81, 149, 163, 250, 226, 250, 137, 157, 108, 146, 55, 243, 244, 162, 10, 250, 44, 141, 222, 98, 147, 207, 1, 134, 17, 80, 143, 236, 84, 21, 47, 14, 56, 92, 193, 27, 151, 168, 40, 116, 122, 238, 79, 70, 33, 95, 71, 195, 218, 124, 127, 185, 50, 3, 146, 215, 206, 52, 11, 250, 141, 125, 73, 79, 37, 28, 62, 40, 18, 19, 53, 119, 95, 60, 252, 174, 201, 60, 155, 113, 81, 203, 25, 129, 36, 185, 24, 55, 49, 132, 83, 34, 164, 193, 19, 168, 142, 132, 200, 24, 62, 179, 178, 227, 100, 159, 193, 160, 142, 233, 127, 84, 187, 113, 80, 133, 135, 255, 190}
	dst = make([]byte, len(src))
	d.XORKeyStream(dst, src)
	fmt.Println(dst)
}