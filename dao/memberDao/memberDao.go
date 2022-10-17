package memberDao

import (
	"errors"

	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"gorm.io/gorm"
)

const table = "member"

type QueryModel struct {
	ID          uint64
	Account     string
	Mail        string
	LineID      string
	CountryCode string
	Phone       string
}

func New(db *gorm.DB, model models.MemberModel) (uint64, error) {
	err := db.Table(table).Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func Get(tx *gorm.DB, query *QueryModel) (*models.MemberModel, error) {
	result := &models.MemberModel{}
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

func Gets(db *gorm.DB, query *QueryModel) ([]models.MemberModel, error) {
	result := make([]models.MemberModel, 0)
	err := db.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []models.MemberModel{}, nil
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
			Scopes(accountEqualScope(query.Account)).
			Scopes(mailEqualScope(query.Mail)).
			Scopes(lineIDEqualScope(query.LineID)).
			Scopes(phoneEqualScope(query.CountryCode, query.Phone))
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

func accountEqualScope(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if account != "" {
			return db.Where(table+".account = ?", account)
		}
		return db
	}
}

func mailEqualScope(mail string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if mail != "" {
			return db.Where(table+".mail = ?", mail)
		}
		return db
	}
}

func lineIDEqualScope(lineID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if lineID != "" {
			return db.Where(table+".line_id = ?", lineID)
		}
		return db
	}
}

func phoneEqualScope(counrtyCode string, phone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if counrtyCode != "" && phone != "" {
			return db.Where(table+".country_code = ? AND "+table+".phone = ?", counrtyCode, phone)
		}
		return db
	}
}
