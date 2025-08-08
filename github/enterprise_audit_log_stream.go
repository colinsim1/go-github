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

const (
	StreamTypeAzureBlob      = "Azure Blob Storage"
	StreamTypeAzureEventHubs = "Azure Event Hubs"
	StreamTypeAmazonS3       = "Amazon S3"
	StreamTypeSplunk         = "Splunk"
	StreamTypeHEC            = "HTTPS Event Collector"
	StreamTypeGCS            = "Google Cloud Storage"
	StreamTypeDatadog        = "Datadog"
)

type (
	AzureBlobConfig struct {
		KeyID           string `json:"key_id"`
		EncryptedSASURL string `json:"encrypted_sas_url"`
	}
	AzureEventHubsConfig struct {
		Name                string `json:"name"`
		EncryptedConnstring string `json:"encrypted_connstring"`
		KeyID               string `json:"key_id"`
	}
	AmazonS3OIDCConfig struct {
		Bucket             string `json:"bucket"`
		Region             string `json:"region"`
		KeyID              string `json:"key_id"`
		AuthenticationType string `json:"authentication_type"`
		ArnRole            string `json:"arn_role"`
	}
	AmazonS3AccessKeysConfig struct {
		Bucket               string `json:"bucket"`
		Region               string `json:"region"`
		KeyID                string `json:"key_id"`
		AuthenticationType   string `json:"authentication_type"`
		EncryptedAccessKeyID string `json:"encrypted_access_key_id"`
		EncryptedSecretKey   string `json:"encrypted_secret_key"`
	}
	SplunkConfig struct {
		Domain         string `json:"domain"`
		Port           int    `json:"port"`
		KeyID          string `json:"key_id"`
		EncryptedToken string `json:"encrypted_token"`
		SSLVerify      bool   `json:"ssl_verify"`
	}
	HECConfig struct {
		Domain         string `json:"domain"`
		Port           int    `json:"port"`
		Path           string `json:"path"`
		KeyID          string `json:"key_id"`
		EncryptedToken string `json:"encrypted_token"`
		SSLVerify      bool   `json:"ssl_verify"`
	}
	GCSConfig struct {
		Bucket                   string `json:"bucket"`
		KeyID                    string `json:"key_id"`
		EncryptedJSONCredentials string `json:"encrypted_json_credentials"`
	}
	DatadogConfig struct {
		KeyID          string `json:"key_id"`
		EncryptedToken string `json:"encrypted_token"`
		Site           string `json:"site"`
	}
)

// AuditLogStreamConfig is the request payload.
//
// VendorSpecific is intentionally interface{} to stay idiomatic with go-github’s
// lightweight API surface. You can pass one of the concrete vendor structs below,
// or your own map[string]any with the expected JSON fields for the chosen stream type.
type AuditLogStreamConfig struct {
	Enabled        bool        `json:"enabled,omitempty"`
	StreamType     string      `json:"stream_type,omitempty"`
	VendorSpecific interface{} `json:"vendor_specific,omitempty"`
}

// AuditLogStreamEntry is the response payload.
//
// VendorSpecific is an interface{} to match the API’s polymorphic shape. Use
// UnmarshalVendorSpecific to decode it into a concrete struct if desired.
type AuditLogStreamEntry struct {
	ID             int         `json:"id,omitempty"`
	StreamType     string      `json:"stream_type"`
	StreamDetails  string      `json:"stream_details,omitempty"`
	Enabled        bool        `json:"enabled"`
	CreatedAt      time.Time   `json:"created_at,omitempty"`
	UpdatedAt      time.Time   `json:"updated_at,omitempty"`
	PausedAt       *time.Time  `json:"paused_at,omitempty"`
	VendorSpecific interface{} `json:"vendor_specific"`
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

	out := new(AuditLogStreamEntry)
	resp, err := s.client.Do(ctx, req, out)
	if err != nil {
		return nil, resp, err
	}
	return out, resp, nil
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

	out := new(AuditLogStreamEntry)
	resp, err := s.client.Do(ctx, req, out)
	if err != nil {
		return nil, resp, err
	}
	return out, resp, nil
}

// DeleteAuditLogStream deletes an audit log stream
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
