provider "example" {
  address = "http://iac-items.getsandbox.com"
  token   = "superSecretToken"
}

resource "example_item" "test" {
  name        = "item_via_terraform_id"
  description = "this is an item created by terraform chenget"

  tags = [
    "hello!",
    "terraform",
  ]
}

resource "example_item" "second" {
  name        = "item_via_terraform_seccont_id"
  description = "this is an item created by terraform changed"

  tags = [
    "hello! 2",
    "terraform 2 ",
  ]
}
