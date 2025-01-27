package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/labstack/echo/v4"
)

func Group(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	log.Print("userid: ", id)
	log.Print("hasUser: ", hasUser)
	if !hasUser {
		return c.Redirect(302, "/")
	}
	cookie, _ := c.Cookie("sid")
	log.Print(id, " ", cookie.Value)
	orgId, _ := strconv.Atoi(c.Param("id"))
	jobs, sqlErr := dao.GetAllJobsByOrgId(orgId)
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}
	activeUser,_ := dao.GetUserById(id)
	activeOrg, daoErr := dao.GetOrgById(orgId)
	if daoErr != nil{
		log.Print("failed to get orgaization from db at group page")
	}
	data := struct {
		User *models.User
		Org *models.Org
		Jobs []models.Job
		JobNum int

	} {
		User: activeUser,
		Jobs: jobs,
		Org: activeOrg,
		JobNum: len(jobs),
	}

	return c.Render(200, "orgPage", data)


}

func NewGroup(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if !hasUser {
		return c.Redirect(302, "/")
	}
	var newGroup models.Org
	newGroup.Name = c.FormValue("name")
	newGroup.Description = c.FormValue("description")
	groupId, err := dao.CreateOrg(&newGroup)
	if err != nil {
		log.Print("failed to create group")
		log.Print(err)
		return c.Redirect(302, "home")
	}
	userErr := dao.AddOrgUser(id, groupId)
	if userErr != nil {
		log.Print("failed to add group")
		log.Print(userErr)
		return c.Redirect(302, "home")
	}
	// data := struct {
	// 	User *models.User
	// 	Org *models.Org
	// 	Jobs []models.Job
	//
	// } {
	//
	// }

	log.Print("Created new group and added user")
	return c.Redirect(302, "/home")


}
