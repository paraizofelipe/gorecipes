// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GifSearcher is an autogenerated mock type for the GifSearcher type
type GifSearcher struct {
	mock.Mock
}

// Search provides a mock function with given fields: title
func (_m *GifSearcher) Search(title string) (string, error) {
	ret := _m.Called(title)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
