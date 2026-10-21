# End-to-end tests

Live-backend tests that exercise the generated Opal SDK against a running
Opal API. Unlike the `TestAcc*` tests in `internal/provider/`, these tests do
not go through Terraform plan/apply — they call the generated SDK methods
directly, which makes them the shortest path to verifying that the openapi
spec and the generated SDK round-trip against the real API.

Suites live alongside their SDK surface. Today:

- `configuration_template_sub_resource_test.go` — covers `GET` and `PUT
  /resources/{resource_id}/configuration-template` ([EPRD-4793]).

Each sub-test creates its own fixtures (a CUSTOM resource, one or two
configuration templates) and tears them down on exit, so the suite is safe to
re-run and does not leave orphaned test data behind.

## Running

Tests are gated by the `e2e` build tag and skipped unless `OPAL_E2E_TOKEN` is
set, so `go test ./...` and the CI unit-test job are unaffected.

```bash
# against the local mprocs backend
export OPAL_E2E_TOKEN=<token>
go test -tags=e2e -v ./internal/e2e/...

# against a remote environment
export OPAL_E2E_BASE_URL=https://api.opal.dev/v1
export OPAL_E2E_TOKEN=<token>
go test -tags=e2e -v ./internal/e2e/...
```

The `Makefile` wraps this as `make e2e`.

## Configuration

| Env var | Default | Purpose |
| --- | --- | --- |
| `OPAL_E2E_TOKEN` | *(required)* | Bearer token for the Opal API. |
| `OPAL_E2E_BASE_URL` | `http://localhost:3000/v1` | Base URL (include the `/v1` segment). |
| `OPAL_E2E_APP_ID` | `849aa50a-5869-4988-8396-e8aee2f4a538` | CUSTOM-type app that owns the test resources. |
| `OPAL_E2E_ADMIN_OWNER_ID` | `ee664aad-b2f5-4011-9e19-32776613825a` | Admin owner used for created templates. |

The defaults match the `baseline-1792208657` CUSTOM app seeded in the
local dev DB. Point `OPAL_E2E_APP_ID` / `OPAL_E2E_ADMIN_OWNER_ID` at an app and
owner you control before running against a shared environment.
