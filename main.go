// Copyright 2025 Sudo Sweden AB
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/sudoswedenab/terraform-provider-sudo/internal/provider"
)

var (
	version string = "dev"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/sudoswedenab/sudo",
		Debug:   false,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
