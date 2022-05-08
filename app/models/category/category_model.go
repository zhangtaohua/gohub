//Package category 模型
package category

import (
	"github.com/zhangtaohua/gohub/app/models"
	"github.com/zhangtaohua/gohub/pkg/database"
)

type Category struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
