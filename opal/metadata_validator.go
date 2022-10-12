package opal

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

func validateGroupMetadata(i interface{}, k string) ([]string, []error) {
	sl := gojsonschema.NewSchemaLoader()

	schemaJson := gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "Group Metadata",
  "properties": {
    "ad_group": {
      "properties": {
        "object_guid": {
          "type": "string"
        }
      },
      "required": ["object_guid"],
      "additionalProperties": false,
      "type": "object",
      "title": "Active Directory Group"
    },
    "duo_group": {
      "properties": {
        "group_id": {
          "type": "string"
        }
      },
      "required": ["group_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "Duo Group"
    },
    "git_hub_team": {
      "properties": {
        "org_name": {
          "type": "string"
        },
        "team_slug": {
          "type": "string"
        }
      },
      "required": ["org_name", "team_slug"],
      "additionalProperties": false,
      "type": "object",
      "title": "GitHub Team"
    },
    "google_groups_group": {
      "properties": {
        "group_id": {
          "type": "string"
        }
      },
      "required": ["group_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "Google Groups Group"
    },
    "ldap_group": {
      "properties": {
        "group_uid": {
          "type": "string"
        }
      },
      "required": ["group_uid"],
      "additionalProperties": false,
      "type": "object",
      "title": "LDAP Group"
    },
    "okta_directory_group": {
      "properties": {
        "group_id": {
          "type": "string"
        }
      },
      "required": ["group_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "Okta Directory Group"
    }
  },
  "additionalProperties": false,
  "minProperties": 1,
  "maxProperties": 1,
  "type": "object"
}`)

	schema, err := sl.Compile(schemaJson)
	if err != nil {
		return nil, []error{err}
	}

	documentJson := gojsonschema.NewStringLoader(i.(string))

	result, err := schema.Validate(documentJson)
	if err != nil {
		return nil, []error{err}
	}

	errs := make([]error, 0, len(result.Errors()))
	for _, err := range result.Errors() {
		errs = append(errs, errors.New(err.String()))
	}
	return nil, errs
}
