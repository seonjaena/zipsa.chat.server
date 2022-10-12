package main

import (
	"net/http"
	"zipsa.chat.server/chat"
	"zipsa.chat.server/zlog"
)

var log = zlog.Instance()

func main() {
	forever := make(chan bool)
	http.HandleFunc("/hello", receiveMail)

	http.ListenAndServe(":8088", nil)
	<-forever
}

func receiveMail(w http.ResponseWriter, req *http.Request) {

	mail := chat.GetChatRequest(w, req)
	log.Infof("%+v", mail)

	log.Infof("%s", req.Method)
	log.Infof("%s", req.Header)
	log.Infof("%s", req.Body)
	log.Infof("%s", req.URL)
	log.Infof("%s", req.RemoteAddr)
	log.Infof("%s", req.RequestURI)
	log.Infof("%s", req.Host)

}
