//  package to avoide naming conflict
// compilation speed high --> only updated packages are reloaded in memory
// do not repeat yourself principle fullfilled

// module intialize , project initialize
//  go mod init "giturl" // conversion github url

// pacakge => folder -> files
// folder name can be anything but convesion is to have both folder and package name same

//  scope -> limited use in only certain set places

package main

import "github.com/piyushmishrax/golang/auth"

func main() {
	auth.LoginWithCredentials("piyush", "password")


}