package seeds

import "gorm.io/gorm"

type Seed struct {
	db *gorm.DB
}

func seed(s Seed, seedMethod string) {

}
