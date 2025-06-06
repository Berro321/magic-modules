package resourcemanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataSourceGoogleOrganizationIamCustomRoles_basic(t *testing.T) {
	t.Parallel()

	orgId := envvar.GetTestOrgFromEnv(t)
	roleId := "tfIamCustomRole" + acctest.RandString(t, 10)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleOrganizationIamCustomRolesConfig(orgId, roleId),
				Check: resource.ComposeTestCheckFunc(
					// We can't guarantee no organization won't have our custom role as first element, so we'll check set-ness rather than correctness
					resource.TestCheckResourceAttrSet("data.google_organization_iam_custom_roles.this", "roles.0.id"),
					resource.TestCheckResourceAttrSet("data.google_organization_iam_custom_roles.this", "roles.0.name"),
					resource.TestCheckResourceAttrSet("data.google_organization_iam_custom_roles.this", "roles.0.role_id"),
					resource.TestCheckResourceAttrSet("data.google_organization_iam_custom_roles.this", "roles.0.stage"),
					resource.TestCheckResourceAttrSet("data.google_organization_iam_custom_roles.this", "roles.0.title"),
				),
			},
		},
	})
}

func testAccCheckGoogleOrganizationIamCustomRolesConfig(orgId string, roleId string) string {
	return fmt.Sprintf(`
resource "google_organization_iam_custom_role" "this" {
  org_id  = "%s"
  role_id = "%s"
  title   = "Terraform Test"

  permissions = [
	"iam.roles.create",
	"iam.roles.delete",
	"iam.roles.list",
  ]
}

data "google_organization_iam_custom_roles" "this" {
  org_id = google_organization_iam_custom_role.this.org_id
}
`, orgId, roleId)
}
