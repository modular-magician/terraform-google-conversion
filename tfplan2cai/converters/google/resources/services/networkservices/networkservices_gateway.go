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

package networkservices

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Checks if there is another gateway under the same location.
func gatewaysSameLocation(d *schema.ResourceData, config *transport_tpg.Config, billingProject, userAgent string) ([]interface{}, error) {
	log.Print("[DEBUG] Looking for gateways under the same location.")
	var gateways []interface{}

	gatewaysUrl, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways")
	if err != nil {
		return gateways, err
	}

	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    gatewaysUrl,
		UserAgent: userAgent,
	})
	if err != nil {
		return gateways, err
	}

	data, ok := resp["gateways"]
	if !ok || data == nil {
		log.Print("[DEBUG] No gateways under the same location found.")
		return gateways, nil
	}

	gateways = data.([]interface{})

	log.Printf("[DEBUG] There are still gateways under the same location: %#v", gateways)

	return gateways, nil
}

// Checks if the given list of gateways contains a gateway of type SECURE_WEB_GATEWAY.
func isLastSWGGateway(gateways []interface{}, network string) bool {
	log.Print("[DEBUG] Checking if this is the last gateway of type SECURE_WEB_GATEWAY.")
	for _, itemRaw := range gateways {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		gType, ok := item["type"]
		if !ok || gType == nil {
			continue
		}

		gNetwork, ok := item["network"]
		if !ok || gNetwork == nil {
			continue
		}

		if gType.(string) == "SECURE_WEB_GATEWAY" && gNetwork.(string) == network {
			return false
		}
	}

	log.Print("[DEBUG] There is no other gateway of type SECURE_WEB_GATEWAY.")
	// no gateways of type SWG found.
	return true
}

// Deletes the swg-autogen-router if the current gateway being deleted is the type of swg so there is no other gateway using it.
func deleteSWGAutoGenRouter(d *schema.ResourceData, config *transport_tpg.Config, billingProject, userAgent string) error {
	log.Printf("[DEBUG] Searching the network id by name %q.", d.Get("network"))

	networkPath := fmt.Sprintf("{{ComputeBasePath}}%s", d.Get("network"))
	networkUrl, err := tpgresource.ReplaceVars(d, config, networkPath)
	if err != nil {
		return err
	}

	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    networkUrl,
		UserAgent: userAgent,
	})
	if err != nil {
		return err
	}

	// The name of swg auto generated router is in the following format: swg-autogen-router-{NETWORK-ID}
	routerId := fmt.Sprintf("swg-autogen-router-%s", resp["id"])
	log.Printf("[DEBUG] Deleting the auto generated router %q.", routerId)

	routerPath := fmt.Sprintf("{{ComputeBasePath}}projects/{{project}}/regions/{{location}}/routers/%s", routerId)
	routerUrl, err := tpgresource.ReplaceVars(d, config, routerPath)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               routerUrl,
		UserAgent:            userAgent,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsSwgAutogenRouterRetryable},
	})
	if err != nil {
		if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
			// The swg auto gen router may have already been deleted.
			// No further action needed.
			return nil
		}

		return err
	}

	return nil
}

const NetworkServicesGatewayAssetType string = "networkservices.googleapis.com/Gateway"

func ResourceConverterNetworkServicesGateway() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: NetworkServicesGatewayAssetType,
		Convert:   GetNetworkServicesGatewayCaiObject,
	}
}

func GetNetworkServicesGatewayCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/{{location}}/gateways/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetNetworkServicesGatewayApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: NetworkServicesGatewayAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "Gateway",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetNetworkServicesGatewayApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandNetworkServicesGatewayLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandNetworkServicesGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	typeProp, err := expandNetworkServicesGatewayType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	portsProp, err := expandNetworkServicesGatewayPorts(d.Get("ports"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ports"); !tpgresource.IsEmptyValue(reflect.ValueOf(portsProp)) && (ok || !reflect.DeepEqual(v, portsProp)) {
		obj["ports"] = portsProp
	}
	scopeProp, err := expandNetworkServicesGatewayScope(d.Get("scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}
	serverTlsPolicyProp, err := expandNetworkServicesGatewayServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverTlsPolicyProp)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}
	addressesProp, err := expandNetworkServicesGatewayAddresses(d.Get("addresses"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("addresses"); !tpgresource.IsEmptyValue(reflect.ValueOf(addressesProp)) && (ok || !reflect.DeepEqual(v, addressesProp)) {
		obj["addresses"] = addressesProp
	}
	subnetworkProp, err := expandNetworkServicesGatewaySubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	networkProp, err := expandNetworkServicesGatewayNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	gatewaySecurityPolicyProp, err := expandNetworkServicesGatewayGatewaySecurityPolicy(d.Get("gateway_security_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gateway_security_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewaySecurityPolicyProp)) && (ok || !reflect.DeepEqual(v, gatewaySecurityPolicyProp)) {
		obj["gatewaySecurityPolicy"] = gatewaySecurityPolicyProp
	}
	certificateUrlsProp, err := expandNetworkServicesGatewayCertificateUrls(d.Get("certificate_urls"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate_urls"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateUrlsProp)) && (ok || !reflect.DeepEqual(v, certificateUrlsProp)) {
		obj["certificateUrls"] = certificateUrlsProp
	}

	return obj, nil
}

func expandNetworkServicesGatewayLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesGatewayDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayServerTlsPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewaySubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayGatewaySecurityPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayCertificateUrls(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
