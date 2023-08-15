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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityServerTlsPolicyAssetType string = "networksecurity.googleapis.com/ServerTlsPolicy"

func ResourceConverterNetworkSecurityServerTlsPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecurityServerTlsPolicyAssetType,
		Convert:   GetNetworkSecurityServerTlsPolicyCaiObject,
	}
}

func GetNetworkSecurityServerTlsPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecurityServerTlsPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecurityServerTlsPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "ServerTlsPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecurityServerTlsPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkSecurityServerTlsPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkSecurityServerTlsPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	allowOpenProp, err := expandNetworkSecurityServerTlsPolicyAllowOpen(d.Get("allow_open"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_open"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowOpenProp)) && (ok || !reflect.DeepEqual(v, allowOpenProp)) {
		obj["allowOpen"] = allowOpenProp
	}
	serverCertificateProp, err := expandNetworkSecurityServerTlsPolicyServerCertificate(d.Get("server_certificate"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("server_certificate"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverCertificateProp)) && (ok || !reflect.DeepEqual(v, serverCertificateProp)) {
		obj["serverCertificate"] = serverCertificateProp
	}
	mtlsPolicyProp, err := expandNetworkSecurityServerTlsPolicyMtlsPolicy(d.Get("mtls_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mtls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(mtlsPolicyProp)) && (ok || !reflect.DeepEqual(v, mtlsPolicyProp)) {
		obj["mtlsPolicy"] = mtlsPolicyProp
	}

	return obj, nil
}

func expandNetworkSecurityServerTlsPolicyLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkSecurityServerTlsPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyAllowOpen(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGrpcEndpoint, err := expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint(original["grpc_endpoint"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["grpcEndpoint"] = transformedGrpcEndpoint
	}

	transformedCertificateProviderInstance, err := expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance(original["certificate_provider_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificateProviderInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["certificateProviderInstance"] = transformedCertificateProviderInstance
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetUri, err := expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointTargetUri(original["target_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetUri"] = transformedTargetUri
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateGrpcEndpointTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPluginInstance, err := expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstancePluginInstance(original["plugin_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPluginInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pluginInstance"] = transformedPluginInstance
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyServerCertificateCertificateProviderInstancePluginInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClientValidationMode, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationMode(original["client_validation_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientValidationMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientValidationMode"] = transformedClientValidationMode
	}

	transformedClientValidationTrustConfig, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationTrustConfig(original["client_validation_trust_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientValidationTrustConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientValidationTrustConfig"] = transformedClientValidationTrustConfig
	}

	transformedClientValidationCa, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa(original["client_validation_ca"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientValidationCa); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientValidationCa"] = transformedClientValidationCa
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationTrustConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCa(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGrpcEndpoint, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(original["grpc_endpoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGrpcEndpoint); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["grpcEndpoint"] = transformedGrpcEndpoint
		}

		transformedCertificateProviderInstance, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(original["certificate_provider_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCertificateProviderInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["certificateProviderInstance"] = transformedCertificateProviderInstance
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetUri, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointTargetUri(original["target_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetUri"] = transformedTargetUri
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointTargetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPluginInstance, err := expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstancePluginInstance(original["plugin_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPluginInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pluginInstance"] = transformedPluginInstance
	}

	return transformed, nil
}

func expandNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstancePluginInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
