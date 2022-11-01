package video

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"video/api"
	proto "video/api/qvbilam/video/v1"
	"video/global"
	"video/validate"
)

func Create(ctx *gin.Context) {
	var userId int64
	userId = 1
	paramDramaId := ctx.Param("id")
	dramaId, _ := strconv.Atoi(paramDramaId)
	request := validate.VideoCreate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	res, err := global.VideoServerClient.Create(context.Background(), &proto.UpdateVideoRequest{
		FileId:          request.FileId,
		CategoryId:      request.CategoryId,
		UserId:          userId,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
		IsVisible:       true,
		DramaId:         int64(dramaId),
		Episode:         request.Episode,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotMessage(ctx, gin.H{"id": res.Id})
}

func Update(ctx *gin.Context) {
	paramDramaId := ctx.Param("id")
	dramaId, _ := strconv.Atoi(paramDramaId)
	paramVideoId := ctx.Param("videoId")
	videoId, _ := strconv.Atoi(paramVideoId)

	request := validate.VideoUpdate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	pr := proto.UpdateVideoRequest{
		Id:              int64(videoId),
		FileId:          request.FileId,
		CategoryId:      request.CategoryId,
		UserId:          0,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
		IsVisible:       true,
		DramaId:         int64(dramaId),
		Episode:         request.Episode,
	}

	if _, err := global.VideoServerClient.Update(context.Background(), &pr); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Delete(ctx *gin.Context) {
	paramDramaId := ctx.Param("id")
	dramaId, _ := strconv.Atoi(paramDramaId)
	paramVideoId := ctx.Param("videoId")
	videoId, _ := strconv.Atoi(paramVideoId)

	if _, err := global.VideoServerClient.Delete(context.Background(), &proto.UpdateVideoRequest{DramaId: int64(dramaId), Id: int64(videoId)}); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotContent(ctx)
}

func Detail(ctx *gin.Context) {
	paramDramaId := ctx.Param("id")
	dramaId, _ := strconv.Atoi(paramDramaId)
	res, err := global.DramaServerClient.Detail(context.Background(), &proto.SearchDramaRequest{
		Id: int64(dramaId),
	})

	api.SuccessNotContent(ctx)
}

func List(ctx *gin.Context) {
	request := validate.VideoList{}
	if err := ctx.BindQuery(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	client := global.VideoServerClient
	rsp, err := client.Get(context.Background(), HandleSearchRequest(&request))

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	fmt.Printf("%+v\n", request)
	ctx.JSON(http.StatusOK, rsp)
}

func HandleSearchRequest(request *validate.VideoList) *proto.SearchVideoRequest {
	r := &proto.SearchVideoRequest{}

	if request.UserId > 0 {
		r.UserId = request.UserId
	}

	if request.RegionId > 0 {
		r.RegionId = request.RegionId
	}

	if request.CategoryId > 0 {
		r.CategoryId = request.CategoryId
	}

	if request.Keyword != "" {
		r.Keyword = request.Keyword
	}

	if request.Sort != "" {
		r.Sort = request.Sort
	}

	if &request.IsRecommend != nil {
		r.IsRecommend = request.IsRecommend
	}

	if &request.IsNew != nil {
		r.IsNew = request.IsNew
	}

	if &request.IsHot != nil {
		r.IsHot = request.IsHot
	}

	if &request.IsEnd != nil {
		r.IsEnd = request.IsEnd
	}

	if &request.IsVisible != nil {
		r.IsVisible = request.IsVisible
	}

	if request.TotalCountMin > 0 {
		r.TotalCountMin = request.TotalCountMin
	}

	if request.TotalCountMax > 0 {
		r.TotalCountMax = request.TotalCountMax
	}

	if request.FavoriteCountMin > 0 {
		r.FavoriteCountMin = request.FavoriteCountMin
	}

	if request.FavoriteCountMax > 0 {
		r.FavoriteCountMax = request.FavoriteCountMax
	}

	if request.LikeCountMin > 0 {
		r.LikeCountMin = request.LikeCountMin
	}

	if request.LikeCountMax > 0 {
		r.LikeCountMax = request.LikeCountMax
	}

	if request.PlayCountMin > 0 {
		r.PlayCountMin = request.PlayCountMin
	}

	if request.PlayCountMax > 0 {
		r.PlayCountMax = request.PlayCountMax
	}

	if request.BarrageCountMin > 0 {
		r.BarrageCountMin = request.BarrageCountMin
	}

	if request.BarrageCountMax > 0 {
		r.BarrageCountMax = request.BarrageCountMax
	}

	if request.Page > 0 {
		r.Page = request.Page
	}

	if request.PlayCountMax > 0 {
		r.PerPage = request.PerPage
	}

	return r
}
