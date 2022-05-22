package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func main() {
	fmt.Println(Warn("This is all about File"))

	err := makeDirIfNotExist("temp")
	checkNilErr(err)

	f, err := os.Create("temp/data")
	checkNilErr(err)

	defer f.Close()

	stringMsg := "hello"
	fmt.Print(Info("String: "))
	fmt.Printf("%v\n", stringMsg)

	data1 := []byte(stringMsg)

	//ASCII for "hello" [104 101 108 108 111]10进制
	fmt.Print(Info("ASCII Decimal: "))
	fmt.Printf("%v\n", data1)

	//ASCII to binary 2进制
	fmt.Print(Info("Binary:  "))
	fmt.Printf("%s\n", binary(stringMsg))

	//hex string 16进制， 16进制 比 10进制 更容易拆解成 2进制
	hexString := hex.EncodeToString(data1[:])
	fmt.Print(Info("Hex:  "))
	fmt.Printf("%s\n", hexString)

	//encode to base64 就是把8位的0和1 变成6位的
	//01101000 01100101 01101100 01101100 01101111 变成了
	//011010 000110 010101 101100 011011 000110 111100 = 不够8位补上0 完全没有的加=
	//26 6 21 44 27 6 60
	//a G v s b G 8 =
	sEnc := base64.StdEncoding.EncodeToString(data1)
	fmt.Print(Info("Encode Base64:  "))
	fmt.Println(sEnc)

	//decode from base64
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Print(Info("Decode Base64:  "))
	fmt.Println(string(sDec))

	//write to file
	n1, err := f.Write(data1)
	checkNilErr(err)

	fmt.Printf("wrote %d bytes\n", n1)

	//we can also write string directly
	n2, err := f.WriteString(" Qi\n")
	checkNilErr(err)
	fmt.Printf("wrote %d bytes\n", n2)

	fi, _ := f.Stat()
	fmt.Printf("file length %v bytes\n", fi.Size())

	// read file
	content, err := os.ReadFile("temp/data")
	checkNilErr(err)
	fmt.Print(string(content))

	// clean "temp" folder
	clearDir("temp")
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
