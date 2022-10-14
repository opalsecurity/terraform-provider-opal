package data_sources

data "opal_user" "alice" {
  email = "alice@mycompany.com"
}

data "opal_user" "bob" {
  id = "e5e5ba2b-e126-4699-a8bc-dc186d490b6e"
}
