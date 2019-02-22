provider "example" {
  address = "http://iac-items.getsandbox.com"
  token   = "superSecretToken"
}

resource "example_item" "kaizen" {
  name        = "item_via_terraform_kaizen"
  description = "this is an item created by terraform "

  tags = [
    "hello!",
    "terraform",
  ]
}

resource "example_item" "second" {
  name        = "item_via_terraform_seccont_id"
  description = "this is an item created by terraform "

  tags = [
    "hello!",
    "terraform",
  ]
}

resource "example_campaign" "piloto" {
  name            = "Piloto WebPremios"
  client_id       = 1
  external_points = 1
}

resource "example_catalog" "catalogopiloto" {
  name           = "Piloto WebPremios Catalogo"
  project_id     = "${example_campaign.piloto.id}"
  conversionRate = 20
}

resource "example_store" "extra" {
  name                     = "Parceiro Extra Catalogo"
  description              = "Parceiro Extra Catalogo Description"
  vendor_id                = 59
  project_configuration_id = "${example_catalog.catalogopiloto.id}"
  visible                  = 1
}
