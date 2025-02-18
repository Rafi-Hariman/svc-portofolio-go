package http

import (
	"github.com/gin-gonic/gin"

	"svc-portofolio-golang/domain/boilerplate"
	"svc-portofolio-golang/utils/middleware"
)

type HttpBoilerplateHandler struct {
	boilerplateUsecase boilerplate.Usecase
}

func NewBoilerplateHttpHandler(boilerplate boilerplate.Usecase, httpRouter *gin.Engine) {
	handler := &HttpBoilerplateHandler{
		boilerplateUsecase: boilerplate,
	}

	public := httpRouter.Group("/public/api/v1")
	public.Use(middleware.PublicMiddleware)
	public.GET("/boilerplate", handler.GetAll)
	public.GET("/boilerplate/:uuid", handler.GetByUUID)
	public.POST("/boilerplate", handler.Store)
	public.PUT("/boilerplate", handler.Update)
	public.DELETE("/boilerplate", handler.Delete)

	private := httpRouter.Group("/private/api/v1")
	private.Use(middleware.PrivateMiddleware)
	private.GET("/boilerplate", handler.GetAll)
	private.GET("/boilerplate/:uuid", handler.GetByUUID)
	private.POST("/boilerplate", handler.Store)
	private.PUT("/boilerplate", handler.Update)
	private.DELETE("/boilerplate", handler.Delete)
}
