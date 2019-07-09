package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/userup", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("users"); os.IsNotExist(err) {
			err = os.Mkdir("users", 777)
			if err != nil {
				fmt.Println(err)
			}
		}
		r.ParseForm()

		name := r.FormValue("name")
		port := r.FormValue("port")
		remote := strings.Split(r.RemoteAddr, ":")
		userip := remote[0] + port

		f, err := os.Create("users/" + name)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = f.WriteString(userip)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, "ok")
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		//list := make(map[string]string)
		r.ParseForm()
		name := r.FormValue("name")
		files, err := ioutil.ReadDir("users")
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range files {
			if file.Name() == name {
				continue
			}
			content, _ := ioutil.ReadFile("user/" + file.Name())
			fmt.Fprintf(w, string(content))
			fmt.Fprintf(w, "\n")
			fmt.Fprintf(w, file.Name())
		}
		//fmt.Fprintf(w, name)
	})
	http.ListenAndServe(":8889", nil)
}
