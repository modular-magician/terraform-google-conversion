// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func rrdatasDnsDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	o, n := d.GetChange("rrdatas")
	if o == nil || n == nil {
		return false
	}

	oList := convertStringArr(o.([]interface{}))
	nList := convertStringArr(n.([]interface{}))

	parseFunc := func(record string) string {
		switch d.Get("type") {
		case "AAAA":
			// parse ipv6 to a key from one list
			return net.ParseIP(record).String()
		case "MX":
			return strings.ToLower(record)
		case "TXT":
			return strings.ToLower(strings.Trim(record, `"`))
		default:
			return record
		}
	}
	return rrdatasListDiffSuppress(oList, nList, parseFunc, d)
}

// suppress on a list when 1) its items have dups that need to be ignored
// and 2) string comparison on the items may need a special parse function
// example of usage can be found ../../../third_party/terraform/tests/resource_dns_record_set_test.go.erb
func rrdatasListDiffSuppress(oldList, newList []string, fun func(x string) string, _ *schema.ResourceData) bool {
	// compare two lists of unordered records
	diff := make(map[string]bool, len(oldList))
	for _, oldRecord := range oldList {
		// set all new IPs to true
		diff[fun(oldRecord)] = true
	}
	for _, newRecord := range newList {
		// set matched IPs to false otherwise can't suppress
		if diff[fun(newRecord)] {
			diff[fun(newRecord)] = false
		} else {
			return false
		}
	}
	// can't suppress if unmatched records are found
	for _, element := range diff {
		if element {
			return false
		}
	}
	return true
}

func GetDNSResourceDnsRecordSetCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dns.googleapis.com/projects/{{project}}/managedZones/{{managed_zone}}/rrsets/{{name}}/{{type}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDNSResourceDnsRecordSetApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "dns.googleapis.com/ResourceDnsRecordSet",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dns/v1/rest",
				DiscoveryName:        "ResourceDnsRecordSet",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDNSResourceDnsRecordSetApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandDNSResourceDnsRecordSetName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	typeProp, err := expandDNSResourceDnsRecordSetType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	ttlProp, err := expandDNSResourceDnsRecordSetTtl(d.Get("ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ttl"); !isEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	rrdatasProp, err := expandDNSResourceDnsRecordSetRrdatas(d.Get("rrdatas"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rrdatas"); !isEmptyValue(reflect.ValueOf(rrdatasProp)) && (ok || !reflect.DeepEqual(v, rrdatasProp)) {
		obj["rrdatas"] = rrdatasProp
	}
	managed_zoneProp, err := expandDNSResourceDnsRecordSetManagedZone(d.Get("managed_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("managed_zone"); !isEmptyValue(reflect.ValueOf(managed_zoneProp)) && (ok || !reflect.DeepEqual(v, managed_zoneProp)) {
		obj["managed_zone"] = managed_zoneProp
	}

	return obj, nil
}

func expandDNSResourceDnsRecordSetName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSResourceDnsRecordSetType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSResourceDnsRecordSetTtl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSResourceDnsRecordSetRrdatas(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDNSResourceDnsRecordSetManagedZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("managedZones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for managed_zone: %s", err)
	}
	return f.RelativeLink(), nil
}
