package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/user"
	userrpcclent "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/user_rpc"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	userrpccl "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userrpcclent.InitUserRpcClient()
	res, err := userrpccl.Register(h.Context, &user_service.RegisterReq{
		Email:           req.Email,
		UserName:        req.UserName,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		UserPermissions: req.UserPermissions,
	})

	if err != nil {
		log.Printf("Failed to register user: %v", err)
		h.RequestContext.JSON(consts.StatusBadRequest, &user.RegisterResp{
			UserID: 0,
			ResponseStatus: &user.ResponseStatus{
				Status:  false,
				Message: err.Error(),
			},
		})
		return
	}

	resp = &user.RegisterResp{
		UserID: res.UserId,
		ResponseStatus: &user.ResponseStatus{
			Status:  res.ResponseStatus.Status,
			Message: res.ResponseStatus.Message,
		},
	}
	return resp, nil
}
