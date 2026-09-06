package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/Shang47/simple-web-server"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}
func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
func TestCLI(t *testing.T) {
	t.Run("finish game with 'Chris' as winner", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.StartedWith != 1 {
			t.Errorf("wanted Start called with 1 but got %d", game.StartedWith)
		}
		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})
	t.Run("record 'cleo' win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		if game.StartedWith != 5 {
			t.Errorf("wanted Start called with 5 but got %d", game.StartedWith)
		}
		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishedWith)
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout.String()

		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
	})
}
func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []scheduledAlert{
			{at: 0 * time.Second, amount: 100},
			{at: 10 * time.Minute, amount: 200},
			{at: 20 * time.Minute, amount: 300},
			{at: 30 * time.Minute, amount: 400},
			{at: 40 * time.Minute, amount: 500},
			{at: 50 * time.Minute, amount: 600},
			{at: 60 * time.Minute, amount: 800},
			{at: 70 * time.Minute, amount: 1000},
			{at: 80 * time.Minute, amount: 2000},
			{at: 90 * time.Minute, amount: 4000},
			{at: 100 * time.Minute, amount: 8000},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedule alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []scheduledAlert{
			{at: 0 * time.Second, amount: 100},
			{at: 12 * time.Minute, amount: 200},
			{at: 24 * time.Minute, amount: 300},
			{at: 36 * time.Minute, amount: 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
}
func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewGame(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}
func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
func checkSchedulingCases(cases []scheduledAlert, t *testing.T, alerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}
func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
