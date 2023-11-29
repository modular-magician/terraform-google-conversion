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

package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeSubnetworkIAMAssetType string = "compute.googleapis.com/Subnetwork"

func ResourceConverterComputeSubnetworkIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeSubnetworkIamPolicy,
	}
}

func ResourceConverterComputeSubnetworkIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamBindingCaiObject,
		FetchFullResource: FetchComputeSubnetworkIamPolicy,
		MergeCreateUpdate: MergeComputeSubnetworkIamBinding,
		MergeDelete:       MergeComputeSubnetworkIamBindingDelete,
	}
}

func ResourceConverterComputeSubnetworkIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamMemberCaiObject,
		FetchFullResource: FetchComputeSubnetworkIamPolicy,
		MergeCreateUpdate: MergeComputeSubnetworkIamMember,
		MergeDelete:       MergeComputeSubnetworkIamMemberDelete,
	}
}

func GetComputeSubnetworkIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetComputeSubnetworkIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetComputeSubnetworkIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeComputeSubnetworkIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeSubnetworkIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeComputeSubnetworkIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeComputeSubnetworkIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeComputeSubnetworkIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newComputeSubnetworkIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ComputeSubnetworkIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeSubnetworkIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("subnetwork"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ComputeSubnetworkIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}",
		ComputeSubnetworkIAMAssetType,
	)
}
