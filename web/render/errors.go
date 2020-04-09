package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

func NotFoundError(ctx *gin.Context, detail string) {
	objects := []*jsonapi.ErrorObject{
		{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusNotFound),
			Detail: detail,
		},
	}

	Errors(ctx, http.StatusNotFound, objects)
}

func InternalServerError(ctx *gin.Context, detail string) {
	objects := []*jsonapi.ErrorObject{
		{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: detail,
		},
	}

	Errors(ctx, http.StatusInternalServerError, objects)
}

func UnprocessableEntity(ctx *gin.Context, detail string) {
	objects := []*jsonapi.ErrorObject{
		{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusUnprocessableEntity),
			Detail: detail,
		},
	}

	Errors(ctx, http.StatusUnprocessableEntity, objects)
}

func Unauthorized(ctx *gin.Context, detail string) {
	objects := []*jsonapi.ErrorObject{
		{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusUnauthorized),
			Detail: detail,
		},
	}

	Errors(ctx, http.StatusUnauthorized, objects)
}

func BadRequest(ctx *gin.Context, detail string) {
	objects := []*jsonapi.ErrorObject{
		{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: detail,
		},
	}

	Errors(ctx, http.StatusBadRequest, objects)
}

func ValidationErrors(ctx *gin.Context, errs []error) {
	objects := []*jsonapi.ErrorObject{}
	for i := range errs {
		o := &jsonapi.ErrorObject{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusUnprocessableEntity),
			Detail: errs[i].Error(),
		}
		objects = append(objects, o)
	}

	Errors(ctx, http.StatusUnprocessableEntity, objects)
}

func ConflictError(ctx *gin.Context, detail string) {
	objects := []*jsonapi.ErrorObject{
		{
			ID:     uuid.New().String(),
			Title:  http.StatusText(http.StatusConflict),
			Detail: detail,
		},
	}

	Errors(ctx, http.StatusConflict, objects)
}
