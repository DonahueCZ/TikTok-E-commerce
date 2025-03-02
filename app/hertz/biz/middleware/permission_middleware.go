package middleware

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/dao"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"strconv"
)

func CheckPermissionMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userIDStr := c.Query("user_id")
		if userIDStr == "" {
			c.String(http.StatusBadRequest, "无效的用户ID")
			c.Abort()
			return
		}

		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的用户ID")
			c.Abort()
			return
		}

		userDAO := dao.GetUserDAO()
		permissions, err := userDAO.GetUserPermissions(ctx, userID)
		if err != nil {
			klog.Error("查询用户权限失败:", err)
			c.String(http.StatusInternalServerError, "查询用户权限失败")
			c.Abort()
			return
		}

		if permissions == 0 {
			c.String(http.StatusForbidden, "没有权限访问")
			c.Abort()
			return
		} else if permissions != 1 {
			c.String(http.StatusForbidden, "无效的用户权限")
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
