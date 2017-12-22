package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"flag"
)

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

//dirents返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}