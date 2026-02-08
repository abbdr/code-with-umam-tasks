package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) TodayReport() (*models.Report, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	now := time.Now()
	today := now.Format("2006-01-02")

	var report models.Report

	rows, err := tx.Query("SELECT a.transaction_id, a.created_at, a.total_amount, b.transaction_detail_id, b.quantity, b.subtotal, c.product_name, b.product_id FROM transactions AS a INNER JOIN transaction_details AS b ON a.transaction_id = b.transaction_id INNER JOIN products AS c ON b.product_id = c.product_id WHERE DATE(a.created_at) = $1", today)
	if err != nil {
		return nil, err
	}
	// fmt.Println(report)
 
	for rows.Next() {
		var t models.Transaction
		var d models.TransactionDetail
		if err := rows.Scan(&t.ID, &t.CreatedAt, &t.TotalAmount, &d.ID, &d.Quantity, &d.Subtotal, &d.ProductName, &d.ProductID); err != nil {
			return nil, err
		}
		t.Details = append(t.Details, d)
		report.TransactionReport = append(report.TransactionReport, t)
	}

	// fmt.Println(report)

	return &report, err
}