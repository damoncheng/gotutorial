package main

import (
	"fmt"
	_ "io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

/*
FileUtil file operate util
*/
type FileUtil struct {
	filename string
	*os.File
	writeChan chan int
	readChan  chan int
	endChan   chan bool
}

func main() {
	log.Println("starting effect")
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("Len(args) is %d, it must least be 2", len(os.Args)))
	}

	filepath := os.Args[1]

	if check := strings.HasPrefix(filepath, "/tmp/"); !check {
		panic("file path must start by /tmp")
	}

	fd, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer func() {
		fd.Close()
		//os.Remove(filepath)
	}()

	oneFileUtil := &FileUtil{
		filepath, fd, make(chan int), make(chan int), make(chan bool)}

	numCPU := runtime.NumCPU()

	for i := 0; i < numCPU; i++ {
		go func(index int) {
			time.Sleep(2000 * time.Millisecond)
			if _, err := oneFileUtil.Write([]byte(fmt.Sprint(index))); err != nil {
				panic(fmt.Sprintf("fail write %d", index))
			}
			oneFileUtil.writeChan <- index
		}(i)
	}

	for i := 0; i < numCPU; i++ {

		<-oneFileUtil.writeChan
	}

}
