package entity

type User struct {
	ID        int    `gorm:"primary_key;column:id;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email;unique"`
	Password  string `gorm:"column:password"`
	Token     string `gorm:"column:token"`
	Phone     string `gorm:"column:phone"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt int64  `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
