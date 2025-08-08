// Copyright 2021 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"time"
)

type AuditLogStreamConfig struct {
	Enabled        bool            `json:"enabled,omitempty"`
	StreamType     string          `json:"stream_type,omitempty"` // e.g. "active" or "inactive"
	VendorSpecific AzureBlobConfig `json:"vendor_specific,omitempty"`
}

type AzureBlobConfig struct {
	KeyID           string `json:"key_id"`
	EncryptedSASURL string `json:"encrypted_sas_url"`
}

type AuditLogStreamEntry struct {
	ID             int             `json:"id,omitempty"`
	StreamType     string          `json:"stream_type"`
	StreamDetails  string          `json:"stream_details,omitempty"`
	Enabled        bool            `json:"enabled"`
	CreatedAt      time.Time       `json:"created_at,omitempty"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty"`
	PausedAt       *time.Time      `json:"paused_at,omitempty"`
	VendorSpecific AzureBlobConfig `json:"vendor_specific"`
}

// CreateAuditLogStream creates an audit log stream
//
// GitHub API docs: https://docs.github.com/enterprise-cloud@latest/rest/enterprise-admin/audit-log#create-an-audit-log-streaming-configuration-for-an-enterprise
//
//meta:operation POST /enterprises/{enterprise}/audit-log/streams
func (s *EnterpriseService) CreateAuditLogStream(ctx context.Context, enterprise string, config *AuditLogStreamConfig) (*AuditLogStreamEntry, *Response, error) {
	u := fmt.Sprintf("enterprises/%v/audit-log/streams", enterprise)

	req, err := s.client.NewRequest("POST", u, config)
	if err != nil {
		return nil, nil, err
	}

	auditLogStreamEntry := new(AuditLogStreamEntry)
	resp, err := s.client.Do(ctx, req, auditLogStreamEntry)
	if err != nil {
		return nil, resp, err
	}
	return auditLogStreamEntry, resp, nil
}

// ListAuditLogStreams lists all audit log streams
//
// GitHub API docs: https://docs.github.com/enterprise-cloud@latest/rest/enterprise-admin/audit-log#list-audit-log-stream-configurations-for-an-enterprise
//
//meta:operation GET /enterprises/{enterprise}/audit-log/streams
func (s *EnterpriseService) ListAuditLogStreams(ctx context.Context, enterprise string) ([]*AuditLogStreamEntry, *Response, error) {
	u := fmt.Sprintf("enterprises/%v/audit-log/streams", enterprise)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var streams []*AuditLogStreamEntry
	resp, err := s.client.Do(ctx, req, &streams)
	if err != nil {
		return nil, resp, err
	}
	return streams, resp, nil
}

// GetAuditLogStream returns a single audit log stream by ID.
//
// GitHub API docs: https://docs.github.com/enterprise-cloud@latest/rest/enterprise-admin/audit-log#list-one-audit-log-streaming-configuration-via-a-stream-id
//
//meta:operation GET /enterprises/{enterprise}/audit-log/streams/{stream_id}
func (s *EnterpriseService) GetAuditLogStream(ctx context.Context, enterprise string, streamID int) (*AuditLogStreamEntry, *Response, error) {
	u := fmt.Sprintf("enterprises/%v/audit-log/streams/%d", enterprise, streamID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	var stream AuditLogStreamEntry
	resp, err := s.client.Do(ctx, req, &stream)
	if err != nil {
		return nil, resp, err
	}
	return &stream, resp, nil
}

// DeleteAuditLogStream deletes an audit log streams
//
// GitHub API docs: https://docs.github.com/enterprise-cloud@latest/rest/enterprise-admin/audit-log#delete-an-audit-log-streaming-configuration-for-an-enterprise
//
//meta:operation DELETE /enterprises/{enterprise}/audit-log/streams/{stream_id}
func (s *EnterpriseService) DeleteAuditLogStream(ctx context.Context, enterprise string, streamID int) (*Response, error) {
	u := fmt.Sprintf("enterprises/%v/audit-log/streams/%d", enterprise, streamID)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
