package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) TodayReport() (*models.Report, error) {
	return s.repo.TodayReport()
}

func (s *ReportService) RangeReport(start, end string) (*models.Report, error) {
	return s.repo.RangeReport(start, end)
}