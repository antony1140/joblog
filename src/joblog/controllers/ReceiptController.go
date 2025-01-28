package controllers

import (
	"log"
	"os"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/antony1140/joblog/service"
	"github.com/labstack/echo/v4"
)

func UploadReceipt(c echo.Context) (error) {
	hasUser, _ := security.GetSession(c)	
	if !hasUser {
		return c.Redirect(302, "/")
	}
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

		db := data.OpenDb()
		tx, err := db.Begin()
		recId, daoErr := dao.CreateReceipt(tx, header.Filename, expId )
		fileKey := "receipts/" + Id + "/" + header.Filename
		if daoErr != nil {
			log.Print(err)
			return c.NoContent(400)
		}

		s3Err := data.UploadS3(client, file, fileKey)
		if s3Err != nil {
			log.Print("err 3, ", s3Err)
		}
		s3Err = nil

		url, s3Err := service.DownloadReceipt(expId)
		if s3Err != nil {
			log.Print("err 4, ", s3Err)
		}

		log.Print("made file, " + header.Filename)


		receipt, err := dao.GetReceiptById(recId)


		data := struct{
			Receipt *models.Receipt
			S3Url string
			ExpenseId int
		} {
			Receipt: receipt,
			S3Url: url.URL,
			ExpenseId: expId,

		}
		return c.Render(200, "uploadDocResponse", data )	

}

func PreviewDocument(c echo.Context) (error) {
	hasUser, _ := security.GetSession(c)
	if !hasUser {
		return c.Redirect(302, "/")
	}

	params := c.ParamValues()
	fileKey := params[0]
	expIdstr := params[1]
	log.Println("debug key, id: ", fileKey, expIdstr)
	expId, err := strconv.Atoi(expIdstr)
	if err != nil {
		log.Println("error getting document", err)
		return c.Redirect(302, "/expense" + expIdstr)
	}
	req, err := service.DownloadReceiptByFileKey(fileKey, expId)
	url := req.URL

	data := struct {
		S3Url string
	} {
		S3Url: url,
	}
	

	return c.Render(200, "preview", data)
}

func DownloadReceipt(c echo.Context) (error) {
	hasUser, id := security.GetSession(c)	
	if hasUser {
		cookie, _ := c.Cookie("sid")
		log.Print(id, " ", cookie.Value)

		expId := c.Param("id")
		expIdStr, err := strconv.Atoi(expId)
		if err != nil {
			log.Println("error at receipt download, ", err)
		}
		request, err := service.DownloadReceipt(expIdStr)

		log.Println(request.URL)
		c.Redirect(301, request.URL)



	}
	return c.Redirect(302, "/")

}

func DeleteReceipt(c echo.Context) (error) {

	hasUser, id := security.GetSession(c)	
	if !hasUser {
		return c.Redirect(302, "/")
	}
	cookie, _ := c.Cookie("sid")
	log.Print(id, " ", cookie.Value)

	// fileKey := c.FormValue("fileKey")
	exp := c.FormValue("expId")
	expId, err := strconv.Atoi(exp)
	if err != nil {
		log.Print("expense id was not an integer")
	}
	s3Err := service.DeleteReceipt(expId)
	if s3Err != nil {
		log.Print("s3 failed to delete receipt", s3Err)
		log.Print("do something here")
	}
	deleteErr := dao.DeleteReceiptByExpenseId(expId)
	if deleteErr != nil {
		log.Print("failed to delete receipt from db", deleteErr)
		log.Print("do something here")

	}


	activeExp, expDaoErr := dao.GetExpenseById(expId)
	if expDaoErr != nil {
		log.Print(expDaoErr)
	}
	var expList []*models.Expense
	expList = append(expList, activeExp)
	newMap := service.GroupExpenseReceipts(expList)

	data := struct {
		ReceiptMap map[*models.Expense]*models.Receipt
	} {
		ReceiptMap: newMap,
	}
	return c.Render(200, "receiptChangeReturn", data)

}


