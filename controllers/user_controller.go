package controller

import (
	"meddit/auth"
	"meddit/config"
	"meddit/controllers/dto"
	"meddit/models"
	"meddit/services"
	"meddit/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserrService
}

func NewUserController(counselorService services.UserrService) UserController {
	return UserController{
		service: userv cService,
	}
}

func (c UserController) RegisterCounselor(ctx *gin.Context) {
	// リクエストボディを取得する
	var userrRegisterRequest dto.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&userrRegisterRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: userrRegisterRequest.Username,
		Email:     userrRegisterRequest.Email,
		Password:  userrRegisterRequest.Password,
	}

	// userを作成する
	createdUser, err := c.service.RegisterUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// JWTトークンを発行する
	token, err := auth.IssueUserToken(createdUser.ID, config.JWTSecretKey, config.JWTDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスを返す
	ctx.JSON(http.StatusCreated, gin.H{
		"user": createdUser,
		"token":     token,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	// リクエストボディを取得する
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// counselorを検索する
	user, err := c.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードをチェックする
	if !utils.CheckPasswordHash(loginRequest.Password, counselor.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}

	// JWTトークンを発行する
	token, err := auth.IssueUserToken(user.ID, config.JWTSecretKey, config.JWTDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスを返す
	ctx.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
	})
}

func (c UserController) GetUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	id, err := strconv.ParseUint(counselorID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.GetUserByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
