package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"unicode/utf8"

	"github.com/go-spectest/spectest"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	"github.com/zeebo/assert"
)

var _ = Describe("ULID api test", func() {
	var (
		t   GinkgoTInterface
		api *API
	)

	BeforeEach(func() {
		t = GinkgoT()
		api = NewAPI()
	})

	Context("Success to get ulid", func() {
		It("generate ulid by server", func() {
			spectest.New().
				Report(spectest.SequenceDiagram(documentDirPath())).
				CustomReportName("ulid_success").
				Handler(api).
				Get("/v1/ulid").
				Expect(t).
				// Assert(jsonpath.Contains(`$.ulid`,"")).
				Status(http.StatusOK).
				End()
		})
	})
})

// This test 
func TestULIDHandler(t *testing.T) {
	t.Parallel()

	t.Run("return ulid response", func(t *testing.T) {
		t.Parallel()

		e := echo.New()
		ctrl := NewULIDController()
		e.GET("/ulid", ctrl.generate)

		req := httptest.NewRequest(http.MethodGet, "/ulid", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := ctrl.generate(c)

		// assert response
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		// This assertion ensures that the response is not nil
		// and contains a string that is 26 characters.
		var got ULIDResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, utf8.RuneCountInString(got.ULID), 26)
	})
}
