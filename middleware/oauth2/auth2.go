package oauth2

import (
	"github.com/kataras/iris"
)

// APIAuthorization 客户端中间件
func APIAuthorization(ctx iris.Context) {
	if ID := ctx.Params().Get("appid"); ID != "" {

		ctx.JSON(iris.StatusOK)
	}
}
