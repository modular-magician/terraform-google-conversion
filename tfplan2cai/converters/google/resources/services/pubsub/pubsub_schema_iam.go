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

package pubsub

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const PubsubSchemaIAMAssetType string = "pubsub.googleapis.com/Schema"

func ResourceConverterPubsubSchemaIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         PubsubSchemaIAMAssetType,
		Convert:           GetPubsubSchemaIamPolicyCaiObject,
		MergeCreateUpdate: MergePubsubSchemaIamPolicy,
	}
}

func ResourceConverterPubsubSchemaIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         PubsubSchemaIAMAssetType,
		Convert:           GetPubsubSchemaIamBindingCaiObject,
		FetchFullResource: FetchPubsubSchemaIamPolicy,
		MergeCreateUpdate: MergePubsubSchemaIamBinding,
		MergeDelete:       MergePubsubSchemaIamBindingDelete,
	}
}

func ResourceConverterPubsubSchemaIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         PubsubSchemaIAMAssetType,
		Convert:           GetPubsubSchemaIamMemberCaiObject,
		FetchFullResource: FetchPubsubSchemaIamPolicy,
		MergeCreateUpdate: MergePubsubSchemaIamMember,
		MergeDelete:       MergePubsubSchemaIamMemberDelete,
	}
}

func GetPubsubSchemaIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newPubsubSchemaIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetPubsubSchemaIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newPubsubSchemaIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetPubsubSchemaIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newPubsubSchemaIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergePubsubSchemaIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergePubsubSchemaIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergePubsubSchemaIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergePubsubSchemaIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergePubsubSchemaIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newPubsubSchemaIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/schemas/{{schema}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: PubsubSchemaIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchPubsubSchemaIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("schema"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		PubsubSchemaIamUpdaterProducer,
		d,
		config,
		"//pubsub.googleapis.com/projects/{{project}}/schemas/{{schema}}",
		PubsubSchemaIAMAssetType,
	)
}
