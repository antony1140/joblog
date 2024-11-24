package service
import (
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/dao"
	"database/sql"
)

type JobService struct {
	Db *sql.DB
}

func GetJobById(id int)(*models.Job, error){
	job, err := dao.GetJobById(id)	
	if err != nil {
		return nil, err
	}
	return job, nil
}
