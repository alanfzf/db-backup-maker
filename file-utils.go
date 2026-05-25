package main

import (
    "compress/gzip"
    "time"
    "os"
    "io"
)

func CompressFile(file string) string {
    src, err := os.Open(file)

    if err != nil {
        panic(err)
    }

    defer src.Close()

    gzFile := file+".gz"

    // 1. create destination .gz file
    dst, err := os.Create(gzFile)
    if err != nil {
        panic(err)
    }
    defer dst.Close()

    // 2. create gzip writer
    gz := gzip.NewWriter(dst)
    defer gz.Close()
    gz.Name = gzFile

    // copy file contents into gzip writer
    _, err = io.Copy(gz, src)
    if err != nil {
        panic(err)
    }

    return gzFile
}

func CreateTmpFile() string {
    ts := time.Now().Format(time.RFC3339)
    f, err := os.CreateTemp("", "dump_"+ts+"_*.sql")

    if err != nil {
        panic(err)
    }

    defer f.Close()

    return f.Name()
}
