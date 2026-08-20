package storage

func (s *AuditBatchStore) Snapshot() []string { return s.committed }
func (s *AuditBatchStore) Count() int         { return len(s.committed) }
