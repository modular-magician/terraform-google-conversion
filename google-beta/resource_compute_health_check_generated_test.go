// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeHealthCheck_HealthCheckBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_HealthCheckBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_health_check.internal-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_HealthCheckBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_health_check" "internal-health-check" {
 name = "internal-service-health-check-%s"

 timeout_sec        = 1
 check_interval_sec = 1

 tcp_health_check {
   port = "80"
 }
}
`, val,
	)
}

func testAccCheckComputeHealthCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_health_check" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/healthChecks/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeHealthCheck still exists at %s", url)
		}
	}

	return nil
}
