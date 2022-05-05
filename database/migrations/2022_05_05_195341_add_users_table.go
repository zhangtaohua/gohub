package migrations

import (
	"database/sql"

	"github.com/zhangtaohua/gohub/app/models"
	"github.com/zhangtaohua/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel

		Name     string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2022_05_05_195341_add_users_table", up, down)
	// FIXME("操作错误，自动生成的 migrate.Add 名称，无需修改！！！")
}
