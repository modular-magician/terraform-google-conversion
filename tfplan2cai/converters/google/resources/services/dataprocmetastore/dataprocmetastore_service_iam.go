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

package dataprocmetastore

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataprocMetastoreServiceIAMAssetType string = "metastore.googleapis.com/Service"

func ResourceConverterDataprocMetastoreServiceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreServiceIAMAssetType,
		Convert:           GetDataprocMetastoreServiceIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataprocMetastoreServiceIamPolicy,
	}
}

func ResourceConverterDataprocMetastoreServiceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreServiceIAMAssetType,
		Convert:           GetDataprocMetastoreServiceIamBindingCaiObject,
		FetchFullResource: FetchDataprocMetastoreServiceIamPolicy,
		MergeCreateUpdate: MergeDataprocMetastoreServiceIamBinding,
		MergeDelete:       MergeDataprocMetastoreServiceIamBindingDelete,
	}
}

func ResourceConverterDataprocMetastoreServiceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreServiceIAMAssetType,
		Convert:           GetDataprocMetastoreServiceIamMemberCaiObject,
		FetchFullResource: FetchDataprocMetastoreServiceIamPolicy,
		MergeCreateUpdate: MergeDataprocMetastoreServiceIamMember,
		MergeDelete:       MergeDataprocMetastoreServiceIamMemberDelete,
	}
}

func GetDataprocMetastoreServiceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreServiceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataprocMetastoreServiceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreServiceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataprocMetastoreServiceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreServiceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataprocMetastoreServiceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataprocMetastoreServiceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataprocMetastoreServiceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataprocMetastoreServiceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataprocMetastoreServiceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataprocMetastoreServiceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataprocMetastoreServiceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataprocMetastoreServiceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("service_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataprocMetastoreServiceIamUpdaterProducer,
		d,
		config,
		"//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service_id}}",
		DataprocMetastoreServiceIAMAssetType,
	)
}
