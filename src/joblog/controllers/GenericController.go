package controllers

import (
	"log"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if !hasUser {

		return c.Redirect(302, "/")

	}

	choice := c.Param("type")
	orgValue := c.FormValue("org-id")
	orgId,_ := strconv.Atoi(orgValue)
	jobValue := c.FormValue("job-id")
	jobId,_ := strconv.Atoi(jobValue)
	activeUser,_ := dao.GetUserById(id)
	activeOrg, orgDaoErr := dao.GetOrgById(orgId)
	if orgDaoErr != nil {
		log.Print("error at org dao, ", orgDaoErr)
	}
	activeJob, jobDaoErr := dao.GetJobById(jobId)
	if jobDaoErr != nil {
		log.Print("error at job dao, ", jobDaoErr)
	}
	activeClient, clientDaoErr := dao.GetClientById(activeJob.ClientId)
	if clientDaoErr != nil {
		log.Print("error at client dao, ", clientDaoErr)
	}
	expenseList, expDaoErr := dao.GetAllExpensesByJobId(jobId)
	if expDaoErr != nil {
		log.Println(expDaoErr)
	}



	data := struct {
		User *models.User
		Choice string
		Job *models.Job
		Org *models.Org
		Client *models.Client
		ExpenseList []*models.Expense

	} {
		User: activeUser,
		Choice: choice,
		Job: activeJob,
		Org: activeOrg,
		Client: activeClient,
		ExpenseList: expenseList,

	}

	render := "create" + choice
	return c.Render(200, render, data)


}
