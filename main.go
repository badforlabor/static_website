/**
 * Auth :   liubo
 * Date :   2018/11/1 23:41
 * Comment: 简单的静态网站服务器
 */

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var uploadFolder = "./upload"

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the download directory of static file to host")
	upfolder := flag.String("u", uploadFolder, "the upload directory of static file to host")
	flag.Parse()

	uploadFolder = *upfolder
	os.MkdirAll(uploadFolder, 0666)

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	http.HandleFunc("/upload", upload)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

