// main is a user's login Service.

package main

import "fmt"

// @Title  main.go
// @Description main.go contains the user login function and logout function
// @Author  XdpCs 2023/3/13 11:14
// @Update  XdpCs 2023/3/13 11:18

// User defines user login info
type User struct {
	UserName string // user's name
	Password string // user's password
}

// IUser defines user function
type IUser interface {
	Login()  // user login into the system
	Logout() // user logout the system
}

var _ IUser = (*User)(nil)

// Login
// @Title       Login
// @Description User login into the system and print success or print failed
// @Author      XdpCs 2023/3/13 11:14
// @Param     	nil
// @Return    	nil
func (user *User) Login() {
	if user.UserName == "XdpCs" && user.Password == "11141118" {
		fmt.Println("登录成功")
		return
	}
	fmt.Println("登录失败")
}

// Logout
// @Title       Logout
// @Description User logout the system and print success
// @Author      XdpCs 2023/3/13 11:14
// @Param     	nil
// @Return    	nil
func (user *User) Logout() {
	fmt.Println("注销成功")
}

// Register
// @Title       Register
// @Description User login into the system and print success or print failed
// @Author      XdpCs 2023/3/13 11:14
// @Param     	userName string "user's name"
// @Param 		password string "user's password"
// @Return      anonymity *User "User object contains userName and password"
func Register(userName string, password string) *User {
	return &User{
		UserName: userName,
		Password: password,
	}
}

func main() {
	user := Register("XdpCs", "11141118")
	user.Login()
	user.Logout()
}
