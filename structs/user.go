package structs

// 存储在数据库中的完整 User 对象
type DatabaseUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	ID        uint32 `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	GroupId   uint8  `json:"group_id"`
	Level     uint8  `json:"level"`
}

// 包含基本信息的 User 对象
type RegisteringUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
}