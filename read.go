package main

import (
	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)

	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}

func processLine(line []byte) {
	if string(line) == "" {
		return
	}
	s := strings.Split(string(line), ",")
	s2 := strings.Split(s[2], "\n")
	v64, err := strconv.ParseInt(s2[0], 10, 64)
	if err != nil {
		return
	}
	if v64 == getSize(s[0]) {
		//fmt.Println(v64, getSize(s[0]))

		file, err := os.Open(s[0])
		if err != nil {
			return
		}
		defer file.Close()
		h := sha1.New()
		_, erro := io.Copy(h, file)
		if erro != nil {
			return
		}
		r := fmt.Sprintf("%x", h.Sum(nil))
		hiz := strings.EqualFold(r, s[1])
		if hiz == true {
			fmt.Printf("Dir Hash Process is Ok!\n")
			return
		} else {
			fmt.Printf("Dir Hash Process failed!\n")
			return
		}

	} else {
		fmt.Printf("Dir Hash Process failed!\n")
		return
	}
	return
}

func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func main() {
	flag.Parse()
	path := flag.Arg(0)
	f, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("input file is not exist!\n")
		return
	}
	defer f.Close()

	fmt.Printf("Now will process this path:%s \n",path)
	ReadLine(path, processLine)
}
