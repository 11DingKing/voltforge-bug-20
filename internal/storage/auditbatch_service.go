package storage

import "errors"

var ErrAuditBatchRejected = errors.New("auditbatch batch rejected")

type AuditBatchStore struct{ committed []string }

func (s *AuditBatchStore) Apply(values []string) error {
	for _, value := range values {
		if value == "bad" {
			return ErrAuditBatchRejected
		}
		s.committed = append(s.committed, value)
	}
	return nil
}
