package database

import (
	"tisea-backend/structs"
	"tisea-backend/utils/security"
)

// 从给定的数据中创建一个 *RegisteringUser 对象
func MakeRegisteringUser(username string, email string, password string) (*structs.RegisterUserRequest, error) {
	hashed, err := security.GenerateHash(password)

	user := new(structs.RegisterUserRequest)

	if err != nil {
		return user, err
	}

	user.Username = username
	user.Email = email
	user.Hash = hashed

	return user, nil
}

// 将所给的 RegisteringUser 对象插入到数据库中
func InsertRegisteringUser(user structs.RegisterUserRequest) error {
	_, err := Exec("INSERT INTO `tisea_users` (username, password_hash, email) VALUES (?, ?, ?)", user.Username, user.Hash, user.Email)

	return err
}

// 使用用户名获取 *DatabaseUser 对象。若不存在指定的用户，返回 nil 和 nil。
func GetUserByUsername(username string) (*structs.DatabaseUser, error) {
	result, err := Query("SELECT email, password_hash, nickname, bio, id, created_at, updated_at, group_id, level FROM `tisea_users` WHERE username=?", username)
	user := new(structs.DatabaseUser)
	if err != nil {
		return user, err
	}
	defer result.Close()

	if result.Next() {
		result.Scan(&user.Email, &user.Hash, &user.Nickname, &user.Bio, &user.ID, &user.CreatedAt, &user.UpdatedAt, &user.GroupId, &user.Level)
	} else {
		return nil, nil
	}

	user.Username = username

	return user, nil
}
