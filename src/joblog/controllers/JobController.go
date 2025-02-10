package controllers

import (
	"log"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/antony1140/joblog/service"
	"github.com/labstack/echo/v4"
)

func Job(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if !hasUser {
		return c.Redirect(302, "/")
	}
	cookie, _ := c.Cookie("sid")
	log.Print(id, " ", cookie.Value)
	jobId, _ := strconv.Atoi(c.Param("id"))

	activeUser,_ := dao.GetUserById(id)
	activeJob, daoErr := dao.GetJobById(jobId)
	activeOrg, orgDaoErr := dao.GetOrgByJobId(jobId)
	if orgDaoErr != nil {
		log.Print(orgDaoErr)
	}
	ClientData,_ := dao.GetClientById(activeJob.ClientId)
	ExpenseList, daoErr := dao.GetAllExpensesByJobId(jobId)

	Expenses := service.GroupExpenseReceipts(ExpenseList)	
	for _, receipt := range Expenses {
		if receipt.Id != 0 {
		}
	}
	if daoErr != nil{
		log.Print("failed to get job from db")
	}
	data := struct {
		User *models.User
		Org *models.Org
		Job *models.Job
		Client *models.Client
		ExpenseList map[*models.Expense] *models.Receipt
	} {
		User: activeUser,
		Job: activeJob,
		Org: activeOrg,
		Client: ClientData,
		ExpenseList: Expenses,
	}

	return c.Render(200, "jobPage", data)


}

func NewJob(c echo.Context) (error) {
	hasUser, _ := security.GetSession(c)
	if !hasUser {
		return c.Redirect(302, "/")
	}
	
	jobName := c.FormValue("job-name")
	
	jobDesc := c.FormValue("job-desc")
	groupId, err := strconv.Atoi(c.FormValue("group-id"))
	if err != nil {
		log.Println("error creating job", err)
		return c.Redirect(302, "/home")
	}
	log.Println("name, desc, id: ", jobName, jobDesc, groupId)
	job := models.Job{
		Title: jobName,
		Description: jobDesc, 
		OrgId: groupId,
		
	}

	err = service.NewJob(&job)
	if err != nil {
		log.Println("error creating job", err)
		return c.Redirect(302, "/home")
	}

	return c.Redirect(302, "/group/" + strconv.Itoa(groupId))
}


