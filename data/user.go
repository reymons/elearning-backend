package data

import (
    "elrng/db"
    "context"
)

type UserDto struct {
    Id uint64
    FirstName string
    LastName string
    Email string
}

type userDao struct{}

var UserDao userDao

func (userDao) GetById(ctx context.Context, id uint64) (*UserDto, error) {
    dto := new(UserDto)
    row := db.DB.QueryRowContext(ctx, "SELECT id, first_name, last_name, email FROM users WHERE id = $1", id)
    err := row.Scan(&dto.Id, &dto.FirstName, &dto.LastName, &dto.Email)
    return dto, err
}

