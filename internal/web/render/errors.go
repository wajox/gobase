package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
)

// NotFoundError renders Not Found error
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

// InternalServerError renders Internal Server error
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

// UnprocessableEntity renders Unprocessable Entity error
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

// Unauthorized renders Unauthorized error
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

// BadRequest renders Bad Request error
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

// ValidationErrors renders validation errors with Unprocessable Entity status
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

// ConflictError renders Conflict error
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
