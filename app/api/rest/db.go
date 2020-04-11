package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/morteza-r/flexdb-server/app/api/rest/req"
	"github.com/morteza-r/flexdb-server/app/application"
	"net/http"
)

type DbController struct {
	DbService application.DbService
}

func (pc *DbController) Query(c *gin.Context) {
	var queryReq req.QueryRequest
	err := c.ShouldBindJSON(&queryReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query, err := queryReq.GetQuery()
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}
	result, err := pc.DbService.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	//TODO check if manual json can save time ! 140 micro second

	c.JSON(200, gin.H{"status": true, "data": result, "took": query.Took.String()})
}
