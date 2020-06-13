package service

import (
	mysql "github.com/koushamad/goro/app/Model/Mysql"
)

type (
	Auth struct {
		Token string  `json:"token"`
	}
)

func (a Auth)Check(sId string) uint {
	user := mysql.Users{}
	if user.CheckToken(a.Token) {
		user.SetSId(sId)
		profile := user.Profile()
		return profile.ID
	}
	return 0
}

func Connect(sId string) bool {
	user := mysql.Users{}
	user.GetBySid(sId)
	return user.SId == sId
}

func Disconnect(sId string) bool {
	user := mysql.Users{}
	user.GetBySid(sId)
	user.UnSetSId()
	return user.SId == ""
}

func Profile(sId string) mysql.Profiles{
	user := mysql.Users{}
	user.GetBySid(sId)
	return user.Profile()
}