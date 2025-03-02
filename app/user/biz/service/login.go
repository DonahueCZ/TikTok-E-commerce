package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/user/biz/dal/dao"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user_service.LoginReq) (resp *user_service.LoginResp, err error) {
	// Finish your business logic.

	userDAO := dao.GetUserDAO()
	if userDAO == nil {
		klog.Error("userDAO 未初始化")
		return &user_service.LoginResp{
			ResponseStatus: buildErrorResponse("系统错误，请稍后重试", false),
		}, nil
	}
	user, err := userDAO.FindByEmail(req.Email)
	if err != nil {
		klog.Error("登录失败，用户不存在：", err)
		return &user_service.LoginResp{
			ResponseStatus: buildErrorResponse("用户不存在", false),
		}, nil
	}
	if user == nil {
		klog.Info("用户不存在")
		return &user_service.LoginResp{
			ResponseStatus: buildErrorResponse("用户不存在", false),
		}, nil
	}

	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password+user.Salt)); err != nil {
		klog.Error("密码验证失败:", err)
		return &user_service.LoginResp{
			ResponseStatus: buildErrorResponse("邮箱或密码不正确", false),
		}, nil
	}
	klog.Info("用户注册成功")
	return &user_service.LoginResp{
		ResponseStatus: &user_service.ResponseStatus{
			Message: "登录成功",
			Status:  true,
		},
	}, nil
}
