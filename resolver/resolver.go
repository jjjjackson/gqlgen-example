package resolver

import (
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}
