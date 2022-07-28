package tests

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/ericklima-ca/formx/models"
	"github.com/stretchr/testify/assert"
)

func TestServeStatic(t *testing.T) {
	Init()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/v1/", nil)
	if err != nil {
		log.Fatalln(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewForm(t *testing.T) {
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
	path_name := fmt.Sprintf("./temp/%s.pdf", formMock.Name)
	defer os.Remove(path_name)

	assert.Equal(t, http.StatusFound, w.Code, "Fail to get success in form")
	assert.FileExists(t, path_name, "Fail to save file!")
}
