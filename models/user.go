package models

import (
	"strings"
  "fmt"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-"` 
	Role     string `json:"role" gorm:"default:user"` 
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Email = strings.ToLower(u.Email) 
	if u.Role == "" {
		u.Role = "user" 
	}
	return u.HashPassword()            // Password hashing
}

// Password Hashing Function
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println("Original Password:", u.Password)              // ✅ Debug Line 1
	fmt.Println("Hashed Password (before save):", string(hashedPassword)) // ✅ Debug Line 2
	u.Password = string(hashedPassword)
	return nil
}


func (u *User) CheckPassword(password string) bool {
	fmt.Println("Stored Hash:", u.Password)  // Debugging
	fmt.Println("User Input:", password)  // Debugging

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
	}
	return err == nil
}








