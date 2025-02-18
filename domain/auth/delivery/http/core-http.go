package http

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"svc-portofolio-golang/utils/config"
	"svc-portofolio-golang/utils/message"
	"svc-portofolio-golang/valueobject"
)

func (handler *HttpAuthHandler) GetAll(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"param":     ctx.Query("param"),
			"param_nil": nil,
			"IN": map[string][]string{
				"column_in": strings.Split(ctx.Query("column_in"), ","),
			},
			"NOT": map[string]interface{}{
				"IN": map[string][]string{
					"column_not_in": strings.Split(ctx.Query("column_not_in"), ","),
				},
				"column_not":     ctx.Query("param_not"),
				"column_not_nil": nil,
			},
			"BETWEEN": map[string][]interface{}{
				"column_between": {"A", "B"},
			},
			"LIKE": map[string]interface{}{
				"column_like": "%" + ctx.Query("param_like") + "%",
			},
			"OR": []map[string]interface{}{
				{
					"column_one": ctx.Query("param"),
					"column_two": ctx.Query("nim_no_tagihan"),
				},
			},
			"CASE WHEN somecolumn IS NULL THEN leftjoin.column IN (?,?) ELSE TRUE END": []interface{}{"query", 234},
		},
		"GROUP": map[string]interface{}{
			"GROUP_BY": []string{"column", "column"},
			"HAVING": [][]interface{}{
				{"column", ">", "unsigned int"},
				{"column", "<", "unsigned int"},
				{"column", "=", "int"},
			},
		},
		"ORDER": map[string]interface{}{
			"ORDER_BY": []string{"column asc", "column desc"},
		},
		"LIMIT": []interface{}{ctx.Query("offset"), ctx.Query("limit")},
	}

	response, err := handler.authUsecase.GetAll(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	count, err := handler.authUsecase.Count(param)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnPagination(ctx, &response, param, count.Count)
}

func (handler *HttpAuthHandler) GetByUUID(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"uuid": ctx.Param("uuid"),
		},
	}

	response, err := handler.authUsecase.GetOne(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnOk(ctx, &response, param)
}

func (handler *HttpAuthHandler) Store(ctx *gin.Context) {
	var payload valueobject.AuthPayloadInsert

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	payload.User = ctx.Request.Header.Get("X-Member")

	feedback, err := handler.authUsecase.Store(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessInsert(ctx, feedback.Data)
}

func (handler *HttpAuthHandler) Update(ctx *gin.Context) {
	var payload valueobject.AuthPayloadUpdate

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	payload.User = ctx.Request.Header.Get("X-Member")

	err = handler.authUsecase.Update(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessUpdate(ctx)
}

func (handler *HttpAuthHandler) Delete(ctx *gin.Context) {
	var payload valueobject.AuthPayloadDelete

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	err = handler.authUsecase.Delete(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessDelete(ctx)
}
