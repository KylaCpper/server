package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func page(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //analysis parameter
	fmt.Println("path", r.URL.Path)
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }

	buf := ReadFile(path)
	setmime(w, path)

	w.Write([]byte(buf)) //print client
}
func router(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	b, _ := json.Marshal(r.Form)
	fmt.Fprintf(w, "%s", string(b))
}
func setmime(w http.ResponseWriter, path string) {
	contentType := "text/plain"
	if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpeg"
		// http.ServeFile(w, r, r.URL.Path)
	} else if strings.HasSuffix(path, ".gif") {
		contentType = "image/gif"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml"
	} else if strings.HasSuffix(path, ".mp4") {
		contentType = "video/mp4"
	} else if strings.HasSuffix(path, ".webm") {
		contentType = "video/webm"
	} else if strings.HasSuffix(path, ".ogg") {
		contentType = "video/ogg"
	} else if strings.HasSuffix(path, ".mp3") {
		contentType = "audio/mp3"
	} else if strings.HasSuffix(path, ".wav") {
		contentType = "audio/wav"
	}
	w.Header().Add("Content-Type", contentType)
}

func main() {

	http.HandleFunc("/", page)
	http.HandleFunc("/aaa", router)          //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func ReadFile(path string) (buf_ string) {
	inputFile := "D:/webwork" + path
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	buf_ = string(buf)

	return buf_

}
