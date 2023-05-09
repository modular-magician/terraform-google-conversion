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

package google

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudRunV2JobIAMAssetType string = "run.googleapis.com/Job"

func resourceConverterCloudRunV2JobIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudRunV2JobIAMAssetType,
		Convert:           GetCloudRunV2JobIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudRunV2JobIamPolicy,
	}
}

func resourceConverterCloudRunV2JobIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudRunV2JobIAMAssetType,
		Convert:           GetCloudRunV2JobIamBindingCaiObject,
		FetchFullResource: FetchCloudRunV2JobIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2JobIamBinding,
		MergeDelete:       MergeCloudRunV2JobIamBindingDelete,
	}
}

func resourceConverterCloudRunV2JobIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudRunV2JobIAMAssetType,
		Convert:           GetCloudRunV2JobIamMemberCaiObject,
		FetchFullResource: FetchCloudRunV2JobIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2JobIamMember,
		MergeDelete:       MergeCloudRunV2JobIamMemberDelete,
	}
}

func GetCloudRunV2JobIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudRunV2JobIamAsset(d, config, expandIamPolicyBindings)
}

func GetCloudRunV2JobIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudRunV2JobIamAsset(d, config, expandIamRoleBindings)
}

func GetCloudRunV2JobIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudRunV2JobIamAsset(d, config, expandIamMemberBindings)
}

func MergeCloudRunV2JobIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudRunV2JobIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeCloudRunV2JobIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeCloudRunV2JobIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeCloudRunV2JobIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newCloudRunV2JobIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//run.googleapis.com/projects/{{project}}/locations/{{location}}/jobs/{{name}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: CloudRunV2JobIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudRunV2JobIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		CloudRunV2JobIamUpdaterProducer,
		d,
		config,
		"//run.googleapis.com/projects/{{project}}/locations/{{location}}/jobs/{{name}}",
		CloudRunV2JobIAMAssetType,
	)
}
