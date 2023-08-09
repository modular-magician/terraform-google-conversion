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

package iap

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const IapTunnelInstanceIAMAssetType string = "iap.googleapis.com/TunnelInstance"

func ResourceConverterIapTunnelInstanceIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         IapTunnelInstanceIAMAssetType,
		Convert:           GetIapTunnelInstanceIamPolicyCaiObject,
		MergeCreateUpdate: MergeIapTunnelInstanceIamPolicy,
	}
}

func ResourceConverterIapTunnelInstanceIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         IapTunnelInstanceIAMAssetType,
		Convert:           GetIapTunnelInstanceIamBindingCaiObject,
		FetchFullResource: FetchIapTunnelInstanceIamPolicy,
		MergeCreateUpdate: MergeIapTunnelInstanceIamBinding,
		MergeDelete:       MergeIapTunnelInstanceIamBindingDelete,
	}
}

func ResourceConverterIapTunnelInstanceIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         IapTunnelInstanceIAMAssetType,
		Convert:           GetIapTunnelInstanceIamMemberCaiObject,
		FetchFullResource: FetchIapTunnelInstanceIamPolicy,
		MergeCreateUpdate: MergeIapTunnelInstanceIamMember,
		MergeDelete:       MergeIapTunnelInstanceIamMemberDelete,
	}
}

func GetIapTunnelInstanceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newIapTunnelInstanceIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetIapTunnelInstanceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newIapTunnelInstanceIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetIapTunnelInstanceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newIapTunnelInstanceIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeIapTunnelInstanceIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapTunnelInstanceIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeIapTunnelInstanceIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeIapTunnelInstanceIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeIapTunnelInstanceIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newIapTunnelInstanceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//iap.googleapis.com/projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{instance}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: IapTunnelInstanceIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapTunnelInstanceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("zone"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("instance"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		IapTunnelInstanceIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{instance}}",
		IapTunnelInstanceIAMAssetType,
	)
}
