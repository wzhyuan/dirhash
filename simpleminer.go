package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	//	"reflect"
)

func main() {
	//sha1
	k := 1
	l := 0
	f := "abcddfg"
	for true {
		h := sha1.New()
		d := strconv.Itoa(k)
		s := f + d
		z := 0
		fmt.Println(s)
		fmt.Println("\n")
		io.WriteString(h, s)
		data := h.Sum(nil)
		fmt.Printf("%x\n", data)
		//fmt.Print(data[0])
		fmt.Println(data[0])
		for l = 0; l < len(data); l++ {
			if data[l] == 0 {
				z++
			}
		}
		if z > 3 {
			fmt.Println(" get it ")
			fmt.Println(data)
			fmt.Println(s)
			break
		}
		k++
		fmt.Println(k)
	}
}
