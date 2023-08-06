package controller

import (
	"net/http"

	"github.com/Fuuma0000/manetabi_api/model"
	"github.com/Fuuma0000/manetabi_api/usecase"
	"github.com/labstack/echo"
)

type IPlanController interface {
	CreatePlan(c echo.Context) error
}

type planController struct {
	pu usecase.IPlanUsecase
}

func NewPlanController(pu usecase.IPlanUsecase) IPlanController {
	return &planController{pu}
}

func (pc *planController) CreatePlan(c echo.Context) error {
	plan := model.Plan{}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	resPlan, err := pc.pu.CreatePlan(plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resPlan)
}
