package mysql

import (
	"go-import-manage/internal/models"
)

func InsertRecord(record models.Record) error {
	query := `INSERT INTO records (
		first_name, last_name, company_name, address, city, county, postal, phone, email, web
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(query,
		record.FirstName, record.LastName, record.CompanyName, record.Address,
		record.City, record.County, record.Postal, record.Phone,
		record.Email, record.Web)
	return err
}

func GetRecords() ([]models.Record, error) {
	query := `SELECT 
		id, first_name, last_name, company_name, address, city, county, postal, phone, email, web 
		FROM records`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record
	for rows.Next() {
		var record models.Record
		if err := rows.Scan(&record.ID, &record.FirstName, &record.LastName, &record.CompanyName,
			&record.Address, &record.City, &record.County, &record.Postal,
			&record.Phone, &record.Email, &record.Web); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func UpdateRecord(id int, record models.Record) error {
	query := `UPDATE records SET 
		first_name=?, last_name=?, company_name=?, address=?, city=?, county=?, 
		postal=?, phone=?, email=?, web=? 
		WHERE id=?`
	_, err := DB.Exec(query,
		record.FirstName, record.LastName, record.CompanyName, record.Address,
		record.City, record.County, record.Postal, record.Phone,
		record.Email, record.Web, id)
	return err
}

func DeleteRecord(id int) error {
	query := `DELETE FROM records WHERE id=?`
	_, err := DB.Exec(query, id)
	return err
}
