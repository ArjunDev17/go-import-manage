package mysql

import (
	"go-import-manage/internal/models"
)

func InsertRecord(record models.Record) error {
	query := `INSERT INTO records (name, email, phone) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, record.Name, record.Email, record.Phone)
	return err
}

func GetRecords() ([]models.Record, error) {
	query := `SELECT id, name, email, phone FROM records`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record
	for rows.Next() {
		var record models.Record
		if err := rows.Scan(&record.ID, &record.Name, &record.Email, &record.Phone); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func UpdateRecord(id int, record models.Record) error {
	query := `UPDATE records SET name=?, email=?, phone=? WHERE id=?`
	_, err := DB.Exec(query, record.Name, record.Email, record.Phone, id)
	return err
}

func DeleteRecord(id int) error {
	query := `DELETE FROM records WHERE id=?`
	_, err := DB.Exec(query, id)
	return err
}
