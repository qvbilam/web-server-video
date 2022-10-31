package episode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"video/api"
	"video/validate"
)

func Create(ctx *gin.Context) {
	request := validate.VideoCreate{}
	if err := ctx.Bind(request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	fmt.Printf("%+v\n", request)

	//global.VideoServerClient.Create(context.Background(), &proto.UpdateVideoRequest{
	//	FileId:         0,
	//	CategoryId:     0,
	//	UserId:         0,
	//	Name:           "",
	//	Introduce:      "",
	//	Url:            "",
	//	Icon:           "",
	//	HorizontalIcon: "",
	//	IsVisible:      false,
	//	DramaId:        0,
	//	Episode:        0,
	//})
}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}

func Detail(ctx *gin.Context) {

}

func List(ctx *gin.Context) {

}
