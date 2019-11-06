package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	db *gorm.DB
	gorm.Model
	Username string
	Password string `json:"-"`
	Phone    string
	Name     string
	Age      uint
}

func NewUser() User {
	db := DB()
	// Migrate the schema
	db.AutoMigrate(&User{})

	model := User{db: db}
	return model
}

// TODO: 暂不支持过滤条件
func (m User) Find(page int, pageSize int) ([]User, int) {
	db := m.db
	var count int
	db.Model(&User{}).Count(&count)

	users := make([]User, 0)
	if page > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset((page - 1) * pageSize)
	}

	db.Find(&users)
	return users, count
}

func (m User) Create(u *User) {
	m.db.Create(u)
}

func (m User) First(id int) *gorm.DB {
	return m.db.First(&User{}, id)
}

// TODO: 目前只支持全量更新
func (m User) Update(u *User) {
	m.db.Save(u)
}

func (m User) Delete(id uint) *gorm.DB {
	u := &User{}
	u.ID = id
	return m.db.Delete(u)
}
