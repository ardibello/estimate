package server_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPostIssues(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// create echo context
		c, rec := newEchoContext(&newEchoContextParams{
			method:  http.MethodPost,
			target:  "/issues",
			payload: fmt.Sprintf(`{"action":"%s"}`, "opened"),
		})

		// create service
		svc, appLayer := newMockService(t)

		// mock
		appLayer.
			EXPECT().
			ProcessNewIssue(c.Request().Context(), gomock.Any()).
			Return(nil)

		// act
		err := svc.PostIssues(c)

		// assert
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("error in the app layer", func(t *testing.T) {
		// create echo context
		c, _ := newEchoContext(&newEchoContextParams{
			method:  http.MethodPost,
			target:  "/issues",
			payload: fmt.Sprintf(`{"action":"%s"}`, "opened"),
		})

		// create service
		svc, appLayer := newMockService(t)

		// mock
		appLayer.
			EXPECT().
			ProcessNewIssue(c.Request().Context(), gomock.Any()).
			Return(errors.New("something terrible happened"))

		// act
		err := svc.PostIssues(c)

		// assert
		assert.Error(t, err)
	})
}
