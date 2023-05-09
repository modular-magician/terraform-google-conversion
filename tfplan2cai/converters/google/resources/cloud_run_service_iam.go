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
const CloudRunServiceIAMAssetType string = "run.googleapis.com/Service"

func resourceConverterCloudRunServiceIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudRunServiceIAMAssetType,
		Convert:           GetCloudRunServiceIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudRunServiceIamPolicy,
	}
}

func resourceConverterCloudRunServiceIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudRunServiceIAMAssetType,
		Convert:           GetCloudRunServiceIamBindingCaiObject,
		FetchFullResource: FetchCloudRunServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunServiceIamBinding,
		MergeDelete:       MergeCloudRunServiceIamBindingDelete,
	}
}

func resourceConverterCloudRunServiceIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudRunServiceIAMAssetType,
		Convert:           GetCloudRunServiceIamMemberCaiObject,
		FetchFullResource: FetchCloudRunServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunServiceIamMember,
		MergeDelete:       MergeCloudRunServiceIamMemberDelete,
	}
}

func GetCloudRunServiceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudRunServiceIamAsset(d, config, expandIamPolicyBindings)
}

func GetCloudRunServiceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudRunServiceIamAsset(d, config, expandIamRoleBindings)
}

func GetCloudRunServiceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudRunServiceIamAsset(d, config, expandIamMemberBindings)
}

func MergeCloudRunServiceIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudRunServiceIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeCloudRunServiceIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeCloudRunServiceIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeCloudRunServiceIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newCloudRunServiceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: CloudRunServiceIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudRunServiceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("service"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		CloudRunServiceIamUpdaterProducer,
		d,
		config,
		"//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service}}",
		CloudRunServiceIAMAssetType,
	)
}
