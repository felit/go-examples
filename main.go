package main

import (
	"os"
	"strings"
	"fmt"
	"time"
	"math/rand"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	who := "world!"

	println("helloworld")
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
	write_file_test()
	//write_buf()
}
func write_file_test() {
	filename := "/important/data/tmp/hehehe.txt"
	file1, error := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeSetuid)
	defer file1.Close()
	os.Chmod(filename, 7666)
	if error != nil {
		panic(error)
	}
	t1 := time.Now().UnixNano()
	num := 1000000
	for a := 0; a < num; a++ {
		//bb := bytes.NewBuffer([]byte("Hello"))
		//bb.WriteTo(io.Writer(file1))
		file1.WriteAt([]byte("Hello"), 1)

	}
	t2 := time.Now().UnixNano()

	fmt.Println("顺序写耗时:", (t2 - t1))
	for a := 0; a < num; a++ {
		file1.WriteAt([]byte("world"), 0)
	}

	t3 := time.Now().UnixNano()
	//496,064,343
	fmt.Println("随机写耗时:", (t3 - t2))
	buf := make([]byte, 5)
	for a := 0; a < num; a++ {
		s := rand.Int63n(90*1024*1024)
		file1.Seek(0,os.SEEK_SET)
		file1.ReadAt(buf, s)

	}
	t4 := time.Now().UnixNano()
	for a := 0; a < num; a++ {
		file1.Read(buf)
	}
	t5 := time.Now().UnixNano()
	fmt.Println("随机读耗时:", (t4 - t3));
	fmt.Println("顺序读耗时:", (t5 - t4));
	fmt.Println("|", string(buf))

}

func write_buf() {
	filename := "/important/data/tmp/buf.txt"
	bufsize := 1024 * 1024 * 1200
	buf := make([]byte, bufsize)
	file1, error := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeSetuid)
	os.Chmod(filename, 7666)
	if error != nil {
		panic(error)
	}
	t1 := time.Now().Nanosecond()
	file1.Write(buf)
	t2 := time.Now().Nanosecond()
	//496,064,343
	fmt.Println("耗时:", (t2 - t1))
	file1.Close()
}
