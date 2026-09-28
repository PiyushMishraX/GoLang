//  package to avoide naming conflict
// compilation speed high --> only updated packages are reloaded in memory
// do not repeat yourself principle fullfilled

// module intialize , project initialize
//  go mod init "giturl" // conversion github url

// pacakge => folder -> files
// folder name can be anything but convesion is to have both folder and package name same

//  scope -> limited use in only certain set places

package main

import (
	"github.com/fatih/color"
	"github.com/piyushmishrax/golang/user"
)

// "fmt"
// "os/user" // no this one

//  using open source package
//  github.com/fatih/color
// go get github.com/fatih/color

func main() {
	// auth.LoginWithCredentials("piyush", "password")

	// session := auth.GetSession()

	// fmt.Println(session)

	user := user.User{
		Email: "user@email.com",
		Name:  "User",
	}
	// fmt.Println(user)
	// fmt.Println(user.Email, user.Name)

	color.Red(user.Email)
	color.Blue(user.Name)



}
