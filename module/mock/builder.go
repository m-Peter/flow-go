// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import flow "github.com/dapperlabs/flow-go/model/flow"
import mock "github.com/stretchr/testify/mock"

// Builder is an autogenerated mock type for the Builder type
type Builder struct {
	mock.Mock
}

// BuildOn provides a mock function with given fields: parentID
func (_m *Builder) BuildOn(parentID flow.Identifier) (*flow.Block, error) {
	ret := _m.Called(parentID)

	var r0 *flow.Block
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.Block); ok {
		r0 = rf(parentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(parentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
