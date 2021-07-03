package gorm_repository

import (
	"gorm.io/gorm"
	"repository-pattern/repository"
)

type GormUnitOfWork struct {
	courseRepo repository.CourseRepository
	authorRepo repository.AuthorRepository
	db         *gorm.DB
}

func NewGormUnitOfWork(db *gorm.DB) repository.IUnitOfWork {
	tx := db.Begin()
	return &GormUnitOfWork{
		courseRepo: NewCourseRepo(tx),
		authorRepo: NewAuthorRepo(tx),
		db:         tx,
	}
}

func (g *GormUnitOfWork) CourseRepo() repository.CourseRepository {
	return g.courseRepo
}

func (g *GormUnitOfWork) AuthorRepo() repository.AuthorRepository {
	return g.authorRepo
}

func (g *GormUnitOfWork) Complete() error {
	err := g.db.Commit().Error
	if err != nil {
		g.db.Rollback()
	}
	return err
}

func (g *GormUnitOfWork) Rollback() error {
	return g.db.Rollback().Error
}

type authorRepo struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) repository.AuthorRepository {
	return &authorRepo{db: db}
}

func (a *authorRepo) Add(author *repository.Author) error {
	return a.db.Create(author).Error
}

func (a *authorRepo) Get(id uint) (*repository.Author, error) {
	author := repository.Author{}
	return &author, a.db.First(&author, id).Error
}

func (a *authorRepo) Remove(author *repository.Author) error {
	return a.db.Delete(author).Error
}

func (a *authorRepo) All() ([]*repository.Author, error) {
	var authors []*repository.Author
	err := a.db.Find(&authors).Error
	return authors, err
}

type courseRepo struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) repository.CourseRepository {
	return &courseRepo{db: db}
}

func (c *courseRepo) Add(course *repository.Course) error {
	return c.db.Create(course).Error
}

func (c *courseRepo) Get(id uint) (*repository.Course, error) {
	course := repository.Course{}
	return &course, c.db.First(&course, id).Error
}

func (c *courseRepo) Remove(course *repository.Course) error {
	return c.db.Delete(course).Error
}

func (c *courseRepo) GetByAuthorID(authorID uint) ([]*repository.Course, error) {
	var courses []*repository.Course
	err := c.db.Where("author_id = ?", authorID).Find(&courses).Error
	return courses, err
}

func (c *courseRepo) All() ([]*repository.Course, error) {
	var courses []*repository.Course
	err := c.db.Find(&courses).Error
	return courses, err
}
