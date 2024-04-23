// Code generated by MockGen. DO NOT EDIT.
// Source: ../../service/taxcsv.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTaxCSVService is a mock of TaxCSVService interface.
type MockTaxCSVService struct {
	ctrl     *gomock.Controller
	recorder *MockTaxCSVServiceMockRecorder
}

// MockTaxCSVServiceMockRecorder is the mock recorder for MockTaxCSVService.
type MockTaxCSVServiceMockRecorder struct {
	mock *MockTaxCSVService
}

// NewMockTaxCSVService creates a new mock instance.
func NewMockTaxCSVService(ctrl *gomock.Controller) *MockTaxCSVService {
	mock := &MockTaxCSVService{ctrl: ctrl}
	mock.recorder = &MockTaxCSVServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaxCSVService) EXPECT() *MockTaxCSVServiceMockRecorder {
	return m.recorder
}

// CalculateTax mocks base method.
func (m *MockTaxCSVService) CalculateTax(totalIncome, wht, donation, kReceipt float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateTax", totalIncome, wht, donation, kReceipt)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateTax indicates an expected call of CalculateTax.
func (mr *MockTaxCSVServiceMockRecorder) CalculateTax(totalIncome, wht, donation, kReceipt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateTax", reflect.TypeOf((*MockTaxCSVService)(nil).CalculateTax), totalIncome, wht, donation, kReceipt)
}

// ImportCSV mocks base method.
func (m *MockTaxCSVService) ImportCSV(reader io.Reader) ([]map[string]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportCSV", reader)
	ret0, _ := ret[0].([]map[string]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportCSV indicates an expected call of ImportCSV.
func (mr *MockTaxCSVServiceMockRecorder) ImportCSV(reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCSV", reflect.TypeOf((*MockTaxCSVService)(nil).ImportCSV), reader)
}
