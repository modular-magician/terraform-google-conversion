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

package identityplatform

import (
	"testing"

	cai2hcl_testing "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/testing"
)

func TestIdentityPlatformTenantInboundSamlConfig(t *testing.T) {
	cai2hcl_testing.AssertTestFiles(
		t,
		UberConverter,
		"./testdata",
		[]string{"identityplatform_tenant_inbound_saml_config"})
}
