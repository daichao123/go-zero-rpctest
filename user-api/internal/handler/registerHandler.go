package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-rpctest/user-api/common/result"
	"go-zero-rpctest/user-api/common/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-rpctest/user-api/internal/logic"
	"go-zero-rpctest/user-api/internal/svc"
	"go-zero-rpctest/user-api/internal/types"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq

		if err := httpx.Parse(r, &req); err != nil {
			httpx.WriteJson(w, http.StatusOK, result.Error(xerr.REUQEST_PARAM_ERROR, "参数错误"))
			return
		}
		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			logx.Info(err)
			httpx.WriteJson(w, http.StatusOK, result.Error(xerr.REUQEST_PARAM_ERROR, "请输入正确的参数"))
			return
		}
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		result.HttpResult(r, w, resp, err)
	}
}
