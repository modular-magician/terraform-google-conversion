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

package securesourcemanager

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const SecureSourceManagerInstanceIAMAssetType string = "securesourcemanager.googleapis.com/Instance"

func ResourceConverterSecureSourceManagerInstanceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecureSourceManagerInstanceIAMAssetType,
		Convert:           GetSecureSourceManagerInstanceIamPolicyCaiObject,
		MergeCreateUpdate: MergeSecureSourceManagerInstanceIamPolicy,
	}
}

func ResourceConverterSecureSourceManagerInstanceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecureSourceManagerInstanceIAMAssetType,
		Convert:           GetSecureSourceManagerInstanceIamBindingCaiObject,
		FetchFullResource: FetchSecureSourceManagerInstanceIamPolicy,
		MergeCreateUpdate: MergeSecureSourceManagerInstanceIamBinding,
		MergeDelete:       MergeSecureSourceManagerInstanceIamBindingDelete,
	}
}

func ResourceConverterSecureSourceManagerInstanceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecureSourceManagerInstanceIAMAssetType,
		Convert:           GetSecureSourceManagerInstanceIamMemberCaiObject,
		FetchFullResource: FetchSecureSourceManagerInstanceIamPolicy,
		MergeCreateUpdate: MergeSecureSourceManagerInstanceIamMember,
		MergeDelete:       MergeSecureSourceManagerInstanceIamMemberDelete,
	}
}

func GetSecureSourceManagerInstanceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecureSourceManagerInstanceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetSecureSourceManagerInstanceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecureSourceManagerInstanceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetSecureSourceManagerInstanceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecureSourceManagerInstanceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeSecureSourceManagerInstanceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeSecureSourceManagerInstanceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeSecureSourceManagerInstanceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeSecureSourceManagerInstanceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeSecureSourceManagerInstanceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newSecureSourceManagerInstanceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//securesourcemanager.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: SecureSourceManagerInstanceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchSecureSourceManagerInstanceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("instance_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		SecureSourceManagerInstanceIamUpdaterProducer,
		d,
		config,
		"//securesourcemanager.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance_id}}",
		SecureSourceManagerInstanceIAMAssetType,
	)
}
