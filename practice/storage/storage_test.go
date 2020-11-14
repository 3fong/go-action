package storage

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	const user = "liu"
	usage[user] = 990000000 // simulate a 980MB-used condition

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}

// A poor assertion function.
func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}
func TestSplit(t *testing.T) {
	words := strings.Split("a:b:c", ":")
	assertEqual(len(words), 3)
	// ...
}

func TestSplit2(t *testing.T) {
	s, sep := "a:b:c:D", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q, %q) returned %d words, want %d",
			s, sep, got, want)
	}
	// ...
}
