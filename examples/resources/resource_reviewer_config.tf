resource "opal_resource" "not_requestable" {
  // ...

  // If you want a resource to not be requestable, you can set `is_requestable` to false and omit the `reviewer_stage` attribute
  is_requestable = false
}

resource "opal_resource" "auto_approval" {
  // ...

  // If you want a resource to be auto-approved, you can set `auto_approval` to true and omit the `reviewer_stage` attribute
  auto_approval = false
}

resource "opal_resource" "basic_reviewer_config" {
  // ...

  // NOTE: operator = "AND" and require_manager_approval = false are the default if not explicitly set
  reviewer_stage {
    reviewer {
      id = opal_owner.security.id
    }
  }
}

resource "opal_resource" "or_reviewer_config" {
  // ...

  // Here the manager of the requesting user or the security owner would need to approve
  reviewer_stage {
    operator = "OR"
    require_manager_approval = true
    reviewer {
      id = opal_owner.security.id
    }
  }
}

resource "opal_resource" "complex_reviewer_config" {
  // ...

  // Here first the manager has to approve. Once the manager has approved, both the security owner and the data owner need to approve
  // NOTE: The ordering determines the ordering of the stages
  reviewer_stage {
    require_manager_approval = true
  }

  reviewer_stage {
    reviewer {
      id = opal_owner.security.id
    }

    reviewer {
      id = opal_owner.data.id
    }
  }
}
