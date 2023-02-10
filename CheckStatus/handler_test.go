package CheckStatus

import (
	"fmt"
	"github.com/RohomRutuja/Go_API/CheckStatus/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	mock "github.com/stretchr/testify/mock"
)

type checkWebsitesHandlerTestSuite struct {
	suite.Suite
	statusChecker *mocks.WebsiteChecker
}

func (suite *checkWebsitesHandlerTestSuite) SetupTest() {
	suite.statusChecker = &mocks.WebsiteChecker{}
}

func TestCheckWebsitesTestSuite(t *testing.T) {
	suite.Run(t, new(checkWebsitesHandlerTestSuite))
}

func (suite *checkWebsitesHandlerTestSuite) TestAddWebsitesHandler() {
	suite.Run("test if POST method returns status OK", func() {
		reqBody := `{
			"websites":["www.google.com"]
		}`
		rr := makeHTTPCall(http.MethodPost, "/websites", reqBody, AddWebsitesHandler(suite.statusChecker))
		assert.Equal(suite.T(), http.StatusOK, rr.Code)
	})

	suite.Run("check if the body is empty", func() {
		reqBody := `{}`
		rr := makeHTTPCall(http.MethodPost, "/websites", reqBody, AddWebsitesHandler(suite.statusChecker))
		assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	})
}

func (suite *checkWebsitesHandlerTestSuite) TestGethandler() {
	suite.Run("checks if GET request returns status BadRequest if wrong request being send", func() {
		suite.statusChecker.On("CreateStatus", mock.Anything, mock.Anything).Return("UP", nil)
		rr := makeHTTPCall(http.MethodGet, "/websites", "", GetWebsitesHandler(suite.statusChecker))
		assert.Equal(suite.T(), http.StatusOK, rr.Code)
		expected := fmt.Sprintf("%v\n", `{"www.google.com":"UP"}`)
		assert.Equal(suite.T(), expected, rr.Body.String())
	})

}

func makeHTTPCall(method, path, body string, handlerFunc http.HandlerFunc) (recorder *httptest.ResponseRecorder) {
	// create a http request using the given parameters
	req, _ := http.NewRequest(method, path, strings.NewReader(body))

	// test recorder created for capturing api responses
	recorder = httptest.NewRecorder()

	// create a router to serve the handler in test with the prepared request
	router := mux.NewRouter()
	router.HandleFunc(path, handlerFunc).Methods(method)

	// serve the request and write the response to recorder
	router.ServeHTTP(recorder, req)
	return
}
