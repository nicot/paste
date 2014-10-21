package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const path = "/"
const port = ":8080"
const dir = "pastes/"

var dirs uint64

func read(key string) string {
	// ReadFile only works in subdirs (we won't read /etc/passwd)
	data, err := ioutil.ReadFile(dir + key)
	if err != nil {
		log.Printf("File not found")
	}
	return string(data)
}

func getName(n uint64) string {
	s := ""
	for (n > 0) || (len(s) == 0) {
		c := string(n%26 + 'a')
		s += c
		n /= 26
	}
	return s
}

func post(content []byte) (string, error) {
	name := getName(dirs)
	err := os.Mkdir(dir, os.ModeDir+0777)
	err = ioutil.WriteFile(dir+name, content, 0644)
	if err != nil {
		log.Println(err)
		log.Fatal("Error writing file")
	}
	dirs += 1
	return name, err
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Path[1:]
	fmt.Fprintf(w, read(k))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	content := make([]byte, r.ContentLength)
	_, err := r.Body.Read(content)
	if err != nil {
		fmt.Fprintf(w, "Error reading POST\n")
		//w.Status = 500
		return
	}
	if len(content) == 0 {
		fmt.Fprintf(w, "Empty response not saved\n")
		//w.Status = 500
		return
	}
	name, err := post(content)
	if err != nil {
		fmt.Fprintf(w, "Error writing content\n")
		//w.Status = 500
		return
	}
	fmt.Fprintf(w, name+"\n")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(path+"{key}", getHandler).Methods("GET")
	r.HandleFunc(path, postHandler).Methods("POST")
	http.Handle(path, r)
	http.ListenAndServe(port, nil)
}
