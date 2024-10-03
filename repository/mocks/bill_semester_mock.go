package mocks

import (
	"Edupay/model"

	"github.com/stretchr/testify/mock"
)

type BillSemesterRepository struct {
	mock.Mock
}

// DeleteBillSemesterByIDRepository mock for the method DeleteBillSemesterByIDRepository
func (_m *BillSemesterRepository) DeleteBillSemesterByIDRepository(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBillSemestersRepository mock for the method GetAllBillSemestersRepository
func (_m *BillSemesterRepository) GetAllBillSemestersRepository(page int, limit int, studentId string) ([]*model.BillSemester, error) {
	ret := _m.Called(page, limit, studentId)

	var r0 []*model.BillSemester
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*model.BillSemester, error)); ok {
		return rf(page, limit, studentId)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*model.BillSemester); ok {
		r0 = rf(page, limit, studentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BillSemester)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(page, limit, studentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBillSemesterByIDRepository mock for the method GetBillSemesterByIDRepository
func (_m *BillSemesterRepository) GetBillSemesterByIDRepository(id string) (*model.BillSemester, error) {
	ret := _m.Called(id)

	var r0 *model.BillSemester
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.BillSemester, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.BillSemester); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BillSemester)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBillSemesterRepository mock for the method InsertBillSemesterRepository
func (_m *BillSemesterRepository) InsertBillSemesterRepository(bill *model.BillSemester) error {
	ret := _m.Called(bill)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.BillSemester) error); ok {
		r0 = rf(bill)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBillSemesterByIDRepository mock for the method UpdateBillSemesterByIDRepository
func (_m *BillSemesterRepository) UpdateBillSemesterByIDRepository(id string, bill *model.BillSemester) (*model.BillSemester, error) {
	ret := _m.Called(id, bill)

	var r0 *model.BillSemester
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.BillSemester) (*model.BillSemester, error)); ok {
		return rf(id, bill)
	}
	if rf, ok := ret.Get(0).(func(string, *model.BillSemester) *model.BillSemester); ok {
		r0 = rf(id, bill)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BillSemester)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.BillSemester) error); ok {
		r1 = rf(id, bill)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBillSemesterRepository creates a new instance of BillSemesterRepository
// The first argument is usually *testing.T
func NewBillSemesterRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BillSemesterRepository {
	mock := &BillSemesterRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
