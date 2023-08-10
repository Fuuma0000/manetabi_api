package infrastructure

import (
	"database/sql"

	"github.com/Fuuma0000/manetabi_api/model"
)

type IPlanInfrastructer interface {
	CreatePlan(plan *model.Plan) error
	GetPlansByUserID(plans *[]model.Plan, userId uint) error
	GetPlanByID(id int) (model.Plan, error)
}

type planInfrastructer struct {
	db *sql.DB
}

func NewPlanInfrastructer(db *sql.DB) IPlanInfrastructer {
	return &planInfrastructer{db}
}

func (pi *planInfrastructer) CreatePlan(plan *model.Plan) error {
	q := `INSERT INTO plans (user_id, title, description, thumbnail_path, cost, start_date, end_date, is_public) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := pi.db.Exec(q, plan.UserID, plan.Title, plan.Description, plan.Thumbnail, plan.Cost, plan.StartDate, plan.EndDate, plan.IsPublic)
	if err != nil {
		return err
	}

	// 最後に生成されたIDを取得
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	plan.PlanID = uint(lastInsertID)

	return nil
}

func (pi *planInfrastructer) GetPlansByUserID(plans *[]model.Plan, userId uint) error {
	q := `SELECT * FROM plans WHERE user_id = ? & is_public = 1`
	rows, err := pi.db.Query(q, userId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var plan model.Plan
		if err := rows.Scan(&plan.PlanID, &plan.UserID, &plan.Title, &plan.Description, &plan.Thumbnail, &plan.Cost, &plan.StartDate, &plan.EndDate, &plan.IsPublic, &plan.CreatedAt, &plan.UpdatedAt); err != nil {
			return err
		}
		*plans = append(*plans, plan)
	}
	return nil
}

func (pi *planInfrastructer) GetPlanByID(id int) (model.Plan, error) {
	plan := model.Plan{}
	q := `SELECT * FROM plans WHERE plan_id = ? LIMIT 1`
	err := pi.db.QueryRow(q, id).Scan(&plan.PlanID, &plan.UserID, &plan.Title, &plan.Description, &plan.Thumbnail, &plan.Cost, &plan.StartDate, &plan.EndDate, &plan.IsPublic, &plan.CreatedAt, &plan.UpdatedAt)
	if err != nil {
		return model.Plan{}, err
	}
	return plan, nil
}
