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
const DataplexDatascanIAMAssetType string = "dataplex.googleapis.com/Datascan"

func resourceConverterDataplexDatascanIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataplexDatascanIAMAssetType,
		Convert:           GetDataplexDatascanIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataplexDatascanIamPolicy,
	}
}

func resourceConverterDataplexDatascanIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataplexDatascanIAMAssetType,
		Convert:           GetDataplexDatascanIamBindingCaiObject,
		FetchFullResource: FetchDataplexDatascanIamPolicy,
		MergeCreateUpdate: MergeDataplexDatascanIamBinding,
		MergeDelete:       MergeDataplexDatascanIamBindingDelete,
	}
}

func resourceConverterDataplexDatascanIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataplexDatascanIAMAssetType,
		Convert:           GetDataplexDatascanIamMemberCaiObject,
		FetchFullResource: FetchDataplexDatascanIamPolicy,
		MergeCreateUpdate: MergeDataplexDatascanIamMember,
		MergeDelete:       MergeDataplexDatascanIamMemberDelete,
	}
}

func GetDataplexDatascanIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDataplexDatascanIamAsset(d, config, expandIamPolicyBindings)
}

func GetDataplexDatascanIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDataplexDatascanIamAsset(d, config, expandIamRoleBindings)
}

func GetDataplexDatascanIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDataplexDatascanIamAsset(d, config, expandIamMemberBindings)
}

func MergeDataplexDatascanIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataplexDatascanIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeDataplexDatascanIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeDataplexDatascanIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeDataplexDatascanIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newDataplexDatascanIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/dataScans/{{data_scan}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: DataplexDatascanIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataplexDatascanIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("data_scan"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		DataplexDatascanIamUpdaterProducer,
		d,
		config,
		"//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/dataScans/{{data_scan}}",
		DataplexDatascanIAMAssetType,
	)
}
