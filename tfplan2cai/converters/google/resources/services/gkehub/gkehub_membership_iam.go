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

package gkehub

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const GKEHubMembershipIAMAssetType string = "gkehub.googleapis.com/Membership"

func ResourceConverterGKEHubMembershipIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEHubMembershipIAMAssetType,
		Convert:           GetGKEHubMembershipIamPolicyCaiObject,
		MergeCreateUpdate: MergeGKEHubMembershipIamPolicy,
	}
}

func ResourceConverterGKEHubMembershipIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEHubMembershipIAMAssetType,
		Convert:           GetGKEHubMembershipIamBindingCaiObject,
		FetchFullResource: FetchGKEHubMembershipIamPolicy,
		MergeCreateUpdate: MergeGKEHubMembershipIamBinding,
		MergeDelete:       MergeGKEHubMembershipIamBindingDelete,
	}
}

func ResourceConverterGKEHubMembershipIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEHubMembershipIAMAssetType,
		Convert:           GetGKEHubMembershipIamMemberCaiObject,
		FetchFullResource: FetchGKEHubMembershipIamPolicy,
		MergeCreateUpdate: MergeGKEHubMembershipIamMember,
		MergeDelete:       MergeGKEHubMembershipIamMemberDelete,
	}
}

func GetGKEHubMembershipIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEHubMembershipIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetGKEHubMembershipIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEHubMembershipIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetGKEHubMembershipIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEHubMembershipIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeGKEHubMembershipIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeGKEHubMembershipIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeGKEHubMembershipIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeGKEHubMembershipIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeGKEHubMembershipIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newGKEHubMembershipIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//gkehub.googleapis.com/projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: GKEHubMembershipIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchGKEHubMembershipIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("membership_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		GKEHubMembershipIamUpdaterProducer,
		d,
		config,
		"//gkehub.googleapis.com/projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}",
		GKEHubMembershipIAMAssetType,
	)
}
