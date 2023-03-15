package admin

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/db/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UpdateRedirectRequest struct {
	NewActiveLink string `json:"active_link" validate:"required"`
}

func UpdateRedirect(c *gin.Context, instance *db.Instance, validate *validator.Validate) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	body := &UpdateRedirectRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	redirect := &models.Redirect{}
	err = instance.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(redirect)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Redirect not found",
		})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"active_link":  body.NewActiveLink,
			"history_link": redirect.ActiveLink,
		},
	}
	_, err = instance.Collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedRedirect := &models.Redirect{}
	err = instance.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(updatedRedirect)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Redirect not found",
		})
		return
	}
	c.JSON(http.StatusOK, updatedRedirect)
}
