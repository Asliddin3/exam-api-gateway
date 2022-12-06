package postgres

import (
	model "github.com/Asliddin3/exam-api-gateway/api/models"
	"github.com/jmoiron/sqlx"
)

type adminRepo struct {
	db *sqlx.DB
}

func NewAdminRepo(db *sqlx.DB) *adminRepo {
	return &adminRepo{db: db}
}

func (r *adminRepo) LoginAdmin(req *model.AdminRequest) (*model.AdminRequest, error) {
	adminResp := model.AdminRequest{}
	err := r.db.QueryRow(`
	select password from admin where username=$1
	`, req.UserName).Scan(&adminResp.PassWord)
	if err != nil {
		return &model.AdminRequest{}, err
	}
	adminResp.UserName = req.UserName
	return &adminResp, nil
}

func (r *adminRepo) LoginModerator(req *model.ModeratorRequest) (*model.ModeratorRequest, error) {
	moderatorResp := model.ModeratorRequest{}
	err := r.db.QueryRow(`
	select id,password from moderator where username=$1
	`, req.UserName).Scan(&moderatorResp.Id, &moderatorResp.PassWord)
	if err != nil {
		return &model.ModeratorRequest{}, err
	}
	moderatorResp.UserName = req.UserName
	return &moderatorResp, nil
}
