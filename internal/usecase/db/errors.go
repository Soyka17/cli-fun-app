package db

import "strconv"

type UserAlreadyExistError struct {
	Id int
}

func (e UserAlreadyExistError) Error() string {
	return "User with id " + strconv.Itoa(e.Id) + " already exist"
}

type NilUserError struct{}

func (e NilUserError) Error() string {
	return "Nil user struct"
}

type UserWithZeroIdError struct{}

func (e UserWithZeroIdError) Error() string {
	return "User with zero id"
}

type UserNotFoundError struct {
	Id int
}

func (e UserNotFoundError) Error() string {
	return "User with id " + strconv.Itoa(e.Id) + " not found"
}
