package api

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"unicode/utf8"

	"github.com/go-spectest/spectest"
	"github.com/zeebo/assert"
)

func TestULID(t *testing.T) { //nolint
	spectest.New().
		Report(spectest.SequenceDiagram(documentDirPath())).
		CustomReportName("ulid_success").
		Handler(NewAPI()).
		Get("/v1/ulid").
		Expect(t).
		Assert(func(res *http.Response, req *http.Request) error {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close() //nolint

			var got ULIDResponse
			if err := json.Unmarshal(body, &got); err != nil {
				t.Error(err)
			}
			assert.Equal(t, utf8.RuneCountInString(got.ULID), 26)
			return nil
		}).
		Status(http.StatusOK).
		End()
}
