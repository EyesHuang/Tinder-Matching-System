package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	person "tinder"
	rMock "tinder/mock"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

var validate *validator.Validate

func setup() (*rMock.MockPersonService, *chi.Mux) {
	if validate == nil {
		validate = validator.New()
	}
	person.Validate = validate

	mockService := new(rMock.MockPersonService)
	server := NewServer(mockService)
	return mockService, server.router
}

func TestHandlerAddSinglePersonAndMatch_Success(t *testing.T) {
	mockService, router := setup()

	mockPerson := person.Person{Name: "John", Gender: "male", Height: 180, WantedDates: 3}
	mockMatches := []*person.Person{{Name: "May", Gender: "female", Height: 160, WantedDates: 4}}
	mockService.On("AddPersonAndMatch", &mockPerson).Return(mockMatches, nil)

	body, err := json.Marshal(mockPerson)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/persons", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var result []person.Person
	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatal("Expected JSON response")
	}
	assert.Equal(t, "May", result[0].Name, "Expected May as a match")
}

func TestHandlerAddSinglePersonAndMatch_InvalidJSON(t *testing.T) {
	_, router := setup()

	body := bytes.NewBufferString("{invalidJson:}")
	req := httptest.NewRequest(http.MethodPost, "/persons", body)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var errResp ErrorResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &errResp); err != nil {
		t.Fatal("Expected JSON error response")
	}

	assert.NotEmpty(t, errResp.Errors, "Expected error messages in response")
}

func TestHandlerAddSinglePersonAndMatch_ValidationFailure(t *testing.T) {
	mockService, router := setup()

	invalidPerson := person.Person{Name: ""}
	body, err := json.Marshal(invalidPerson)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/persons", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	mockService.On("AddPersonAndMatch", &invalidPerson).Return(nil, errors.New("validation failed"))

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandlerRemoveSinglePerson_Success(t *testing.T) {
	mockService, router := setup()

	mockService.On("RemovePerson", "John").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/persons?name=John", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestHandlerRemoveSinglePerson_MissingQueryParam(t *testing.T) {
	_, router := setup()

	req := httptest.NewRequest(http.MethodDelete, "/persons", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandlerRemoveSinglePerson_NotFound(t *testing.T) {
	mockService, router := setup()

	mockService.On("RemovePerson", "Unknown").Return(errors.New(person.NotFoundStr))

	req := httptest.NewRequest(http.MethodDelete, "/persons?name=Unknown", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestHandlerQuerySinglePeople_Success(t *testing.T) {
	mockService, router := setup()

	mockPeople := []*person.Person{{Name: "John", Gender: "male", Height: 180, WantedDates: 3}}
	mockService.On("QuerySinglePeople", 1).Return(mockPeople, nil)

	req := httptest.NewRequest(http.MethodGet, "/persons?n=1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var result []person.Person
	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatal("Expected JSON response")
	}

	assert.Equal(t, "John", result[0].Name, "Expected John in the list")
}

func TestHandlerQuerySinglePeople_InvalidQueryParam(t *testing.T) {
	_, router := setup()

	req := httptest.NewRequest(http.MethodGet, "/persons?n=abc", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestHandlerQuerySinglePeople_ServiceError(t *testing.T) {
	mockService, router := setup()

	mockService.On("QuerySinglePeople", 2).Return(([]*person.Person)(nil), errors.New("internal server error"))

	req := httptest.NewRequest(http.MethodGet, "/persons?n=2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}
