// Code generated by MockGen. DO NOT EDIT.
// Source: blep.ai/data/connectors/source/uniswapv2 (interfaces: TokenCacheInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	evm "blep.ai/data/connectors/source/evm"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockTokenCacheInterface is a mock of TokenCacheInterface interface.
type MockTokenCacheInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCacheInterfaceMockRecorder
}

// MockTokenCacheInterfaceMockRecorder is the mock recorder for MockTokenCacheInterface.
type MockTokenCacheInterfaceMockRecorder struct {
	mock *MockTokenCacheInterface
}

// NewMockTokenCacheInterface creates a new mock instance.
func NewMockTokenCacheInterface(ctrl *gomock.Controller) *MockTokenCacheInterface {
	mock := &MockTokenCacheInterface{ctrl: ctrl}
	mock.recorder = &MockTokenCacheInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenCacheInterface) EXPECT() *MockTokenCacheInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTokenCacheInterface) Get(arg0 common.Address) (*evm.ERC20, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*evm.ERC20)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTokenCacheInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTokenCacheInterface)(nil).Get), arg0)
}
