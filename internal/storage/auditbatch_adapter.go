package storage

func (s *AuditBatchStore) Snapshot() []string {
	out := make([]string, len(s.committed))
	copy(out, s.committed)
	return out
}
func (s *AuditBatchStore) Count() int { return len(s.committed) }
