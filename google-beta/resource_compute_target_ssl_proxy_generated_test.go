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

func TestAccComputeTargetSslProxy_TargetSslProxyBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeTargetSslProxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetSslProxy_TargetSslProxyBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_target_ssl_proxy.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeTargetSslProxy_TargetSslProxyBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_target_ssl_proxy" "default" {
  name             = "test-proxy-%s"
  backend_service  = "${google_compute_backend_service.default.self_link}"
  ssl_certificates = ["${google_compute_ssl_certificate.default.self_link}"]
}

resource "google_compute_ssl_certificate" "default" {
  name        = "default-cert-%s"
  private_key = "${file("test-fixtures/ssl_cert/test.key")}"
  certificate = "${file("test-fixtures/ssl_cert/test.crt")}"
}

resource "google_compute_backend_service" "default" {
  name          = "backend-service-%s"
  protocol      = "SSL"
  health_checks = ["${google_compute_health_check.default.self_link}"]
}

resource "google_compute_health_check" "default" {
  name               = "health-check-%s"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "443"
  }
}
`, val, val, val, val,
	)
}

func testAccCheckComputeTargetSslProxyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_target_ssl_proxy" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/targetSslProxies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeTargetSslProxy still exists at %s", url)
		}
	}

	return nil
}
