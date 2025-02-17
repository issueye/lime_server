package initialize

import (
	"lime/internal/app/admin/logic"
	adminModel "lime/internal/app/admin/model"
	projectModel "lime/internal/app/project/model"
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
	db.AutoMigrate(&adminModel.User{})
	db.AutoMigrate(&adminModel.Role{})
	db.AutoMigrate(&adminModel.UserRole{})
	db.AutoMigrate(&adminModel.RoleMenu{})
	db.AutoMigrate(&adminModel.Menu{})
	db.AutoMigrate(&adminModel.DictsInfo{})
	db.AutoMigrate(&adminModel.DictDetail{})

	db.AutoMigrate(&projectModel.ProjectInfo{})
	db.AutoMigrate(&projectModel.BranchInfo{})
	db.AutoMigrate(&projectModel.TagInfo{})
	db.AutoMigrate(&projectModel.VersionInfo{})
	db.AutoMigrate(&projectModel.CompileInfo{})
	db.AutoMigrate(&projectModel.PackageInfo{})

	// admin
	logic.InitRoles()
	logic.InitRoleMenus()
	logic.InitUserRole()
	logic.InitAdminUser()
	logic.InitMenus()
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
