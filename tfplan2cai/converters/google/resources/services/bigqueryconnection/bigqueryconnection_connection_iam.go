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

package bigqueryconnection

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const BigqueryConnectionConnectionIAMAssetType string = "bigqueryconnection.googleapis.com/Connection"

func ResourceConverterBigqueryConnectionConnectionIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         BigqueryConnectionConnectionIAMAssetType,
		Convert:           GetBigqueryConnectionConnectionIamPolicyCaiObject,
		MergeCreateUpdate: MergeBigqueryConnectionConnectionIamPolicy,
	}
}

func ResourceConverterBigqueryConnectionConnectionIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         BigqueryConnectionConnectionIAMAssetType,
		Convert:           GetBigqueryConnectionConnectionIamBindingCaiObject,
		FetchFullResource: FetchBigqueryConnectionConnectionIamPolicy,
		MergeCreateUpdate: MergeBigqueryConnectionConnectionIamBinding,
		MergeDelete:       MergeBigqueryConnectionConnectionIamBindingDelete,
	}
}

func ResourceConverterBigqueryConnectionConnectionIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         BigqueryConnectionConnectionIAMAssetType,
		Convert:           GetBigqueryConnectionConnectionIamMemberCaiObject,
		FetchFullResource: FetchBigqueryConnectionConnectionIamPolicy,
		MergeCreateUpdate: MergeBigqueryConnectionConnectionIamMember,
		MergeDelete:       MergeBigqueryConnectionConnectionIamMemberDelete,
	}
}

func GetBigqueryConnectionConnectionIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newBigqueryConnectionConnectionIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetBigqueryConnectionConnectionIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newBigqueryConnectionConnectionIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetBigqueryConnectionConnectionIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newBigqueryConnectionConnectionIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeBigqueryConnectionConnectionIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeBigqueryConnectionConnectionIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeBigqueryConnectionConnectionIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeBigqueryConnectionConnectionIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeBigqueryConnectionConnectionIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newBigqueryConnectionConnectionIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//bigqueryconnection.googleapis.com/projects/{{project}}/locations/{{location}}/connections/{{connection_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: BigqueryConnectionConnectionIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchBigqueryConnectionConnectionIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("connection_id"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		BigqueryConnectionConnectionIamUpdaterProducer,
		d,
		config,
		"//bigqueryconnection.googleapis.com/projects/{{project}}/locations/{{location}}/connections/{{connection_id}}",
		BigqueryConnectionConnectionIAMAssetType,
	)
}
