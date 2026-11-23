resource "opal_request_template" "my_requesttemplate" {
  custom_fields = [
    {
      description = "...my_description..."
      metadata = {
        callout_data = {
          severity = "WARNING"
          text     = "This role grants access to production customer data."
        }
        multi_choice_data = {
          options = [
            "Incident response",
            "Scheduled maintenance",
          ]
        }
      }
      name     = "Why do you need this access?"
      required = true
      type     = "CALLOUT"
    }
  ]
  name = "Production access questions"
}