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

import "reflect"

func GetAccessContextManagerServicePerimeterCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//accesscontextmanager.googleapis.com/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetAccessContextManagerServicePerimeterApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "accesscontextmanager.googleapis.com/ServicePerimeter",
			Resource: &AssetResource{
				Version:              "",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accesscontextmanager//rest",
				DiscoveryName:        "ServicePerimeter",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetAccessContextManagerServicePerimeterApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerServicePerimeterTitle(d.Get("title"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerServicePerimeterDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	perimeterTypeProp, err := expandAccessContextManagerServicePerimeterPerimeterType(d.Get("perimeter_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("perimeter_type"); !isEmptyValue(reflect.ValueOf(perimeterTypeProp)) && (ok || !reflect.DeepEqual(v, perimeterTypeProp)) {
		obj["perimeterType"] = perimeterTypeProp
	}
	statusProp, err := expandAccessContextManagerServicePerimeterStatus(d.Get("status"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("status"); !isEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	parentProp, err := expandAccessContextManagerServicePerimeterParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerServicePerimeterName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return resourceAccessContextManagerServicePerimeterEncoder(d, config, obj)
}

func resourceAccessContextManagerServicePerimeterEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}

func expandAccessContextManagerServicePerimeterTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterPerimeterType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatus(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResources, err := expandAccessContextManagerServicePerimeterStatusResources(original["resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !isEmptyValue(val) {
		transformed["resources"] = transformedResources
	}

	transformedAccessLevels, err := expandAccessContextManagerServicePerimeterStatusAccessLevels(original["access_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["accessLevels"] = transformedAccessLevels
	}

	transformedRestrictedServices, err := expandAccessContextManagerServicePerimeterStatusRestrictedServices(original["restricted_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRestrictedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["restrictedServices"] = transformedRestrictedServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimeterStatusResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterStatusRestrictedServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimeterName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
