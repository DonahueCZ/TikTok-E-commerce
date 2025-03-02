package router

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"net/http"
)

func CustomizedRegister(r *server.Hertz) {
	r.GET("/permission", middleware.CheckPermissionMiddleware(), func(ctx context.Context, c *app.RequestContext) {
		c.String(http.StatusOK, "欢迎")
	})
}
