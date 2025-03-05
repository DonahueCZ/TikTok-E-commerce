package router

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func CustomizedRegister(r *server.Hertz) {
	r.GET("/permission", user.CheckPermissionMiddleware(), func(ctx context.Context, c *app.RequestContext) {
	})
}
