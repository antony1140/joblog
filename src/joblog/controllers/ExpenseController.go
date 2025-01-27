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

func Expense(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	log.Print("hasUser: ", hasUser)
	if !hasUser {
		return c.Redirect(302, "/")
	}
	cookie, _ := c.Cookie("sid")
	log.Print(id, " ", cookie.Value)
	exp := c.Param("id")
	expId,_ := strconv.Atoi(exp)

	activeUser,_ := dao.GetUserById(id)
	expense, expDaoErr := dao.GetExpenseById(expId)

	if expDaoErr != nil {
		log.Print("err at get expense/id Error: expDaoErr", expDaoErr)
	}
	var expToList []*models.Expense

	tempExp := expense
	expToList = append(expToList, tempExp)
	receiptMap := service.GroupExpenseReceipts(expToList)
	docList, err := dao.GetReceiptsByExpense(expense)
	if err != nil {
		log.Println("error at docList", err)
	}
	docs := len(docList)

	job, jobDaoErr := dao.GetJobById(expense.JobId)
	if jobDaoErr != nil {
		log.Print("err at get expense/id Error: jobDaoErr", jobDaoErr)
	}
	ClientData,_ := dao.GetClientById(job.ClientId)
	activeOrg, orgDaoErr := dao.GetOrgByJobId(job.Id)

	if orgDaoErr != nil{
		log.Print("err at get expense/id Error: orgDaoErr", orgDaoErr)
	}


	data := struct {
		Error []error
		User *models.User
		Job *models.Job
		Org *models.Org
		Client *models.Client
		Expense *models.Expense
		ReceiptMap map[*models.Expense] *models.Receipt
		DocList int

	} {
		User: activeUser,
		Job: job,
		Org: activeOrg,
		Client: ClientData,
		Expense: expense,
		ReceiptMap: receiptMap,
		DocList: docs,

	}

	return c.Render(200, "expense", data)




}

func NewExpense(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)

		var newExp models.Expense
		expName := c.FormValue("exp-name")
		newExp.Name = expName
		expCost := c.FormValue("exp-cost")
		newExp.Cost = expCost
		expDesc := c.FormValue("exp-description")
		newExp.Description = expDesc

		jobId, convErr := strconv.Atoi(c.FormValue("job-id"))
		newExp.JobId = jobId
		if convErr != nil {
			log.Println(convErr)
		}
		_, err := dao.CreateExpense(&newExp)
		if err != nil {
			log.Println(err)
			return c.NoContent(400)
		}

		return c.Redirect(302, "/job/" + strconv.Itoa(jobId))
	} 
	return c.Redirect(302, "/")
}

func EditExpense(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if !hasUser {
		return c.Redirect(302, "/")
	}
	cookie, _ := c.Cookie("sid")
	log.Print(id, " ", cookie.Value)
	expId, err := strconv.Atoi(c.Param("id"))
	var errs []error
	if err != nil {
		log.Print(err)
		errs = append(errs, err)
	}
	expense, expDaoErr := dao.GetExpenseById(expId)
	expense.Id = expId
	expense.Name = c.FormValue("exp-name")
	expense.Cost = c.FormValue("exp-cost")
	dao.UpdateExpenseById(expense)
	if expDaoErr != nil {
		log.Print(expDaoErr)
		errs = append(errs, expDaoErr)
	}

	return c.Redirect(302, "/expense/" + strconv.Itoa(expId))


}
