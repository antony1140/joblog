package main

import (



	"html/template"
	// "os"

	// "context"
	"io"

	"github.com/antony1140/joblog/controllers"
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


	// endpoint boiler plate
	// e.POST("", func(c echo.Context) error {
	// 	return nil
	// })


	e.GET("/", func(c echo.Context) error {
		return controllers.Index(c)
	})



	e.POST("/invoice", func(c echo.Context) error {
		return controllers.NewInvoice(c)	
	})


	e.GET("/succesesinvoice", func(c echo.Context) error {

		return nil
	})
	

	e.POST("/login", func(c echo.Context) error {
		return controllers.Login(c)
	})
	
	e.GET("/logout", func(c echo.Context) error {
		return controllers.Logout(c)
	})

	e.GET("/home", func(c echo.Context) error {
		return controllers.Home(c)

	})

	e.POST("/create/:type", func(c echo.Context) error{
		return controllers.Create(c)
	})


	e.POST("/expense/edit/:id", func(c echo.Context) error {
		return controllers.EditExpense(c)
	})


	e.GET("/expense/:id", func(c echo.Context) error {
		return controllers.Expense(c)
	})

	e.POST("/newexpense", func(c echo.Context) error {
		return controllers.NewExpense(c)
	})

	e.POST("/upload/receipt/:id", func(c echo.Context) error {
		return controllers.UploadReceipt(c)
	})


	e.POST("/delete/receipt", func(c echo.Context) error {
		return controllers.DeleteReceipt(c)
	})

	e.GET("/download/receipt/:id", func(c echo.Context) error {
		return controllers.DownloadReceipt(c)
	})

	e.GET("/preview/:filekey/:expid", func(c echo.Context) error {
		return controllers.PreviewDocument(c)
	})

	e.POST("/newGroup", func(c echo.Context) error {
		return controllers.NewGroup(c)	
	})
	e.GET("/group/:id", func(c echo.Context) error {
		return controllers.Group(c)	
	})

	e.GET("/job/:id", func(c echo.Context) error {
		return controllers.Job(c)	
	})

	// e.POST("/newjob", func(c echo.Context) error {
	//
	// 	hasUser, id := security.GetSession(c)	
	// 	log.Print("userid: ", id)
	// 	log.Print("hasUser: ", hasUser)
	// 	if hasUser {
	// 	cookie, _ := c.Cookie("sid")
	// 	log.Print(id, " ", cookie.Value)
	//
	// 	// activeUser,_ := dao.GetUserById(id)
	//
	// 	orgIdString := c.FormValue("org-id")
	// 	orgId, convErr := strconv.Atoi(orgIdString)
	// 	if convErr != nil {
	// 		fmt.Println("error at newjob id to int conversion, need to do something here")
	// 	}
	//
	// 	activeOrg, orgDaoErr := dao.GetOrgById(orgId)
	// 	if orgDaoErr != nil {
	// 		log.Print(orgDaoErr)
	// 	}
	// 	data := struct {
	//
	// 	}
	// }
	// 	return c.Redirect(302, "/")
	// })




	

	e.Logger.Fatal(e.Start(":3333"))


}

