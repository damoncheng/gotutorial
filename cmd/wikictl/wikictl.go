package main

import (
	"github.com/damoncheng/gotutorial/pkg/wiki"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	//p1 := Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//p1.save()
	//p2, _ := LoadPage("TestPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/", wiki.Handler)
	http.HandleFunc("/view/", wiki.ViewHandler)
	http.HandleFunc("/edit/", wiki.EditHandler)
	http.HandleFunc("/save/", wiki.SaveHandler)

	go func() {

	}()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
