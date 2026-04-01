package mso

// Note: User association tests are skipped because the mso_user data source uses
// api/v1/users (or api/v2/users on ND), which no longer works on ND 4.1+ where
// the endpoint changed to /api/v1/infra/aaa/localUsers.
//
// Note: Cloud site association tests (AWS/Azure/GCP) are skipped because they
// require real cloud account credentials and vendor-specific configuration that
// is not available in the standard test environment.

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccMSOTenantResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMsoTenantDestroy,
		Steps: []resource.TestStep{
			{
				PreConfig: func() { fmt.Println("Test: Create Tenant") },
				Config:    testAccMSOTenantConfigCreate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mso_tenant.tenant", "name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "display_name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "description", "Terraform test tenant"),
				),
			},
			{
				PreConfig: func() { fmt.Println("Test: Update Tenant basic fields") },
				Config:    testAccMSOTenantConfigUpdateBasicFields(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mso_tenant.tenant", "name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "display_name", msoTenantName+" updated"),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "description", "Terraform test tenant updated"),
				),
			},
			{
				PreConfig: func() { fmt.Println("Test: Remove Tenant description") },
				Config:    testAccMSOTenantConfigRemoveDescription(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mso_tenant.tenant", "name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "display_name", msoTenantName+" updated"),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "description", ""),
				),
			},
			{
				PreConfig:         func() { fmt.Println("Test: Import Tenant") },
				ResourceName:      "mso_tenant.tenant",
				ImportState:       true,
				ImportStateVerify: true,
				// orchestrator_only is client-side only (controls delete behavior) and is not returned by the API.
				ImportStateVerifyIgnore: []string{"orchestrator_only"},
			},
			{
				PreConfig: func() { fmt.Println("Test: Add site association") },
				Config:    testAccMSOTenantConfigAddSiteAssociation(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mso_tenant.tenant", "name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "site_associations.#", "1"),
				),
			},
			{
				PreConfig: func() { fmt.Println("Test: Add extra site association") },
				Config:    testAccMSOTenantConfigAddExtraSiteAssociation(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mso_tenant.tenant", "name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "site_associations.#", "2"),
				),
			},
			{
				PreConfig: func() { fmt.Println("Test: Remove extra site association") },
				Config:    testAccMSOTenantConfigRemoveSiteAssociation(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mso_tenant.tenant", "name", msoTenantName),
					resource.TestCheckResourceAttr("mso_tenant.tenant", "site_associations.#", "1"),
				),
			},
			{
				PreConfig:         func() { fmt.Println("Test: Import Tenant with site associations") },
				ResourceName:      "mso_tenant.tenant",
				ImportState:       true,
				ImportStateVerify: true,
				// orchestrator_only is client-side only (controls delete behavior) and is not returned by the API.
				ImportStateVerifyIgnore: []string{"orchestrator_only"},
			},
		},
	})
}

func testAccMSOTenantConfigCreate() string {
	return fmt.Sprintf(`
	resource "mso_tenant" "tenant" {
		name         = "%s"
		display_name = "%s"
		description  = "Terraform test tenant"
	}`, msoTenantName, msoTenantName)
}

func testAccMSOTenantConfigUpdateBasicFields() string {
	return fmt.Sprintf(`
	resource "mso_tenant" "tenant" {
		name         = "%s"
		display_name = "%s updated"
		description  = "Terraform test tenant updated"
	}`, msoTenantName, msoTenantName)
}

func testAccMSOTenantConfigRemoveDescription() string {
	return fmt.Sprintf(`
	resource "mso_tenant" "tenant" {
		name         = "%s"
		display_name = "%s updated"
	}`, msoTenantName, msoTenantName)
}

func testAccMSOTenantConfigAddSiteAssociation() string {
	return fmt.Sprintf(`%s
	resource "mso_tenant" "tenant" {
		name         = "%s"
		display_name = "%s updated"
		site_associations {
			site_id = data.mso_site.%s.id
		}
	}`, testSiteConfigAnsibleTest(), msoTenantName, msoTenantName, msoTemplateSiteName1)
}

func testAccMSOTenantConfigAddExtraSiteAssociation() string {
	return fmt.Sprintf(`%s%s
	resource "mso_tenant" "tenant" {
		name         = "%s"
		display_name = "%s updated"
		site_associations {
			site_id = data.mso_site.%s.id
		}
		site_associations {
			site_id = data.mso_site.%s.id
		}
	}`, testSiteConfigAnsibleTest(), testSiteConfigAnsibleTest2(), msoTenantName, msoTenantName, msoTemplateSiteName1, msoTemplateSiteName2)
}

func testAccMSOTenantConfigRemoveSiteAssociation() string {
	return fmt.Sprintf(`%s
	resource "mso_tenant" "tenant" {
		name         = "%s"
		display_name = "%s updated"
		site_associations {
			site_id = data.mso_site.%s.id
		}
	}`, testSiteConfigAnsibleTest(), msoTenantName, msoTenantName, msoTemplateSiteName1)
}

// testAccCheckMsoTenantDestroy verifies that the tenant is deleted from MSO.
// The generic testCheckResourceDestroyPolicy helpers cannot be used here because
// they query the policy/template API (api/v1/templates/objects), whereas tenants
// use a separate API endpoint (api/v1/tenants).
func testAccCheckMsoTenantDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "mso_tenant" {
			_, err := client.GetViaURL("api/v1/tenants/" + rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Tenant still exists")
			}
		} else {
			continue
		}
	}
	return nil
}
