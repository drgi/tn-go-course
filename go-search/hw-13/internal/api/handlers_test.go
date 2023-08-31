package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tn-go-course/go-search/hw-13/pkg/crawler"
)

var (
	api *Api

	testDocument = []crawler.Document{
		{
			ID:    1,
			Title: "Doc 1",
		},
		{
			ID:    2,
			Title: "Doc 2",
		},
		{
			ID:    3,
			Title: "Doc 3",
		},
	}
)

type AppMock struct{}

func (a *AppMock) Search(_ context.Context, query string) ([]crawler.Document, error) {
	if len(testDocument) == 0 {
		return nil, errors.New("empty document list")
	}
	return testDocument, nil
}

func (a *AppMock) CreateDocument(ctx context.Context, doc crawler.Document) (int, error) {
	doc.ID = len(testDocument) + 1
	testDocument = append(testDocument, doc)
	return doc.ID, nil
}
func (a *AppMock) UpdateDocument(ctx context.Context, doc crawler.Document) error {
	return nil
}
func (a *AppMock) DeleteDocument(ctx context.Context, id int) error {
	return nil
}

func TestMain(m *testing.M) {
	api = New(&AppMock{})
	api.RegisterHandlers()
	m.Run()
}

func Test_WebServerSearch(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/search?query=test", nil)
	res := httptest.NewRecorder()
	api.router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Wrong http status in response. Got %d, want %d", res.Code, http.StatusOK)
	}

	b, _ := json.Marshal(&ResponseEnvelope{Result: testDocument})
	want := string(b)
	got := string(res.Body.Bytes())

	if got != want {
		t.Errorf("Wrong body in response. Got [%v], want [%v]", got, want)
	}
}

func Test_WebServerCreate(t *testing.T) {
	payload := &crawler.Document{
		Title: "NewDocument",
	}
	body, _ := json.Marshal(&payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/document", bytes.NewReader(body))
	res := httptest.NewRecorder()
	api.router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Wrong http status in response. Got %d, want %d", res.Code, http.StatusOK)
	}

	b, _ := json.Marshal(&ResponseEnvelope{Result: len(testDocument)})
	want := string(b)
	got := string(res.Body.Bytes())

	if got != want {
		t.Errorf("Wrong body in response. Got [%v], want [%v]", got, want)
	}
}

func Test_WebServerUpdate(t *testing.T) {
	payload := &crawler.Document{
		ID:    1,
		Title: "UpTittle",
	}
	body, _ := json.Marshal(&payload)
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/document", bytes.NewReader(body))
	res := httptest.NewRecorder()
	api.router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Wrong http status in response. Got %d, want %d", res.Code, http.StatusOK)
	}

	b, _ := json.Marshal(&ResponseEnvelope{Result: true})
	want := string(b)
	got := string(res.Body.Bytes())

	if got != want {
		t.Errorf("Wrong body in response. Got [%v], want [%v]", got, want)
	}
}

func Test_WebServerDelete(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/document/1", nil)
	res := httptest.NewRecorder()
	api.router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Wrong http status in response. Got %d, want %d", res.Code, http.StatusOK)
	}

	b, _ := json.Marshal(&ResponseEnvelope{Result: true})
	want := string(b)
	got := string(res.Body.Bytes())

	if got != want {
		t.Errorf("Wrong body in response. Got [%v], want [%v]", got, want)
	}
}
