package mysql

type (
	Users struct {
		ID    uint   `gorm:"primary_key"`
		Name  string `gorm:"column:name"`
		Phone string `gorm:"column:phone"`
		Token string `gorm:"column:token"`
		Email string `gorm:"column:email"`
		SId   string `gorm:"column:sid"`
	}
)

func (user Users)Profile() Profiles {
	profile := Profiles{}
	profile.GetByUserId(user.ID)
	return profile
}

func (user *Users) CheckToken(token string) bool {
	mysql.Where(Users{Token: token}).First(user)
	return user.Token == token
}

func (user *Users) GetBySid(sId string) {
	mysql.Where(Users{SId: sId}).First(user)
}

func (user *Users) SetSId(sId string) {
	mysql.Model(user).Updates(Users{SId: sId})
}

func (user *Users) UnSetSId() {
	mysql.Model(user).Updates(Users{SId: ""})
}
