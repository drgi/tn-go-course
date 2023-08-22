package webapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testWebApp *WebServer

	indexWant     = map[string]interface{}{"key": "value"}
	documentsWant = []struct {
		Id    int64  `json:"id"`
		Title string `json:"title"`
	}{{
		Id:    1,
		Title: "go",
	},
		{
			Id:    2,
			Title: "test",
		},
	}
)

type SearcherMock struct{}

func (m *SearcherMock) List() (interface{}, error) {
	return indexWant, nil
}
func (m *SearcherMock) Documents() (interface{}, error) {
	return documentsWant, nil
}

func TestMain(m *testing.M) {
	testWebApp = New(&SearcherMock{})
	testWebApp.RegisterHandlers()
	m.Run()
}

func Test_WebServerIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	res := httptest.NewRecorder()
	testWebApp.router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Wrong http status in response. Got %d, want %d", res.Code, http.StatusOK)
	}

	b, _ := json.MarshalIndent(indexWant, "", "  ")
	want := string(b)
	got := string(res.Body.Bytes())

	if got != want {
		t.Errorf("Wrong http status in response. Got [%v], want [%v]", got, want)
	}
}

func Test_WebServerDocuments(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	res := httptest.NewRecorder()
	testWebApp.router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Wrong http status in response. Got %d, want %d", res.Code, http.StatusOK)
	}

	b, _ := json.MarshalIndent(documentsWant, "", "  ")
	want := string(b)
	got := string(res.Body.Bytes())

	if got != want {
		t.Errorf("Wrong http status in response. Got [%v], want [%v]", got, want)
	}
}
