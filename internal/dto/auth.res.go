package dto

type UserInfoResponse struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name,omitempty"`
	Username string   `json:"username,omitempty"`
	Email    string   `json:"email,omitempty"`
	Permits  []string `json:"permits,omitempty"`
	Foto     string   `json:"foto,omitempty"`
	Role     string   `json:"role,omitempty"`
}
