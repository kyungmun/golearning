package main

import (
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"
	//sqlite3 "modernc.org/sqlite/lib"
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
)

type Book struct {
	ID     int       `json:"id"`
	Title  string    `json:"name"`
	Author string    `json:"author"`
	Rating int       `json:"rating"`
	Date   time.Time `json:"date" gorm:"column:date"`
}

type Email struct {
	Sender string `json:"sender"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//func (Email) TableName() string {
//	return "email"
//}

func benchmarkInsert(b *testing.B, d gorm.Dialector) {
	// open DB connection
	db, err := gorm.Open(d, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// CreateBatchSize:        100,
	})
	if err != nil {
		panic("failed to connect database")
	}
	/*
		// migrate DB
		if err := db.Migrator().DropTable(&Book{}); err != nil {
			b.Fatal(err)
		}
		if err := db.Migrator().AutoMigrate(&Book{}); err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			book := new(Book)
			book.Author = "tolkein"
			book.Title = "lord of the rings"
			book.Rating = 5
			book.Date = time.Now()
			if err := db.Create(&book).Error; err != nil {
				b.Fatal(err)
			}
		}
	*/

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		email := new(Email)
		email.Sender = "kyung mun" + strconv.Itoa(i)
		email.Title = "title of mail"
		email.Body = "body test" + strconv.Itoa(i)
		if err := db.Create(&email).Error; err != nil {
			b.Fatal(err)
		}
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

}

func BenchmarkInsertPureGo(b *testing.B) {
	//fmt.Println(b.N)
	benchmarkInsert(b, sqlite.Open("test_pure.db?_pragma=journal_mode(wal)"))
}

func BenchmarkInsertCGo(b *testing.B) {
	//fmt.Println(b.N)
	//benchmarkInsert(b, sqlite3.Open("test_cgo.db?_journal_mode=WAL"))
}
