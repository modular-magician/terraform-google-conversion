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

package servicemanagement

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ServiceManagementServiceConsumersIAMAssetType string = "servicemanagement.googleapis.com/ServiceConsumers"

func ResourceConverterServiceManagementServiceConsumersIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ServiceManagementServiceConsumersIAMAssetType,
		Convert:           GetServiceManagementServiceConsumersIamPolicyCaiObject,
		MergeCreateUpdate: MergeServiceManagementServiceConsumersIamPolicy,
	}
}

func ResourceConverterServiceManagementServiceConsumersIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ServiceManagementServiceConsumersIAMAssetType,
		Convert:           GetServiceManagementServiceConsumersIamBindingCaiObject,
		FetchFullResource: FetchServiceManagementServiceConsumersIamPolicy,
		MergeCreateUpdate: MergeServiceManagementServiceConsumersIamBinding,
		MergeDelete:       MergeServiceManagementServiceConsumersIamBindingDelete,
	}
}

func ResourceConverterServiceManagementServiceConsumersIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ServiceManagementServiceConsumersIAMAssetType,
		Convert:           GetServiceManagementServiceConsumersIamMemberCaiObject,
		FetchFullResource: FetchServiceManagementServiceConsumersIamPolicy,
		MergeCreateUpdate: MergeServiceManagementServiceConsumersIamMember,
		MergeDelete:       MergeServiceManagementServiceConsumersIamMemberDelete,
	}
}

func GetServiceManagementServiceConsumersIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newServiceManagementServiceConsumersIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetServiceManagementServiceConsumersIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newServiceManagementServiceConsumersIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetServiceManagementServiceConsumersIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newServiceManagementServiceConsumersIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeServiceManagementServiceConsumersIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeServiceManagementServiceConsumersIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeServiceManagementServiceConsumersIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeServiceManagementServiceConsumersIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeServiceManagementServiceConsumersIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newServiceManagementServiceConsumersIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//servicemanagement.googleapis.com/services/{{service_name}}/consumers/{{consumer_project}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ServiceManagementServiceConsumersIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchServiceManagementServiceConsumersIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("service_name"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("consumer_project"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ServiceManagementServiceConsumersIamUpdaterProducer,
		d,
		config,
		"//servicemanagement.googleapis.com/services/{{service_name}}/consumers/{{consumer_project}}",
		ServiceManagementServiceConsumersIAMAssetType,
	)
}
