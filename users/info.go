package users

import (
	"time"
	zp "zipsa.chat.server/properties"
)

type UserInfo struct {
	UserId       string
	UserNickname string
}

func GetUserInfo() {
	for {
		select {
		case <-time.After(time.Millisecond * time.Duration(zp.GetDBUserInfoFethCycleMS())):

		}
	}
}
