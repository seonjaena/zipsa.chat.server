package chat

import (
	"encoding/json"
	"net/http"
	"strings"
	"zipsa.chat.server/zlog"
)

var log = zlog.Instance()

type Mail struct {
	SenderId         string `json:"senderId"`
	ReceiverNickname string `json:"receiverNickname""`
	Title            string `json:"title""`
	Body             string `json:"body""`
	MailIdx          uint   `json:"mailIdx""`
}

func GetChatRequest(w http.ResponseWriter, req *http.Request) *Mail {

	/*
		[ Chat server ]
		- 요청을 받는다.
		- 형식 검사
		- rabbitmq에 저장

		- 필요한 정보: 보내는 사용자 ID, 받는 사용자 닉네임, 제목, 내용, 메일 번호 (idx)
		- 서버에서 추가해야하는 정보: 보내지는 시간, 사용자 ID로 찾은 닉네임
		- 사용자 닉네임은 DB에서 로컬로 주기적으로 가져온다.
	*/

	method := req.Method

	if strings.Compare(method, http.MethodPost) != 0 {
		writeResponse(w, "connect only POST method is allowed", http.StatusMethodNotAllowed)
		return nil
	}

	headerContentTtype := req.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		writeResponse(w, "allow only application/json Content Type", http.StatusBadRequest)
		return nil
	}

	var mail Mail

	err := json.NewDecoder(req.Body).Decode(&mail)
	if err != nil {
		writeResponse(w, "bad data format", http.StatusBadRequest)
		return nil
	}

	writeResponse(w, "Success", http.StatusOK)
	return &mail
}

func writeResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
