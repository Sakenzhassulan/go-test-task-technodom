package admin

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func DeleteRedirect(c *gin.Context, instance *db.Instance) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	_, err = instance.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	c.JSON(http.StatusOK, gin.H{
		"message": "Redirect deleted successfully",
	})
}
