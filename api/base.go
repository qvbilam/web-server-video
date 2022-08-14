package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"video/global"
)

// HandleGrpcErrorToHttp 处理 grpc 错误转换成 http 错误
func HandleGrpcErrorToHttp(ctx *gin.Context, err error) {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": s.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": fmt.Sprintf("服务错误: %d; 信息: %s", s.Code(), err.Error()),
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": fmt.Sprintf("参数错误: %d; 信息: %s", s.Code(), err.Error()),
				})
			case codes.Unavailable:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": fmt.Sprintf("用户服务: %d; 信息: %s", s.Code(), err.Error()),
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": fmt.Sprintf("其他错误: %d; 信息: %s", s.Code(), err.Error()),
				})
			}

			return
		}
	}
}

// HandleValidateError 处理参数验证错误
func HandleValidateError(ctx *gin.Context, err error) {
	// 如果翻译错误
	errs, ok := err.(validator.ValidationErrors)
	if !ok { // 直接返回错误信息
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 翻译成功后返回翻译之后的错误信息
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)), // 使用全局变量 trans 翻译中文内容
	})
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}

	// 循环处理传进来的错误信息
	for field, err := range fields {
		// 获取第一个 . 出现的位置
		position := strings.Index(field, ".")
		fmt.Println(field, position)

		// 获取截取后的 key 名
		key := field[position+1:]

		// 追加返回结果
		res[key] = err
	}
	return res
}
