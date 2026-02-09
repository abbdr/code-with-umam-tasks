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

	rows, err := tx.Query("SELECT a.transaction_id, a.created_at, a.total_amount, b.transaction_detail_id, b.quantity, b.subtotal, c.product_name, b.product_id, b.transaction_id FROM transactions AS a INNER JOIN transaction_details AS b ON a.transaction_id = b.transaction_id INNER JOIN products AS c ON b.product_id = c.product_id WHERE DATE(a.created_at) = $1", today)
	if err != nil {
		return nil, err
	}
	
	var t models.Transaction
	var d []models.TransactionDetail
	var t_temp models.Transaction
	var d_temp models.TransactionDetail
	// i := 0
	for rows.Next() {

		if err := rows.Scan(&t_temp.ID, &t_temp.CreatedAt, &t_temp.TotalAmount, &d_temp.ID, &d_temp.Quantity, &d_temp.Subtotal, &d_temp.ProductName, &d_temp.ProductID, &d_temp.TransactionID); err != nil {
			return nil, err
		}

		// fmt.Println(d_temp.TransactionID, t.ID)
		// fmt.Println()
		// fmt.Println(d_temp.TransactionID != t.ID, t.ID != 0)
		if d_temp.TransactionID != t.ID && t.ID != 0 {
			// fmt.Println("report added")
			report.TransactionReport = append(report.TransactionReport, t)
			t = models.Transaction{}
			d = []models.TransactionDetail{}
		}

		if t.ID != t_temp.ID {
			t.ID = t_temp.ID
			t.CreatedAt = t_temp.CreatedAt
			t.TotalAmount = t_temp.TotalAmount
		}
		d = append(d, d_temp)
		t.Details = d

		// i += 1
		// fmt.Println(i)
		// fmt.Println(report, "\n")
	}
	report.TransactionReport = append(report.TransactionReport, t)

	// fmt.Println(report)

	return &report, err
}

func (repo *ReportRepository) RangeReport(start, end string) (*models.Report, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var report models.Report

	rows, err := tx.Query("SELECT a.transaction_id, a.created_at, a.total_amount, b.transaction_detail_id, b.quantity, b.subtotal, c.product_name, b.product_id, b.transaction_id FROM transactions AS a INNER JOIN transaction_details AS b ON a.transaction_id = b.transaction_id INNER JOIN products AS c ON b.product_id = c.product_id WHERE a.created_at::DATE BETWEEN $1::DATE AND $2::DATE", start, end)
	if err != nil {
		return nil, err
	}
	
	var t models.Transaction
	var d []models.TransactionDetail
	var t_temp models.Transaction
	var d_temp models.TransactionDetail
	// i := 0
	for rows.Next() {

		if err := rows.Scan(&t_temp.ID, &t_temp.CreatedAt, &t_temp.TotalAmount, &d_temp.ID, &d_temp.Quantity, &d_temp.Subtotal, &d_temp.ProductName, &d_temp.ProductID, &d_temp.TransactionID); err != nil {
			return nil, err
		}

		// fmt.Println(d_temp.TransactionID, t.ID)
		// fmt.Println()
		// fmt.Println(d_temp.TransactionID != t.ID, t.ID != 0)
		if d_temp.TransactionID != t.ID && t.ID != 0 {
			// fmt.Println("report added")
			report.TransactionReport = append(report.TransactionReport, t)
			t = models.Transaction{}
			d = []models.TransactionDetail{}
		}

		if t.ID != t_temp.ID {
			t.ID = t_temp.ID
			t.CreatedAt = t_temp.CreatedAt
			t.TotalAmount = t_temp.TotalAmount
		}
		d = append(d, d_temp)
		t.Details = d

		// i += 1
		// fmt.Println(i)
		// fmt.Println(report, "\n")
	}
	report.TransactionReport = append(report.TransactionReport, t)

	// fmt.Println(report)

	return &report, err
}


