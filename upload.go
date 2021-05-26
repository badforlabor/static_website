/**
 * Auth :   liubo
 * Date :   2021/5/26 17:37
 * Comment: 上传文件
 */

package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func uploadTemplate() *template.Template {

	//t, _ := template.ParseFiles("upload.gtpl")
	var t = template.New("upload.gtpl")
	t.Parse(`<html>
<head>
       <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
</form>
</body>
</html>`)

	return t
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		var t = uploadTemplate()
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(filepath.Join(uploadFolder, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		var _, e = io.Copy(f, file)

		if e == nil {
			w.Write([]byte("ok"))
		} else {
			w.Write([]byte(e.Error()))
		}
	}
}