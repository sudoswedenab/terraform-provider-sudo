package main

import (
	"context"
	"fmt"
	"os"

	"bitbucket.org/sudosweden/terraform-provider-sudo/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/sudosweden/sudo",
		Debug:   false,
	}

	err := providerserver.Serve(context.Background(), provider.New("0.1.0"), opts)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
