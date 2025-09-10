package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/xiao-hub-create/vblog/apps/token"
	"github.com/xiao-hub-create/vblog/apps/token/impl"
)

func NewTokenApiHandler() *TokenApiHandler {
	return &TokenApiHandler{
		token: impl.TokenService,
	}
}

type TokenApiHandler struct {
	token token.UserService
}

func (h *TokenApiHandler) Registry(r *gin.Engine) {
	router := r.Group("/vblog/api/v1/tokens")
	router.POST("", h.IssueToken)
	router.DELETE("", h.RevolkToken)
}

func (h *TokenApiHandler) IssueToken(ctx *gin.Context) {
	in := token.NewIssueTokenRequest("", "")
	if err := ctx.BindJSON(in); err != nil {

		response.Failed(ctx, err)
		return
	}
	ins, err := h.token.IssueToken(ctx.Request.Context(), in)
	if err != nil {

		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *TokenApiHandler) RevolkToken(ctx *gin.Context) {
	in := &token.RevokeTokenRequest{}
	if err := ctx.BindJSON(in); err != nil {

		response.Failed(ctx, err)
		return
	}
	ins, err := h.token.RevokeToken(ctx, in)
	if err != nil {

		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}
