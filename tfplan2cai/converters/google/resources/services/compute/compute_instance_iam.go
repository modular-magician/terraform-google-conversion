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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeInstanceIAMAssetType string = "compute.googleapis.com/Instance"

func ResourceConverterComputeInstanceIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeInstanceIamPolicy,
	}
}

func ResourceConverterComputeInstanceIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamBindingCaiObject,
		FetchFullResource: FetchComputeInstanceIamPolicy,
		MergeCreateUpdate: MergeComputeInstanceIamBinding,
		MergeDelete:       MergeComputeInstanceIamBindingDelete,
	}
}

func ResourceConverterComputeInstanceIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamMemberCaiObject,
		FetchFullResource: FetchComputeInstanceIamPolicy,
		MergeCreateUpdate: MergeComputeInstanceIamMember,
		MergeDelete:       MergeComputeInstanceIamMemberDelete,
	}
}

func GetComputeInstanceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newComputeInstanceIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetComputeInstanceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newComputeInstanceIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetComputeInstanceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newComputeInstanceIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeComputeInstanceIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeInstanceIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeComputeInstanceIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeComputeInstanceIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeComputeInstanceIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newComputeInstanceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instances/{{instance_name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ComputeInstanceIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeInstanceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("zone"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("instance_name"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ComputeInstanceIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instances/{{instance_name}}",
		ComputeInstanceIAMAssetType,
	)
}
