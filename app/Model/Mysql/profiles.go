package mysql

type (
	Profiles struct {
		ID       uint   `gorm:"primary_key"`
		UserId   uint   `gorm:"column:userId"`
		FullName string `gorm:"column:fullName"`
	}
)

func (profiles *Profiles) GetByUserId(userID uint) {
	mysql.Where(Profiles{UserId: userID}).First(profiles)
}
