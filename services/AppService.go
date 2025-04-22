package services

import (
	"book-store-server/internal/storage/DatabaseService"
)

type AppService struct {
	DBService *DatabaseService.DBService
}

func NewAppService(dbService *DatabaseService.DBService) *AppService {
	return &AppService{
		DBService: dbService,
	}
}
