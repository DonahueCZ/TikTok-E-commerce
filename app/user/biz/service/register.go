package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/usermd"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user_service.RegisterReq) (resp *user_service.RegisterResp, err error) {
	// Finish your business logic.

	// 检查用户名和邮箱是否已存在
	userDAO := dao.GetUserDAO()
	if existingUser, err := userDAO.FindByEmail(req.Email); err != nil {
		klog.Error("检查邮箱是否存在时出错:", err)
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("注册失败，请稍后重试", false),
		}, nil
	} else if existingUser != nil {
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("该邮箱已注册", false),
		}, nil
	}

	if existingUser, err := userDAO.FindByUsername(req.UserName); err != nil {
		klog.Error("检查用户名是否存在时出错:", err)
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("注册失败，请稍后重试", false),
		}, nil
	} else if existingUser != nil {
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("该用户名已存在", false),
		}, nil
	}

	// 验证密码和确认密码是否一致
	if req.Password != req.ConfirmPassword {
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("密码和确认密码不一致", false),
		}, nil
	}

	// 生成随机盐值
	salt, err := generateSalt(16)
	if err != nil {
		klog.Error("生成盐值失败:", err)
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("注册失败，请稍后重试", false),
		}, nil
	}
	//加密密码
	hashedPassword, err := hashPasswordWithSalt(req.Password, salt)
	if err != nil {
		klog.Error("密码加密失败:", err)
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("注册失败，请稍后重试", false),
		}, nil
	}

	data := &usermd.User{
		Email:           req.Email,
		Username:        req.UserName,
		Password:        hashedPassword,
		Salt:            salt,
		UserPermissions: req.UserPermissions,
		Created_at:      time.Now(),
		Updated_at:      time.Now(),
	}

	if err := userDAO.Insert(data); err != nil {
		klog.Error("用户注册失败:", err)
		return &user_service.RegisterResp{
			ResponseStatus: buildErrorResponse("注册失败", false),
		}, nil
	}

	return &user_service.RegisterResp{
		UserId:         data.UserId,
		ResponseStatus: buildErrorResponse("注册成功", true),
	}, nil
}
