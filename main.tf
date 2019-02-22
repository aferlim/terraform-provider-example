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
  name             = "Piloto WebPremios Catalogo"
  project_id       = "${example_campaign.piloto.id}"
  conversion_rate  = 20
  external_payment = true
}

resource "example_store" "extra" {
  name                     = "Parceiro Extra Catalogo"
  description              = "Parceiro Extra Catalogo Description"
  vendor_id                = 59
  project_configuration_id = "${example_catalog.catalogopiloto.id}"
  visible                  = 1
}

resource "example_store" "pontofrio" {
  name                     = "Parceiro Ponto Frio Catalogo"
  description              = "Parceiro Ponto Frio Catalogo Description"
  vendor_id                = 58
  project_configuration_id = "${example_catalog.catalogopiloto.id}"
  visible                  = 1
}

resource "example_store" "casasbahia" {
  name                     = "Parceiro CasasBahia Catalogo"
  description              = "Parceiro Casas Bahia Catalogo Description"
  vendor_id                = 60
  project_configuration_id = "${example_catalog.catalogopiloto.id}"
  visible                  = 1
}

resource "example_participant" "bufoni" {
  name                     = "Vinicius Bufoni"
  login                    = "vinicius.bufoni"
  email                    = "vinicius.bufoni@grupoltm.com.br"
  password                 = "123456"
  customer_id              = 88
  project_id               = "${example_campaign.piloto.id}"
  project_configuration_id = "${example_catalog.catalogopiloto.id}"
  active                   = true
}

resource "example_participant" "naldo" {
  name                     = "Andre Lima"
  login                    = "andre.lima"
  email                    = "andre.lima@grupoltm.com.br"
  password                 = "123456"
  customer_id              = 88
  project_id               = "${example_campaign.piloto.id}"
  project_configuration_id = "${example_catalog.catalogopiloto.id}"
  active                   = true
}
