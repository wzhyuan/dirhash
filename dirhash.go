package main

import (
//	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
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





func main() {
	dpath := flag.String("D", "/Users/wzhyuan/Downloads/doc/", "-D set search path")
	pattern := flag.String("P", "(~*)123", "-P set skip  repexp")
	logfile := flag.String("L", "/Users/wzhyuan/Downloads/doc/dirhash.log", "-L set logfile")
	flag.Parse()
       //	fmt.Println(*dpath, *pattern)
	log, err := os.Create(*logfile)
	if err != nil {
	//	fmt.Println(logfile, err)
		return
	}
	defer log.Close()

	filepath.Walk(*dpath, func(dpath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		//		fmt.Println(f)
		if f.IsDir() {
	//		fmt.Println("dir return")
			return nil
		}
		reg, err := regexp.Compile(*pattern)
		if err != nil {
	//		fmt.Println(err)
			return nil
		}
		matched := reg.MatchString(dpath)

		if matched {
			println("skip this file or path")
		} else {

			file, err := os.Open(dpath)
			if err != nil {
				return nil
			}
			defer file.Close()
			h := sha1.New()
			_, erro := io.Copy(h, file)
			if erro != nil {
				return nil
			}
                        fmt.Fprintf(log,"%s,%x,%d\n",dpath,h.Sum(nil),getSize(dpath))

		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
