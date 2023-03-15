package user

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/db/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Redirect(c *gin.Context, instance *db.Instance) {
	redirect := &models.Redirect{}
	link := c.DefaultQuery("link", "")
	if isExists, _ := instance.IsActiveLinkExists(link); isExists {
		c.Status(http.StatusOK)
		return
	}
	err := instance.Collection.FindOne(context.Background(), bson.M{"history_link": link}).Decode(redirect)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Redirect not found",
		})
		return
	}
	c.JSON(http.StatusMovedPermanently, gin.H{
		"active_link": redirect.ActiveLink,
	})
}
