package postgres

import (
	"context"

	"github.com/techerfan/2DCH7-20059/entity"
)

func (p *PostgresDB) AddTable(ctx context.Context, table entity.Table) (entity.Table, error) {
	model := mapTableEntitytoTable(table)

	if err := p.db.WithContext(ctx).Create(&model).Error; err != nil {
		return entity.Table{}, err
	}

	return mapTableToTableEntity(model), nil
}

func (p *PostgresDB) RemoveTableByID(ctx context.Context, id uint) error {
	if err := p.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&Table{}).Error; err != nil {
		return err
	}
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
