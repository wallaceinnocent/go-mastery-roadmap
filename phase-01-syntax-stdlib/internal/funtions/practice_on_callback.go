package functions

import (
	"fmt"
	"time"
)

func UpdateUsers(onComplete func(string)) {
	fmt.Println("Update users started ")
	time.Sleep(10 * time.Second)
	res := fmt.Sprintf("All user profiles updated: %d users processed", 10)
	onComplete(res)
}
