package app

type User struct {
	Id        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Tweet struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
