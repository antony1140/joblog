package controllers

import (
	"fmt"
	"log"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/antony1140/joblog/service"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	log.Print("hasUser: ", hasUser)
	if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)

		log.Print("got to render home")
		return c.Redirect(302, "home")

	}


	return c.Render(200, "index","" )
}

func Login(c echo.Context) (error) {
	username := c.FormValue("username")
	pass := c.FormValue("password")
	user, err := service.LoginUser(username, pass)
	if err != nil {
		fmt.Println(err)
		return c.Render(404, "index", "Invalid Credentials, Try again.")
	}
	sessionCookie := security.CreateSession(true, user.Id)
	c.SetCookie(sessionCookie)


	return c.Redirect(302, "/home")
}

func Logout(c echo.Context) (error) {
	hasUser, _ := security.GetSession(c)	
	if hasUser {
		cookie, _ := c.Cookie("sid")
		err := security.DestroySession(cookie.Value)
		if err != nil {
			return c.Redirect(302, "/")
		}

	}

	return c.Redirect(302, "/")
}

func Home(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if !hasUser {
		return c.Redirect(302, "/")
	}

	orgs, sqlErr := dao.GetAllOrgsByUserId(id)

	jobs, err := dao.GetAllJobsByUserId(id)

	if err != nil{
		log.Print(err)

	}
	if sqlErr != nil {
		log.Print(sqlErr)
	}
	activeUser,_ := dao.GetUserById(id)
	data := struct {
		User *models.User
		Orgs []models.Org
		Jobs []models.Job

	} {
		User: activeUser,
		Orgs: orgs,
		Jobs: jobs,
	}

	return c.Render(200, "home", data)


}
