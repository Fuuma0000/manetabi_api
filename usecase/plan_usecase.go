package usecase

import (
	"github.com/Fuuma0000/manetabi_api/infrastructure"
	"github.com/Fuuma0000/manetabi_api/model"
)

type IPlanUsecase interface {
	CreatePlan(plan model.Plan) (model.PlanResponse, error)
	GetPlansByUserID(userId uint) ([]model.PlanResponse, error)
	GetPlanByID(id int) (model.PlanResponse, error)
	UpdatePlan(plan model.Plan) (model.PlanResponse, error)
	DeletePlan(id int, userId uint) error
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
		PlanID:      plan.PlanID,
		UserID:      plan.UserID,
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

func (pu *planUsecase) GetPlansByUserID(userId uint) ([]model.PlanResponse, error) {
	plans := []model.Plan{}
	if err := pu.pi.GetPlansByUserID(&plans, userId); err != nil {
		return []model.PlanResponse{}, err
	}
	resPlans := []model.PlanResponse{}
	for _, plan := range plans {
		resPlan := model.PlanResponse{
			PlanID:      plan.PlanID,
			UserID:      plan.UserID,
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
		resPlans = append(resPlans, resPlan)
	}
	return resPlans, nil
}

func (pu *planUsecase) GetPlanByID(id int) (model.PlanResponse, error) {
	plan, err := pu.pi.GetPlanByID(id)
	if err != nil {
		return model.PlanResponse{}, err
	}
	resPlan := model.PlanResponse{
		PlanID:      plan.PlanID,
		UserID:      plan.UserID,
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

func (pu *planUsecase) UpdatePlan(plan model.Plan) (model.PlanResponse, error) {
	if err := pu.pi.UpdatePlan(&plan); err != nil {
		return model.PlanResponse{}, err
	}
	resPlan := model.PlanResponse{
		PlanID:      plan.PlanID,
		UserID:      plan.UserID,
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

func (pu *planUsecase) DeletePlan(planId int, userId uint) error {
	if err := pu.pi.DeletePlan(planId, userId); err != nil {
		return err
	}
	return nil
}
