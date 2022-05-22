package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("This is all about File")

	err := makeDirIfNotExist("temp")
	checkNilErr(err)

	f, err := os.Create("temp/data")
	checkNilErr(err)

	defer f.Close()

	data1 := []byte("hello")
	//ASCII for "hello" [104 101 108 108 111] 10进制
	fmt.Printf("%v\n", data1)

	//ASCII to binary 2进制
	fmt.Printf("%s\n", binary("hello"))

	//hex string 16进制， 16进制 比 10进制 更容易拆解成 2进制
	hexString := hex.EncodeToString(data1[:])
	fmt.Printf("%s\n", hexString)

	//encode to base64 就是把8位的0和1 变成6位的
	//01101000 01100101 01101100 01101100 01101111 变成了
	//011010 000110 010101 101100 011011 000110 111100 = 不够8位补上0 完全没有的加=
	//26 6 21 44 27 6 60
	//a G v s b G 8 =
	sEnc := base64.StdEncoding.EncodeToString(data1)
	fmt.Println(sEnc)

	n1, err := f.Write(data1)
	checkNilErr(err)

	fmt.Printf("wrote %d bytes\n", n1)

	fi, _ := f.Stat()

	fmt.Printf("file length %v\n bytes", fi.Size())

	// n, err := f.WriteString("data\n")
	// checkNilErr(err)
	// fmt.Printf("wrote %d bytes\n", n)

}

//clearDir
func clearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

//makeDirIfNotExist
func makeDirIfNotExist(dirpath string) error {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		err := os.MkdirAll(dirpath, os.ModeDir|os.ModePerm)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		err := clearDir(dirpath)
		return err
	}
	return nil
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatalln("error: ", err)
	}
}

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b ", res, c)
	}
	return res
}
