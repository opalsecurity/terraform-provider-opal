package data_sources

data "opal_owner" "design" {
  name = "Design Owner"
}

data "opal_owner" "devops" {
  id = "e5e5ba2b-e126-4699-a8bc-dc186d490b6e"
}
