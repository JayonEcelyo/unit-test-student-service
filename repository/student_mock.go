package repository

import "session-9/model"

type MockStudentRepository struct {
	Students []model.Student
	ErrGet error
	ErrSave error
}

func (mockStudentRepository *MockStudentRepository) GetAll() ([]model.Student, error) {
	 return mockStudentRepository.Students,mockStudentRepository.ErrGet
}
func (mockStudentRepository *MockStudentRepository) SaveAll(Students []model.Student) error {
	mockStudentRepository.Students=Students
	return mockStudentRepository.ErrSave
}