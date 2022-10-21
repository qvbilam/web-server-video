package drama

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"video/api"
	proto "video/api/qvbilam/video/v1"
	"video/global"
	"video/validate"
)

func List(ctx *gin.Context) {
	request := validate.DramaSearch{}
	if err := ctx.BindQuery(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	//global.VideoServerCl(context.Background(), &proto.Dr)
	requestProto := searchRequestToProto(&request)
	response, err := global.DramaServerClient.Get(context.Background(), requestProto)
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total": response.Total,
		"data":  response.Drama,
	})
}

func searchRequestToProto(search *validate.DramaSearch) *proto.SearchDramaRequest {
	p := &proto.SearchDramaRequest{}
	if search.Keyword != "" {
		p.Keyword = search.Keyword
	}
	if search.Sort != "" {
		p.Sort = search.Sort
	}
	if search.IsVisible != nil {
		p.IsVisible = *search.IsVisible
	}
	if search.Page != 0 {
		p.Page.Page = search.Page
	}
	if search.PerPage != 0 {
		p.Page.PerPage = search.PerPage
	}
	return p
}
