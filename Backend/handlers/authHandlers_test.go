package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Set up expectations for the SQL query
	mock.ExpectQuery(`SELECT uid,password FROM users WHERE email=?`).WithArgs("test1@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"uid", "password"}).AddRow(1, "test1@12345"))

	mock.ExpectQuery(`SELECT uid,password FROM users WHERE email=?`).WithArgs("newUser@gmail.com").
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery(`SELECT uid,password FROM users WHERE email=?`).WithArgs("test1@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"uid", "password"}).AddRow(1, "test1@12345"))

	h := NewBaseHandler(db)

	testCases := []struct {
		testName string
		email    string
		password string
		Code     int
	}{
		{
			testName: "Correct Details",
			email:    "test1@gmail.com",
			password: "test1@12345",
			Code:     http.StatusOK,
		},
		{
			testName: "New User",
			email:    "newUser@gmail.com",
			password: "test1@12345",
			Code:     http.StatusNotFound,
		},
		{
			testName: "Wrong Password",
			email:    "test1@gmail.com",
			password: "wrongPassword",
			Code:     http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			requestBody := map[string]string{
				"email":    tc.email,
				"password": tc.password,
			}
			jsonBody, err := json.Marshal(requestBody)
			if err != nil {
				t.Fatal(err)
			}

			// Create a new mock request
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the function to be tested
			h.Login(rr, req)

			if rr.Code != tc.Code {
				t.Errorf("Expected status code %d, got %d", tc.Code, rr.Code)
			}
		})
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %s", err)
	}
}
