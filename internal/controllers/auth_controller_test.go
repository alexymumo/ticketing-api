package controllers_test

import (
	"database/sql"
	"events/internal/controllers"
	"events/pkg/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	route := gin.Default()
	route.DELETE("/user/:userid", controllers.DeleteUser(database.Connect()))
	return route
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' occurred when opening stub connection", err)
	}
	defer db.Close()
	mock.ExpectExec("DELETE FROM user WHERE userid = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	router := SetupRouter(db)

	r, _ := http.NewRequest("DELETE", "/user/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"successfuly deleted\"}", w.Body.String())
}

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' occurred when opening stub connection", err)
	}
	defer db.Close()
	mock.ExpectExec("INSERT INTO user").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s' occurred when opening stub connection", err)
	}
	defer db.Close()
	mock.ExpectExec("SELECT * FROM user")
}

func TestUpdateUser(t *testing.T) {

}
