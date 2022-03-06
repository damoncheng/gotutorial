package main

import (
	"fmt"
	_ "io"
	"log"
	"os"
)

/*
FileUtil file operate util
*/
type FileUtil struct {
	filename string
	*os.File
}

func main() {
	log.Println("starting effect")
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("Len(args) is %d, it must least be 2", len(os.Args)))
	}

	filepath := os.Args[1]

	fd, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer func() {
		fd.Close()
		os.Remove(filepath)
	}()

	oneFileUtil := &FileUtil{filepath, fd}
	oneFileUtil.Write([]byte("456"))

}
