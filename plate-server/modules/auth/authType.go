package auth

type Register struct {
	UserName  string `json:"userName"`
	SecretKey string `json:"secretKey"`
	UserRole  string `json:"userRole"`
}

type Login struct {
	UserName  string `json:"userName"`
	SecretKey string `json:"secretKey"`
}
