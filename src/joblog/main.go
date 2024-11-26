package main

import (
	// "errors"
	"fmt"
	"net/http"

	// "os"
	"html/template"
	// "context"
	"io"
	"log"
	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
	"github.com/antony1140/joblog/security"
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
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/", "views")
	e.Use(middleware.CORS())


	e.GET("/", func(c echo.Context) error {
		haveSesh, _ := security.GetSession(c)
		if haveSesh {
			// userId := id	
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
		log.Print("session created for user", user.Id, "session: ", sessionCookie.Value)
		c.SetCookie(sessionCookie)

		
			// return c.Render(404, "home", "")
			log.Print("got to redirect /home")
			return c.Redirect(302, "/home/")
	})

	e.GET("/home/", func(c echo.Context) error {
		hasUser, id := security.GetSession(c)	
		log.Print("hasUser: ", hasUser)
		if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)

			log.Print("got to render home")
		return c.Render(200, "home", "")

		}
			log.Print("got to redirect /")
		
		return c.Redirect(302, "/")


	})

	e.GET("/getJob", func(c echo.Context) error {
		 job, err := service.GetJobById(1)
		 if err != nil {
			 fmt.Println("error", err)
		 }
		 resp := models.PrintJob(job)
		 fmt.Print(resp)
		return c.String(http.StatusOK, resp)
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
		return c.Render(200, "home", orgs)
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

