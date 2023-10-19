package api

import (
	"net/http"

	"github.com/go-spectest/spectest"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Network api test", func() {
	var (
		t   GinkgoTInterface
		api *API
	)

	BeforeEach(func() {
		t = GinkgoT()
		api = NewAPI()
	})

	Context("Success to GET /v1/ip", func() {
		It("get client global ip", func() {
			spectest.New().
				Report(spectest.SequenceDiagram(documentDirPath())).
				CustomReportName("ip_success").
				Handler(api).
				Get("/v1/ip").
				Headers(map[string]string{"X-Real-IP": "127.0.0.1"}).
				Expect(t).
				Body(`{"global_ip": "127.0.0.1"}`).
				Status(http.StatusOK).
				End()
		})
	})
})
