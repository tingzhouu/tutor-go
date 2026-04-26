package handler

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	strLen := 6
	str := randomString(strLen)
	if len(str) != strLen {
		t.Fatalf("Unexpceted string length %d - %s ", len(str), str)
	}

	for _, c := range str {
		if !strings.ContainsRune(Charset, c) {
			t.Fatalf("char %v is not in charset %s", c, Charset)
		}
	}
}
