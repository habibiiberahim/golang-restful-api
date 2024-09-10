package test

import "golang-restful-api/internal/entity"

func ClearAll() {
	ClearUsers()
}

func ClearUsers() {
	err := db.Where("email is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}
