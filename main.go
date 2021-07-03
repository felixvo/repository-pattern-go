package main

import (
	"fmt"
	"repository-pattern/gorm_repository"
	"repository-pattern/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&repository.Author{}, &repository.Course{})
	if err != nil {
		panic(err)
	}
	func() { // create author and courses
		work := gorm_repository.NewGormUnitOfWork(db)
		author := repository.Author{Name: "felix", Email: "abcd@gmail.com"}

		err = work.AuthorRepo().Add(&author)
		if err != nil {
			work.Rollback()
			return
		}
		course := repository.Course{
			Name:     "Database 1",
			AuthorID: author.ID,
		}
		err = work.CourseRepo().Add(&course)
		if err != nil {
			work.Rollback()
			return
		}
		err = work.Complete()
		if err != nil {
			panic(err)
		}
	}()
	fmt.Println("Authors:")
	authorRepo := gorm_repository.NewAuthorRepo(db)
	authors, _ := authorRepo.All()
	for _, c := range authors {
		fmt.Println(c)
	}
	fmt.Println("Courses:")
	courseRepo := gorm_repository.NewCourseRepo(db)
	courses, _ := courseRepo.All()
	for _, c := range courses {
		fmt.Println(c)
	}
}
