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

const ComputeRegionCommitmentAssetType string = "compute.googleapis.com/RegionCommitment"

func ResourceConverterComputeRegionCommitment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeRegionCommitmentAssetType,
		Convert:   GetComputeRegionCommitmentCaiObject,
	}
}

func GetComputeRegionCommitmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/commitments/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeRegionCommitmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeRegionCommitmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionCommitment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeRegionCommitmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionCommitmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionCommitmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	planProp, err := expandComputeRegionCommitmentPlan(d.Get("plan"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("plan"); !tpgresource.IsEmptyValue(reflect.ValueOf(planProp)) && (ok || !reflect.DeepEqual(v, planProp)) {
		obj["plan"] = planProp
	}
	resourcesProp, err := expandComputeRegionCommitmentResources(d.Get("resources"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourcesProp)) && (ok || !reflect.DeepEqual(v, resourcesProp)) {
		obj["resources"] = resourcesProp
	}
	typeProp, err := expandComputeRegionCommitmentType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	categoryProp, err := expandComputeRegionCommitmentCategory(d.Get("category"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("category"); !tpgresource.IsEmptyValue(reflect.ValueOf(categoryProp)) && (ok || !reflect.DeepEqual(v, categoryProp)) {
		obj["category"] = categoryProp
	}
	licenseResourceProp, err := expandComputeRegionCommitmentLicenseResource(d.Get("license_resource"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("license_resource"); !tpgresource.IsEmptyValue(reflect.ValueOf(licenseResourceProp)) && (ok || !reflect.DeepEqual(v, licenseResourceProp)) {
		obj["licenseResource"] = licenseResourceProp
	}
	autoRenewProp, err := expandComputeRegionCommitmentAutoRenew(d.Get("auto_renew"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("auto_renew"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoRenewProp)) && (ok || !reflect.DeepEqual(v, autoRenewProp)) {
		obj["autoRenew"] = autoRenewProp
	}
	existingReservationsProp, err := expandComputeRegionCommitmentExistingReservations(d.Get("existing_reservations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("existing_reservations"); !tpgresource.IsEmptyValue(reflect.ValueOf(existingReservationsProp)) && (ok || !reflect.DeepEqual(v, existingReservationsProp)) {
		obj["existingReservations"] = existingReservationsProp
	}
	regionProp, err := expandComputeRegionCommitmentRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeRegionCommitmentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentPlan(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandComputeRegionCommitmentResourcesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedAmount, err := expandComputeRegionCommitmentResourcesAmount(original["amount"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAmount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["amount"] = transformedAmount
		}

		transformedAcceleratorType, err := expandComputeRegionCommitmentResourcesAcceleratorType(original["accelerator_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAcceleratorType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["acceleratorType"] = transformedAcceleratorType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionCommitmentResourcesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentResourcesAmount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentResourcesAcceleratorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentCategory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentLicenseResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLicense, err := expandComputeRegionCommitmentLicenseResourceLicense(original["license"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLicense); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["license"] = transformedLicense
	}

	transformedAmount, err := expandComputeRegionCommitmentLicenseResourceAmount(original["amount"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAmount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["amount"] = transformedAmount
	}

	transformedCoresPerLicense, err := expandComputeRegionCommitmentLicenseResourceCoresPerLicense(original["cores_per_license"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCoresPerLicense); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["coresPerLicense"] = transformedCoresPerLicense
	}

	return transformed, nil
}

func expandComputeRegionCommitmentLicenseResourceLicense(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentLicenseResourceAmount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentLicenseResourceCoresPerLicense(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentAutoRenew(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentExistingReservations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionCommitmentRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
