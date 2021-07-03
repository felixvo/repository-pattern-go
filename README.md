# Repository Pattern in Go with [gorm](https://gorm.io/)

The sample implementation of Repository Pattern in go with gorm

Let's first understand what is Repository Pattern

[![IMAGE ALT TEXT](http://img.youtube.com/vi/rtXpYpZdOzM/0.jpg)](http://www.youtube.com/watch?v=rtXpYpZdOzM "Repository Pattern with C# and Entity Framework, Done Right | Mosh")

## Key Notes from [P of EAA Catalog - Repository](https://martinfowler.com/eaaCatalog/repository.html)

- isolates domain objects from details of the database access code
- acting like an in-memory domain object collection

## UnitOfWork usage sample

```go
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
err = work.Complete() // commit changes
if err != nil {
    panic(err)
}
```

