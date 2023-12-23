package handlers

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GLVSKiriti/MiniRoute/utils"
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

			req, rr, err := utils.MockHttpFunc("/login", "POST", requestBody)
			if err != nil {
				t.Fatal(err)
			}
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

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE email=\$1`).WithArgs("test1@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM users WHERE email=\$1`).WithArgs("newUser@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	mock.ExpectQuery(`INSERT INTO users \(email,password\) VALUES\(\$1,\$2\) RETURNING uid`).
		WithArgs("newUser@gmail.com", "newPassword").
		WillReturnRows(sqlmock.NewRows([]string{"uid"}).AddRow(2))

	h := NewBaseHandler(db)

	testCases := []struct {
		testname string
		email    string
		password string
		Code     int
	}{
		{
			testname: "User Already Exists",
			email:    "test1@gmail.com",
			password: "anyPassword",
			Code:     http.StatusConflict,
		},
		{
			testname: "New User So Register",
			email:    "newUser@gmail.com",
			password: "newPassword",
			Code:     http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testname, func(t *testing.T) {
			requestBody := map[string]string{
				"email":    tc.email,
				"password": tc.password,
			}

			req, rr, err := utils.MockHttpFunc("/register", "POST", requestBody)
			if err != nil {
				t.Fatal(err)
			}
			// Call the function to be tested
			h.Register(rr, req)

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
