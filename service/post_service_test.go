package service

import (
	"fmt"
	"testing"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post entity.Post) ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	err := NewPostService(nil).Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "This is a test post"}
	err := NewPostService(nil).Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{Id: 1, Title: "A", Text: "B"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{Id: 3, Title: "A", Text: "B"}
	exitsting_post := []entity.Post{{Id: 1, Title: "C", Text: "D"}, {Id: 2, Title: "E", Text: "F"}}
	exitsting_post = append(exitsting_post, post)

	mockRepo.On("Save").Return(exitsting_post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(post)

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	fmt.Println(result)
	assert.Equal(t, exitsting_post, result)
	assert.Nil(t, err)
}
