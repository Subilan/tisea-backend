package database

import (
	"tisea-backend/structs"
	"tisea-backend/utils/security"
)

func MakeRegisteringUser(username string, email string, password string) (*structs.RegisteringUser, error) {
	hashed, err := security.GenerateHash(password)

	user := new(structs.RegisteringUser)

	if err != nil {
		return user, err
	}

	user.Username = username
	user.Email = email
	user.Hash = hashed

	return user, nil
}

func InsertRegisteringUser(user structs.RegisteringUser) error {
	_, err := Exec("INSERT INTO `tisea_users` (username, password_hash, email) VALUES (?, ?, ?)", user.Username, user.Email, user.Hash)
	
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*structs.DatabaseUser, error) {
	result, err := Query("SELECT (email, hash, nickname, bio, id, created_at, updated_at, group_id, level) FROM `tisea_users` WHERE username=?", username)
	user := new(structs.DatabaseUser)
	if err != nil {
		return user, err
	}
	defer result.Close()

	for result.Next() {
		result.Scan(&user.Email, &user.Hash, &user.Nickname, &user.Bio, &user.ID, &user.CreatedAt, &user.UpdatedAt, &user.GroupId, &user.Level)
	}

	user.Username = username

	return user, nil
}