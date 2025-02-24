package http

import (
	"github.com/gin-gonic/gin"

	"svc-portofolio-golang/domain/auth"
	"svc-portofolio-golang/utils/middleware"
)

type HttpAuthHandler struct {
	authUsecase auth.Usecase
}

func NewAuthHttpHandler(auth auth.Usecase, httpRouter *gin.Engine) {
	handler := &HttpAuthHandler{
		authUsecase: auth,
	}

	public := httpRouter.Group("/public/api/v1")
	public.Use(middleware.PublicMiddleware)
	public.GET("/auth", handler.GetAll)
	public.GET("/auth/:uuid", handler.GetByUUID)
	public.POST("/auth", handler.Store)
	public.PUT("/auth", handler.Update)
	public.DELETE("/auth", handler.Delete)

	// new route

	public.GET("/auth/login-detail/:uuid", handler.GetOneUserLogin)
	public.POST("/auth/login", handler.StoreLogin)
	public.GET("/auth/login/list", handler.GetAllUserLogin)

	// new route

	private := httpRouter.Group("/private/api/v1")
	private.Use(middleware.PrivateMiddleware)
	private.GET("/auth", handler.GetAll)
	private.GET("/auth/:uuid", handler.GetByUUID)
	private.POST("/auth", handler.Store)
	private.PUT("/auth", handler.Update)
	private.DELETE("/auth", handler.Delete)
}
