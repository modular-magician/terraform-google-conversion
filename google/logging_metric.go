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

func GetLoggingMetricCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//logging.googleapis.com/projects/{{project}}/metrics/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetLoggingMetricApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "logging.googleapis.com/Metric",
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/logging/v2/rest",
				DiscoveryName:        "Metric",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetLoggingMetricApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandLoggingMetricName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandLoggingMetricDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	filterProp, err := expandLoggingMetricFilter(d.Get("filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	metricDescriptorProp, err := expandLoggingMetricMetricDescriptor(d.Get("metric_descriptor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metric_descriptor"); !isEmptyValue(reflect.ValueOf(metricDescriptorProp)) && (ok || !reflect.DeepEqual(v, metricDescriptorProp)) {
		obj["metricDescriptor"] = metricDescriptorProp
	}
	labelExtractorsProp, err := expandLoggingMetricLabelExtractors(d.Get("label_extractors"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_extractors"); !isEmptyValue(reflect.ValueOf(labelExtractorsProp)) && (ok || !reflect.DeepEqual(v, labelExtractorsProp)) {
		obj["labelExtractors"] = labelExtractorsProp
	}
	valueExtractorProp, err := expandLoggingMetricValueExtractor(d.Get("value_extractor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("value_extractor"); !isEmptyValue(reflect.ValueOf(valueExtractorProp)) && (ok || !reflect.DeepEqual(v, valueExtractorProp)) {
		obj["valueExtractor"] = valueExtractorProp
	}
	bucketOptionsProp, err := expandLoggingMetricBucketOptions(d.Get("bucket_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket_options"); !isEmptyValue(reflect.ValueOf(bucketOptionsProp)) && (ok || !reflect.DeepEqual(v, bucketOptionsProp)) {
		obj["bucketOptions"] = bucketOptionsProp
	}

	return obj, nil
}

func expandLoggingMetricName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricFilter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUnit, err := expandLoggingMetricMetricDescriptorUnit(original["unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnit); val.IsValid() && !isEmptyValue(val) {
		transformed["unit"] = transformedUnit
	}

	transformedValueType, err := expandLoggingMetricMetricDescriptorValueType(original["value_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !isEmptyValue(val) {
		transformed["valueType"] = transformedValueType
	}

	transformedMetricKind, err := expandLoggingMetricMetricDescriptorMetricKind(original["metric_kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetricKind); val.IsValid() && !isEmptyValue(val) {
		transformed["metricKind"] = transformedMetricKind
	}

	transformedLabels, err := expandLoggingMetricMetricDescriptorLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandLoggingMetricMetricDescriptorUnit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorMetricKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandLoggingMetricMetricDescriptorLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedDescription, err := expandLoggingMetricMetricDescriptorLabelsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedValueType, err := expandLoggingMetricMetricDescriptorLabelsValueType(original["value_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !isEmptyValue(val) {
			transformed["valueType"] = transformedValueType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandLoggingMetricMetricDescriptorLabelsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabelsDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabelsValueType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricLabelExtractors(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandLoggingMetricValueExtractor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLinearBuckets, err := expandLoggingMetricBucketOptionsLinearBuckets(original["linear_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLinearBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["linearBuckets"] = transformedLinearBuckets
	}

	transformedExponentialBuckets, err := expandLoggingMetricBucketOptionsExponentialBuckets(original["exponential_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExponentialBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["exponentialBuckets"] = transformedExponentialBuckets
	}

	transformedExplicitBuckets, err := expandLoggingMetricBucketOptionsExplicitBuckets(original["explicit_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExplicitBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["explicitBuckets"] = transformedExplicitBuckets
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsLinearBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNumFiniteBuckets, err := expandLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(original["num_finite_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumFiniteBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["numFiniteBuckets"] = transformedNumFiniteBuckets
	}

	transformedWidth, err := expandLoggingMetricBucketOptionsLinearBucketsWidth(original["width"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWidth); val.IsValid() && !isEmptyValue(val) {
		transformed["width"] = transformedWidth
	}

	transformedOffset, err := expandLoggingMetricBucketOptionsLinearBucketsOffset(original["offset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOffset); val.IsValid() && !isEmptyValue(val) {
		transformed["offset"] = transformedOffset
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsWidth(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsOffset(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNumFiniteBuckets, err := expandLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(original["num_finite_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumFiniteBuckets); val.IsValid() && !isEmptyValue(val) {
		transformed["numFiniteBuckets"] = transformedNumFiniteBuckets
	}

	transformedGrowthFactor, err := expandLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(original["growth_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrowthFactor); val.IsValid() && !isEmptyValue(val) {
		transformed["growthFactor"] = transformedGrowthFactor
	}

	transformedScale, err := expandLoggingMetricBucketOptionsExponentialBucketsScale(original["scale"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScale); val.IsValid() && !isEmptyValue(val) {
		transformed["scale"] = transformedScale
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsScale(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExplicitBuckets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBounds, err := expandLoggingMetricBucketOptionsExplicitBucketsBounds(original["bounds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBounds); val.IsValid() && !isEmptyValue(val) {
		transformed["bounds"] = transformedBounds
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsExplicitBucketsBounds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
