package migrate

import "gohub/pkg/database"

// Migration 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// GetLastMigration 获取最后一条迁移记录
func getLastMigration() (lastMigration Migration) {
	database.DB.Order("id DESC").First(&lastMigration)
	return
}

// GetLastMigration 获取迁移记录列表
func getLastMigrationFind() (migrations []Migration) {
	database.DB.Order("id DESC").Find(&migrations)
	return
}

// getByBatch 根据 batch 获取迁移记录列表
func getByBatch(batch int) (migrations []Migration) {
	database.DB.Where("batch = ?", batch).Order("id DESC").Find(&migrations)
	return
}
