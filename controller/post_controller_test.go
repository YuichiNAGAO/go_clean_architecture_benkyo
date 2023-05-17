package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/entity"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/repository"
	"github.com/YuichiNAGAO/go_clean_architecture_benkyo/service"
	"github.com/stretchr/testify/assert"
)

var (
	postRepository repository.PostRepository = repository.NewPostRepository()
	postSvc        service.PostService       = service.NewPostService(postRepository)
	postConrtoller PostCntroller             = NewPostController(postSvc)
)

// func TestGetPosts(t *testing.T) {

// }

func TestAddPost(t *testing.T) {
	// HTTP POSTリクエストを作成する
	json := []byte(`{"id":1, "title":"タイトル", "text":"本文"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(json))

	// ハンドラののアサイン
	handler := http.HandlerFunc(postConrtoller.AddPost)

	// レスポンスをレコーダーを作成する
	response := httptest.NewRecorder()

	// requestをdispatchする
	handler.ServeHTTP(response, req)

	// レスポンスステータスコードのアサーション
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// レスポンスボディのデコード
	var posts []entity.Post
	_ = json.NewDecoder(response.Body).Decode(&posts)
	// レスポンスボディのアサーション
	assert.Equal(t, posts[0].Id, int64(1))
}
