package postgres

import (
	"context"
	"errors"

	"github.com/techerfan/2DCH7-20059/entity"
	"gorm.io/gorm"
)

func (p *PostgresDB) AddTable(ctx context.Context, table entity.Table) (entity.Table, error) {
	model := mapTableEntitytoTable(table)

	if err := p.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.Table{}, err
	}

	return mapTableToTableEntity(model), nil
}

func (p *PostgresDB) RemoveTableByID(ctx context.Context, id uint) error {
	// Start a transaction
	tx := p.db.WithContext(ctx).Begin()

	// Fetch the table
	var table Table
	if err := tx.Preload("Reservations").First(&table).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Reservations related to this table must be deleted too
	for _, reservation := range table.Reservations {
		if err := tx.Unscoped().Delete(&reservation).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Delete the table itself
	if err := tx.Unscoped().Delete(&table).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (p *PostgresDB) FindTables(ctx context.Context) ([]entity.Table, error) {
	var tables []Table

	if err := p.db.WithContext(ctx).Find(&tables).Error; err != nil {
		return nil, err
	}

	resp := make([]entity.Table, 0, len(tables))
	for _, table := range tables {
		resp = append(resp, mapTableToTableEntity(table))
	}

	return resp, nil
}

func (p *PostgresDB) FindTableByID(ctx context.Context, id uint) (entity.Table, error) {
	var table Table
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&table).Error; err != nil {
		return entity.Table{}, err
	}

	return mapTableToTableEntity(table), nil
}

func (p *PostgresDB) DoesTableExist(id uint) (bool, error) {
	var table Table
	if err := p.db.Where("id = ?", id).First(&table).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *PostgresDB) DoesTableExistByTableNum(no uint8) (bool, error) {
	var table Table
	if err := p.db.Where("number = ?", no).First(&table).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
