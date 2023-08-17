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

package vertexai

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const VertexAIEndpointIAMAssetType string = "{{region}}-aiplatform.googleapis.com/Endpoint"

func ResourceConverterVertexAIEndpointIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIEndpointIAMAssetType,
		Convert:           GetVertexAIEndpointIamPolicyCaiObject,
		MergeCreateUpdate: MergeVertexAIEndpointIamPolicy,
	}
}

func ResourceConverterVertexAIEndpointIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIEndpointIAMAssetType,
		Convert:           GetVertexAIEndpointIamBindingCaiObject,
		FetchFullResource: FetchVertexAIEndpointIamPolicy,
		MergeCreateUpdate: MergeVertexAIEndpointIamBinding,
		MergeDelete:       MergeVertexAIEndpointIamBindingDelete,
	}
}

func ResourceConverterVertexAIEndpointIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIEndpointIAMAssetType,
		Convert:           GetVertexAIEndpointIamMemberCaiObject,
		FetchFullResource: FetchVertexAIEndpointIamPolicy,
		MergeCreateUpdate: MergeVertexAIEndpointIamMember,
		MergeDelete:       MergeVertexAIEndpointIamMemberDelete,
	}
}

func GetVertexAIEndpointIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIEndpointIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetVertexAIEndpointIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIEndpointIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetVertexAIEndpointIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIEndpointIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeVertexAIEndpointIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeVertexAIEndpointIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeVertexAIEndpointIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeVertexAIEndpointIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeVertexAIEndpointIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newVertexAIEndpointIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//{{region}}-aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/endpoints/{{endpoint}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VertexAIEndpointIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchVertexAIEndpointIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("endpoint"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		VertexAIEndpointIamUpdaterProducer,
		d,
		config,
		"//{{region}}-aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/endpoints/{{endpoint}}",
		VertexAIEndpointIAMAssetType,
	)
}
