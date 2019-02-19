provider "example" {
  address = "http://iac-items.getsandbox.com"
  token   = "superSecretToken"
}

resource "example_item" "test" {
  name        = "item_via_terraform"
  description = "this is an item created by terraform changed"

  tags = [
    "hello!",
    "terraform",
  ]
}

resource "example_item" "second" {
  name        = "item_via_terraform"
  description = "this is an item created by terraform changed"

  tags = [
    "hello!",
    "terraform",
  ]
}
