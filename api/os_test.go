package api

import (
	"net/http"
	"path/filepath"

	"github.com/go-spectest/spectest"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("OS api test", func() {
	var (
		t   GinkgoTInterface
		api *API
	)

	BeforeEach(func() {
		t = GinkgoT()
		api = NewAPI()
	})

	Context("Success to GET /v1/os/android", func() {
		It("get version, code name, release date, api level", func() {
			spectest.New().
				Report(spectest.SequenceDiagram(documentDirPath())).
				CustomReportName("android_success").
				Report(spectest.SequenceReport(spectest.ReportFormatterConfig{
					Path: filepath.Join("..", "docs", "os"),
					Kind: spectest.ReportKindMarkdown,
				})).
				Handler(api).
				Get("/v1/os/android").
				Expect(t).
				BodyFromFile(filepath.Join("testdata", "android.json")).
				Status(http.StatusOK).
				End()
		})
	})

	Context("Success to GET /v1/os/debian", func() {
		It("get version, code name, release date", func() {
			spectest.New().
				Report(spectest.SequenceDiagram(documentDirPath())).
				CustomReportName("debian_success").
				Report(spectest.SequenceReport(spectest.ReportFormatterConfig{
					Path: filepath.Join("..", "docs", "os"),
					Kind: spectest.ReportKindMarkdown,
				})).
				Handler(api).
				Get("/v1/os/debian").
				Expect(t).
				BodyFromFile(filepath.Join("testdata", "debian.json")).
				Status(http.StatusOK).
				End()
		})
	})
})
