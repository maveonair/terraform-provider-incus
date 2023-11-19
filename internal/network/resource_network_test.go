package network_test

import (
	"fmt"
	"testing"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/terraform-lxd/terraform-provider-lxd/internal/acctest"
)

func TestAccNetwork_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetwork_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "type", "bridge"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "managed", "true"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "description", ""),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.%", "1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv4.address", "10.150.19.1/24"),
				),
			},
		},
	})
}

func TestAccNetwork_description(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetwork_desc(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "type", "bridge"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "description", "My network"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.%", "2"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv4.address", "10.150.19.1/24"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv6.address", "fd42:474b:622d:259d::1/64"),
				),
			},
		},
	})
}

func TestAccNetwork_attach(t *testing.T) {
	profileName := petname.Generate(2, "-")
	instanceName := petname.Generate(2, "-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetwork_attach(profileName, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_profile.profile1", "name", profileName),
					resource.TestCheckResourceAttr("lxd_profile.profile1", "device.#", "1"),
					resource.TestCheckResourceAttr("lxd_profile.profile1", "device.0.name", "eth1"),
					resource.TestCheckResourceAttr("lxd_profile.profile1", "device.0.type", "nic"),
					resource.TestCheckResourceAttr("lxd_profile.profile1", "device.0.properties.parent", "eth1"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "name", instanceName),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "status", "Running"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "profiles.#", "2"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "profiles.0", "default"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "profiles.1", profileName),
					resource.TestCheckResourceAttrSet("lxd_instance.instance1", "ipv4_address"),
					resource.TestCheckResourceAttrSet("lxd_instance.instance1", "ipv6_address"),
					resource.TestCheckResourceAttrSet("lxd_instance.instance1", "mac_address"),
				),
			},
		},
	})
}

func TestAccNetwork_updateConfig(t *testing.T) {
	instanceName := petname.Generate(2, "-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetwork_updateConfig_1(instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv4.address", "10.150.19.1/24"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv4.nat", "true"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "name", instanceName),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "status", "Running"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "device.#", "1"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "device.0.properties.parent", "eth1"),
					resource.TestCheckResourceAttrSet("lxd_instance.instance1", "ipv4_address"),
					resource.TestCheckResourceAttrSet("lxd_instance.instance1", "mac_address"),
				),
			},
			{
				Config: testAccNetwork_updateConfig_2(instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv4.address", "10.150.21.1/24"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.ipv4.nat", "false"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "name", instanceName),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "status", "Running"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "device.#", "1"),
					resource.TestCheckResourceAttr("lxd_instance.instance1", "device.0.properties.parent", "eth1"),
				),
			},
		},
	})
}

func TestAccNetwork_typeMacvlan(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetwork_typeMacvlan(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "type", "macvlan"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "config.parent", "lxdbr0"),
				),
			},
		},
	})
}

// TODO:
//   - Precheck for clustered mode.
// func TestAccNetwork_target(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccNetwork_target(),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("lxd_network.cluster_network_node1", "name", "cluster_network"),
// 					resource.TestCheckResourceAttr("lxd_network.cluster_network_node2", "name", "cluster_network"),
// 					resource.TestCheckResourceAttr("lxd_network.cluster_network", "name", "cluster_network"),
// 					resource.TestCheckResourceAttr("lxd_network.cluster_network", "type", "bridge"),
// 					resource.TestCheckResourceAttr("lxd_network.cluster_network", "config.ipv4.address", "10.150.19.1/24"),
// 				),
// 			},
// 		},
// 	})
// }

func TestAccNetwork_project(t *testing.T) {
	projectName := petname.Name()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetwork_project(projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_network.eth1", "name", "eth1"),
					resource.TestCheckResourceAttr("lxd_network.eth1", "project", projectName),
					resource.TestCheckResourceAttr("lxd_project.project1", "name", projectName),
				),
			},
		},
	})
}

func testAccNetwork_basic() string {
	return `
resource "lxd_network" "eth1" {
  name = "eth1"
  config = {
    "ipv4.address" = "10.150.19.1/24"
  }
}
`
}

func testAccNetwork_desc() string {
	return `
resource "lxd_network" "eth1" {
  name        = "eth1"
  description = "My network"
  config = {
    "ipv4.address" = "10.150.19.1/24"
    "ipv6.address" = "fd42:474b:622d:259d::1/64"
  }
}
`
}

func testAccNetwork_attach(profileName, instanceName string) string {
	return fmt.Sprintf(`
resource "lxd_network" "eth1" {
  name = "eth1"
  config = {
    "ipv4.address" = "10.150.19.1/24"
  }
}

resource "lxd_profile" "profile1" {
  name = "%s"

  device {
    name = "eth1"
    type = "nic"
    properties = {
      nictype = "bridged"
      parent  = lxd_network.eth1.name
    }
  }
}

resource "lxd_instance" "instance1" {
  name     = "%s"
  image    = "images:alpine/3.18"
  profiles = ["default", lxd_profile.profile1.name]
}
`, profileName, instanceName)
}

func testAccNetwork_updateConfig_1(name string) string {
	return fmt.Sprintf(`
resource "lxd_network" "eth1" {
  name = "eth1"
  config = {
    "ipv4.address" = "10.150.19.1/24"
    "ipv4.nat"     = true
  }
}

# We do need an instance here to ensure the network cannot
# be deleted, but must be updated in-place.
resource "lxd_instance" "instance1" {
  name             = "%s"
  image            = "images:alpine/3.18"
  wait_for_network = false

  device {
    name = "eth0"
    type = "nic"
    properties = {
      nictype = "bridged"
      parent  = lxd_network.eth1.name
    }
  }
}
  `, name)
}

func testAccNetwork_updateConfig_2(name string) string {
	return fmt.Sprintf(`
resource "lxd_network" "eth1" {
  name = "eth1"

  config = {
    "ipv4.address" = "10.150.21.1/24"
    "ipv4.nat"     = false
  }
}

# We do need an instance here to ensure the network cannot
# be deleted, but must be updated in-place.
resource "lxd_instance" "instance1" {
  name             = "%s"
  image            = "images:alpine/3.18"
  wait_for_network = false

  device {
    name = "eth0"
    type = "nic"
    properties = {
      nictype = "bridged"
      parent  = lxd_network.eth1.name
    }
  }
}
  `, name)
}

func testAccNetwork_typeMacvlan() string {
	return `
resource "lxd_network" "eth1" {
  name = "eth1"
  type = "macvlan"

  config = {
    "parent" = "lxdbr0"
  }
}
`
}

func testAccNetwork_target() string {
	return `
resource "lxd_network" "cluster_network_node1" {
  name   = "cluster_network"
  target = "node1"

  config = {
    "bridge.external_interfaces" = "nosuchint"
  }
}

resource "lxd_network" "cluster_network_node2" {
  name   = "cluster_network"
  target = "node2"

  config = {
    "bridge.external_interfaces" = "nosuchint"
  }
}

resource "lxd_network" "cluster_network" {
  depends_on = [
    "lxd_network.cluster_network_node1",
    "lxd_network.cluster_network_node2",
  ]

  name = lxd_network.cluster_network_node1.name
  config = {
    "ipv4.address" = "10.150.19.1/24"
  }
}
`
}

func testAccNetwork_project(project string) string {
	return fmt.Sprintf(`
resource "lxd_project" "project1" {
  name        = "%s"
  description = "Terraform provider test project"
}

resource "lxd_network" "eth1" {
  name    = "eth1"
  type    = "bridge"
  project = lxd_project.project1.name
}
	`, project)
}
