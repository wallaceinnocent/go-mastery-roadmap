package structs

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"isActive"`
}
type Admin struct {
	User
	Position string `json:"position"`
}

func InitialiseUser() {
	var user1 User
	user1.Age = 20
	user1.Username = "Job"
	user1.IsActive = false
	user1.Email = "job@ex.com"
	user2 := User{
		Username: "wallace_k",
		Email:    "wallace@example.com",
		Age:      22,
		IsActive: true,
	}
	user3 := User{"john_doe", "john@example.com", 25, false}

	admin := Admin{user2, "executive"}
	fmt.Println(user1, "\t", user2, "\t", user3)
	jsonData, _ := json.MarshalIndent(admin, "", " ")

	fmt.Println(string(jsonData))
}
