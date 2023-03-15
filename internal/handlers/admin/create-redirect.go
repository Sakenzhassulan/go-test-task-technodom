package admin

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/db/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CreateRedirectRequest struct {
	ActiveLink  string `json:"active_link" validate:"required"`
	HistoryLink string `json:"history_link" validate:"required"`
}

func CreateRedirect(c *gin.Context, instance *db.Instance, validate *validator.Validate) {
	body := &CreateRedirectRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	id := primitive.NewObjectID()
	if isExists, _ := instance.IsActiveLinkExists(body.ActiveLink); isExists {
		c.JSON(http.StatusConflict, gin.H{"error": "Such Active Link already exists"})
		return
	}
	if isExists, _ := instance.IsHistoryLinkExists(body.HistoryLink); isExists {
		c.JSON(http.StatusConflict, gin.H{"error": "Such History Link already exists"})
		return
	}
	redirect := &models.Redirect{
		Id:          id,
		ActiveLink:  body.ActiveLink,
		HistoryLink: body.HistoryLink,
	}
	_, err := instance.Collection.InsertOne(context.Background(), redirect)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, redirect)

}
