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
                //              fmt.Println(f)
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
"f.go" 79L, 1509C
