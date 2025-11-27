package service

import (
	"session-9/model"
	"session-9/repository"
	"testing"
)

func newTestService(Students []model.Student) (*StudentService, *repository.MockStudentRepository) {
	repo := &repository.MockStudentRepository{Students: Students}
	service := NewStudentService(repo)
	return service,repo
}

func TestStudent_Create(t *testing.T){
	service,repo:=newTestService([]model.Student{})

	created,err:=service.Create(model.Student{
		Name:"Rudi",
		Age:20,
	})

	if err!=nil{
		t.Fatalf("created return error: %v",err)
	}

	if created.ID!=1{
		t.Errorf("Expected ID 1, got: %d",created.ID)
	}

	if created.Name!="Rudi"{
		t.Errorf("Expected Name Budi, got: %s",created.Name)
	}

	if len(repo.Students)!=1{
		t.Errorf("Expected repo have to 1 students, got: %d",len(repo.Students))
	}
}