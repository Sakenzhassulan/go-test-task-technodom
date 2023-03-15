package admin

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/db/models"
	"github.com/Sakenzhassulan/go-test-task-technodom/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"net/http"
	"strconv"
)

func GetRedirects(c *gin.Context, instance *db.Instance, cache *store.Store) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset := (page - 1) * limit
	cursor, err := instance.Collection.Find(context.Background(), bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var redirects []models.Redirect
	if err = cursor.All(context.Background(), &redirects); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalCount, err := instance.Collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pageCount := int(math.Ceil(float64(totalCount) / float64(limit)))
	nextPage := page + 1
	if nextPage > pageCount {
		nextPage = -1
	}

	c.JSON(http.StatusOK, gin.H{
		"data": redirects,
		"meta": gin.H{
			"totalCount":  totalCount,
			"pageCount":   pageCount,
			"currentPage": page,
			"nextPage":    nextPage,
		},
	})
}
