package models

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
	CreatedAt   string `json:"created_at"`
}

type RoleResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type RoleRequest struct {
	Name string `json:"name"`
}
