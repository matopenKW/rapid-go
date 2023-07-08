// Code generated by MockGen. DO NOT EDIT.
// Source: asset_path.go

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/abyssparanoia/rapid-go/internal/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAssetPath is a mock of AssetPath interface.
type MockAssetPath struct {
	ctrl     *gomock.Controller
	recorder *MockAssetPathMockRecorder
}

// MockAssetPathMockRecorder is the mock recorder for MockAssetPath.
type MockAssetPathMockRecorder struct {
	mock *MockAssetPath
}

// NewMockAssetPath creates a new mock instance.
func NewMockAssetPath(ctrl *gomock.Controller) *MockAssetPath {
	mock := &MockAssetPath{ctrl: ctrl}
	mock.recorder = &MockAssetPathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssetPath) EXPECT() *MockAssetPathMockRecorder {
	return m.recorder
}

// GetWithValidate mocks base method.
func (m *MockAssetPath) GetWithValidate(ctx context.Context, assetType model.AssetType, assetKey string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithValidate", ctx, assetType, assetKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithValidate indicates an expected call of GetWithValidate.
func (mr *MockAssetPathMockRecorder) GetWithValidate(ctx, assetType, assetKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithValidate", reflect.TypeOf((*MockAssetPath)(nil).GetWithValidate), ctx, assetType, assetKey)
}

// Set mocks base method.
func (m *MockAssetPath) Set(ctx context.Context, asset *model.Asset, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, asset, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockAssetPathMockRecorder) Set(ctx, asset, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockAssetPath)(nil).Set), ctx, asset, expiration)
}