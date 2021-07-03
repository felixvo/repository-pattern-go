package repository

import (
	"fmt"
	"time"
)

type Author struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name    string
	Email   string
	Courses []Course
}

func (a *Author) String() string {
	return fmt.Sprintf("Author ID: %d Name: %s Email:%s", a.ID, a.Name, a.Email)
}

type Course struct {
	ID        uint
	AuthorID uint
	Name     string
	Length   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Course) String() string {
	return fmt.Sprintf("Course ID:%d Name: %s", c.ID, c.Name)
}

type AuthorRepository interface {
	Add(author *Author) error
	All() ([]*Author, error)
	Get(id uint) (*Author, error)
	Remove(author *Author) error
}

type CourseRepository interface {
	Add(course *Course) error
	Get(id uint) (*Course, error)
	All() ([]*Course, error)
	Remove(course *Course) error
	GetByAuthorID(authorID uint) ([]*Course, error)
}

type IUnitOfWork interface {
	CourseRepo() CourseRepository
	AuthorRepo() AuthorRepository
	Complete() error
	Rollback() error
}
