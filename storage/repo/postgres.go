package repo

import (
	"github.com/Asliddin3/exam-api-gateway/api/models"
)

type AdminRepo interface {
	LoginAdmin(*models.AdminRequest) (*models.AdminRequest, error)
}
