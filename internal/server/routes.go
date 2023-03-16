package server

import (
	"github.com/Sakenzhassulan/go-test-task-technodom/internal/handlers/admin"
	"github.com/Sakenzhassulan/go-test-task-technodom/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, client *Client) {
	r.GET("/admin/redirects", client.GetRedirects)
	r.GET("/admin/redirects/:id", client.GetRedirectById)
	r.POST("/admin/redirects", client.CreateRedirect)
	r.PATCH("/admin/redirects/:id", client.UpdateRedirect)
	r.DELETE("/admin/redirects/:id", client.DeleteRedirect)
	r.GET("/redirects", client.Redirect)
}

func (client *Client) GetRedirects(ctx *gin.Context) {
	admin.GetRedirects(ctx, client.DB, client.Cache)
}

func (client *Client) GetRedirectById(ctx *gin.Context) {
	admin.GetRedirectById(ctx, client.DB)
}

func (client *Client) CreateRedirect(ctx *gin.Context) {
	admin.CreateRedirect(ctx, client.DB, client.Validator)
}

func (client *Client) UpdateRedirect(ctx *gin.Context) {
	admin.UpdateRedirect(ctx, client.DB, client.Validator)
}

func (client *Client) DeleteRedirect(ctx *gin.Context) {
	admin.DeleteRedirect(ctx, client.DB)
}

func (client *Client) Redirect(ctx *gin.Context) {
	user.Redirect(ctx, client.DB, client.Cache)
}
