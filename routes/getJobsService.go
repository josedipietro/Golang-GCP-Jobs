package routes

import (
	"log"
	"net/http"

	"teamcore/test/firestore"

	"github.com/gin-gonic/gin"
)

func getJob(c *gin.Context) {
	id := c.Param("id")

	job, err := firestore.GetJob(id)

	if (err != nil) {
		log.Printf("%s", err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, job)
}
