package main

import (
	"github.com/steinfletcher/api-test"
	"net/http"
	"testing"
)

func TestGetUser_Success(t *testing.T) {
	apitest.New("gets the user").
		Mocks(getPreferencesMock, getUserMock).
		Handler(newApp().Router).
		Get("/user").
		Host("user-service").
		Query("name", "jan").
		QueryCollection(map[string][]string{"hobbies": {"gardening", "knitting"}}).
		Expect(t).
		Status(http.StatusOK).
		Header("Content-Type", "application/json").
		Body(`{"name": "jon", "is_contactable": true}`).
		Report()
}

var getPreferencesMock = apitest.NewMock().
	Get("http://3rdparty.com/preferences/12345").
	RespondWith().
	Body(`{"is_contactable": true}`).
	Status(http.StatusOK).
	End()

var getUserMock = apitest.NewMock().
	Get("http://3rdparty.com/user/12345").
	RespondWith().
	Body(`{"name": "jon", "id": "1234"}`).
	Status(http.StatusAccepted).
	End()