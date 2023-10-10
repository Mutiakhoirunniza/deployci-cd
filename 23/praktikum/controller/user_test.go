package controller_test

import (
    "20/controller"
    "20/model"
    "20/repository"
    "errors"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func TestGetAllUsersSuccess(t *testing.T) {
    // Inisialisasi Echo
    e := echo.New()

    // Membuat mock repository dan controller
    mockUserRepo := &repository.MockUserRepository{}
    uController := controller.NewUserController(mockUserRepo)

    // Membuat permintaan HTTP GET palsu
    req := httptest.NewRequest(http.MethodGet, "/users", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Mengatur mock repository untuk mengembalikan beberapa data pengguna
    users := []model.User{}
    mockUserRepo.On("Find").Return(users, nil)

    // Menjalankan fungsi pengujian
    err := uController.GetAllUsers(c)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    mockUserRepo.AssertExpectations(t)
}

func TestCreateUserValidJSON(t *testing.T) {
    // JSON body yang valid
    reqbody := `{ "email": "Mutia123@gmail.com", "password": "mutia1234fyu" }`

    // Membuat mock repository dan controller
    mockUserRepo := &repository.MockUserRepository{}
    uController := controller.NewUserController(mockUserRepo)

    // Membuat pengguna yang diharapkan untuk dibuat
    expectedUser := model.User{
        Email:    "Mutia123@gmail.com",
        Password: "mutia1234fyu",
    }

    mockUserRepo.On("Create", expectedUser).Return(nil)

    // Membuat permintaan HTTP POST palsu dengan JSON body
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    e := echo.New()
    c := e.NewContext(req, rec)

    // Menjalankan fungsi pengujian
    err := uController.CreateUser(c)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    
    mockUserRepo.AssertExpectations(t)
}

func TestCreateUserOtherErrors(t *testing.T) {
    // JSON body yang valid
    reqbody := `{ "email": "Mutia123@gmail.com"}`

    // Membuat mock repository dan controller
    mockUserRepo := &repository.MockUserRepository{}
    uController := controller.NewUserController(mockUserRepo)

    // Membuat pengguna yang diharapkan untuk dibuat
    expectedUser := model.User{
        Email:    "Mutia123@gmail.com",
    }

    // Mengatur ekspektasi panggilan Create pada mock repository untuk mengembalikan kesalahan
    mockUserRepo.On("Create", expectedUser).Return(errors.New("some other error"))

    // Membuat permintaan HTTP POST palsu dengan JSON body
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    e := echo.New()
    c := e.NewContext(req, rec)

    // Menjalankan fungsi pengujian
    if assert.NoError(t, uController.CreateUser(c)){
        assert.Equal(t, http.StatusInternalServerError, rec.Code)
    }

    mockUserRepo.AssertExpectations(t)
}
