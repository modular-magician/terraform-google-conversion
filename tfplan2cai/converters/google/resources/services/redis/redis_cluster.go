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

package redis

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const RedisClusterAssetType string = "redis.googleapis.com/Cluster"

func ResourceConverterRedisCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: RedisClusterAssetType,
		Convert:   GetRedisClusterCaiObject,
	}
}

func GetRedisClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//redis.googleapis.com/projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetRedisClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: RedisClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/redis/v1beta1/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetRedisClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	authorizationModeProp, err := expandRedisClusterAuthorizationMode(d.Get("authorization_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorization_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationModeProp)) && (ok || !reflect.DeepEqual(v, authorizationModeProp)) {
		obj["authorizationMode"] = authorizationModeProp
	}
	transitEncryptionModeProp, err := expandRedisClusterTransitEncryptionMode(d.Get("transit_encryption_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transit_encryption_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(transitEncryptionModeProp)) && (ok || !reflect.DeepEqual(v, transitEncryptionModeProp)) {
		obj["transitEncryptionMode"] = transitEncryptionModeProp
	}
	nodeTypeProp, err := expandRedisClusterNodeType(d.Get("node_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeProp)) && (ok || !reflect.DeepEqual(v, nodeTypeProp)) {
		obj["nodeType"] = nodeTypeProp
	}
	zoneDistributionConfigProp, err := expandRedisClusterZoneDistributionConfig(d.Get("zone_distribution_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone_distribution_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneDistributionConfigProp)) && (ok || !reflect.DeepEqual(v, zoneDistributionConfigProp)) {
		obj["zoneDistributionConfig"] = zoneDistributionConfigProp
	}
	pscConfigsProp, err := expandRedisClusterPscConfigs(d.Get("psc_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("psc_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscConfigsProp)) && (ok || !reflect.DeepEqual(v, pscConfigsProp)) {
		obj["pscConfigs"] = pscConfigsProp
	}
	replicaCountProp, err := expandRedisClusterReplicaCount(d.Get("replica_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replica_count"); ok || !reflect.DeepEqual(v, replicaCountProp) {
		obj["replicaCount"] = replicaCountProp
	}
	shardCountProp, err := expandRedisClusterShardCount(d.Get("shard_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shard_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(shardCountProp)) && (ok || !reflect.DeepEqual(v, shardCountProp)) {
		obj["shardCount"] = shardCountProp
	}
	redisConfigsProp, err := expandRedisClusterRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redis_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}

	return obj, nil
}

func expandRedisClusterAuthorizationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterTransitEncryptionMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterNodeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterZoneDistributionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandRedisClusterZoneDistributionConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedZone, err := expandRedisClusterZoneDistributionConfigZone(original["zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["zone"] = transformedZone
	}

	return transformed, nil
}

func expandRedisClusterZoneDistributionConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterZoneDistributionConfigZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterPscConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandRedisClusterPscConfigsNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterPscConfigsNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterShardCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterRedisConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
