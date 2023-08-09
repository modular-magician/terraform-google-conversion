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

package containeranalysis

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ContainerAnalysisNoteIAMAssetType string = "containeranalysis.googleapis.com/Note"

func ResourceConverterContainerAnalysisNoteIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ContainerAnalysisNoteIAMAssetType,
		Convert:           GetContainerAnalysisNoteIamPolicyCaiObject,
		MergeCreateUpdate: MergeContainerAnalysisNoteIamPolicy,
	}
}

func ResourceConverterContainerAnalysisNoteIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ContainerAnalysisNoteIAMAssetType,
		Convert:           GetContainerAnalysisNoteIamBindingCaiObject,
		FetchFullResource: FetchContainerAnalysisNoteIamPolicy,
		MergeCreateUpdate: MergeContainerAnalysisNoteIamBinding,
		MergeDelete:       MergeContainerAnalysisNoteIamBindingDelete,
	}
}

func ResourceConverterContainerAnalysisNoteIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ContainerAnalysisNoteIAMAssetType,
		Convert:           GetContainerAnalysisNoteIamMemberCaiObject,
		FetchFullResource: FetchContainerAnalysisNoteIamPolicy,
		MergeCreateUpdate: MergeContainerAnalysisNoteIamMember,
		MergeDelete:       MergeContainerAnalysisNoteIamMemberDelete,
	}
}

func GetContainerAnalysisNoteIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newContainerAnalysisNoteIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetContainerAnalysisNoteIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newContainerAnalysisNoteIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetContainerAnalysisNoteIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newContainerAnalysisNoteIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeContainerAnalysisNoteIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeContainerAnalysisNoteIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeContainerAnalysisNoteIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeContainerAnalysisNoteIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeContainerAnalysisNoteIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newContainerAnalysisNoteIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//containeranalysis.googleapis.com/projects/{{project}}/notes/{{note}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ContainerAnalysisNoteIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchContainerAnalysisNoteIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("note"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ContainerAnalysisNoteIamUpdaterProducer,
		d,
		config,
		"//containeranalysis.googleapis.com/projects/{{project}}/notes/{{note}}",
		ContainerAnalysisNoteIAMAssetType,
	)
}
