package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/user"
	userrpcclent "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/user_rpc"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	userrpccl "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserService(Context context.Context, RequestContext *app.RequestContext) *GetUserService {
	return &GetUserService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userrpcclent.InitUserRpcClient()
	res, err := userrpccl.GetUser(h.Context, &user_service.GetUserReq{
		UserId: req.UserID,
	})
	if err != nil {
		h.RequestContext.String(consts.StatusBadRequest, err.Error())
		return
	}
	responseStatus := &user.ResponseStatus{
		Status:  res.ResponseStatus.Status,
		Message: res.ResponseStatus.Message,
	}
	resp = &user.GetUserResp{
		UserID:          res.UserId,
		UserName:        res.UserName,
		Email:           res.Email,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       res.UpdatedAt,
		UserPermissions: res.UserPermissions,
		ResponseStatus:  responseStatus,
	}
	return resp, nil
}
