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

package accesscontextmanager

import (
	"testing"

	cai2hcl_testing "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/testing"
)

func TestAccessContextManagerEgressPolicy(t *testing.T) {
	cai2hcl_testing.AssertTestFiles(
		t,
		UberConverter,
		"./testdata",
		[]string{"accesscontextmanager_egress_policy"})
}
