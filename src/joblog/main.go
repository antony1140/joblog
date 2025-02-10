package main

import (
	"html/template"
	"log"
	"os"

	"io"

	"github.com/antony1140/joblog/controllers"
	"github.com/antony1140/joblog/data"
	"github.com/joho/godotenv"
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

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}


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

	
	e.GET("/register", func(c echo.Context) error {
		return c.Render(200, "register", "")
	})

	e.POST("/newuser", func(c echo.Context) error {
		return controllers.Register(c)
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

	e.POST("/newjob", func(c echo.Context) error {
		return controllers.NewJob(c)
	})




	// vars := os.Environ()	



	data.InitDb()
	if os.Getenv("MODE") == "prod" {
		log.Print(`///////////// 
		PRODUCTION  
		//////////`)
	} 
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))


}

