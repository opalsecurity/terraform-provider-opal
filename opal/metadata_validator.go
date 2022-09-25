package opal

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

func validateResourceMetadata(i interface{}, k string) ([]string, []error) {
	sl := gojsonschema.NewSchemaLoader()

	schemaJson := gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "Resource Metadata",
  "properties": {
    "aws_ec2_instance": {
      "properties": {
        "instance_id": {
          "type": "string"
        },
        "region": {
          "type": "string"
        }
      },
      "required": ["instance_id", "region"],
      "additionalProperties": false,
      "type": "object",
      "title": "AWS EC2 Instance"
    },
    "aws_eks_cluster": {
      "properties": {
        "cluster_name": {
          "type": "string"
        },
        "cluster_region": {
          "type": "string"
        },
        "cluster_arn": {
          "type": "string"
        }
      },
      "required": ["cluster_name", "cluster_region", "cluster_arn"],
      "additionalProperties": false,
      "type": "object",
      "title": "AWS EKS Cluster"
    },
    "aws_rds_instance": {
      "properties": {
        "instance_id": {
          "type": "string"
        },
        "engine": {
          "type": "string"
        },
        "region": {
          "type": "string"
        },
        "resource_id": {
          "type": "string"
        },
        "database_name": {
          "type": "string"
        }
      },
      "required": [
        "instance_id",
        "engine",
        "region",
        "resource_id",
        "database_name"
      ],
      "additionalProperties": false,
      "type": "object",
      "title": "AWS RDS Instance"
    },
    "aws_role": {
      "properties": {
        "arn": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      },
      "required": ["arn", "name"],
      "additionalProperties": false,
      "type": "object",
      "title": "AWS Role"
    },
    "gcp_bucket": {
      "properties": {
        "bucket_id": {
          "type": "string"
        }
      },
      "required": ["bucket_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "GCP Bucket"
    },
    "gcp_compute_instance": {
      "properties": {
        "instance_id": {
          "type": "string"
        },
        "project_id": {
          "type": "string"
        },
        "zone": {
          "type": "string"
        }
      },
      "required": ["instance_id", "project_id", "zone"],
      "additionalProperties": false,
      "type": "object",
      "title": "GCP Compute Instance"
    },
    "gcp_folder": {
      "properties": {
        "folder_id": {
          "type": "string"
        }
      },
      "required": ["folder_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "GCP Folder"
    },
    "gcp_gke_cluster": {
      "properties": {
        "cluster_name": {
          "type": "string"
        }
      },
      "required": ["cluster_name"],
      "additionalProperties": false,
      "type": "object",
      "title": "GCP GKE Cluster"
    },
    "gcp_project": {
      "properties": {
        "project_id": {
          "type": "string"
        }
      },
      "required": ["project_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "GCP Project"
    },
    "gcp_sql_instance": {
      "properties": {
        "instance_id": {
          "type": "string"
        },
        "project_id": {
          "type": "string"
        }
      },
      "required": ["instance_id", "project_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "GCP SQL Instance"
    },
    "git_hub_repo": {
      "properties": {
        "org_name": {
          "type": "string"
        },
        "repo_name": {
          "type": "string"
        }
      },
      "required": ["org_name", "repo_name"],
      "additionalProperties": false,
      "type": "object",
      "title": "GitHub Repo"
    },
    "okta_directory_app": {
      "properties": {
        "app_id": {
          "type": "string"
        },
        "logo_url": {
          "type": "string"
        }
      },
      "required": ["app_id", "logo_url"],
      "additionalProperties": false,
      "type": "object",
      "title": "Okta Directory App"
    },
    "okta_directory_role": {
      "properties": {
        "role_type": {
          "type": "string"
        },
        "role_id": {
          "type": "string"
        }
      },
      "required": ["role_type", "role_id"],
      "additionalProperties": false,
      "type": "object",
      "title": "Okta Directory Role"
    },
    "salesforce_profile": {
      "properties": {
        "user_license": {
          "type": "string"
        }
      },
      "required": ["user_license"],
      "additionalProperties": false,
      "type": "object",
      "title": "Salesforce Profile"
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
