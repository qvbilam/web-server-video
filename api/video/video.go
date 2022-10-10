package video

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"video/api"
	proto "video/api/qvbilam/video/v1"
	"video/global"
	"video/validate"
)

func Create(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}

func Detail(ctx *gin.Context) {

}

func List(ctx *gin.Context) {
	request := validate.VideoList{}
	if err := ctx.BindQuery(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	fmt.Println("排序请求: ", request.Sort)
	client := global.VideoServerClient
	rsp, err := client.Get(context.Background(), &proto.SearchVideoRequest{
		UserId:           request.UserId,
		RegionId:         request.RegionId,
		CategoryId:       request.CategoryId,
		Keyword:          request.Keyword,
		Sort:             request.Sort,
		IsRecommend:      request.IsRecommend,
		IsNew:            request.IsNew,
		IsHot:            request.IsHot,
		IsEnd:            request.IsEnd,
		IsVisible:        request.IsVisible,
		TotalCountMin:    request.TotalCountMin,
		TotalCountMax:    request.TotalCountMax,
		FavoriteCountMin: request.FavoriteCountMin,
		FavoriteCountMax: request.FavoriteCountMax,
		LikeCountMin:     request.LikeCountMin,
		LikeCountMax:     request.LikeCountMax,
		PlayCountMin:     request.PlayCountMin,
		PlayCountMax:     request.PlayCountMax,
		BarrageCountMin:  request.BarrageCountMin,
		BarrageCountMax:  request.BarrageCountMax,
		Page:             request.Page,
		PerPage:          request.PerPage,
	})

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	fmt.Printf("%+v\n", request)
	ctx.JSON(http.StatusOK, rsp)
}
