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

package dns

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DNSManagedZoneAssetType string = "dns.googleapis.com/ManagedZone"

func ResourceConverterDNSManagedZone() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DNSManagedZoneAssetType,
		Convert:   GetDNSManagedZoneCaiObject,
	}
}

func GetDNSManagedZoneCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dns.googleapis.com/projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDNSManagedZoneApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DNSManagedZoneAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dns/v1beta2/rest",
				DiscoveryName:        "ManagedZone",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDNSManagedZoneApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDNSManagedZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dnsNameProp, err := expandDNSManagedZoneDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dns_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnsNameProp)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	dnssecConfigProp, err := expandDNSManagedZoneDnssecConfig(d.Get("dnssec_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dnssec_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(dnssecConfigProp)) && (ok || !reflect.DeepEqual(v, dnssecConfigProp)) {
		obj["dnssecConfig"] = dnssecConfigProp
	}
	nameProp, err := expandDNSManagedZoneName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	visibilityProp, err := expandDNSManagedZoneVisibility(d.Get("visibility"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("visibility"); !tpgresource.IsEmptyValue(reflect.ValueOf(visibilityProp)) && (ok || !reflect.DeepEqual(v, visibilityProp)) {
		obj["visibility"] = visibilityProp
	}
	privateVisibilityConfigProp, err := expandDNSManagedZonePrivateVisibilityConfig(d.Get("private_visibility_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_visibility_config"); ok || !reflect.DeepEqual(v, privateVisibilityConfigProp) {
		obj["privateVisibilityConfig"] = privateVisibilityConfigProp
	}
	forwardingConfigProp, err := expandDNSManagedZoneForwardingConfig(d.Get("forwarding_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forwarding_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardingConfigProp)) && (ok || !reflect.DeepEqual(v, forwardingConfigProp)) {
		obj["forwardingConfig"] = forwardingConfigProp
	}
	peeringConfigProp, err := expandDNSManagedZonePeeringConfig(d.Get("peering_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(peeringConfigProp)) && (ok || !reflect.DeepEqual(v, peeringConfigProp)) {
		obj["peeringConfig"] = peeringConfigProp
	}
	reverseLookupConfigProp, err := expandDNSManagedZoneReverseLookup(d.Get("reverse_lookup"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reverse_lookup"); !tpgresource.IsEmptyValue(reflect.ValueOf(reverseLookupConfigProp)) && (ok || !reflect.DeepEqual(v, reverseLookupConfigProp)) {
		obj["reverseLookupConfig"] = reverseLookupConfigProp
	}
	serviceDirectoryConfigProp, err := expandDNSManagedZoneServiceDirectoryConfig(d.Get("service_directory_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_directory_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceDirectoryConfigProp)) && (ok || !reflect.DeepEqual(v, serviceDirectoryConfigProp)) {
		obj["serviceDirectoryConfig"] = serviceDirectoryConfigProp
	}
	cloudLoggingConfigProp, err := expandDNSManagedZoneCloudLoggingConfig(d.Get("cloud_logging_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cloud_logging_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloudLoggingConfigProp)) && (ok || !reflect.DeepEqual(v, cloudLoggingConfigProp)) {
		obj["cloudLoggingConfig"] = cloudLoggingConfigProp
	}
	labelsProp, err := expandDNSManagedZoneEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDNSManagedZoneDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKind, err := expandDNSManagedZoneDnssecConfigKind(original["kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kind"] = transformedKind
	}

	transformedNonExistence, err := expandDNSManagedZoneDnssecConfigNonExistence(original["non_existence"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNonExistence); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nonExistence"] = transformedNonExistence
	}

	transformedState, err := expandDNSManagedZoneDnssecConfigState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedDefaultKeySpecs, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecs(original["default_key_specs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultKeySpecs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultKeySpecs"] = transformedDefaultKeySpecs
	}

	return transformed, nil
}

func expandDNSManagedZoneDnssecConfigKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigNonExistence(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAlgorithm, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(original["algorithm"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["algorithm"] = transformedAlgorithm
		}

		transformedKeyLength, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(original["key_length"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyLength); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["keyLength"] = transformedKeyLength
		}

		transformedKeyType, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(original["key_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKeyType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["keyType"] = transformedKeyType
		}

		transformedKind, err := expandDNSManagedZoneDnssecConfigDefaultKeySpecsKind(original["kind"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKind); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["kind"] = transformedKind
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKeyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneDnssecConfigDefaultKeySpecsKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneVisibility(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZonePrivateVisibilityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		// The API won't remove the the field unless an empty network array is sent.
		transformed := make(map[string]interface{})
		emptyNetwork := make([]interface{}, 0)
		transformed["networks"] = emptyNetwork
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGkeClusters, err := expandDNSManagedZonePrivateVisibilityConfigGkeClusters(original["gke_clusters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGkeClusters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gkeClusters"] = transformedGkeClusters
	}

	transformedNetworks, err := expandDNSManagedZonePrivateVisibilityConfigNetworks(original["networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networks"] = transformedNetworks
	}

	return transformed, nil
}

func expandDNSManagedZonePrivateVisibilityConfigNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZonePrivateVisibilityConfigGkeClusters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGkeClusterName, err := expandDNSManagedZonePrivateVisibilityConfigGkeClustersGkeClusterName(original["gke_cluster_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGkeClusterName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["gkeClusterName"] = transformedGkeClusterName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZonePrivateVisibilityConfigNetworksNetworkUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return tpgresource.ConvertSelfLinkToV1(url), nil
}

func expandDNSManagedZonePrivateVisibilityConfigGkeClustersGkeClusterName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneForwardingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNameServers, err := expandDNSManagedZoneForwardingConfigTargetNameServers(original["target_name_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNameServers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetNameServers"] = transformedTargetNameServers
	}

	return transformed, nil
}

func expandDNSManagedZoneForwardingConfigTargetNameServers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpv4Address, err := expandDNSManagedZoneForwardingConfigTargetNameServersIpv4Address(original["ipv4_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv4Address); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipv4Address"] = transformedIpv4Address
		}

		transformedForwardingPath, err := expandDNSManagedZoneForwardingConfigTargetNameServersForwardingPath(original["forwarding_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedForwardingPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["forwardingPath"] = transformedForwardingPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSManagedZoneForwardingConfigTargetNameServersIpv4Address(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneForwardingConfigTargetNameServersForwardingPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZonePeeringConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNetwork, err := expandDNSManagedZonePeeringConfigTargetNetwork(original["target_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetNetwork"] = transformedTargetNetwork
	}

	return transformed, nil
}

func expandDNSManagedZonePeeringConfigTargetNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkUrl, err := expandDNSManagedZonePeeringConfigTargetNetworkNetworkUrl(original["network_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkUrl"] = transformedNetworkUrl
	}

	return transformed, nil
}

func expandDNSManagedZonePeeringConfigTargetNetworkNetworkUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return tpgresource.ConvertSelfLinkToV1(url), nil
}

func expandDNSManagedZoneReverseLookup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return struct{}{}, nil
}

func expandDNSManagedZoneServiceDirectoryConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNamespace, err := expandDNSManagedZoneServiceDirectoryConfigNamespace(original["namespace"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespace"] = transformedNamespace
	}

	return transformed, nil
}

func expandDNSManagedZoneServiceDirectoryConfigNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNamespaceUrl, err := expandDNSManagedZoneServiceDirectoryConfigNamespaceNamespaceUrl(original["namespace_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespaceUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespaceUrl"] = transformedNamespaceUrl
	}

	return transformed, nil
}

func expandDNSManagedZoneServiceDirectoryConfigNamespaceNamespaceUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := tpgresource.ReplaceVars(d, config, "{{ServiceDirectoryBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return url, nil
}

func expandDNSManagedZoneCloudLoggingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableLogging, err := expandDNSManagedZoneCloudLoggingConfigEnableLogging(original["enable_logging"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableLogging); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableLogging"] = transformedEnableLogging
	}

	return transformed, nil
}

func expandDNSManagedZoneCloudLoggingConfigEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSManagedZoneEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
