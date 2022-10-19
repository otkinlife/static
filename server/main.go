package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	dir := flag.String("dir", "./", "server path")
	port := flag.String("port", "8999", "port")
	flag.Parse()
	if ret, _ := pathExists(*dir); !ret {
		fmt.Println("dir not exists")
		os.Exit(1)
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*dir))))
	fmt.Println("server start at port ", *port)
	err := http.ListenAndServe(":"+*port, mux)
	if err != nil {
		os.Exit(1)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
