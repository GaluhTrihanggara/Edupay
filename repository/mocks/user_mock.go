package mocks

import (
	"Edupay/model"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

// DeleteUserByIDRepository mock untuk method DeleteUserByIDRepository
func (_m *UserRepository) DeleteUserByIDRepository(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsersRepository mock untuk method GetAllUsersRepository
func (_m *UserRepository) GetAllUsersRepository(page int, limit int, name string) ([]*model.User, error) {
	ret := _m.Called(page, limit, name)

	var r0 []*model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*model.User, error)); ok {
		return rf(page, limit, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*model.User); ok {
		r0 = rf(page, limit, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(page, limit, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail mock untuk method GetUserByEmail
func (_m *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByIDRepository mock untuk method GetUserByIDRepository
func (_m *UserRepository) GetUserByIDRepository(id string) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByPhone mock untuk method GetUserByPhone
func (_m *UserRepository) GetUserByPhone(phone string) (*model.User, error) {
	ret := _m.Called(phone)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(phone)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(phone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByQueryRepository mock untuk method GetUserByQueryRepository
func (_m *UserRepository) GetUserByQueryRepository(query string, page int, limit int) ([]*model.User, error) {
	ret := _m.Called(query, page, limit)

	var r0 []*model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*model.User, error)); ok {
		return rf(query, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.User); ok {
		r0 = rf(query, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(query, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertAmountByUserIDRepository mock untuk method InsertAmountByUserIDRepository
func (_m *UserRepository) InsertAmountByUserIDRepository(userID string, amount float64) error {
	ret := _m.Called(userID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64) error); ok {
		r0 = rf(userID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserAmountByIDRepository mock untuk method UpdateUserAmountByIDRepository
func (_m *UserRepository) UpdateUserAmountByIDRepository(id string, user *model.User) (*model.User, error) {
	ret := _m.Called(id, user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.User) (*model.User, error)); ok {
		return rf(id, user)
	}
	if rf, ok := ret.Get(0).(func(string, *model.User) *model.User); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.User) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByIDRepository mock untuk method UpdateUserByIDRepository
func (_m *UserRepository) UpdateUserByIDRepository(id string, user *model.User) (*model.User, error) {
	ret := _m.Called(id, user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.User) (*model.User, error)); ok {
		return rf(id, user)
	}
	if rf, ok := ret.Get(0).(func(string, *model.User) *model.User); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.User) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository membuat instance baru dari UserRepository
// Argument pertama biasanya adalah *testing.T
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
