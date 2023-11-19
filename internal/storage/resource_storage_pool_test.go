package storage_test

import (
	"fmt"
	"testing"

	"github.com/canonical/lxd/shared/api"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/terraform-lxd/terraform-provider-lxd/internal/acctest"
)

func TestAccStoragePool_dir(t *testing.T) {
	poolName := petname.Generate(2, "-")
	driverName := "dir"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
					// Ensure computed keys are not tracked.
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.source"),
				),
			},
			{
				// Ensure no error is thrown on update.
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
				),
			},
		},
	})
}

func TestAccStoragePool_zfs(t *testing.T) {
	poolName := petname.Generate(2, "-")
	driverName := "zfs"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
					// Ensure computed keys are not tracked.
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.zfs.pool_name"),
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.size"),
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.source"),
				),
			},
			{
				// Ensure no error is thrown on update.
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
				),
			},
		},
	})
}

func TestAccStoragePool_lvm(t *testing.T) {
	poolName := petname.Generate(2, "-")
	driverName := "lvm"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
					// Ensure computed keys are not tracked.
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.lvm.vg_name"),
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.lvm.thinpool_name"),
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.size"),
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.source"),
				),
			},
			{
				// Ensure no error is thrown on update.
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
				),
			},
		},
	})
}

func TestAccStoragePool_btrfs(t *testing.T) {
	poolName := petname.Generate(2, "-")
	driverName := "btrfs"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
					// Ensure computed keys are not tracked.
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.size"),
					resource.TestCheckNoResourceAttr("lxd_storage_pool.storage_pool1", "config.source"),
				),
			},
			{
				// Ensure no error is thrown on update.
				Config: testAccStoragePool(poolName, driverName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "0"),
				),
			},
		},
	})
}

func TestAccStoragePool_size(t *testing.T) {
	poolName := petname.Generate(2, "-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool_size(poolName, "zfs"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", "zfs"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "1"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.size", "128MiB"),
				),
			},
			{
				Config: testAccStoragePool_size(poolName, "lvm"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", "lvm"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "1"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.size", "128MiB"),
				),
			},
			{
				Config: testAccStoragePool_size(poolName, "btrfs"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", "btrfs"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.%", "1"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "config.size", "128MiB"),
				),
			},
		},
	})
}

func TestAccStoragePool_project(t *testing.T) {
	poolName := petname.Generate(2, "-")
	projectName := petname.Name()
	driverName := "dir"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool_project(poolName, driverName, projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("lxd_project.project1", "name", projectName),
					resource.TestCheckResourceAttr("lxd_project.project1", "config.%", "1"),
					resource.TestCheckResourceAttr("lxd_project.project1", "config.features.storage.volumes", "false"),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "project", projectName),
				),
			},
		},
	})
}

// TODO:
//   - requires clustering precheck
// func TestAccStoragePool_target(t *testing.T) {
// 	poolName := petname.Generate(2, "-")
// 	driverName := "dir"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccStoragePool_target(poolName, driverName),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "name", poolName),
// 					resource.TestCheckResourceAttr("lxd_storage_pool.storage_pool1", "driver", driverName),
// 				),
// 			},
// 		},
// 	})
// }

func testAccStoragePoolConfig(pool *api.StoragePool, k, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if pool.Config == nil {
			return fmt.Errorf("No config")
		}

		for key, value := range pool.Config {
			if k != key {
				continue
			}

			if v == value {
				return nil
			}

			return fmt.Errorf("Bad value for %s: %s", k, value)
		}

		return fmt.Errorf("Config not found: %s", k)
	}
}

func testAccStoragePool(name string, driver string) string {
	return fmt.Sprintf(`
resource "lxd_storage_pool" "storage_pool1" {
  name   = "%s"
  driver = "%s"
}
	`, name, driver)
}

func testAccStoragePool_size(name string, driver string) string {
	return fmt.Sprintf(`
resource "lxd_storage_pool" "storage_pool1" {
  name   = "%s"
  driver = "%s"
  config = {
	size = "128MiB"
  }
}
	`, name, driver)
}

func testAccStoragePool_project(name string, driver string, project string) string {
	return fmt.Sprintf(`
resource "lxd_project" "project1" {
  name        = "%s"
  description = "Terraform provider test project"
  config = {
	"features.storage.volumes" = false
  }
}

resource "lxd_storage_pool" "storage_pool1" {
  name    = "%s"
  driver  = "%s"
  project = lxd_project.project1.name
}
	`, project, name, driver)
}

func testAccStoragePool_target(name, driver string) string {
	return fmt.Sprintf(`
resource "lxd_storage_pool" "storage_pool1_node1" {
  name   = "%[1]s"
  driver = "%[2]s"
  target = "node1"
}

resource "lxd_storage_pool" "storage_pool1_node2" {
  name   = "%[1]s"
  driver = "%[2]s"
  target = "node2"
}

resource "lxd_storage_pool" "storage_pool1" {
  depends_on = [
    "lxd_storage_pool.storage_pool1_node1",
    "lxd_storage_pool.storage_pool1_node2",
  ]

  name   = "%[1]s"
  driver = "%[2]s"
}
	`, name, driver)
}
