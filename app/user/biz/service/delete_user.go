package service

import (
	"context"
	"fmt"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/dao"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user_service.DeleteUserReq) (resp *user_service.DeleteUserResp, err error) {
	// Finish your business logic.

	userDAO := dao.GetUserDAO()
	err = userDAO.Delete(req.UserId)
	if err != nil {
		klog.Error("删除用户失败：", err)
		return &user_service.DeleteUserResp{
			ResponseStatus: buildErrorResponse("删除用户失败", false),
		}, nil
	}

	// 同步删除 Redis 缓存
	cacheKey := fmt.Sprintf("users:%d", req.UserId)
	if err := userDAO.Cache().Del(s.ctx, cacheKey).Err(); err != nil {
		klog.Error("删除 Redis 缓存失败：", err)
	}
	klog.Info("删除用户成功")
	return &user_service.DeleteUserResp{
		ResponseStatus: buildErrorResponse("删除用户成功", true),
	}, nil
}
