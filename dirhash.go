package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	//        "strings"
	"regexp"
)

func getSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)

	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}

func getFilelist(path string, skip string, pattern string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		//		fmt.Println(f)
		if f.IsDir() {
			fmt.Println("dir return")
			return nil
		}
		if path == skip {
			return nil
		}
		reg, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		matched := reg.MatchString(path)

		if matched {
			println("skip this file or path")
		} else {

			println(path, getSize(path))
			file, err := os.Open(path)
			if err != nil {
				return nil
			}
			defer file.Close()
			h := sha1.New()
			_, erro := io.Copy(h, file)
			if erro != nil {
				return nil
			}
			//println(path, getSize(path),h.Sum(nil))
			fmt.Printf("%x\n\n", h.Sum(nil))
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	dpath := flag.String("D", "/Users/wzhyuan/Downloads/doc/", "search path")
	passpath := flag.String("P", "/Users/wzhyuan/Downloads/doc/", "pass path")
	flag.Parse()
	//root := flag.Arg(0)
	//skip := flag.Arg(1)
	//pattern := flag.Arg(2)
	//fmt.Println(root, skip)
	// fmt.Println(*passpath,*dpath)
	//getFilelist(root, skip, pattern)
}
