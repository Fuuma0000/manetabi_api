package infrastructure

import (
	"database/sql"

	"github.com/Fuuma0000/manetabi_api/model"
)

type IPlanInfrastructer interface {
	CreatePlan(plan *model.Plan) error
}

type planInfrastructer struct {
	db *sql.DB
}

func NewPlanInfrastructer(db *sql.DB) IPlanInfrastructer {
	return &planInfrastructer{db}
}

func (pi *planInfrastructer) CreatePlan(plan *model.Plan) error {
	// TODO: plan作成処理
	return nil
}
