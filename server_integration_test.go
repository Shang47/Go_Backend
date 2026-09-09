package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/Shang47/simple-web-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	//t.Run("test with in memory player store", func(t *testing.T) {
	store := poker.NewInMemoryPlayerStore()
	server := mustMakePlayerServer(t, store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(player))
		poker.AssertStatus(t, response.Code, http.StatusOK)

		poker.AssertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewLeagueRequest())
		poker.AssertStatus(t, response.Code, http.StatusOK)

		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
	//})
	/*t.Run("test with player store in MariaDB", func(t *testing.T) {
		store := NewMariaPlayerStore()
		server := PlayerServer{store}
		player := "Pepper"

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "21")

		// Test insert new player
		player = "Steve"
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response = httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "1")
	})*/
}
