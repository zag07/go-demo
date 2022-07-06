package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bits-and-blooms/bitset"
)

func main() {
	filepath := "/Users/zhangsheng/Desktop/test.txt"
	f, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b bitset.BitSet

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		// hash
		if b.Test(String(Reverse(line))) {
			fmt.Println("反转字符串：", line)
		} else {
			b.Set(String(line))
		}

		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				log.Fatal(err)
			}
		}
	}
}

func String(s string) uint {
	v := uint(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
