package controller

import (
	"net/http"
	"strconv"

	"github.com/Fuuma0000/manetabi_api/model"
	"github.com/Fuuma0000/manetabi_api/usecase"
	"github.com/labstack/echo"
)

type IPlanController interface {
	CreatePlan(c echo.Context) error
	GetPlansByUserID(c echo.Context) error
	GetPlanByID(c echo.Context) error
	UpdatePlan(c echo.Context) error
	DeletePlan(c echo.Context) error
}

type planController struct {
	pu usecase.IPlanUsecase
}

func NewPlanController(pu usecase.IPlanUsecase) IPlanController {
	return &planController{pu}
}

func (pc *planController) CreatePlan(c echo.Context) error {
	userID := c.Get("userID")
	plan := model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	plan.UserID = userID.(uint)
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

func (pc *planController) UpdatePlan(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	userID := c.Get("userID")
	plan := model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	plan.PlanID = uint(id)
	plan.UserID = userID.(uint)
	resPlan, err := pc.pu.UpdatePlan(plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resPlan)
}

func (pc *planController) DeletePlan(c echo.Context) error {
	PlanIdStr := c.Param("id")
	planId, err := strconv.Atoi(PlanIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	userId := c.Get("userID").(uint)
	if err := pc.pu.DeletePlan(planId, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deleted")
}
