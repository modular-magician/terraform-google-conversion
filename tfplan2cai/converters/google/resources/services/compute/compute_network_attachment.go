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

package compute

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeNetworkAttachmentAssetType string = "compute.googleapis.com/NetworkAttachment"

func ResourceConverterComputeNetworkAttachment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNetworkAttachmentAssetType,
		Convert:   GetComputeNetworkAttachmentCaiObject,
	}
}

func GetComputeNetworkAttachmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/networkAttachments/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNetworkAttachmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNetworkAttachmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NetworkAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNetworkAttachmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNetworkAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	connectionPreferenceProp, err := expandComputeNetworkAttachmentConnectionPreference(d.Get("connection_preference"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connection_preference"); !tpgresource.IsEmptyValue(reflect.ValueOf(connectionPreferenceProp)) && (ok || !reflect.DeepEqual(v, connectionPreferenceProp)) {
		obj["connectionPreference"] = connectionPreferenceProp
	}
	subnetworksProp, err := expandComputeNetworkAttachmentSubnetworks(d.Get("subnetworks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetworks"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworksProp)) && (ok || !reflect.DeepEqual(v, subnetworksProp)) {
		obj["subnetworks"] = subnetworksProp
	}
	producerRejectListsProp, err := expandComputeNetworkAttachmentProducerRejectLists(d.Get("producer_reject_lists"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("producer_reject_lists"); !tpgresource.IsEmptyValue(reflect.ValueOf(producerRejectListsProp)) && (ok || !reflect.DeepEqual(v, producerRejectListsProp)) {
		obj["producerRejectLists"] = producerRejectListsProp
	}
	producerAcceptListsProp, err := expandComputeNetworkAttachmentProducerAcceptLists(d.Get("producer_accept_lists"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("producer_accept_lists"); !tpgresource.IsEmptyValue(reflect.ValueOf(producerAcceptListsProp)) && (ok || !reflect.DeepEqual(v, producerAcceptListsProp)) {
		obj["producerAcceptLists"] = producerAcceptListsProp
	}
	fingerprintProp, err := expandComputeNetworkAttachmentFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	nameProp, err := expandComputeNetworkAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	regionProp, err := expandComputeNetworkAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeNetworkAttachmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentConnectionPreference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentSubnetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for subnetworks: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("subnetworks", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for subnetworks: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeNetworkAttachmentProducerRejectLists(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentProducerAcceptLists(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAttachmentRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
