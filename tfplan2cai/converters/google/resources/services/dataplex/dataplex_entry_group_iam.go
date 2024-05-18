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

package dataplex

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataplexEntryGroupIAMAssetType string = "dataplex.googleapis.com/EntryGroup"

func ResourceConverterDataplexEntryGroupIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexEntryGroupIAMAssetType,
		Convert:           GetDataplexEntryGroupIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataplexEntryGroupIamPolicy,
	}
}

func ResourceConverterDataplexEntryGroupIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexEntryGroupIAMAssetType,
		Convert:           GetDataplexEntryGroupIamBindingCaiObject,
		FetchFullResource: FetchDataplexEntryGroupIamPolicy,
		MergeCreateUpdate: MergeDataplexEntryGroupIamBinding,
		MergeDelete:       MergeDataplexEntryGroupIamBindingDelete,
	}
}

func ResourceConverterDataplexEntryGroupIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexEntryGroupIAMAssetType,
		Convert:           GetDataplexEntryGroupIamMemberCaiObject,
		FetchFullResource: FetchDataplexEntryGroupIamPolicy,
		MergeCreateUpdate: MergeDataplexEntryGroupIamMember,
		MergeDelete:       MergeDataplexEntryGroupIamMemberDelete,
	}
}

func GetDataplexEntryGroupIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexEntryGroupIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataplexEntryGroupIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexEntryGroupIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataplexEntryGroupIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexEntryGroupIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataplexEntryGroupIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataplexEntryGroupIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataplexEntryGroupIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataplexEntryGroupIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataplexEntryGroupIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataplexEntryGroupIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataplexEntryGroupIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataplexEntryGroupIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("entry_group_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataplexEntryGroupIamUpdaterProducer,
		d,
		config,
		"//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/entryGroups/{{entry_group_id}}",
		DataplexEntryGroupIAMAssetType,
	)
}
