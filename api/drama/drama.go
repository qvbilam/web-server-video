package drama

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"video/api"
	PageProto "video/api/qvbilam/page/v1"
	proto "video/api/qvbilam/video/v1"
	"video/global"
	"video/resource"
	"video/validate"
)

func List(ctx *gin.Context) {
	request := validate.DramaSearch{}
	if err := ctx.BindQuery(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	requestProto := searchRequestToProto(&request)
	response, err := global.DramaServerClient.Get(context.Background(), requestProto)
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	r := resource.DramaResource{}
	api.SuccessList(ctx, r.Collection(response), response.Total)
}

func Create(ctx *gin.Context) {
	request := validate.DramaCreate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	res, err := global.DramaServerClient.Create(context.Background(), &proto.UpdateDramaRequest{
		CategoryId:      request.CategoryId,
		RegionId:        request.RegionId,
		Name:            request.Name,
		Introduce:       request.Introduce,
		Cover:           request.Cover,
		HorizontalCover: request.HorizontalCover,
		TotalCount:      request.TotalCount,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotMessage(ctx, gin.H{"id": res.Id})
}

func Update(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	request := validate.DramaUpdate{}
	if err := ctx.Bind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	updateRequest := UpdateRequestToProto(&request)
	updateRequest.Id = int64(id)

	if _, err := global.DramaServerClient.Update(context.Background(), updateRequest); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func searchRequestToProto(search *validate.DramaSearch) *proto.SearchDramaRequest {
	r := &proto.SearchDramaRequest{}
	if search.Keyword != "" {
		r.Keyword = search.Keyword
	}
	if search.Sort != "" {
		r.Sort = search.Sort
	}
	if search.IsVisible != nil {
		r.IsVisible = *search.IsVisible
	}

	p := &PageProto.PageRequest{
		Page:    1,
		PerPage: 100,
	}
	if search.Page != 0 {
		p.Page = search.Page
	}
	if search.PerPage != 0 {
		p.PerPage = search.PerPage
	}
	r.Page = p
	return r
}

func UpdateRequestToProto(update *validate.DramaUpdate) *proto.UpdateDramaRequest {
	r := &proto.UpdateDramaRequest{}

	if update.CategoryId != nil {
		r.CategoryId = *update.CategoryId
	}

	if update.RegionId != nil {
		r.RegionId = *update.RegionId
	}

	if update.Name != nil {
		r.Name = *update.Name
	}

	if update.Introduce != nil {
		r.Introduce = *update.Introduce
	}
	if update.Cover != nil {
		r.Cover = *update.Cover
	}

	if update.HorizontalCover != nil {
		r.HorizontalCover = *update.HorizontalCover
	}

	if update.TotalCount != nil {
		r.TotalCount = *update.TotalCount
	}

	return r
}
