package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fuuma0000/manetabi_api/model"
	"github.com/Fuuma0000/manetabi_api/usecase"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type IPlanController interface {
	CreatePlan(c echo.Context) error
	GetPlansByUserID(c echo.Context) error
	GetPlanByID(c echo.Context) error
}

type planController struct {
	pu usecase.IPlanUsecase
}

func NewPlanController(pu usecase.IPlanUsecase) IPlanController {
	return &planController{pu}
}

func (pc *planController) CreatePlan(c echo.Context) error {
	fmt.Println("CreatePlan")
	fmt.Println(c)
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println("claims")
	fmt.Println(claims)
	userId := claims["user_id"]
	fmt.Println("userId")
	fmt.Println(userId)
	plan := model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	plan.UserID = uint(userId.(float64))
	fmt.Println(plan.UserID)
	resPlan, err := pc.pu.CreatePlan(plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resPlan)
}

func (pc *planController) GetPlansByUserID(c echo.Context) error {
	userIDStr := c.QueryParam("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid userID")
	}

	resPlans, err := pc.pu.GetPlansByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resPlans)
}

func (pc *planController) GetPlanByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	resPlan, err := pc.pu.GetPlanByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resPlan)
}
