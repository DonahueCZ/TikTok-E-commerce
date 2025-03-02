package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/dao"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user_service.UpdateUserReq) (resp *user_service.UpdateUserResp, err error) {
	// Finish your business logic.
	userDAO := dao.GetUserDAO()
	user, err := userDAO.FindOne(s.ctx, req.UserId)
	if err != nil {
		klog.Info("用户不存在：", err)
		return &user_service.UpdateUserResp{
			ResponseStatus: buildErrorResponse("查询用户失败，请稍后重试", false),
		}, nil
	}
	if user == nil {
		klog.Info("用户不存在")
		return &user_service.UpdateUserResp{
			ResponseStatus: buildErrorResponse("用户不存在", false),
		}, nil
	}
	// 验证当前密码
	saltedCurrentPassword := req.CurrentPassword + user.Salt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(saltedCurrentPassword)); err != nil {
		klog.Error("当前密码错误：", err)
		return &user_service.UpdateUserResp{
			ResponseStatus: buildErrorResponse("当前密码错误", false),
		}, nil
	}

	//更新邮箱
	if req.NewEmail != "" {
		user.Email = req.NewEmail
	}

	//更新密码
	if req.NewPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			klog.Error("密码加密失败:", err)
			return &user_service.UpdateUserResp{
				ResponseStatus: buildErrorResponse("更新失败，请稍后重试", false),
			}, nil
		}
		user.Password = string(hashedPassword)
	}

	//更新用户名
	if req.NewUserName != "" {
		user.Username = req.NewUserName
	}

	//更新数据库和Redis缓存
	user.Updated_at = time.Now()
	if err := userDAO.Update(user); err != nil {
		klog.Error("更新用户失败：", err)
		return &user_service.UpdateUserResp{
			ResponseStatus: buildErrorResponse("更新用户失败", false),
		}, nil
	}

	// 同步更新 Redis 缓存
	cacheKey := user.GetCacheKey()
	if err := userDAO.Cache().Set(s.ctx, cacheKey, user, time.Hour).Err(); err != nil {
		klog.Error("更新 Redis 缓存失败：", err)
	}
	klog.Info("更新成功")
	return &user_service.UpdateUserResp{
		ResponseStatus: buildErrorResponse("更新成功", true),
	}, nil
}
