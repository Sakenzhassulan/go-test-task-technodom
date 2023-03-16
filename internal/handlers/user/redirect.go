package user

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/db/models"
	"github.com/Sakenzhassulan/go-test-task-technodom/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Redirect(c *gin.Context, instance *db.Instance, cache *store.Store) {
	redirect := &models.Redirect{}
	link := c.DefaultQuery("link", "")
	// if link exists in cache get it from there
	if value, ok := cache.Get(link); ok {
		// in cache: key=active_link, value=active_link
		if value == link {
			c.Status(http.StatusOK)
			return
		}
		// in cache: key=history_link, value=active_link
		c.JSON(http.StatusMovedPermanently, gin.H{
			"active_link": value,
		})
		return
	}

	if isExists, _ := instance.IsActiveLinkExists(link); isExists {
		cache.Add(link, link)
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
	//if not write link to cache
	cache.Add(link, redirect.ActiveLink)
	c.JSON(http.StatusMovedPermanently, gin.H{
		"active_link": redirect.ActiveLink,
	})
}
