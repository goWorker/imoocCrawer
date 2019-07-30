package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/util"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Info("User careate function called.", lager.Data{"X-Requesr-Id":util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind,nil)
	}

	u := model.UserModel {
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation,nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrValidation,nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrValidation,nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}
	handler.SendResponse(c,nil,rsp)
}