package http

import (
	"github.com/gin-gonic/gin"

	"svc-portofolio-golang/domain/portofolio"
	"svc-portofolio-golang/utils/middleware"
)

type HttpPortofolioHandler struct {
	portofolioUsecase portofolio.Usecase
}

func NewPortofolioHttpHandler(portofolio portofolio.Usecase, httpRouter *gin.Engine) {
	handler := &HttpPortofolioHandler{
		portofolioUsecase: portofolio,
	}

	public := httpRouter.Group("/public/api/v1")
	public.Use(middleware.PublicMiddleware)
	public.GET("/portofolio", handler.GetAll)
	public.GET("/portofolio/:uuid", handler.GetByUUID)
	public.POST("/portofolio", handler.Store)
	public.PUT("/portofolio", handler.Update)
	public.DELETE("/portofolio", handler.Delete)

	private := httpRouter.Group("/private/api/v1")
	private.Use(middleware.PrivateMiddleware)
	private.GET("/portofolio", handler.GetAll)
	private.GET("/portofolio/:uuid", handler.GetByUUID)
	private.POST("/portofolio", handler.Store)
	private.PUT("/portofolio", handler.Update)
	private.DELETE("/portofolio", handler.Delete)
}
