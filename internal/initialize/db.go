package initialize

import (
	"lime/internal/app/admin/initialize"
	"lime/internal/global"
	"lime/pkg/db"
	"path/filepath"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"gorm.io/gorm"
)

func InitDB() {
	path := filepath.Join(global.ROOT_PATH, "data", "data.db")
	global.DB = db.InitSqlite(path, global.Logger.Sugar())

	InitDATA(global.DB)
}

func InitDATA(db *gorm.DB) {
	initialize.InitAdminData(db)
	// admin
	// logic.InitRoleMenus()
	// logic.InitUserRole()
	// logic.InitAdminUser()
}

func FreeDB() {
	sqldb, err := global.DB.DB()
	if err != nil {
		global.Logger.Sugar().Errorf("close db error: %v", err)
	}

	if err = sqldb.Close(); err != nil {
		global.Logger.Sugar().Errorf("close db error: %v", err)
	}
}
