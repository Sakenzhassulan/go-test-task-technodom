package admin

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/db/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetRedirectById(c *gin.Context, instance *db.Instance) {
	redirect := &models.Redirect{}
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	filter := bson.M{"_id": id}
	err = instance.Collection.FindOne(context.Background(), filter).Decode(redirect)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Redirect not found",
		})
		return
	}
	c.JSON(http.StatusOK, redirect)
}
