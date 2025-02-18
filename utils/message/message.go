package message

import (
	"net/http"
	"reflect"
	"svc-portofolio-golang/utils/slices"

	"github.com/gin-gonic/gin"
)

const PUBLIC = "public"
const ERROR_STATUS = "Layanan sedang mengalami gangguan"
const SUCCESS_UPDATE_STATUS = "Berhasil melakukan pembaharuan data"

func measureLength(data interface{}) int {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Slice:
		return v.Len()
	case reflect.Struct:
		return 1
	}
	return 0
}

func castNilSliceToArray(data interface{}) interface{} {
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		if reflect.ValueOf(data).Len() < 1 {
			return []interface{}{}
		}
	}
	return data
}

func ReturnOk(ctx *gin.Context, data interface{}, param interface{}) {
	requestSource, exists := ctx.Get("requestSource")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Request source not found"})
		return
	}
	if requestSource == PUBLIC {
		slices.RemoveIDFields(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"param": param,
		"count": measureLength(data),
		"data":  castNilSliceToArray(data),
	})
}

func ReturnPagination(ctx *gin.Context, data interface{}, param interface{}, count int) {
	requestSource, exists := ctx.Get("requestSource")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Request source not found"})
		return
	}
	if requestSource == PUBLIC {
		slices.RemoveIDFields(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"param": param,
		"count": count,
		"data":  castNilSliceToArray(data),
	})
}

func ReturnBadRequest(ctx *gin.Context, message interface{}, param interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"param":   param,
	})
}

func ReturnInternalServerError(ctx *gin.Context, msg ...string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message":               ERROR_STATUS,
		"message for developer": msg,
	})
}

func ReturnSuccessInsert(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func ReturnSuccessDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func ReturnSuccessUpdate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": SUCCESS_UPDATE_STATUS,
	})
}
