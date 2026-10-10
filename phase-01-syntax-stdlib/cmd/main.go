package main

import (
	"fmt"

	functions "github.com/wallaceinnocent/go-mastery-roadmap/phase-01-syntax-stdlib/internal/funtions"
	"github.com/wallaceinnocent/go-mastery-roadmap/phase-01-syntax-stdlib/internal/interfaces"
	"github.com/wallaceinnocent/go-mastery-roadmap/phase-01-syntax-stdlib/internal/structs"
)

func main() {
	runFunctions()
	runStructs()
	runInterfaces()

}
func runStructs() {
	structs.InitialiseUser()
}
func runInterfaces() {
	emailClient := interfaces.EmailService{Email: "user@example.com"}
	smsClient := interfaces.SMSService{Phone: "+254712345678"}

	emailClient.SendMessage("Your order has shipped!")
	interfaces.AlertUser(emailClient, "New order alert")
	interfaces.AlertUser(smsClient, "Your verification code is 4829")
}
func runFunctions() {
	functions.HelloWorld()
	functions.Functions()
	func() {
		fmt.Println("Updating users . . . .")
		go functions.UpdateUsers(func(s string) {
			fmt.Println(s)
		})
		fmt.Println("Updating users finished")
	}()
}
