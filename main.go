package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"router"
)

type myStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func main() {
	var bb string
	_ = bb
	fmt.Println("test-----------")
}

func init() {

	// http.HandleFunc("/", handler)
	http.HandleFunc("/json", myJson)

	http.HandleFunc("/json/insertUser/postformurl/main", postformurl)
	http.HandleFunc("/", router.Route)

	// 设置静态目录
	fsh := http.FileServer(http.Dir("D:/workspace/go-learn/learn2/src/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println(111)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// if r.URL.Path == "/json" {
	// 	myJson(w, r)
	// } else {
	// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// }
}

func myJson(w http.ResponseWriter, request *http.Request) {
	result := myStruct{
		"canpong",
		10,
		20,
	}

	j, _ := json.Marshal(result)
	_ = j

	w.Write(j)

	fmt.Fprintf(w, "%s %s %s\n", request.Method, request.URL, request.Proto)

	for k, v := range request.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", request.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	s, _ := ioutil.ReadAll(request.Body) //把  body 内容读入字符串 s
	fmt.Fprintf(w, "%s", s)              //在返回页面中显示内容。

	// fmt.Fprintf(w, "URL.Path = %s\n", j)

}

func postformurl(w http.ResponseWriter, request *http.Request) {
	userInfo := model.UserInfo{
		// Username:   queryForm["userName"][0],
		// Departname: queryForm["departname"][0],
		// Password:   queryForm["password"][0],
		// Uid:        uuid.Must(uuid.NewV4()).String(),
		// Created:    time.Now(),
	}

	// request.ParseMultipartForm(128)

	request.ParseForm()
	fmt.Println(request.URL)
	for k, v := range request.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	fmt.Println("r.Form:         ", request.Form)

	fmt.Println("r.PostForm:     ", request.PostForm)

	fmt.Println("r.MultiPartForm:", request.MultipartForm)

	userName := request.PostFormValue("username")
	if userName != "" {
		userInfo.Username = userName
	}
	departname := request.PostFormValue("departname")
	if departname != "" {
		userInfo.Departname = departname
	}
	password := request.PostFormValue("password")
	if password != "" {
		userInfo.Password = password
	}

	// service.InsertUserService(userInfo)
}
