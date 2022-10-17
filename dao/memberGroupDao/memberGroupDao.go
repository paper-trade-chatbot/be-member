package memberDao

import (
	"errors"

	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"gorm.io/gorm"
)

const table = "member_group"

type QueryModel struct {
	ID   uint64
	Name string
}

func New(db *gorm.DB, model models.MemberGroupModel) (uint64, error) {
	err := db.Table(table).Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func Get(tx *gorm.DB, query *QueryModel) (*models.MemberGroupModel, error) {
	result := &models.MemberGroupModel{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Gets(db *gorm.DB, query *QueryModel) ([]models.MemberGroupModel, error) {
	result := make([]models.MemberGroupModel, 0)
	err := db.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.MemberGroupModel{}, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Delete(db *gorm.DB, query *QueryModel) error {
	return db.Table(table).
		Scopes(queryChain(query)).
		Delete(&models.MemberModel{}).Error
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(idEqualScope(query.ID)).
			Scopes(nameEqualScope(query.Name))
	}
}

func idEqualScope(id uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id != 0 {
			return db.Where(table+".id = ?", id)
		}
		return db
	}
}

func nameEqualScope(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			return db.Where(table+".name = ?", name)
		}
		return db
	}
}
