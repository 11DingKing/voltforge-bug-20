package storage

import "testing"

func TestVoltForge20(t *testing.T) {
	store := &AuditBatchStore{}
	if err := store.Apply([]string{"session-1", "bad", "session-3"}); err == nil {
		t.Fatal("expected rejection")
	}
	if got := store.Count(); got != 0 {
		t.Fatalf("partial commit leaked %d records", got)
	}
}
