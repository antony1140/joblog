package main

import (
	// "errors"
	"fmt"
	"os"
	// "net/http"

	// "os"
	"html/template"
	// "context"
	"io"
	"log"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/antony1140/joblog/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type Templates struct {
	templates *template.Template

}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Templates) Render (w io.Writer, name string, data interface{}, c echo.Context)error{ 
	return t.templates.ExecuteTemplate(w, name, data)
}

type orgList struct {
	list []models.Org
}


type errList = []LoginErr

type LoginErr struct {
	Err bool
}

type indexData struct {
	errs []LoginErr
}

func newIndexData(loginErr []LoginErr)(indexData){
	return indexData{
		errs: loginErr,
	}
}



func main(){
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"*"},
      AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/", "views")
	e.Static("/views", "icons")

	service.DownloadReceipt(1)




	e.GET("/", func(c echo.Context) error {
		// haveSesh, _ := security.GetSession(c)
		// if haveSesh {
			// userId := id	
		// }
		hasUser, id := security.GetSession(c)	
		log.Print("hasUser: ", hasUser)
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)

			log.Print("got to render home")
		return c.Redirect(302, "home")

		}
		

		return c.Render(200, "index","" )
	})

	e.GET("/testauth", func(c echo.Context) error {
		user := models.User{
			Name: "rick james",
			Username: "ricky123",
			Password: "54321",
		}

		id, _ := service.CreateUser(&user)
		fmt.Println("new id is :",id ) 

		return nil
	})

	e.GET("/testLogin", func(c echo.Context) error {
		user, err := service.LoginUser("antony1140", "12345")
		if err != nil {
			return err
		}
		fmt.Println(user.Name, user.Username)
		return nil
	})
	e.POST("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		pass := c.FormValue("password")
		user, err := service.LoginUser(username, pass)
		if err != nil {
			return c.Render(404, "index", "Invalid Credentials, Try again.")
		}
		fmt.Println(user.Name, user.Username)
		fmt.Println(err)
		sessionCookie := security.CreateSession(true, user.Id)
		log.Print("session created for user ", user.Id, " session: ", sessionCookie.Value)
		c.SetCookie(sessionCookie)

		
			// return c.Render(404, "home", "")
			log.Print("got to redirect /home")
			return c.Redirect(302, "/home")
	})
	
	e.GET("/logout", func(c echo.Context) error {
		hasUser, id := security.GetSession(c)	
		log.Print("userid: ", id)
		log.Print("hasUser: ", hasUser)
		if hasUser {
			cookie, _ := c.Cookie("sid")
			log.Print(id, " ", cookie.Value)
			err := security.DestroySession(cookie.Value)
			if err != nil {
				return c.Redirect(302, "/")
			}

		}

			return c.Redirect(302, "/")
	})

	e.GET("/home", func(c echo.Context) error {
		hasUser, id := security.GetSession(c)	
		log.Print("userid: ", id)
		log.Print("hasUser: ", hasUser)
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)

			// log.Print("got to render home")
		// return c.Render(200, "home", "")
		//
		orgs, sqlErr := dao.GetAllOrgsByUserId(id)
		jobs, err := dao.GetAllJobsByUserId(id)
		if err != nil{
			log.Print(err)

		}
		if sqlErr != nil {
			log.Print(sqlErr)
		}
		log.Print("orgs list: ", len(orgs))
			
		for _, org := range orgs {
			fmt.Println("what is this")
			fmt.Println(org.Name, "id: ", org.Id)
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

			log.Print("got to render home")
		return c.Render(200, "home", data)
		///

		}
			log.Print("got to redirect /")
		
		return c.Redirect(302, "/")


	})

	e.GET("/create/:type", func(c echo.Context) error{

		hasUser, id := security.GetSession(c)	
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)
		choice := c.Param("type")
		activeUser,_ := dao.GetUserById(id)
		data := struct {
			User *models.User
			Choice string

		} {
			User: activeUser,
			Choice: choice,

		}

		return c.Render(200, "create" + choice, data)

		}
		
		return c.Redirect(302, "/")
	})


	//TODO: FINISH 
	e.POST("/editExp/:id", func(c echo.Context) error {

		hasUser, id := security.GetSession(c)	
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)
		expId, err := strconv.Atoi(c.Param("id"))
		var errs []error
		code := 200
		if err != nil {
			log.Print(err)
			errs = append(errs, err)
			code = 500
		}
		activeUser,_ := dao.GetUserById(id)
		expense, expDaoErr := dao.GetExpenseById(expId)
		expense.Id = expId
		expense.Name = c.FormValue("expName")
		expense.Cost = c.FormValue("expCost")
		dao.UpdateExpenseById(expense)
		job, jobDaoErr := dao.GetJobById(expense.JobId)
		expenseList, expListErr := dao.GetAllExpensesByJobId(expense.JobId)
		client, clientDaoErr := dao.GetClientById(job.ClientId)
		if expListErr != nil {
			log.Print(expListErr)
			errs = append(errs, expListErr)
			code = 500
		}
		if jobDaoErr != nil {
			log.Print(jobDaoErr)
			errs = append(errs, jobDaoErr)
			code = 500
		}
		if expDaoErr != nil {
			log.Print(expDaoErr)
			errs = append(errs, expDaoErr)
			code = 500
		}
		if clientDaoErr != nil {
			log.Print(jobDaoErr)
			errs = append(errs, jobDaoErr)
			code = 500
		}
		
		data := struct {
			Error []error
			User *models.User
			Job *models.Job
			Client *models.Client
			ExpenseList []models.Expense

		} {
			Error: errs,
			User: activeUser,
			Job: job,
			Client: client,
			ExpenseList: expenseList,
		}

			return c.Render(code, "jobPage", data)
		}
		
		return c.Redirect(302, "/")
	})


	e.GET("/expense/:id", func(c echo.Context) error {
		// haveSesh, _ := security.GetSession(c)
		// if haveSesh {
			// userId := id	
		// }
		hasUser, id := security.GetSession(c)	
		log.Print("hasUser: ", hasUser)
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)
		exp := c.Param("id")
		expId,_ := strconv.Atoi(exp)
		fmt.Println("debug expId parameter as integer, ", expId)

		activeUser,_ := dao.GetUserById(id)
		expense, expDaoErr := dao.GetExpenseById(expId)
		fmt.Println("debug expId after dao, ", expense.Id)
		
		if expDaoErr != nil {
			log.Print("err at get expense/id Error: expDaoErr", expDaoErr)
		}
		var expToList []models.Expense
		
		tempExp := *expense
		expToList = append(expToList, tempExp)
		fmt.Println("debug expToList id and name, ", tempExp.Id, tempExp.Name)
		receiptMap := service.GroupExpenseReceipts(expToList)

		job, jobDaoErr := dao.GetJobById(expense.JobId)
		if jobDaoErr != nil {
			log.Print("err at get expense/id Error: jobDaoErr", jobDaoErr)
		}
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

		} {
			User: activeUser,
			Job: job,
			Org: activeOrg,
			Expense: expense,
			ReceiptMap: receiptMap,

		}

			log.Print("got to render home")
		return c.Render(200, "expense", data)

		}
		

		return c.Render(200, "index","" )
	})

	e.POST("/upload/receipt/:id", func(c echo.Context) error {
		hasUser, id := security.GetSession(c)	
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)
		client := data.InitS3()	
		Id := c.FormValue("expId")
		expId,_ := strconv.Atoi(Id)
		file, header, err := c.Request().FormFile("file")
		if err != nil {
			log.Print("err 1", err)
			
		return c.NoContent(400)	
		}

		_,osErr := os.Create("./assets/" + header.Filename)
		if osErr != nil {
			log.Print("err 2, ", err)
		return c.NoContent(400)	
		}

		_, daoErr := dao.CreateReceipt(header.Filename, expId )
		fileKey := "receipts/" + Id + "/" + header.Filename
		if daoErr != nil {
			log.Print(err)
			return c.NoContent(400)
		}
		
		s3Err := data.UploadS3(client, file, fileKey)
		if s3Err != nil {
			log.Print("err 3, ", s3Err)
		}

		log.Print("made file, " + header.Filename)

		html := `
							<td class="upload-rec-row">
							<a href={{$receipt.S3Url}}><button> Dowload/view </button></a>
							</td>`
							

		return c.HTML(200, html)	
	}

		return c.Redirect(302, "/")
	})

	// e.GET("/download/receipt/:id", func(c echo.Context) error {
	//
	// 	hasUser, id := security.GetSession(c)	
	// 	if hasUser {
	// 		cookie, _ := c.Cookie("sid")
	// 		log.Print(id, " ", cookie.Value)
	//
	// 		expId := c.Param("id")
	// 		expIdStr, err := strconv.Atoi(expId)
	// 		if err != nil {
	// 			log.Println("error at receipt download, ", err)
	// 		}
	// 		request, err := service.DownloadReceipt(expIdStr)
	//
	// 		log.Println(request.URL)
	// 		c.Redirect(301, request.URL)
	//
	//
	//
	// 	}
	// 	return c.Redirect(302, "/")
	// })

	e.POST("/newGroup", func(c echo.Context) error {
		hasUser, id := security.GetSession(c)	
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)
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
		
		return c.Redirect(302, "/")
		
	})
	e.GET("/group/:id", func(c echo.Context) error {
		hasUser, id := security.GetSession(c)	
		log.Print("userid: ", id)
		log.Print("hasUser: ", hasUser)
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)
		orgId, _ := strconv.Atoi(c.Param("id"))
		jobs, sqlErr := dao.GetAllJobsByOrgId(orgId)
		if sqlErr != nil {
			fmt.Println(sqlErr)
		}
		log.Print("jobs list: ", len(jobs))
			
		for _, job := range jobs {
			fmt.Println("what is this")
			fmt.Println(job.Title)
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

		} {
			User: activeUser,
			Jobs: jobs,
			Org: activeOrg,
		}

			log.Print("got to render orgPage")
		return c.Render(200, "orgPage", data)

		}
			log.Print("got to redirect /")
		
		return c.Redirect(302, "/")
		
	})

	e.GET("/job/:id", func(c echo.Context) error {
		log.Print("reached job/id")
		hasUser, id := security.GetSession(c)	
		log.Print("userid: ", id)
		log.Print("hasUser: ", hasUser)
		if hasUser {
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

			fmt.Println("debug receipt", receipt.S3Url)
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
		
		return c.Redirect(302, "/")
		
	})

	e.GET("/orgs/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		orgs, sqlErr := dao.GetAllOrgsByUserId(id)
		if sqlErr != nil {
			fmt.Println(sqlErr)
		}
			
		for _, org := range orgs {
			fmt.Println("what is this")
			fmt.Println(org.Name)
		}
		activeUser,_ := dao.GetUserById(id)
		data := struct {
			User *models.User
			Orgs []models.Org

		} {
			User: activeUser,
			Orgs: orgs,
		}

		return c.Render(200, "home", data)
	})




	//organizations


	var test models.Job 
	test.Title = "test title"
	test.Description = "test Description"
	dao.CreateJob(&test)
	client := data.InitS3()
	data.DownloadS3(client, "testDownload")

	

	println("server running")
	e.Logger.Fatal(e.Start(":3333"))
	// serverErr := http.ListenAndServe("127.0.0.1:3333", nil)
	//
	// if errors.Is(serverErr, http.ErrServerClosed){
	// 	fmt.Printf("server closed\n")
	// }else if serverErr != nil {
	// 	fmt.Printf("error starting server: %s\n", serverErr)
	// 	os.Exit(1)
	// }


}

