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

package cloudfunctions2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const Cloudfunctions2functionIAMAssetType string = "cloudfunctions.googleapis.com/function"

func ResourceConverterCloudfunctions2functionIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudfunctions2functionIamPolicy,
	}
}

func ResourceConverterCloudfunctions2functionIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamBindingCaiObject,
		FetchFullResource: FetchCloudfunctions2functionIamPolicy,
		MergeCreateUpdate: MergeCloudfunctions2functionIamBinding,
		MergeDelete:       MergeCloudfunctions2functionIamBindingDelete,
	}
}

func ResourceConverterCloudfunctions2functionIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamMemberCaiObject,
		FetchFullResource: FetchCloudfunctions2functionIamPolicy,
		MergeCreateUpdate: MergeCloudfunctions2functionIamMember,
		MergeDelete:       MergeCloudfunctions2functionIamMemberDelete,
	}
}

func GetCloudfunctions2functionIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetCloudfunctions2functionIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetCloudfunctions2functionIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeCloudfunctions2functionIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudfunctions2functionIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeCloudfunctions2functionIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeCloudfunctions2functionIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeCloudfunctions2functionIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newCloudfunctions2functionIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: Cloudfunctions2functionIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudfunctions2functionIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("cloud_function"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		Cloudfunctions2functionIamUpdaterProducer,
		d,
		config,
		"//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}",
		Cloudfunctions2functionIAMAssetType,
	)
}
