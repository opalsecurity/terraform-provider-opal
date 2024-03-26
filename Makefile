.PHONY: *

all: speakeasy docs

docs:
	go generate ./...


speakeasy: check-speakeasy openapi.yaml
	speakeasy generate sdk --lang terraform -o . -s ./openapi.yaml

speakeasy-validate: check-speakeasy
	speakeasy validate openapi -s ./openapi.yaml

terraform_overlay.yaml: check-speakeasy
	speakeasy overlay compare -s ./openapi_original.yaml -s ./openapi.yaml > ./terraform_overlay.yaml

openapi.yaml: check-speakeasy
	speakeasy overlay apply -s ./openapi_original.yaml -o ./terraform_overlay.yaml > ./openapi.yaml

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }
