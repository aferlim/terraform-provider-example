provider "example" {
  address = "http://iac-items.getsandbox.com"
  port    = 80
  token   = "superSecretToken"
}

resource "example_item" "test" {
  name        = "item_via_terraform"
  description = "this is an item created by terraform"

  tags = [
    "hello!",
    "terraform",
  ]
}
