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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ComputeServiceAttachmentAssetType string = "compute.googleapis.com/ServiceAttachment"

func ResourceConverterComputeServiceAttachment() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: ComputeServiceAttachmentAssetType,
		Convert:   GetComputeServiceAttachmentCaiObject,
	}
}

func GetComputeServiceAttachmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetComputeServiceAttachmentApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: ComputeServiceAttachmentAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "ServiceAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetComputeServiceAttachmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeServiceAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeServiceAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeServiceAttachmentFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	connectionPreferenceProp, err := expandComputeServiceAttachmentConnectionPreference(d.Get("connection_preference"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connection_preference"); !tpgresource.IsEmptyValue(reflect.ValueOf(connectionPreferenceProp)) && (ok || !reflect.DeepEqual(v, connectionPreferenceProp)) {
		obj["connectionPreference"] = connectionPreferenceProp
	}
	targetServiceProp, err := expandComputeServiceAttachmentTargetService(d.Get("target_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetServiceProp)) && (ok || !reflect.DeepEqual(v, targetServiceProp)) {
		obj["targetService"] = targetServiceProp
	}
	natSubnetsProp, err := expandComputeServiceAttachmentNatSubnets(d.Get("nat_subnets"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("nat_subnets"); ok || !reflect.DeepEqual(v, natSubnetsProp) {
		obj["natSubnets"] = natSubnetsProp
	}
	enableProxyProtocolProp, err := expandComputeServiceAttachmentEnableProxyProtocol(d.Get("enable_proxy_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_proxy_protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableProxyProtocolProp)) && (ok || !reflect.DeepEqual(v, enableProxyProtocolProp)) {
		obj["enableProxyProtocol"] = enableProxyProtocolProp
	}
	domainNamesProp, err := expandComputeServiceAttachmentDomainNames(d.Get("domain_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("domain_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(domainNamesProp)) && (ok || !reflect.DeepEqual(v, domainNamesProp)) {
		obj["domainNames"] = domainNamesProp
	}
	consumerRejectListsProp, err := expandComputeServiceAttachmentConsumerRejectLists(d.Get("consumer_reject_lists"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_reject_lists"); ok || !reflect.DeepEqual(v, consumerRejectListsProp) {
		obj["consumerRejectLists"] = consumerRejectListsProp
	}
	consumerAcceptListsProp, err := expandComputeServiceAttachmentConsumerAcceptLists(d.Get("consumer_accept_lists"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_accept_lists"); ok || !reflect.DeepEqual(v, consumerAcceptListsProp) {
		obj["consumerAcceptLists"] = consumerAcceptListsProp
	}
	regionProp, err := expandComputeServiceAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeServiceAttachmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConnectionPreference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentTargetService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("forwardingRules", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeServiceAttachmentNatSubnets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for nat_subnets: nil")
		}
		f, err := tpgresource.ParseRegionalFieldValue("subnetworks", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for nat_subnets: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeServiceAttachmentEnableProxyProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentDomainNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerRejectLists(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerAcceptLists(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectIdOrNum, err := expandComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(original["project_id_or_num"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectIdOrNum); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["projectIdOrNum"] = transformedProjectIdOrNum
		}

		transformedConnectionLimit, err := expandComputeServiceAttachmentConsumerAcceptListsConnectionLimit(original["connection_limit"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConnectionLimit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["connectionLimit"] = transformedConnectionLimit
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerAcceptListsConnectionLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
