// Code generated by MockGen. DO NOT EDIT.
// Source: database/database.go

// Package database is a generated GoMock package.
package database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDatabase) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDatabaseMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDatabase)(nil).Create), value)
}

// Delete mocks base method.
func (m *MockDatabase) Delete(value interface{}, where ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{value}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDatabaseMockRecorder) Delete(value interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{value}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabase)(nil).Delete), varargs...)
}

// Find mocks base method.
func (m *MockDatabase) Find(out interface{}, where ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockDatabaseMockRecorder) Find(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDatabase)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockDatabase) First(out interface{}, where ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockDatabaseMockRecorder) First(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockDatabase)(nil).First), varargs...)
}

// Save mocks base method.
func (m *MockDatabase) Save(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockDatabaseMockRecorder) Save(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDatabase)(nil).Save), value)
}

// Where mocks base method.
func (m *MockDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockDatabaseMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockDatabase)(nil).Where), varargs...)
}