package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ericklima-ca/formx/models"
	"github.com/stretchr/testify/assert"
)

func TestFormPost(t *testing.T) {
	Init()
	router := setupRouter()
	formMock := models.Form{
		Name:  "Test",
		Email: "email@email.com",
		Phone: "+5592999999999",
	}
	w := httptest.NewRecorder()
	formData := url.Values{}
	formData.Set("name", formMock.Name)
	formData.Set("email", formMock.Email)
	formData.Set("phone", formMock.Phone)
	query := strings.NewReader(formData.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/v1/form_post", query)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code, "Fail to get success in form")
	assert.FileExists(t, "./temp/_tmp.pdf", "Fail to save file!")
}
