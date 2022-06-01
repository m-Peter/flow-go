// Code generated by mockery v2.12.1. DO NOT EDIT.

package mock

import (
	fvm "github.com/onflow/flow-go/fvm"
	mock "github.com/stretchr/testify/mock"

	programs "github.com/onflow/flow-go/fvm/programs"

	state "github.com/onflow/flow-go/fvm/state"

	testing "testing"
)

// ScriptProcessor is an autogenerated mock type for the ScriptProcessor type
type ScriptProcessor struct {
	mock.Mock
}

// Process provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *ScriptProcessor) Process(_a0 *fvm.VirtualMachine, _a1 fvm.Context, _a2 *fvm.ScriptProcedure, _a3 *state.StateHolder, _a4 *programs.Programs) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fvm.VirtualMachine, fvm.Context, *fvm.ScriptProcedure, *state.StateHolder, *programs.Programs) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewScriptProcessor creates a new instance of ScriptProcessor. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewScriptProcessor(t testing.TB) *ScriptProcessor {
	mock := &ScriptProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
