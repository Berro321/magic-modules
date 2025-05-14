package beyondcorp_test

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"testing"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccDataSourceGoogleBeyondcorpSecurityGatewayApplication_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpSecurityGatewayApplicationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBeyondcorpSecurityGatewayApplication_basic(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_beyondcorp_security_gateway_application.foo", "google_beyondcorp_security_gateway_application.foo"),
				),
			},
		},
	})
}

// func TestAccDataSourceGoogleBeyondcorpSecurityGatewayApplication_full(t *testing.T) {
// 	t.Parallel()

// 	context := map[string]interface{}{
// 		"random_suffix": acctest.RandString(t, 10),
// 	}

// 	acctest.VcrTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
// 		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
// 		CheckDestroy:             testAccCheckBeyondcorpSecurityGatewayApplicationDestroyProducer(t),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDataSourceGoogleBeyondcorpSecurityGatewayApplication_full(context),
// 				Check: resource.ComposeTestCheckFunc(
// 					acctest.CheckDataSourceStateMatchesResourceState("data.google_beyondcorp_security_gateway_application.foo", "google_beyondcorp_security_gateway_application.foo"),
// 				),
// 			},
// 		},
// 	})
// }

func testAccDataSourceGoogleBeyondcorpSecurityGatewayApplication_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "default" {
  security_gateway_id = "default-foo-sg-basic-%{random_suffix}"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}

resource "google_beyondcorp_security_gateway_application" "foo" {
  security_gateways_id = google_beyondcorp_security_gateway.default.security_gateway_id
  application_id = "default-foo-sga-basic-%{random_suffix}"
  endpoint_matchers {
    hostname = "google.com"
  }
}

data "google_beyondcorp_security_gateway_application" "foo" {
	application_id = google_beyondcorp_security_gateway_application.foo.application_id
	security_gateways_id = google_beyondcorp_security_gateway_application.foo.security_gateways_id
}
`, context)
}

// func testAccDataSourceGoogleBeyondcorpSecurityGateway_full(context map[string]interface{}) string {
// 	return acctest.Nprintf(`
// resource "google_beyondcorp_security_gateway" "foo" {
//   security_gateway_id = "default-foo-sg-full-%{random_suffix}"
//   display_name = "My Security Gateway resource"
//   hubs { region = "us-central1" }
// }
// data "google_beyondcorp_security_gateway" "foo" {
// 	security_gateway_id = google_beyondcorp_security_gateway.foo.security_gateway_id
// 	project = google_beyondcorp_security_gateway.foo.project
// }
// `, context)
// }
