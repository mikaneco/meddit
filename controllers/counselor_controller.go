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

type CounselorController struct {
	service services.CounselorService
}

func NewCounselorController(counselorService services.CounselorService) CounselorController {
	return CounselorController{
		service: counselorService,
	}
}

func (c CounselorController) RegisterCounselor(ctx *gin.Context) {
	// リクエストボディを取得する
	var counselorRegisterRequest dto.CounselorRegisterRequest

	if err := ctx.ShouldBindJSON(&counselorRegisterRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	counselor := models.Counselor{
		FirstName: counselorRegisterRequest.FirstName,
		LastName:  counselorRegisterRequest.LastName,
		Email:     counselorRegisterRequest.Email,
		Password:  counselorRegisterRequest.Password,
	}

	// counselorを作成する
	createdCounselor, err := c.service.RegisterCounselor(&counselor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// JWTトークンを発行する
	token, err := auth.IssueCounselorToken(counselor.ID, config.JWTSecretKey, config.JWTDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスを返す
	ctx.JSON(http.StatusCreated, gin.H{
		"counselor": createdCounselor,
		"token":     token,
	})
}

func (c *CounselorController) Login(ctx *gin.Context) {
	// リクエストボディを取得する
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// counselorを検索する
	counselor, err := c.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードをチェックする
	if !utils.CheckPasswordHash(loginRequest.Password, counselor.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	}

	// JWTトークンを発行する
	token, err := auth.IssueCounselorToken(counselor.ID, config.JWTSecretKey, config.JWTDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスを返す
	ctx.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
	})
}

func (c CounselorController) GetCounselor(ctx *gin.Context) {
	counselorID := ctx.Param("counselorID")
	id, err := strconv.ParseUint(counselorID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	counselor, err := c.service.GetCounselorByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, counselor)
}
