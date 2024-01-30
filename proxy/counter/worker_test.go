package counter

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockFileWriter - мок для эмуляции записи файла
type MockFileWriter struct {
	mock.Mock
}

func (m *MockFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	args := m.Called(filename, data, perm)
	return args.Error(0)
}

func TestCounterWorker(t *testing.T) {
	mockFileWriter := new(MockFileWriter)

	// Устанавливаем ожидание вызова WriteFile
	mockFileWriter.On("WriteFile", appPathToCounterFile, mock.Anything, os.FileMode(0644)).Return(nil)

	// Заменяем функцию записи файла в CounterWorker на мок-функцию
	oldFileWriter := fileWriter
	fileWriter = mockFileWriter
	defer func() { fileWriter = oldFileWriter }()

	// Запускаем CounterWorker в отдельной горутине
	go CounterWorker()

	// Ждем некоторое время для выполнения нескольких циклов
	time.Sleep(11 * time.Second)

	// Проверяем, что WriteFile был вызван ожидаемое количество раз
	mockFileWriter.AssertNumberOfCalls(t, "WriteFile", 2)

	// Проверяем, что все ожидаемые вызовы были сделаны
	mockFileWriter.AssertExpectations(t)
}
