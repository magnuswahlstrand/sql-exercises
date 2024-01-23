package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	url2 "net/url"
	"testing"
)

func requireBody(t *testing.T, expectedBody string, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	require.EqualValues(t, expectedBody, string(body))
}

func CheckQuery(t *testing.T, exerciseId string, query string, app *fiber.App) error {
	url := fmt.Sprintf("/check/%s?query=%s", exerciseId, url2.QueryEscape(query))
	req := httptest.NewRequest("GET", url, nil)

	// Act
	resp, err := app.Test(req, 1)
	require.Nil(t, err)

	// Assert
	requireBody(t, "foo", resp)
	require.Equal(t, 200, resp.StatusCode, resp.Body)
	return err
}

func TestApp(t *testing.T) {
	app := SetupApp()
	query := "SELECT * FROM employees"
	if err := CheckQuery(t, "select_all", query, app); err != nil {
		panic(err)
	}

}
