package http

import (
	"net/http"

	"github.com/JIeeiroSst/go-app/domain"
	"github.com/JIeeiroSst/go-app/log"
	"github.com/JIeeiroSst/go-app/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service domain.Service
}

var (
	empty          string = "empty"
	create_message string = "create success"
	delete_message string = "delete success"
	update_message string = "update success"
	create_message_err string = "create failed"
	delete_message_err string = "delete failed"
	update_message_err string = "update failed"
)

func NewHandler(e *echo.Echo,service domain.Service) {
	hander:=&Handler{
		service: service,
	}
	u:=e.Group("/user")
	u.GET("/",hander.UserAll)
	u.GET("/:id",hander.UserById)
	u.POST("",hander.CreateUser)
	u.PUT("/:id",hander.UpdateUser)
	u.DELETE("/:id",hander.DeleteUser)

	p:=e.Group("/profile")
	p.GET("/",hander.ProfileAll)
	p.GET("/:id",hander.ProfileById)
	p.POST("",hander.CreateEmail)
	p.PUT("/:id",hander.UpdateEmail)
	p.DELETE("/:id",hander.DeleteEmail)
}

func(h *Handler) UserAll(e echo.Context) error {
	users,err:=h.service.UserAll()
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,users)
}

func (h *Handler) UserById(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	user,err:=h.service.UserById(id)
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,user)
}

func (h *Handler) CreateUser(e echo.Context) error {
	user:=domain.User{
		Name: e.FormValue("name"),
		Address: e.FormValue("address"),
	}
	err:=h.service.CreateUser(user)
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,create_message_err)
	}
	log.InitZapLog().Info("server running")
	return e.JSON(http.StatusOK,create_message)
}

func (h *Handler) UpdateUser(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	user:=domain.User{
		Name: e.FormValue("name"),
		Address: e.FormValue("address"),
	}
	err:=h.service.UpdateUser(id,user)
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,update_message_err)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,update_message)
}

func (h *Handler) DeleteUser(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	err:=h.service.DeleteUser(id)
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,delete_message_err)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,delete_message)
}

func (h *Handler) ProfileAll(e echo.Context) error {
	profiles,err:=h.service.ProfileAll()
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,profiles)
}

func (h *Handler) ProfileById(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	profile,err:=h.service.ProfileById(id)
	if err!=nil{
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,profile)
}

func (h *Handler) UpdateEmail(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	ctx := e.Request().Context()
	profile:=domain.Profile{
		Email: e.FormValue("email"),
	}
	ok,message:=h.service.UpdateEmail(ctx,id,profile)
	if !ok {
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,message)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,message)
}

func (h *Handler) CreateEmail(e echo.Context) error {
	profile:=domain.Profile{
		Name: e.FormValue("name"),
		Email: e.FormValue("email"),
		UserId: utils.StringToInt(e.Param("user_id")),
	}
	ctx := e.Request().Context()
	ok,message:=h.service.CreateEmail(ctx,profile)
	if !ok {
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,message)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,message)
}

func (h *Handler) DeleteEmail(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	ctx := e.Request().Context()
	ok,message:= h.service.DeleteEmail(ctx,id)
	if !ok {
		log.InitZapLog().Error("server running error")
		return e.String(http.StatusInternalServerError,message)
	}
	log.InitZapLog().Error("server running")
	return e.JSON(http.StatusOK,message)
}