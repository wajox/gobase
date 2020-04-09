package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/rs/zerolog/log"
)

const (
	ContentTypeHeader = "Content-Type"
)

func EmptyOK(ctx *gin.Context) {
	Empty(ctx, http.StatusOK)
}

func Empty(ctx *gin.Context, statusCode int) {
	ctx.AbortWithStatus(statusCode)
}

func NoContent(ctx *gin.Context) {
	Empty(ctx, http.StatusNoContent)
}

func JSONAPIPayload(ctx *gin.Context, statusCode int, payload interface{}) {
	ctx.Header(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)

	if err := jsonapi.MarshalPayload(ctx.Writer, payload); err != nil {
		log.Error().Err(err).Msg("jsonapi.MarshalPayload failed")

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// JSONAPIPayloader - renders internal already marshaled jsonapi struct
func JSONAPIPayloader(ctx *gin.Context, statusCode int, payload jsonapi.Payloader) {
	ctx.Header(ContentTypeHeader, jsonapi.MediaType)

	ctx.JSON(statusCode, payload)
}

func Errors(ctx *gin.Context, statusCode int, errs []*jsonapi.ErrorObject) {
	ctx.Header(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)

	if err := jsonapi.MarshalErrors(ctx.Writer, errs); err != nil {
		log.Error().Err(err).Msg("jsonapi.MarshalErrors failed")

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
