// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package monitoring

import (
	"testing"

	cai2hcl_testing "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/testing"
)

func TestMonitoringGenericService(t *testing.T) {
	cai2hcl_testing.AssertTestFiles(
		t,
		UberConverter,
		"./testdata",
		[]string{"monitoring_generic_service"})
}
