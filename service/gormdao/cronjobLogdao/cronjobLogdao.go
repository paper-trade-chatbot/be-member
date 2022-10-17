package cronjobLogdao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const table = "cronjob_log"

type QueryModel struct {
	Id        int
	Code      string
	CreatedBy string
}

// New a row
func New(tx *gorm.DB, role *models.CronjobModel) {
	err := tx.Table(table).
		Create(role).Error

	if err != nil {
		panic(err)
	}
}

// Modify a row
func Modify(tx *gorm.DB, cronJobLog *models.CronjobModel) {
	attrs := map[string]interface{}{
		"log": cronJobLog.Log,
	}

	err := tx.Table(table).
		Model(models.CronjobModel{}).
		Where("id = ?", cronJobLog.ID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// Get return a record as raw-data-form
func Get(tx *gorm.DB, query *QueryModel) *models.CronjobModel {
	result := models.CronjobModel{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return &result
}

// Gets return records as raw-data-form
func Gets(tx *gorm.DB, query *QueryModel) []models.CronjobModel {
	result := []models.CronjobModel{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]models.CronjobModel, 0)
	}
	if err != nil {
		panic(err)
	}

	return result
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(idEqualScope(query.Id))

	}
}

func idEqualScope(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id != 0 {
			return db.Where(table+".id = ?", id)
		}
		return db
	}
}
