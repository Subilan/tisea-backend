package structs

type DUser struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	ID        uint32 `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	GroupId   uint8  `json:"group_id"`
	Level     uint8  `json:"level"`
}