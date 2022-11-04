package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/pkg/handler"
	"github.com/Ckala62rus/go/pkg/services"
	mock_services "github.com/Ckala62rus/go/pkg/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_services.MockAuthorization, user handler.CreateAuthUser)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           handler.CreateAuthUser
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "",
			inputBody: `{"email":"agr.akyla@mail.ru", "password":"123123"}`,
			inputUser: handler.CreateAuthUser{
				Email:    "agr.akyla@mail.ru",
				Password: "123123",
			},
			mockBehavior: func(s *mock_services.MockAuthorization, user handler.CreateAuthUser) {
				s.EXPECT().CreateUser(domain.User{
					Email:    user.Email,
					Password: user.Password,
				}).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":true,"message":"User success created","data":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)

			// всегда писать, это требование пакета.
			// important write ( defer c.Finish() ) it because need for package
			defer c.Finish()

			auth := mock_services.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &services.Service{Authorization: auth}
			handler := handler.NewHandler(services)

			// Test Server
			r := gin.New()
			r.POST("/sign-up", handler.SignUp)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(
				"POST",
				"/sign-up",
				bytes.NewBufferString(testCase.inputBody),
			)

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
