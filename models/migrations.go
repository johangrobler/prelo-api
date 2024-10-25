package models

import "prelo/database"

func DBMigration() {
	db := database.DB.Db

	db.Exec("GRANT CREATE ON DATABASE prelo TO prelo;")

	//db.AutoMigrate(&Brand{})
	//db.AutoMigrate(&Category{})
	db.AutoMigrate(&User{})
	/*
	   db.AutoMigrate(&User{})
	   db.AutoMigrate(&Space{})
	   db.AutoMigrate(&Booking{})
	   db.AutoMigrate(&Review{})
	   db.AutoMigrate(&Feature{})
	   db.AutoMigrate(&Transaction{})
	   db.AutoMigrate(&Card{})
	   db.AutoMigrate(&Bank{})
	   db.AutoMigrate(&Sponsorship{})
	*/
	// db.AutoMigrate(&City{})
	// db.AutoMigrate(&Review{})

}
