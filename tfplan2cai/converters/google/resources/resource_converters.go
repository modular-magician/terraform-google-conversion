// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//	This file is automatically generated by Magic Modules and manual
//	changes will be clobbered when the file is regenerated.
//
//	Please read more about how to change this file in
//	.github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/accesscontextmanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/apigee"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/artifactregistry"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/bigquery"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/bigqueryanalyticshub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/bigqueryconnection"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/bigquerydatapolicy"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/binaryauthorization"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/cloudfunctions"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/cloudfunctions2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/cloudrun"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/cloudrunv2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/cloudtasks"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/containeranalysis"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/datacatalog"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/datafusion"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/dataplex"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/dataproc"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/dataprocmetastore"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/dns"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/filestore"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/gkebackup"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/gkehub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/gkehub2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/healthcare"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/iap"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/kms"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/logging"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/monitoring"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/notebooks"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/privateca"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/pubsub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/pubsublite"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/redis"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/resourcemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/secretmanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/securitycenter"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/servicemanagement"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/spanner"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/sql"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/vertexai"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/vpcaccess"
)

// ResourceConverter returns a map of terraform resource types (i.e. `google_project`)
// to a slice of ResourceConverters.
//
// Modelling of relationships:
// terraform resources to CAI assets as []cai.ResourceConverter:
// 1:1 = [ResourceConverter{Convert: convertAbc}]                  (len=1)
// 1:N = [ResourceConverter{Convert: convertAbc}, ...]             (len=N)
// N:1 = [ResourceConverter{Convert: convertAbc, merge: mergeAbc}] (len=1)
func ResourceConverters() map[string][]cai.ResourceConverter {
	return map[string][]cai.ResourceConverter{
		"google_compute_address":                                  {compute.ResourceConverterComputeAddress()},
		"google_compute_firewall":                                 {compute.ResourceConverterComputeFirewall()},
		"google_compute_disk":                                     {compute.ResourceConverterComputeDisk()},
		"google_compute_forwarding_rule":                          {compute.ResourceConverterComputeForwardingRule()},
		"google_compute_global_address":                           {compute.ResourceConverterComputeGlobalAddress()},
		"google_compute_global_forwarding_rule":                   {compute.ResourceConverterComputeGlobalForwardingRule()},
		"google_compute_instance":                                 {compute.ResourceConverterComputeInstance()},
		"google_compute_network":                                  {compute.ResourceConverterComputeNetwork()},
		"google_compute_security_policy":                          {resourceConverterComputeSecurityPolicy()},
		"google_compute_snapshot":                                 {compute.ResourceConverterComputeSnapshot()},
		"google_compute_subnetwork":                               {compute.ResourceConverterComputeSubnetwork()},
		"google_compute_ssl_policy":                               {compute.ResourceConverterComputeSslPolicy()},
		"google_compute_target_https_proxy":                       {compute.ResourceConverterComputeTargetHttpsProxy()},
		"google_compute_target_ssl_proxy":                         {compute.ResourceConverterComputeTargetSslProxy()},
		"google_dns_managed_zone":                                 {dns.ResourceConverterDNSManagedZone()},
		"google_dns_policy":                                       {dns.ResourceConverterDNSPolicy()},
		"google_storage_bucket":                                   {resourceConverterStorageBucket()},
		"google_sql_database_instance":                            {resourceConverterSQLDatabaseInstance()},
		"google_sql_database":                                     {sql.ResourceConverterSQLDatabase()},
		"google_container_cluster":                                {resourceConverterContainerCluster()},
		"google_container_node_pool":                              {resourceConverterContainerNodePool()},
		"google_bigquery_dataset":                                 {bigquery.ResourceConverterBigQueryDataset()},
		"google_bigquery_dataset_iam_policy":                      {bigquery.ResourceConverterBigqueryDatasetIamPolicy()},
		"google_bigquery_dataset_iam_binding":                     {bigquery.ResourceConverterBigqueryDatasetIamBinding()},
		"google_bigquery_dataset_iam_member":                      {bigquery.ResourceConverterBigqueryDatasetIamMember()},
		"google_bigquery_table":                                   {resourceConverterBigQueryTable()},
		"google_org_policy_policy":                                {resourceConverterOrgPolicyPolicy()},
		"google_redis_instance":                                   {redis.ResourceConverterRedisInstance()},
		"google_spanner_database":                                 {spanner.ResourceConverterSpannerDatabase()},
		"google_spanner_database_iam_policy":                      {spanner.ResourceConverterSpannerDatabaseIamPolicy()},
		"google_spanner_database_iam_binding":                     {spanner.ResourceConverterSpannerDatabaseIamBinding()},
		"google_spanner_database_iam_member":                      {spanner.ResourceConverterSpannerDatabaseIamMember()},
		"google_spanner_instance":                                 {spanner.ResourceConverterSpannerInstance()},
		"google_spanner_instance_iam_policy":                      {spanner.ResourceConverterSpannerInstanceIamPolicy()},
		"google_spanner_instance_iam_binding":                     {spanner.ResourceConverterSpannerInstanceIamBinding()},
		"google_spanner_instance_iam_member":                      {spanner.ResourceConverterSpannerInstanceIamMember()},
		"google_project_service":                                  {resourceConverterServiceUsage()},
		"google_pubsub_lite_reservation":                          {pubsublite.ResourceConverterPubsubLiteReservation()},
		"google_pubsub_lite_subscription":                         {pubsublite.ResourceConverterPubsubLiteSubscription()},
		"google_pubsub_lite_topic":                                {pubsublite.ResourceConverterPubsubLiteTopic()},
		"google_pubsub_schema":                                    {pubsub.ResourceConverterPubsubSchema()},
		"google_pubsub_subscription":                              {pubsub.ResourceConverterPubsubSubscription()},
		"google_pubsub_subscription_iam_policy":                   {pubsub.ResourceConverterPubsubSubscriptionIamPolicy()},
		"google_pubsub_subscription_iam_binding":                  {pubsub.ResourceConverterPubsubSubscriptionIamBinding()},
		"google_pubsub_subscription_iam_member":                   {pubsub.ResourceConverterPubsubSubscriptionIamMember()},
		"google_storage_bucket_iam_policy":                        {resourceConverterStorageBucketIamPolicy()},
		"google_storage_bucket_iam_binding":                       {resourceConverterStorageBucketIamBinding()},
		"google_storage_bucket_iam_member":                        {resourceConverterStorageBucketIamMember()},
		"google_pubsub_topic":                                     {pubsub.ResourceConverterPubsubTopic()},
		"google_kms_crypto_key":                                   {kms.ResourceConverterKMSCryptoKey()},
		"google_kms_key_ring":                                     {kms.ResourceConverterKMSKeyRing()},
		"google_filestore_instance":                               {filestore.ResourceConverterFilestoreInstance()},
		"google_access_context_manager_service_perimeter":         {accesscontextmanager.ResourceConverterAccessContextManagerServicePerimeter()},
		"google_access_context_manager_access_policy":             {accesscontextmanager.ResourceConverterAccessContextManagerAccessPolicy()},
		"google_cloud_run_service":                                {cloudrun.ResourceConverterCloudRunService()},
		"google_cloud_run_domain_mapping":                         {cloudrun.ResourceConverterCloudRunDomainMapping()},
		"google_cloudfunctions_function":                          {resourceConverterCloudFunctionsCloudFunction()},
		"google_monitoring_notification_channel":                  {monitoring.ResourceConverterMonitoringNotificationChannel()},
		"google_monitoring_alert_policy":                          {monitoring.ResourceConverterMonitoringAlertPolicy()},
		"google_access_context_manager_access_policy_iam_policy":  {accesscontextmanager.ResourceConverterAccessContextManagerAccessPolicyIamPolicy()},
		"google_access_context_manager_access_policy_iam_binding": {accesscontextmanager.ResourceConverterAccessContextManagerAccessPolicyIamBinding()},
		"google_access_context_manager_access_policy_iam_member":  {accesscontextmanager.ResourceConverterAccessContextManagerAccessPolicyIamMember()},
		"google_apigee_environment_iam_policy":                    {apigee.ResourceConverterApigeeEnvironmentIamPolicy()},
		"google_apigee_environment_iam_binding":                   {apigee.ResourceConverterApigeeEnvironmentIamBinding()},
		"google_apigee_environment_iam_member":                    {apigee.ResourceConverterApigeeEnvironmentIamMember()},
		"google_artifact_registry_repository_iam_policy":          {artifactregistry.ResourceConverterArtifactRegistryRepositoryIamPolicy()},
		"google_artifact_registry_repository_iam_binding":         {artifactregistry.ResourceConverterArtifactRegistryRepositoryIamBinding()},
		"google_artifact_registry_repository_iam_member":          {artifactregistry.ResourceConverterArtifactRegistryRepositoryIamMember()},
		"google_bigquery_table_iam_policy":                        {bigquery.ResourceConverterBigQueryTableIamPolicy()},
		"google_bigquery_table_iam_binding":                       {bigquery.ResourceConverterBigQueryTableIamBinding()},
		"google_bigquery_table_iam_member":                        {bigquery.ResourceConverterBigQueryTableIamMember()},
		"google_bigquery_analytics_hub_data_exchange_iam_policy":  {bigqueryanalyticshub.ResourceConverterBigqueryAnalyticsHubDataExchangeIamPolicy()},
		"google_bigquery_analytics_hub_data_exchange_iam_binding": {bigqueryanalyticshub.ResourceConverterBigqueryAnalyticsHubDataExchangeIamBinding()},
		"google_bigquery_analytics_hub_data_exchange_iam_member":  {bigqueryanalyticshub.ResourceConverterBigqueryAnalyticsHubDataExchangeIamMember()},
		"google_bigquery_analytics_hub_listing_iam_policy":        {bigqueryanalyticshub.ResourceConverterBigqueryAnalyticsHubListingIamPolicy()},
		"google_bigquery_analytics_hub_listing_iam_binding":       {bigqueryanalyticshub.ResourceConverterBigqueryAnalyticsHubListingIamBinding()},
		"google_bigquery_analytics_hub_listing_iam_member":        {bigqueryanalyticshub.ResourceConverterBigqueryAnalyticsHubListingIamMember()},
		"google_bigquery_connection_iam_policy":                   {bigqueryconnection.ResourceConverterBigqueryConnectionConnectionIamPolicy()},
		"google_bigquery_connection_iam_binding":                  {bigqueryconnection.ResourceConverterBigqueryConnectionConnectionIamBinding()},
		"google_bigquery_connection_iam_member":                   {bigqueryconnection.ResourceConverterBigqueryConnectionConnectionIamMember()},
		"google_bigquery_datapolicy_data_policy_iam_policy":       {bigquerydatapolicy.ResourceConverterBigqueryDatapolicyDataPolicyIamPolicy()},
		"google_bigquery_datapolicy_data_policy_iam_binding":      {bigquerydatapolicy.ResourceConverterBigqueryDatapolicyDataPolicyIamBinding()},
		"google_bigquery_datapolicy_data_policy_iam_member":       {bigquerydatapolicy.ResourceConverterBigqueryDatapolicyDataPolicyIamMember()},
		"google_binary_authorization_attestor_iam_policy":         {binaryauthorization.ResourceConverterBinaryAuthorizationAttestorIamPolicy()},
		"google_binary_authorization_attestor_iam_binding":        {binaryauthorization.ResourceConverterBinaryAuthorizationAttestorIamBinding()},
		"google_binary_authorization_attestor_iam_member":         {binaryauthorization.ResourceConverterBinaryAuthorizationAttestorIamMember()},
		"google_cloudfunctions_function_iam_policy":               {cloudfunctions.ResourceConverterCloudFunctionsCloudFunctionIamPolicy()},
		"google_cloudfunctions_function_iam_binding":              {cloudfunctions.ResourceConverterCloudFunctionsCloudFunctionIamBinding()},
		"google_cloudfunctions_function_iam_member":               {cloudfunctions.ResourceConverterCloudFunctionsCloudFunctionIamMember()},
		"google_cloudfunctions2_function_iam_policy":              {cloudfunctions2.ResourceConverterCloudfunctions2functionIamPolicy()},
		"google_cloudfunctions2_function_iam_binding":             {cloudfunctions2.ResourceConverterCloudfunctions2functionIamBinding()},
		"google_cloudfunctions2_function_iam_member":              {cloudfunctions2.ResourceConverterCloudfunctions2functionIamMember()},
		"google_cloud_run_service_iam_policy":                     {cloudrun.ResourceConverterCloudRunServiceIamPolicy()},
		"google_cloud_run_service_iam_binding":                    {cloudrun.ResourceConverterCloudRunServiceIamBinding()},
		"google_cloud_run_service_iam_member":                     {cloudrun.ResourceConverterCloudRunServiceIamMember()},
		"google_cloud_run_v2_job_iam_policy":                      {cloudrunv2.ResourceConverterCloudRunV2JobIamPolicy()},
		"google_cloud_run_v2_job_iam_binding":                     {cloudrunv2.ResourceConverterCloudRunV2JobIamBinding()},
		"google_cloud_run_v2_job_iam_member":                      {cloudrunv2.ResourceConverterCloudRunV2JobIamMember()},
		"google_cloud_run_v2_service_iam_policy":                  {cloudrunv2.ResourceConverterCloudRunV2ServiceIamPolicy()},
		"google_cloud_run_v2_service_iam_binding":                 {cloudrunv2.ResourceConverterCloudRunV2ServiceIamBinding()},
		"google_cloud_run_v2_service_iam_member":                  {cloudrunv2.ResourceConverterCloudRunV2ServiceIamMember()},
		"google_cloud_tasks_queue_iam_policy":                     {cloudtasks.ResourceConverterCloudTasksQueueIamPolicy()},
		"google_cloud_tasks_queue_iam_binding":                    {cloudtasks.ResourceConverterCloudTasksQueueIamBinding()},
		"google_cloud_tasks_queue_iam_member":                     {cloudtasks.ResourceConverterCloudTasksQueueIamMember()},
		"google_compute_backend_bucket_iam_policy":                {compute.ResourceConverterComputeBackendBucketIamPolicy()},
		"google_compute_backend_bucket_iam_binding":               {compute.ResourceConverterComputeBackendBucketIamBinding()},
		"google_compute_backend_bucket_iam_member":                {compute.ResourceConverterComputeBackendBucketIamMember()},
		"google_compute_backend_service_iam_policy":               {compute.ResourceConverterComputeBackendServiceIamPolicy()},
		"google_compute_backend_service_iam_binding":              {compute.ResourceConverterComputeBackendServiceIamBinding()},
		"google_compute_backend_service_iam_member":               {compute.ResourceConverterComputeBackendServiceIamMember()},
		"google_compute_disk_iam_policy":                          {compute.ResourceConverterComputeDiskIamPolicy()},
		"google_compute_disk_iam_binding":                         {compute.ResourceConverterComputeDiskIamBinding()},
		"google_compute_disk_iam_member":                          {compute.ResourceConverterComputeDiskIamMember()},
		"google_compute_image_iam_policy":                         {compute.ResourceConverterComputeImageIamPolicy()},
		"google_compute_image_iam_binding":                        {compute.ResourceConverterComputeImageIamBinding()},
		"google_compute_image_iam_member":                         {compute.ResourceConverterComputeImageIamMember()},
		"google_compute_instance_iam_policy":                      {compute.ResourceConverterComputeInstanceIamPolicy()},
		"google_compute_instance_iam_binding":                     {compute.ResourceConverterComputeInstanceIamBinding()},
		"google_compute_instance_iam_member":                      {compute.ResourceConverterComputeInstanceIamMember()},
		"google_compute_region_backend_service_iam_policy":        {compute.ResourceConverterComputeRegionBackendServiceIamPolicy()},
		"google_compute_region_backend_service_iam_binding":       {compute.ResourceConverterComputeRegionBackendServiceIamBinding()},
		"google_compute_region_backend_service_iam_member":        {compute.ResourceConverterComputeRegionBackendServiceIamMember()},
		"google_compute_region_disk_iam_policy":                   {compute.ResourceConverterComputeRegionDiskIamPolicy()},
		"google_compute_region_disk_iam_binding":                  {compute.ResourceConverterComputeRegionDiskIamBinding()},
		"google_compute_region_disk_iam_member":                   {compute.ResourceConverterComputeRegionDiskIamMember()},
		"google_compute_snapshot_iam_policy":                      {compute.ResourceConverterComputeSnapshotIamPolicy()},
		"google_compute_snapshot_iam_binding":                     {compute.ResourceConverterComputeSnapshotIamBinding()},
		"google_compute_snapshot_iam_member":                      {compute.ResourceConverterComputeSnapshotIamMember()},
		"google_compute_subnetwork_iam_policy":                    {compute.ResourceConverterComputeSubnetworkIamPolicy()},
		"google_compute_subnetwork_iam_binding":                   {compute.ResourceConverterComputeSubnetworkIamBinding()},
		"google_compute_subnetwork_iam_member":                    {compute.ResourceConverterComputeSubnetworkIamMember()},
		"google_container_analysis_note_iam_policy":               {containeranalysis.ResourceConverterContainerAnalysisNoteIamPolicy()},
		"google_container_analysis_note_iam_binding":              {containeranalysis.ResourceConverterContainerAnalysisNoteIamBinding()},
		"google_container_analysis_note_iam_member":               {containeranalysis.ResourceConverterContainerAnalysisNoteIamMember()},
		"google_data_catalog_entry_group_iam_policy":              {datacatalog.ResourceConverterDataCatalogEntryGroupIamPolicy()},
		"google_data_catalog_entry_group_iam_binding":             {datacatalog.ResourceConverterDataCatalogEntryGroupIamBinding()},
		"google_data_catalog_entry_group_iam_member":              {datacatalog.ResourceConverterDataCatalogEntryGroupIamMember()},
		"google_data_catalog_policy_tag_iam_policy":               {datacatalog.ResourceConverterDataCatalogPolicyTagIamPolicy()},
		"google_data_catalog_policy_tag_iam_binding":              {datacatalog.ResourceConverterDataCatalogPolicyTagIamBinding()},
		"google_data_catalog_policy_tag_iam_member":               {datacatalog.ResourceConverterDataCatalogPolicyTagIamMember()},
		"google_data_catalog_tag_template_iam_policy":             {datacatalog.ResourceConverterDataCatalogTagTemplateIamPolicy()},
		"google_data_catalog_tag_template_iam_binding":            {datacatalog.ResourceConverterDataCatalogTagTemplateIamBinding()},
		"google_data_catalog_tag_template_iam_member":             {datacatalog.ResourceConverterDataCatalogTagTemplateIamMember()},
		"google_data_catalog_taxonomy_iam_policy":                 {datacatalog.ResourceConverterDataCatalogTaxonomyIamPolicy()},
		"google_data_catalog_taxonomy_iam_binding":                {datacatalog.ResourceConverterDataCatalogTaxonomyIamBinding()},
		"google_data_catalog_taxonomy_iam_member":                 {datacatalog.ResourceConverterDataCatalogTaxonomyIamMember()},
		"google_data_fusion_instance_iam_policy":                  {datafusion.ResourceConverterDataFusionInstanceIamPolicy()},
		"google_data_fusion_instance_iam_binding":                 {datafusion.ResourceConverterDataFusionInstanceIamBinding()},
		"google_data_fusion_instance_iam_member":                  {datafusion.ResourceConverterDataFusionInstanceIamMember()},
		"google_dataplex_asset_iam_policy":                        {dataplex.ResourceConverterDataplexAssetIamPolicy()},
		"google_dataplex_asset_iam_binding":                       {dataplex.ResourceConverterDataplexAssetIamBinding()},
		"google_dataplex_asset_iam_member":                        {dataplex.ResourceConverterDataplexAssetIamMember()},
		"google_dataplex_datascan_iam_policy":                     {dataplex.ResourceConverterDataplexDatascanIamPolicy()},
		"google_dataplex_datascan_iam_binding":                    {dataplex.ResourceConverterDataplexDatascanIamBinding()},
		"google_dataplex_datascan_iam_member":                     {dataplex.ResourceConverterDataplexDatascanIamMember()},
		"google_dataplex_lake_iam_policy":                         {dataplex.ResourceConverterDataplexLakeIamPolicy()},
		"google_dataplex_lake_iam_binding":                        {dataplex.ResourceConverterDataplexLakeIamBinding()},
		"google_dataplex_lake_iam_member":                         {dataplex.ResourceConverterDataplexLakeIamMember()},
		"google_dataplex_task_iam_policy":                         {dataplex.ResourceConverterDataplexTaskIamPolicy()},
		"google_dataplex_task_iam_binding":                        {dataplex.ResourceConverterDataplexTaskIamBinding()},
		"google_dataplex_task_iam_member":                         {dataplex.ResourceConverterDataplexTaskIamMember()},
		"google_dataplex_zone_iam_policy":                         {dataplex.ResourceConverterDataplexZoneIamPolicy()},
		"google_dataplex_zone_iam_binding":                        {dataplex.ResourceConverterDataplexZoneIamBinding()},
		"google_dataplex_zone_iam_member":                         {dataplex.ResourceConverterDataplexZoneIamMember()},
		"google_dataproc_autoscaling_policy_iam_policy":           {dataproc.ResourceConverterDataprocAutoscalingPolicyIamPolicy()},
		"google_dataproc_autoscaling_policy_iam_binding":          {dataproc.ResourceConverterDataprocAutoscalingPolicyIamBinding()},
		"google_dataproc_autoscaling_policy_iam_member":           {dataproc.ResourceConverterDataprocAutoscalingPolicyIamMember()},
		"google_dataproc_metastore_service_iam_policy":            {dataprocmetastore.ResourceConverterDataprocMetastoreServiceIamPolicy()},
		"google_dataproc_metastore_service_iam_binding":           {dataprocmetastore.ResourceConverterDataprocMetastoreServiceIamBinding()},
		"google_dataproc_metastore_service_iam_member":            {dataprocmetastore.ResourceConverterDataprocMetastoreServiceIamMember()},
		"google_dns_managed_zone_iam_policy":                      {dns.ResourceConverterDNSManagedZoneIamPolicy()},
		"google_dns_managed_zone_iam_binding":                     {dns.ResourceConverterDNSManagedZoneIamBinding()},
		"google_dns_managed_zone_iam_member":                      {dns.ResourceConverterDNSManagedZoneIamMember()},
		"google_gke_backup_backup_plan_iam_policy":                {gkebackup.ResourceConverterGKEBackupBackupPlanIamPolicy()},
		"google_gke_backup_backup_plan_iam_binding":               {gkebackup.ResourceConverterGKEBackupBackupPlanIamBinding()},
		"google_gke_backup_backup_plan_iam_member":                {gkebackup.ResourceConverterGKEBackupBackupPlanIamMember()},
		"google_gke_backup_restore_plan_iam_policy":               {gkebackup.ResourceConverterGKEBackupRestorePlanIamPolicy()},
		"google_gke_backup_restore_plan_iam_binding":              {gkebackup.ResourceConverterGKEBackupRestorePlanIamBinding()},
		"google_gke_backup_restore_plan_iam_member":               {gkebackup.ResourceConverterGKEBackupRestorePlanIamMember()},
		"google_gke_hub_membership_iam_policy":                    {gkehub.ResourceConverterGKEHubMembershipIamPolicy()},
		"google_gke_hub_membership_iam_binding":                   {gkehub.ResourceConverterGKEHubMembershipIamBinding()},
		"google_gke_hub_membership_iam_member":                    {gkehub.ResourceConverterGKEHubMembershipIamMember()},
		"google_gke_hub_feature_iam_policy":                       {gkehub2.ResourceConverterGKEHub2FeatureIamPolicy()},
		"google_gke_hub_feature_iam_binding":                      {gkehub2.ResourceConverterGKEHub2FeatureIamBinding()},
		"google_gke_hub_feature_iam_member":                       {gkehub2.ResourceConverterGKEHub2FeatureIamMember()},
		"google_gke_hub_scope_iam_policy":                         {gkehub2.ResourceConverterGKEHub2ScopeIamPolicy()},
		"google_gke_hub_scope_iam_binding":                        {gkehub2.ResourceConverterGKEHub2ScopeIamBinding()},
		"google_gke_hub_scope_iam_member":                         {gkehub2.ResourceConverterGKEHub2ScopeIamMember()},
		"google_healthcare_consent_store_iam_policy":              {healthcare.ResourceConverterHealthcareConsentStoreIamPolicy()},
		"google_healthcare_consent_store_iam_binding":             {healthcare.ResourceConverterHealthcareConsentStoreIamBinding()},
		"google_healthcare_consent_store_iam_member":              {healthcare.ResourceConverterHealthcareConsentStoreIamMember()},
		"google_iap_tunnel_iam_policy":                            {iap.ResourceConverterIapTunnelIamPolicy()},
		"google_iap_tunnel_iam_binding":                           {iap.ResourceConverterIapTunnelIamBinding()},
		"google_iap_tunnel_iam_member":                            {iap.ResourceConverterIapTunnelIamMember()},
		"google_iap_tunnel_instance_iam_policy":                   {iap.ResourceConverterIapTunnelInstanceIamPolicy()},
		"google_iap_tunnel_instance_iam_binding":                  {iap.ResourceConverterIapTunnelInstanceIamBinding()},
		"google_iap_tunnel_instance_iam_member":                   {iap.ResourceConverterIapTunnelInstanceIamMember()},
		"google_iap_web_iam_policy":                               {iap.ResourceConverterIapWebIamPolicy()},
		"google_iap_web_iam_binding":                              {iap.ResourceConverterIapWebIamBinding()},
		"google_iap_web_iam_member":                               {iap.ResourceConverterIapWebIamMember()},
		"google_notebooks_instance_iam_policy":                    {notebooks.ResourceConverterNotebooksInstanceIamPolicy()},
		"google_notebooks_instance_iam_binding":                   {notebooks.ResourceConverterNotebooksInstanceIamBinding()},
		"google_notebooks_instance_iam_member":                    {notebooks.ResourceConverterNotebooksInstanceIamMember()},
		"google_notebooks_runtime_iam_policy":                     {notebooks.ResourceConverterNotebooksRuntimeIamPolicy()},
		"google_notebooks_runtime_iam_binding":                    {notebooks.ResourceConverterNotebooksRuntimeIamBinding()},
		"google_notebooks_runtime_iam_member":                     {notebooks.ResourceConverterNotebooksRuntimeIamMember()},
		"google_privateca_ca_pool_iam_policy":                     {privateca.ResourceConverterPrivatecaCaPoolIamPolicy()},
		"google_privateca_ca_pool_iam_binding":                    {privateca.ResourceConverterPrivatecaCaPoolIamBinding()},
		"google_privateca_ca_pool_iam_member":                     {privateca.ResourceConverterPrivatecaCaPoolIamMember()},
		"google_privateca_certificate_template_iam_policy":        {privateca.ResourceConverterPrivatecaCertificateTemplateIamPolicy()},
		"google_privateca_certificate_template_iam_binding":       {privateca.ResourceConverterPrivatecaCertificateTemplateIamBinding()},
		"google_privateca_certificate_template_iam_member":        {privateca.ResourceConverterPrivatecaCertificateTemplateIamMember()},
		"google_pubsub_topic_iam_policy":                          {pubsub.ResourceConverterPubsubTopicIamPolicy()},
		"google_pubsub_topic_iam_binding":                         {pubsub.ResourceConverterPubsubTopicIamBinding()},
		"google_pubsub_topic_iam_member":                          {pubsub.ResourceConverterPubsubTopicIamMember()},
		"google_secret_manager_secret_iam_policy":                 {secretmanager.ResourceConverterSecretManagerSecretIamPolicy()},
		"google_secret_manager_secret_iam_binding":                {secretmanager.ResourceConverterSecretManagerSecretIamBinding()},
		"google_secret_manager_secret_iam_member":                 {secretmanager.ResourceConverterSecretManagerSecretIamMember()},
		"google_scc_source_iam_policy":                            {securitycenter.ResourceConverterSecurityCenterSourceIamPolicy()},
		"google_scc_source_iam_binding":                           {securitycenter.ResourceConverterSecurityCenterSourceIamBinding()},
		"google_scc_source_iam_member":                            {securitycenter.ResourceConverterSecurityCenterSourceIamMember()},
		"google_endpoints_service_iam_policy":                     {servicemanagement.ResourceConverterServiceManagementServiceIamPolicy()},
		"google_endpoints_service_iam_binding":                    {servicemanagement.ResourceConverterServiceManagementServiceIamBinding()},
		"google_endpoints_service_iam_member":                     {servicemanagement.ResourceConverterServiceManagementServiceIamMember()},
		"google_endpoints_service_consumers_iam_policy":           {servicemanagement.ResourceConverterServiceManagementServiceConsumersIamPolicy()},
		"google_endpoints_service_consumers_iam_binding":          {servicemanagement.ResourceConverterServiceManagementServiceConsumersIamBinding()},
		"google_endpoints_service_consumers_iam_member":           {servicemanagement.ResourceConverterServiceManagementServiceConsumersIamMember()},
		"google_vertex_ai_featurestore_iam_policy":                {vertexai.ResourceConverterVertexAIFeaturestoreIamPolicy()},
		"google_vertex_ai_featurestore_iam_binding":               {vertexai.ResourceConverterVertexAIFeaturestoreIamBinding()},
		"google_vertex_ai_featurestore_iam_member":                {vertexai.ResourceConverterVertexAIFeaturestoreIamMember()},
		"google_vertex_ai_featurestore_entitytype_iam_policy":     {vertexai.ResourceConverterVertexAIFeaturestoreEntitytypeIamPolicy()},
		"google_vertex_ai_featurestore_entitytype_iam_binding":    {vertexai.ResourceConverterVertexAIFeaturestoreEntitytypeIamBinding()},
		"google_vertex_ai_featurestore_entitytype_iam_member":     {vertexai.ResourceConverterVertexAIFeaturestoreEntitytypeIamMember()},
		"google_project": {
			resourceConverterProject(),
			resourceConverterProjectBillingInfo(),
		},
		"google_bigtable_instance": {
			resourceConverterBigtableInstance(),
			resourceConverterBigtableCluster(),
		},
		"google_organization_iam_policy":      {resourcemanager.ResourceConverterOrganizationIamPolicy()},
		"google_organization_iam_binding":     {resourcemanager.ResourceConverterOrganizationIamBinding()},
		"google_organization_iam_member":      {resourcemanager.ResourceConverterOrganizationIamMember()},
		"google_organization_policy":          {resourceConverterOrganizationPolicy()},
		"google_project_organization_policy":  {resourceConverterProjectOrgPolicy()},
		"google_folder":                       {resourceConverterFolder()},
		"google_folder_iam_policy":            {resourcemanager.ResourceConverterFolderIamPolicy()},
		"google_folder_iam_binding":           {resourcemanager.ResourceConverterFolderIamBinding()},
		"google_folder_iam_member":            {resourcemanager.ResourceConverterFolderIamMember()},
		"google_folder_organization_policy":   {resourceConverterFolderOrgPolicy()},
		"google_kms_crypto_key_iam_policy":    {resourceConverterKmsCryptoKeyIamPolicy()},
		"google_kms_crypto_key_iam_binding":   {resourceConverterKmsCryptoKeyIamBinding()},
		"google_kms_crypto_key_iam_member":    {resourceConverterKmsCryptoKeyIamMember()},
		"google_kms_key_ring_iam_policy":      {resourceConverterKmsKeyRingIamPolicy()},
		"google_kms_key_ring_iam_binding":     {resourceConverterKmsKeyRingIamBinding()},
		"google_kms_key_ring_iam_member":      {resourceConverterKmsKeyRingIamMember()},
		"google_project_iam_policy":           {resourcemanager.ResourceConverterProjectIamPolicy()},
		"google_project_iam_binding":          {resourcemanager.ResourceConverterProjectIamBinding()},
		"google_project_iam_member":           {resourcemanager.ResourceConverterProjectIamMember()},
		"google_project_iam_custom_role":      {resourceConverterProjectIAMCustomRole()},
		"google_organization_iam_custom_role": {resourceConverterOrganizationIAMCustomRole()},
		"google_vpc_access_connector":         {vpcaccess.ResourceConverterVPCAccessConnector()},
		"google_logging_metric":               {logging.ResourceConverterLoggingMetric()},
		"google_service_account":              {resourceConverterServiceAccount()},
	}
}
