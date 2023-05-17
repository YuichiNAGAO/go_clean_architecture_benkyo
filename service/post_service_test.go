package service

import (
	"testing"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/stretchr/testify/assert"
)

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
