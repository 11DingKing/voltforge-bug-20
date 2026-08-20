package storage

import "errors"

var ErrAuditBatchRejected = errors.New("auditbatch batch rejected")

type AuditBatchStore struct{ committed []string }

func (s *AuditBatchStore) Apply(values []string) error {
	pending := make([]string, 0, len(values))
	for _, value := range values {
		if value == "bad" {
			return ErrAuditBatchRejected
		}
		pending = append(pending, value)
	}
	s.committed = append(s.committed, pending...)
	return nil
}
