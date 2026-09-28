package auth

import "fmt"

// mostly the fn name starts with capital letter in go 
// bcz small letter only can be accessed in same pkg
// capital ones can be exported used
// func loginWithCredentials(username string, password string) {
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using ", username, password) // or the login logic 
}