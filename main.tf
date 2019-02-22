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

# resource "example_campaign" "piloto" {
#   name           = "Piloto WebPremios"
#   clientId       = 1
#   externalPoints = 1
# }


# resource "example_catalog" "catalogopiloto" {
#   name           = "Piloto WebPremios Catalogo"
#   projectId      = "${example_campaign.piloto.id}"
#   conversionRate = 20
# }


# resource "example_store" "extra" {
#   name                   = "Parceiro Extra Catalogo"
#   description            = "Parceiro Extra Catalogo"
#   vendorId               = 59
#   projectConfigurationId = "${example_catalog.catalogopiloto.id}"
#   visible                = 1
# }

