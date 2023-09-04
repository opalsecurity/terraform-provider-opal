resource "opal_group" "not_requestable" {
  // ...

  // If you want a group to not be requestable, you can set `is_requestable` to false and omit the `reviewer_stage` attribute
  request_configuration {
    is_requestable = false
  }
}

resource "opal_group" "auto_approval" {
  // ...

  // If you want a group to be auto-approved, you can set `auto_approval` to true and omit the `reviewer_stage` attribute
  request_configuration {
    auto_approval = false
  }
}

resource "opal_group" "basic_reviewer_config" {
  // ...

  // NOTE: operator = "AND" and require_manager_approval = false are the default if not explicitly set
  request_configuration {
    reviewer_stage {
      reviewer {
        id = opal_owner.security.id
      }
    }
  }
}

resource "opal_group" "or_reviewer_config" {
  // ...

  // Here the manager of the requesting user or the security owner would need to approve
  request_configuration {
    reviewer_stage {
      operator = "OR"
      require_manager_approval = true
      reviewer {
        id = opal_owner.security.id
      }
    }
  }
}

resource "opal_group" "complex_reviewer_config" {
  // ...

  // Here first the manager has to approve. Once the manager has approved, both the security owner and the data owner need to approve
  // NOTE: The ordering determines the ordering of the stages

  request_configuration {
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
}
