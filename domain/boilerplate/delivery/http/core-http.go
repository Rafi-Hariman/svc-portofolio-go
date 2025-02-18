package http

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"svc-portofolio-golang/utils/config"
	"svc-portofolio-golang/utils/message"
	"svc-portofolio-golang/valueobject"
)

func (handler *HttpBoilerplateHandler) GetAll(ctx *gin.Context) {
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

	response, err := handler.boilerplateUsecase.GetAll(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), param)
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	count, err := handler.boilerplateUsecase.Count(param)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnPagination(ctx, &response, param, count.Count)
}

func (handler *HttpBoilerplateHandler) GetByUUID(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"uuid": ctx.Param("uuid"),
		},
	}

	response, err := handler.boilerplateUsecase.GetOne(param)

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

func (handler *HttpBoilerplateHandler) Store(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadInsert

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	payload.User = ctx.Request.Header.Get("X-Member")

	feedback, err := handler.boilerplateUsecase.Store(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessInsert(ctx, feedback.Data)
}

func (handler *HttpBoilerplateHandler) Update(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadUpdate

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	payload.User = ctx.Request.Header.Get("X-Member")

	err = handler.boilerplateUsecase.Update(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessUpdate(ctx)
}

func (handler *HttpBoilerplateHandler) Delete(ctx *gin.Context) {
	var payload valueobject.BoilerplatePayloadDelete

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	err = handler.boilerplateUsecase.Delete(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessDelete(ctx)
}
