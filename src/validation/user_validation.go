package validation

import (
	"github.com/go-playground/validator/v10"
)

// ===================== STRUCT =====================
type CreateUserRequest struct {
	Username      string `json:"username" validate:"required,min=3"`
	FullName      string `json:"fullName" validate:"required,min=3"`
	Password      string `json:"password" validate:"required,min=6"`
	Role          string `json:"role"`
	EsignID       string `json:"esignId"`
	EsignStatusID string `json:"esignStatusId"`
	RegisterDate  string `json:"registerDate"`
}

type UpdateUserRequest struct {
	Username      string `json:"username" validate:"omitempty,min=3"`
	FullName      string `json:"fullName" validate:"omitempty,min=3"`
	Password      string `json:"password" validate:"omitempty,min=6"`
	Role          string `json:"role" validate:"omitempty"`
	EsignID       string `json:"esignId" validate:"omitempty"`
	EsignStatusID string `json:"esignStatusId" validate:"omitempty"`
}

// ===================== VALIDATION =====================
var validate = validator.New()

// Custom pesan error untuk Create/Update User
var customMessages = map[string]string{
	"CreateUserRequest.Username.required": "Username wajib diisi",
	"CreateUserRequest.Username.min":      "Username minimal 3 karakter",
	"CreateUserRequest.FullName.required": "Nama lengkap wajib diisi",
	"CreateUserRequest.FullName.min":      "Nama lengkap minimal 3 karakter",
	"CreateUserRequest.Password.required": "Password wajib diisi",
	"CreateUserRequest.Password.min":      "Password minimal 6 karakter",

	"UpdateUserRequest.Username.min": "Username minimal 3 karakter",
	"UpdateUserRequest.FullName.min": "Nama lengkap minimal 3 karakter",
	"UpdateUserRequest.Password.min": "Password minimal 6 karakter",
}

// ValidateStruct memvalidasi struct dan mengembalikan error pertama
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		e := errs[0]
		key := e.StructNamespace() + "." + e.Tag()

		if msg, ok := customMessages[key]; ok {
			return &ValidationError{Message: msg}
		}

		// fallback
		return &ValidationError{Message: e.Field() + " tidak valid"}
	}

	return err
}

// ValidationError struct custom
type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}
