package errors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AcceptError(ctx *gin.Context, path string) {
	type mesg struct {
		Message string `json:"message"`
	}

	var errMesg mesg
	if err := ctx.BindJSON(&errMesg); err != nil {
		log.Printf("error from `BindJSON` method, package `gin`: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	} else {
		log.Printf("ERROR : %s : %v\n", path, errMesg.Message)
		ctx.JSON(http.StatusForbidden, gin.H{"message": "error was accepted"})
	}
}

func ManageResponseStatusCodes(ctx *gin.Context) {
	code := ctx.Query("code")
	switch code {
	case "500":
		RenderError500Page(ctx)
	case "403":
		RenderError403Page(ctx)
	}
}

func RenderError403Page(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": "403: Forbidden"})
}

func RenderError500Page(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": "500: Internal Server Error"})
}
