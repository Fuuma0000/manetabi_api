package usecase

import (
	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/model"
)

type IPlanUsecase interface {
	CreatePlan(plan model.Plan) (model.PlanResponse, error)
}

type planUsecase struct {
	pi infrastructure.IPlanInfrastructer
}

func NewPlanUsecase(pi infrastructure.IPlanInfrastructer) IPlanUsecase {
	return &planUsecase{pi}
}

func (pu *planUsecase) CreatePlan(plan model.Plan) (model.PlanResponse, error) {
	if err := pu.pi.CreatePlan(&plan); err != nil {
		return model.PlanResponse{}, err
	}
	resPlan := model.PlanResponse{
		ID:          plan.ID,
		Title:       plan.Title,
		Description: plan.Description,
		Thumbnail:   plan.Thumbnail,
		Cost:        plan.Cost,
		StartDate:   plan.StartDate,
		EndDate:     plan.EndDate,
		IsPublic:    plan.IsPublic,
		CreatedAt:   plan.CreatedAt,
		UpdatedAt:   plan.UpdatedAt,
	}
	return resPlan, nil
}
