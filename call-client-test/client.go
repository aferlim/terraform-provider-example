package main

import iacitem "github.com/aferlim/terraform-provider-example/client/iac-item"

func main() {

	address := " http://iac-items.getsandbox.com/"
	port := 80
	token := "test"

	var newClient iacitem.Client = iacitem.NewClient(address, port, token)

	//newClient.

}
