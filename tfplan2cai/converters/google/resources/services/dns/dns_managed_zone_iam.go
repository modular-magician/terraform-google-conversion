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

package dns

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DNSManagedZoneIAMAssetType string = "dns.googleapis.com/ManagedZone"

func ResourceConverterDNSManagedZoneIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DNSManagedZoneIAMAssetType,
		Convert:           GetDNSManagedZoneIamPolicyCaiObject,
		MergeCreateUpdate: MergeDNSManagedZoneIamPolicy,
	}
}

func ResourceConverterDNSManagedZoneIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DNSManagedZoneIAMAssetType,
		Convert:           GetDNSManagedZoneIamBindingCaiObject,
		FetchFullResource: FetchDNSManagedZoneIamPolicy,
		MergeCreateUpdate: MergeDNSManagedZoneIamBinding,
		MergeDelete:       MergeDNSManagedZoneIamBindingDelete,
	}
}

func ResourceConverterDNSManagedZoneIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DNSManagedZoneIAMAssetType,
		Convert:           GetDNSManagedZoneIamMemberCaiObject,
		FetchFullResource: FetchDNSManagedZoneIamPolicy,
		MergeCreateUpdate: MergeDNSManagedZoneIamMember,
		MergeDelete:       MergeDNSManagedZoneIamMemberDelete,
	}
}

func GetDNSManagedZoneIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDNSManagedZoneIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetDNSManagedZoneIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDNSManagedZoneIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetDNSManagedZoneIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDNSManagedZoneIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeDNSManagedZoneIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDNSManagedZoneIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeDNSManagedZoneIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeDNSManagedZoneIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeDNSManagedZoneIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newDNSManagedZoneIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//dns.googleapis.com/projects/{{project}}/managedZones/{{managed_zone}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: DNSManagedZoneIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDNSManagedZoneIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("managed_zone"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		DNSManagedZoneIamUpdaterProducer,
		d,
		config,
		"//dns.googleapis.com/projects/{{project}}/managedZones/{{managed_zone}}",
		DNSManagedZoneIAMAssetType,
	)
}
