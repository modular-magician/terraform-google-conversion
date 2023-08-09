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

package servicedirectory

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ServiceDirectoryServiceIAMAssetType string = "servicedirectory.googleapis.com/Service"

func ResourceConverterServiceDirectoryServiceIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ServiceDirectoryServiceIAMAssetType,
		Convert:           GetServiceDirectoryServiceIamPolicyCaiObject,
		MergeCreateUpdate: MergeServiceDirectoryServiceIamPolicy,
	}
}

func ResourceConverterServiceDirectoryServiceIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ServiceDirectoryServiceIAMAssetType,
		Convert:           GetServiceDirectoryServiceIamBindingCaiObject,
		FetchFullResource: FetchServiceDirectoryServiceIamPolicy,
		MergeCreateUpdate: MergeServiceDirectoryServiceIamBinding,
		MergeDelete:       MergeServiceDirectoryServiceIamBindingDelete,
	}
}

func ResourceConverterServiceDirectoryServiceIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ServiceDirectoryServiceIAMAssetType,
		Convert:           GetServiceDirectoryServiceIamMemberCaiObject,
		FetchFullResource: FetchServiceDirectoryServiceIamPolicy,
		MergeCreateUpdate: MergeServiceDirectoryServiceIamMember,
		MergeDelete:       MergeServiceDirectoryServiceIamMemberDelete,
	}
}

func GetServiceDirectoryServiceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newServiceDirectoryServiceIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetServiceDirectoryServiceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newServiceDirectoryServiceIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetServiceDirectoryServiceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newServiceDirectoryServiceIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeServiceDirectoryServiceIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeServiceDirectoryServiceIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeServiceDirectoryServiceIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeServiceDirectoryServiceIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeServiceDirectoryServiceIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newServiceDirectoryServiceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//servicedirectory.googleapis.com/projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ServiceDirectoryServiceIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchServiceDirectoryServiceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("name"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ServiceDirectoryServiceIamUpdaterProducer,
		d,
		config,
		"//servicedirectory.googleapis.com/projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}/services/{{service_id}}",
		ServiceDirectoryServiceIAMAssetType,
	)
}
