package repository

func GetUser() User {
	db := Connection()
	var user User

	result := db.Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}
