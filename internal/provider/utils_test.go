package provider

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"strings"
	"testing"
)

// TestDebugResponseRedactsSensitiveHeaders ensures debugResponse never emits
// the plaintext value of any header listed in sensitiveHeaders (e.g. the
// Cloudflare Access service token headers, Pylon #604 / ENG-2597).
func TestDebugResponseRedactsSensitiveHeaders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://example.com/resource", nil)
	req.Header.Set("Authorization", "Bearer super-secret-token")
	req.Header.Set("Cf-Access-Client-Id", "cf-client-id-value")
	req.Header.Set("Cf-Access-Client-Secret", "cf-client-secret-value")
	req.Header.Set("X-Other-Header", "not-sensitive")

	res := &http.Response{
		Status:     "404 Not Found",
		StatusCode: http.StatusNotFound,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
		Body:       http.NoBody,
		Request:    req,
	}

	out := debugResponse(res)

	for _, secret := range []string{"super-secret-token", "cf-client-id-value", "cf-client-secret-value"} {
		if strings.Contains(out, secret) {
			t.Errorf("debugResponse output leaked sensitive value %q:\n%s", secret, out)
		}
	}
	if !strings.Contains(out, "not-sensitive") {
		t.Errorf("debugResponse unexpectedly redacted a non-sensitive header:\n%s", out)
	}
	if got := req.Header.Get("Authorization"); got != "(sensitive)" {
		t.Errorf("Authorization header not redacted in-place, got %q", got)
	}
	if got := req.Header.Get("Cf-Access-Client-Id"); got != "(sensitive)" {
		t.Errorf("Cf-Access-Client-Id header not redacted in-place, got %q", got)
	}
	if got := req.Header.Get("Cf-Access-Client-Secret"); got != "(sensitive)" {
		t.Errorf("Cf-Access-Client-Secret header not redacted in-place, got %q", got)
	}
}

// TestFieldHeadersFromRequestReaderRedactsSensitiveHeaders ensures the
// tflog debug fields never carry the plaintext value of a sensitive header.
func TestFieldHeadersFromRequestReaderRedactsSensitiveHeaders(t *testing.T) {
	raw := "GET /resource HTTP/1.1\r\n" +
		"Authorization: Bearer super-secret-token\r\n" +
		"Cf-Access-Client-Id: cf-client-id-value\r\n" +
		"Cf-Access-Client-Secret: cf-client-secret-value\r\n" +
		"X-Other-Header: not-sensitive\r\n" +
		"\r\n"

	reader := textproto.NewReader(bufio.NewReader(bytes.NewReader([]byte(raw))))
	fields := make(map[string]interface{})

	if err := fieldHeadersFromRequestReader(reader, fields); err != nil {
		t.Fatalf("fieldHeadersFromRequestReader returned error: %v", err)
	}

	for _, h := range []string{"Authorization", "Cf-Access-Client-Id", "Cf-Access-Client-Secret"} {
		if got, _ := fields[h].(string); got != "(sensitive)" {
			t.Errorf("field %q not redacted, got %q", h, got)
		}
	}
	if got, _ := fields["X-Other-Header"].(string); got != "not-sensitive" {
		t.Errorf("field X-Other-Header unexpectedly redacted, got %q", got)
	}
}
