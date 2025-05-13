package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/sudoswedenab/terraform-provider-sudo/internal/provider"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/sudoswedenab/sudo",
		Debug:   false,
	}

	err := providerserver.Serve(context.Background(), provider.New("0.1.0"), opts)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
