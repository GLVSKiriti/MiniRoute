package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GLVSKiriti/MiniRoute/utils"
	"github.com/gorilla/mux"
)

func TestShorten(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT COALESCE\(MAX\(id\), 0\) FROM urlmappings WHERE uid=\$1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"coalesce"}).AddRow(9))

	mock.ExpectExec(`INSERT INTO urlmappings \(uid,id,longurl,shorturl\) VALUES \(\$1,\$2,\$3,\$4\)`).
		WithArgs(1, 10, "https://github.com/GLVSKiriti/MiniRoute", "1-10").WillReturnResult(sqlmock.NewResult(10, 1))

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM urlmappings WHERE shorturl = \$1`).WithArgs("repo").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	mock.ExpectQuery(`SELECT COALESCE\(MAX\(id\), 0\) FROM urlmappings WHERE uid=\$1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"coalesce"}).AddRow(9))

	mock.ExpectExec(`INSERT INTO urlmappings \(uid,id,longurl,shorturl\) VALUES \(\$1,\$2,\$3,\$4\)`).
		WithArgs(1, 10, "https://github.com/GLVSKiriti/MiniRoute", "repo").WillReturnResult(sqlmock.NewResult(10, 1))

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM urlmappings WHERE shorturl = \$1`).WithArgs("repo").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	h := NewBaseHandler(db)
	a := "repo"
	testCases := []struct {
		testName       string
		LongUrl        string
		Code           int
		CustomShortUrl *string //optional parameter
	}{
		{
			testName: "Only LongUrl is given",
			LongUrl:  "https://github.com/GLVSKiriti/MiniRoute",
			Code:     http.StatusOK,
		},
		{
			testName:       "ShortCode is given",
			LongUrl:        "https://github.com/GLVSKiriti/MiniRoute",
			Code:           http.StatusOK,
			CustomShortUrl: &a,
		},
		{
			testName:       "ShortCode already exists",
			LongUrl:        "https://github.com/GLVSKiriti/MiniRoute",
			Code:           http.StatusConflict,
			CustomShortUrl: &a,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			requestBody := map[string]string{
				"longUrl": tc.LongUrl,
			}
			if tc.CustomShortUrl != nil {
				requestBody["shortUrl"] = *tc.CustomShortUrl
			}

			req, rr, err := utils.MockHttpFunc("/shorten", "POST", requestBody)
			if err != nil {
				t.Fatal(err)
			}

			ctx := context.WithValue(req.Context(), "Uid", 1.0)
			// Create a new request with the updated context
			req = req.WithContext(ctx)

			// Call the function to be tested
			h.Shorten(rr, req)

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

func TestRedirectToOriginalUrl(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT longurl FROM urlmappings WHERE shorturl = \$1`).WithArgs("repo").
		WillReturnRows(sqlmock.NewRows([]string{"longurl"}).AddRow("https://github.com/GLVSKiriti/MiniRoute"))

	mock.ExpectQuery(`SELECT longurl FROM urlmappings WHERE shorturl = \$1`).WithArgs("nonShort").
		WillReturnError(sql.ErrNoRows)

	h := NewBaseHandler(db)

	testCases := []struct {
		testname  string
		shortCode string
		Code      int
	}{
		{
			testname:  "ShortCode valid",
			shortCode: "repo",
			Code:      http.StatusSeeOther,
		},
		{
			testname:  "ShortCode Invalid",
			shortCode: "nonShort",
			Code:      http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testname, func(t *testing.T) {
			req, rr, err := utils.MockHttpFunc(`/redirect/`+tc.shortCode, "GET", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Call the function to be tested
			router := mux.NewRouter()
			router.HandleFunc("/redirect/{shortCode}", h.RedirectToOriginalUrl)
			router.ServeHTTP(rr, req)

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
