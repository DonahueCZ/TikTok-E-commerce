package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/user"
	userrpcclent "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/user_rpc"
	user_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"

	userrpccl "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userrpcclent.InitUserRpcClient()

	res, err := userrpccl.DeleteUser(h.Context, &user_service.DeleteUserReq{
		UserId: req.UserID,
	})
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		h.RequestContext.JSON(consts.StatusBadRequest, &user.DeleteUserResp{
			ResponseStatus: &user.ResponseStatus{
				Status:  false,
				Message: err.Error(),
			},
		})
		return
	}

	resp = &user.DeleteUserResp{
		ResponseStatus: &user.ResponseStatus{
			Status:  res.ResponseStatus.Status,
			Message: res.ResponseStatus.Message,
		},
	}
	return resp, nil
}
