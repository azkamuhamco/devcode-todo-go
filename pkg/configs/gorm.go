package configs

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormConfig() *gorm.Config {
	return &gorm.Config{
		PrepareStmt:          true,
		DisableAutomaticPing: true,
		Logger:               logger.Default.LogMode(logger.Silent),
	}
}
