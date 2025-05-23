package service

import (
	"context"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user_service.LogoutReq) (resp *user_service.LogoutResp, err error) {
	// Finish your business logic.

	klog.Info("退出登录成功")
	return &user_service.LogoutResp{
		ResponseStatus: buildErrorResponse("退出登录成功", true),
	}, nil
}
