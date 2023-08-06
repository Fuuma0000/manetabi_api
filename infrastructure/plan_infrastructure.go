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
	q := `INSERT INTO plans (user_id, title, description, thumbnail_path, cost, start_date, end_date, is_public) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := pi.db.Exec(q, plan.UserID, plan.Title, plan.Description, plan.Thumbnail, plan.Cost, plan.StartDate, plan.EndDate, plan.IsPublic)
	if err != nil {
		return err
	}
	return nil
}
