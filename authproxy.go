package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/linuzilla/http_authproxy"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ", os.Args[0], "[jsonfile]")
	} else if b, err := ioutil.ReadFile(os.Args[1]); err != nil {
		fmt.Println(err)
	} else {
		var conf http_authproxy.Config

		if err := json.Unmarshal(b, &conf); err != nil {
			fmt.Println(err)
			os.Exit(3)
		} else {
			proxy := http_authproxy.New(&conf)

			s := &http.Server{
				Addr:           ":8080",
				Handler:        proxy,
				ReadTimeout:    10 * time.Second,
				WriteTimeout:   10 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}

			err := s.ListenAndServe()
			if err != nil {
				fmt.Printf("Server failed: ", err.Error())
			}
		}
	}
}
