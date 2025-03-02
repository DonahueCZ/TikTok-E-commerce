package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/dao"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user_service.GetUserReq) (resp *user_service.GetUserResp, err error) {
	// Finish your business logic.

	userDAO := dao.GetUserDAO()
	user, err := userDAO.FindOne(s.ctx, req.UserId)
	if err != nil {
		klog.Error("用户不存在：", err)
		return &user_service.GetUserResp{
			ResponseStatus: buildErrorResponse("用户不存在", false),
		}, nil
	}
	if user == nil {
		return &user_service.GetUserResp{
			ResponseStatus: buildErrorResponse("用户不存在", false),
		}, nil
	}
	return &user_service.GetUserResp{
		UserId:          user.UserId,
		Email:           user.Email,
		UserName:        user.Username,
		UserPermissions: user.UserPermissions,
		CreatedAt:       user.Created_at.Format("2006-01-02 15:04:05"),
		UpdatedAt:       user.Updated_at.Format("2006-01-02 15:04:05"),
		ResponseStatus:  buildErrorResponse("用户获取成功", true),
	}, nil
}
