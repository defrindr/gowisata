// repositories/db.go

package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/defrindr/gowisata/config"
)

var DB *gorm.DB

func InitDB() {
	connStr := config.Config.DBUser + ":" + config.Config.DBPassword + "@tcp(" + config.Config.DBHost + ":" + config.Config.DBPort + ")/" + config.Config.DBName + "?parseTime=true"
	var err error
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic("Failed to connect to database")
	}
}
