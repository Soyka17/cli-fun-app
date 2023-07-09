package db

import (
	"ExperienceBank/internal/controller/logger"
	"ExperienceBank/internal/entity"
)

type MapDatabaseRepositoryImpl struct {
	db     map[int]*entity.User
	logger logger.LoggerRepository
	newId  int
}

func NewDB(l logger.LoggerRepository) *MapDatabaseRepositoryImpl {
	db := make(map[int]*entity.User)
	admin := entity.NewUser(1, "SEREGA", -1, "1")
	db[1] = admin
	db[2] = entity.NewUser(2, "NOT SEREGA", 1, "2")
	return &MapDatabaseRepositoryImpl{db: db, logger: l, newId: len(db) + 1}
}

func (m *MapDatabaseRepositoryImpl) SaveNewUser(user *entity.User) error {
	if user == nil {
		e := NilUserError{}
		m.logger.Warn("Can't save new user in db:" + e.Error())
		return e
	}
	if user.Id == 0 {
		e := UserWithZeroIdError{}
		m.logger.Warn("Can't save new user in db:" + e.Error())
		return e
	}
	if _, ok := m.db[user.Id]; ok {
		e := UserAlreadyExistError{}
		m.logger.Warn("Can't save new user in db:" + e.Error())
		return e
	}
	m.db[user.Id] = user
	m.newId++
	return nil
}

func (m *MapDatabaseRepositoryImpl) GetNewId() int {
	return m.newId
}

func (m *MapDatabaseRepositoryImpl) GetUserById(id int) (*entity.User, error) {
	if id == 0 {
		e := &UserWithZeroIdError{}
		m.logger.Warn("Can't get user by id in db:" + e.Error())
		return nil, e
	}
	if v, ok := m.db[id]; ok {
		return v, nil
	}
	return nil, &UserNotFoundError{Id: id}
}

func (m *MapDatabaseRepositoryImpl) UpdateUserInfo(user *entity.User) error {
	if user == nil {
		e := NilUserError{}
		m.logger.Warn("Can't update user info in db:" + e.Error())
		return e
	}
	if user.Id == 0 {
		e := UserWithZeroIdError{}
		m.logger.Warn("Can't update user info in db:" + e.Error())
		return e
	}
	if _, ok := m.db[user.Id]; !ok {
		e := UserNotFoundError{user.Id}
		m.logger.Warn("Can't update user info in db:" + e.Error())
		return e
	}

	delete(m.db, user.Id)
	m.db[user.Id] = user
	return nil
}
