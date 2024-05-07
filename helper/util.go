package helper

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPagenationPoint(c *gin.Context) (pageSize int, pageNo int) {

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 0 {
		pageSize = 10
	}

	pageNo, err = strconv.Atoi(c.Query("pageNo"))
	if err != nil || pageNo < 0 {
		pageNo = 0
	}
	return pageSize, pageNo

}

func GetCurentTime() (current_time time.Time) {
	current_time, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Panicln("error while parsing date")
	}
	return current_time
}
