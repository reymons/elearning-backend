package data

import (
    "context"
    "elrng/db"
)

type CourseDto struct {
    Id uint64
    Name string
    Description string
}

type courseDao struct{}

var CourseDao courseDao

func (courseDao) GetById(ctx context.Context, id uint64) (*CourseDto, error) {
    dto := new(CourseDto)
    row := db.DB.QueryRowContext(ctx, "SELECT id, name, description FROM courses WHERE id = $1", id)
    err := row.Scan(&dto.Id, &dto.Name, &dto.Description)
    return dto, err
}

