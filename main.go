package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "main.h"
*/
import "C"

import (
	"net"
	"strings"
)

var urlListener chan string = make(chan string)

func main() {
	go C.RunApp()

	appSchema := "app://"
	rawURL := <-urlListener

	if !strings.HasPrefix(appSchema, appSchema) {
		panic("not a app:// URL")
	}
	rawURL = strings.TrimPrefix(rawURL, appSchema)
	payloadStart := strings.Index(rawURL, "/")
	if payloadStart == -1 {
		panic("not a valid app://host:port/payload URL")
	}
	host, payload := rawURL[0:payloadStart], rawURL[payloadStart+1:]

	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte(payload)); err != nil {
		panic(err)
	}
	if _, err := conn.Write([]byte("\n")); err != nil {
		panic(err)
	}
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
