package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fonteeBoa/go-log-guardian/internal/services"
	"github.com/fonteeBoa/go-log-guardian/internal/testUtils/mocks"
	"github.com/fonteeBoa/go-log-guardian/pkg/domain"
)

func TestFunction(t *testing.T) {
	priority := domain.LOG_INFO
	functionName := "TestFunction"

	logFunction := services.Function(priority, functionName, mocks.MockGeneralError, mocks.MockSpecificError)

	assert.Equal(t, priority, logFunction.Priority)
	assert.Equal(t, "LOG_INFO", logFunction.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, logFunction.Timestamp)          // Verifica se o Timestamp não é zero
	assert.Equal(t, mocks.MockGeneralError, logFunction.GenericErrorMessage)
	assert.Equal(t, mocks.MockSpecificError, logFunction.ErrorMessage)
	assert.Equal(t, functionName, logFunction.FunctionName)
}

func TestDatabase(t *testing.T) {
	priority := domain.LOG_DEBUG
	tableName := mocks.MockTable
	query := "SELECT * FROM TestTable"

	databaseLog := services.Database(priority, tableName, query, mocks.MockGeneralError, mocks.MockSpecificError)

	assert.Equal(t, priority, databaseLog.Priority)
	assert.Equal(t, "LOG_DEBUG", databaseLog.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, databaseLog.Timestamp)           // Verifica se o Timestamp não é zero
	assert.Equal(t, mocks.MockGeneralError, databaseLog.GenericErrorMessage)
	assert.Equal(t, mocks.MockSpecificError, databaseLog.ErrorMessage)
	assert.Equal(t, tableName, databaseLog.TableName)
	assert.Equal(t, query, databaseLog.Query)
}

func TestRequest(t *testing.T) {
	priority := domain.LOG_NOTICE
	method := "POST"
	statusCode := 500
	path := "/test/path"
	responseSize := 1024

	requestLog := services.Request(priority, method, statusCode, path, responseSize, mocks.MockGeneralError, mocks.MockSpecificError)

	assert.Equal(t, priority, requestLog.Priority)
	assert.Equal(t, "LOG_NOTICE", requestLog.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, requestLog.Timestamp)            // Verifica se o Timestamp não é zero
	assert.Equal(t, mocks.MockGeneralError, requestLog.GenericErrorMessage)
	assert.Equal(t, mocks.MockSpecificError, requestLog.ErrorMessage)
	assert.Equal(t, method, requestLog.Method)
	assert.Equal(t, statusCode, requestLog.StatusCode)
	assert.Equal(t, path, requestLog.Path)
	assert.Equal(t, responseSize, requestLog.ResponseSize)
}

func TestDetails(t *testing.T) {
	priority := domain.LOG_CRIT

	logDetails := services.Details(priority, mocks.MockGeneralError, mocks.MockSpecificError)

	assert.Equal(t, priority, logDetails.Priority)
	assert.Equal(t, "LOG_CRIT", logDetails.LogLevel) // Verifica se a string do LogLevel está correta
	assert.NotZero(t, logDetails.Timestamp)          // Verifica se o Timestamp não é zero
	assert.Equal(t, mocks.MockGeneralError, logDetails.GenericErrorMessage)
	assert.Equal(t, mocks.MockSpecificError, logDetails.ErrorMessage)
}
