package network_test

import (
	"fmt"
	"testing"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/lxc/terraform-provider-incus/internal/acctest"
)

func TestAccNetworkForward_basic(t *testing.T) {
	networkName := getNetworkName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkForward(networkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("incus_network_forward.forward", "listen_address", "10.150.19.10"),
					resource.TestCheckResourceAttr("incus_network_forward.forward", "description", "Network Forward"),
					resource.TestCheckResourceAttr("incus_network_forward.forward", "config.target_address", "10.150.19.111"),
				),
			},
		},
	})
}

func TestAccNetworkForward_Ports(t *testing.T) {
	networkName := getNetworkName()

	entry1 := map[string]string{
		"description":    "SSH",
		"protocol":       "tcp",
		"listen_port":    "22",
		"target_port":    "2022",
		"target_address": "10.150.19.112",
	}

	entry2 := map[string]string{
		"description":    "HTTP",
		"protocol":       "tcp",
		"listen_port":    "80",
		"target_port":    "8080",
		"target_address": "10.150.19.112",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkForward_withPorts(networkName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("incus_network_forward.forward", "listen_address", "10.150.19.10"),
					resource.TestCheckResourceAttr("incus_network_forward.forward", "description", "Network Forward"),
					resource.TestCheckResourceAttr("incus_network_forward.forward", "config.target_address", "10.150.19.111"),
					resource.TestCheckTypeSetElemNestedAttrs("incus_network_forward.forward", "ports.*", entry1),
					resource.TestCheckTypeSetElemNestedAttrs("incus_network_forward.forward", "ports.*", entry2),
				),
			},
		},
	})
}

func testAccNetworkForward(networkName string) string {
	return fmt.Sprintf(`
resource "incus_network" "forward" {
  name = "%s"

  config = {
    "ipv4.address" = "10.150.19.1/24"
    "ipv4.nat"     = "true"
    "ipv6.address" = "fd42:474b:622d:259d::1/64"
    "ipv6.nat"     = "true"
  }
}

resource "incus_network_forward" "forward" {
  network        = incus_network.forward.name
  description    = "Network Forward"
  listen_address = "10.150.19.10"

  config = {
	target_address = "10.150.19.111"
  }
}
`, networkName)
}

func testAccNetworkForward_withPorts(networkName string) string {
	return fmt.Sprintf(`
resource "incus_network" "forward" {
  name = "%s"

  config = {
    "ipv4.address" = "10.150.19.1/24"
    "ipv4.nat"     = "true"
    "ipv6.address" = "fd42:474b:622d:259d::1/64"
    "ipv6.nat"     = "true"
  }
}

resource "incus_network_forward" "forward" {
  network        = incus_network.forward.name
  description    = "Network Forward"
  listen_address = "10.150.19.10"

  config = {
	target_address = "10.150.19.111"
  }

  ports = [
    {
      description    = "SSH"
      protocol       = "tcp"
      listen_port    = "22"
      target_port    = "2022"
      target_address = "10.150.19.112"
    },
    {
      description    = "HTTP"
      protocol       = "tcp"
      listen_port    = "80"
      target_port    = "8080"
      target_address = "10.150.19.112"
    }
  ]
}
`, networkName)
}

func getNetworkName() string {
	maxLength := 15
	networkName := petname.Generate(2, "-")

	if len(networkName) > maxLength {
		return networkName[:maxLength]
	}

	return networkName
}
