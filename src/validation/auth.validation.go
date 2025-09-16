package validation

type RegisterRequest struct {
	Username string `json:"username" validate:"required" message:"Username is required"`
	FullName string `json:"full_name" validate:"required" message:"Full name is required"`
	Password string `json:"password" validate:"required,min=6" message:"Password is required and must be at least 6 characters"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required" message:"Username is required"`
	Password string `json:"password" validate:"required" message:"Password is required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required" message:"Refresh token is required"`
}
