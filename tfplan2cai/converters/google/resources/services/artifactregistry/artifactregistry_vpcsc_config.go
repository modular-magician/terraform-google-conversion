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

package artifactregistry

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ArtifactRegistryVPCSCConfigAssetType string = "artifactregistry.googleapis.com/VPCSCConfig"

func ResourceConverterArtifactRegistryVPCSCConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ArtifactRegistryVPCSCConfigAssetType,
		Convert:   GetArtifactRegistryVPCSCConfigCaiObject,
	}
}

func GetArtifactRegistryVPCSCConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//artifactregistry.googleapis.com/projects/{{project}}/locations/{{location}}/vpcscConfig")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetArtifactRegistryVPCSCConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ArtifactRegistryVPCSCConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/artifactregistry/v1/rest",
				DiscoveryName:        "VPCSCConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetArtifactRegistryVPCSCConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	vpcscPolicyProp, err := expandArtifactRegistryVPCSCConfigVpcscPolicy(d.Get("vpcsc_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpcsc_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcscPolicyProp)) && (ok || !reflect.DeepEqual(v, vpcscPolicyProp)) {
		obj["vpcscPolicy"] = vpcscPolicyProp
	}

	return resourceArtifactRegistryVPCSCConfigEncoder(d, config, obj)
}

func resourceArtifactRegistryVPCSCConfigEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	if _, ok := d.GetOk("location"); !ok {
		location, err := tpgresource.GetRegionFromSchema("region", "zone", d, config)
		if err != nil {
			return nil, fmt.Errorf("Cannot determine location: set in this resource, or set provider-level 'region' or 'zone'.")
		}
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	return obj, nil
}

func expandArtifactRegistryVPCSCConfigVpcscPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
