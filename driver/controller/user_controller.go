package controller

import (
	"errors"
	"net/http"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userUC         usecase.UserUC
	rg             *gin.RouterGroup
	middlewareAuth middleware.JwtMiddleware
}

func (c *userController) updateUser(ctx *gin.Context) {
	var userDto dto.UserUpdateDto

	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	userDto.Id = ctx.MustGet("userId").(string)

	user, err := c.userUC.UpdateUser(userDto)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, user)
}

func (c *userController) findUserId(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, err := c.userUC.FindById(userId)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, user)
}

func (c *userController) findUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := c.userUC.FindUserByEmail(email)

	if err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, user)
}

func (c *userController) changePaswordUser(ctx *gin.Context) {
	var userDto dto.UserChangePassDto

	id := ctx.MustGet("userId").(string)

	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	user, err := c.userUC.ChangePaswordUser(userDto.Password, id)
	if err != nil {
		switch err.(type) {
		case common.InvalidError:
			common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
			return
		default:
			common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, user)
}

func (u *userController) deleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	_, err := u.userUC.FindById(userId)

	if err != nil {
		common.ResponseError(ctx, "can't delete users! id not found", http.StatusBadRequest)
		return
	}

	err = u.userUC.DeleteUserById(userId)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "successfully delete users", http.StatusOK, nil)
}

func (c *userController) RouterUser() {
	v2 := c.rg.Group(config.UserUpdatePath)
	v2.PUT("", c.middlewareAuth.AuthMiddleware("Admin", "Participant", "Trainer"), c.updateUser)
	v2.GET(config.UserGEtByEmail, c.middlewareAuth.AuthMiddleware("Admin", "Trainer"), c.findUserByEmail)
	v2.GET(config.UserGetById, c.middlewareAuth.AuthMiddleware("Admin", "Trainer"), c.findUserId)
	v2.PUT(config.UserChangePassword, c.middlewareAuth.AuthMiddleware("Admin", "Trainer", "Participant"), c.changePaswordUser)
	v2.DELETE(config.UserDeleteId, c.middlewareAuth.AuthMiddleware("Admin"), c.deleteUser)
}

func NewUserController(userUC usecase.UserUC, rg *gin.RouterGroup, middlewareAuth middleware.JwtMiddleware) *userController {
	return &userController{
		userUC:         userUC,
		rg:             rg,
		middlewareAuth: middlewareAuth,
	}
}
