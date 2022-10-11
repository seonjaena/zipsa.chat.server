package main

import (
	"net/http"
	"strings"
	"zipsa.chat.server/zlog"
)

var log = zlog.Instance()

func main() {
	forever := make(chan bool)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8088", nil)
	<-forever
}

func hello(w http.ResponseWriter, req *http.Request) {
	if strings.Compare("GET", req.Method) != 0 {
		w.Write([]byte("Connection Error!!!"))
		return
	}
	log.Infof("%s", req.Method)
	log.Infof("%s", req.Header)
	log.Infof("%s", req.Body)
	log.Infof("%s", req.URL)
	log.Infof("%s", req.RemoteAddr)
	log.Infof("%s", req.RequestURI)
	log.Infof("%s", req.Host)
	w.Write([]byte("Hello World"))
}
