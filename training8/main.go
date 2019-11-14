package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// func main() {

// 	e := echo.New()
// 	e.GET("/", HelloController)
// 	e.Start(":8000")

// }

// func HelloController(c echo.Context) error {

// 	return c.String(http.StatusOK, "Hello World")

// type User struct {
// 	Name  string
// 	Email string
// }

// func GetUser(c echo.Context) error {
// 	user := User{Name: "Nurdin", Email: "nurdin@alterra.id"}
// 	return c.JSON(http.StatusOK, user)
// }

// func main() {
// 	e := echo.New()
// 	e.GET("/user", GetUser)
// 	e.Start(":8000")

// }

// type User struct {
// 	Id    int
// 	Name  string
// 	Email string
// }

// func GetUserController(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	if id == 4 {
// 		user := User{Id: id, Name: "didin", Email: "nurdin@alterra.id"}
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"user": user,
// 		})
// 	}
// 	user := User{Id: id, Name: "Nurdin", Email: "nurdin@alterra.id"}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"user": user,
// 	})
// }

// func main() {
// 	e := echo.New()

// 	e.GET("/user/:id", GetUserController)

// 	e.Start(":8000")
// }
// func UserSearchController(c echo.Context) error {

// 	match := c.QueryParam("match")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"match":  match,
// 		"result": []string{"adi", "aan", "asif"},
// 	})
// }

// func main() {
// 	e := echo.New()

// 	e.GET("users", UserSearchController)

// 	e.Start(":8000")
// }

type User struct {
	Id       int    `json :"id" form : "id"`
	Name     string `json : "name" form : "name"`
	Email    string `json: "email" form :"email"`
	Password string `json :"password" form :"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func CreateUserController(c echo.Context) error {

	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func GetUserController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	for_, temp := range users {
		if temp.Id == id {
			return c.JSON(http.StatusOK, map[string]interface
				{}{
			"message" : "success get data",
			"users" : temp,
		})
	}
} 


func main() {

	e := echo.New()

	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/user/:id", GetUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
