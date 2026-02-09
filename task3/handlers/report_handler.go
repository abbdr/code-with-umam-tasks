package handlers

import (
	"encoding/json"
	"net/http"

	"kasir-api/services"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleTodayReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		report, err := h.service.TodayReport()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(report)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) HandleRangeReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		start := r.URL.Query().Get("start_date")
		end := r.URL.Query().Get("end_date")
		report, err := h.service.RangeReport(start, end)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(report)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}