package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	rb := bytes.NewBuffer([]byte("hello world"))
	r := bufio.NewReader(rb)
	var buf [128]byte
	n, _ := r.Read(buf[:])
	fmt.Println(string(buf[:n]))

}
