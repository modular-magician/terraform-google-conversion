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
const PubsubTopicIAMAssetType string = "pubsub.googleapis.com/Topic"

func ResourceConverterPubsubTopicIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         PubsubTopicIAMAssetType,
		Convert:           GetPubsubTopicIamPolicyCaiObject,
		MergeCreateUpdate: MergePubsubTopicIamPolicy,
	}
}

func ResourceConverterPubsubTopicIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         PubsubTopicIAMAssetType,
		Convert:           GetPubsubTopicIamBindingCaiObject,
		FetchFullResource: FetchPubsubTopicIamPolicy,
		MergeCreateUpdate: MergePubsubTopicIamBinding,
		MergeDelete:       MergePubsubTopicIamBindingDelete,
	}
}

func ResourceConverterPubsubTopicIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         PubsubTopicIAMAssetType,
		Convert:           GetPubsubTopicIamMemberCaiObject,
		FetchFullResource: FetchPubsubTopicIamPolicy,
		MergeCreateUpdate: MergePubsubTopicIamMember,
		MergeDelete:       MergePubsubTopicIamMemberDelete,
	}
}

func GetPubsubTopicIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newPubsubTopicIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetPubsubTopicIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newPubsubTopicIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetPubsubTopicIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newPubsubTopicIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergePubsubTopicIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergePubsubTopicIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergePubsubTopicIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergePubsubTopicIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergePubsubTopicIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newPubsubTopicIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/topics/{{topic}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: PubsubTopicIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchPubsubTopicIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("topic"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		PubsubTopicIamUpdaterProducer,
		d,
		config,
		"//pubsub.googleapis.com/projects/{{project}}/topics/{{topic}}",
		PubsubTopicIAMAssetType,
	)
}
