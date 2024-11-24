package service
import (
	// "github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/dao"
	"database/sql"
	"github.com/labstack/echo/v4"
	"fmt"
	"strconv"
)

type OrgService struct {
	Db *sql.DB
}

func CreateOrg(org *models.Org)(error){
	
	return nil
}

func GetAllOrgsByUserId(c echo.Context) (int, string, interface{}, error) {
	var data interface{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 400, "index", data, err
	}

	orgs, sqlErr := dao.GetAllOrgsByUserId(id)
	if sqlErr != nil {
		fmt.Println(sqlErr)
	}

	status := 200
	if len(orgs) < 1 {
		status = 400
	}

	return status, "index", orgs, nil
	
}
