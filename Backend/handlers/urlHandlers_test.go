package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

			jsonBody, err := json.Marshal(requestBody)
			if err != nil {
				t.Fatal(err)
			}

			// Create a new mock request
			req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.WithValue(req.Context(), "Uid", 1.0)

			// Create a new request with the updated context
			req = req.WithContext(ctx)

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

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
