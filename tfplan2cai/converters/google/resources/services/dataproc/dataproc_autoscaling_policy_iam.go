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

package dataproc

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataprocAutoscalingPolicyIAMAssetType string = "dataproc.googleapis.com/AutoscalingPolicy"

func ResourceConverterDataprocAutoscalingPolicyIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataprocAutoscalingPolicyIAMAssetType,
		Convert:           GetDataprocAutoscalingPolicyIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataprocAutoscalingPolicyIamPolicy,
	}
}

func ResourceConverterDataprocAutoscalingPolicyIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataprocAutoscalingPolicyIAMAssetType,
		Convert:           GetDataprocAutoscalingPolicyIamBindingCaiObject,
		FetchFullResource: FetchDataprocAutoscalingPolicyIamPolicy,
		MergeCreateUpdate: MergeDataprocAutoscalingPolicyIamBinding,
		MergeDelete:       MergeDataprocAutoscalingPolicyIamBindingDelete,
	}
}

func ResourceConverterDataprocAutoscalingPolicyIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataprocAutoscalingPolicyIAMAssetType,
		Convert:           GetDataprocAutoscalingPolicyIamMemberCaiObject,
		FetchFullResource: FetchDataprocAutoscalingPolicyIamPolicy,
		MergeCreateUpdate: MergeDataprocAutoscalingPolicyIamMember,
		MergeDelete:       MergeDataprocAutoscalingPolicyIamMemberDelete,
	}
}

func GetDataprocAutoscalingPolicyIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataprocAutoscalingPolicyIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetDataprocAutoscalingPolicyIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataprocAutoscalingPolicyIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetDataprocAutoscalingPolicyIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataprocAutoscalingPolicyIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeDataprocAutoscalingPolicyIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataprocAutoscalingPolicyIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeDataprocAutoscalingPolicyIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeDataprocAutoscalingPolicyIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeDataprocAutoscalingPolicyIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newDataprocAutoscalingPolicyIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//dataproc.googleapis.com/projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: DataprocAutoscalingPolicyIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataprocAutoscalingPolicyIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("policy_id"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		DataprocAutoscalingPolicyIamUpdaterProducer,
		d,
		config,
		"//dataproc.googleapis.com/projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}",
		DataprocAutoscalingPolicyIAMAssetType,
	)
}
