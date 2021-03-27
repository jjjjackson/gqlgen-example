package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Environment struct {
	DatabaseURL string `envconfig:"DB_URL"`
}

const (
	userUID01 = "4f3973d6-2680-402b-a01b-592083224102"
)

func main() {
	var environments Environment
	if err := envconfig.Process("", &environments); err != nil {
		log.Fatal(err.Error())
	}

	db, err := gorm.Open(postgres.Open(environments.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	context := db.Begin()

	for _, sql := range migrations() {
		if err := context.Exec(sql).Error; err != nil {
			context.Rollback()
			log.Printf("%+v\n", err)
			return
		}
	}

	if err := context.Commit().Error; err != nil {
		fmt.Printf("😨 Seed Data Setting Was Failed: %+v \n", err)
		context.Rollback()
		return
	}

	fmt.Println("🚀 Seed Data Set!")
}

func migrations() []string {
	account001 := createUser(userUID01, "user001")

	return []string{
		account001,
	}
}

func createUser(uid, account string) string {
	email := account + "@demo.com"
	encryptedPassword := "$2a$07$RMzL7mjAcE8gExcLqUy8k.bVo9ziVA/aPX/r3aC8.j1ZssHH3k0zO" // demo1234
	salt := "1ZssHH"

	return fmt.Sprintf("INSERT INTO users (uid, email, password, salt, created_at, updated_at) VALUES ('%s', '%s', '%s', '%s', now(), now())",
		uid, email, encryptedPassword, salt)
}
