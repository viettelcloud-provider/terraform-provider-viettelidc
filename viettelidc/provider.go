package viettelidc

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/utils/terraform/auth"
	"github.com/gophercloud/utils/terraform/mutexkv"
)

// Use openstackbase.Config as the base/foundation of this provider's
// Config struct.
type Config struct {
	auth.Config
}

// Provider returns a schema.Provider for OpenStack.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", ""),
				Description: descriptions["auth_url"],
			},

			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["region"],
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},

			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_id"],
			},

			"application_credential_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_ID", ""),
				Description: descriptions["application_credential_id"],
			},

			"application_credential_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_NAME", ""),
				Description: descriptions["application_credential_name"],
			},

			"application_credential_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_APPLICATION_CREDENTIAL_SECRET", ""),
				Description: descriptions["application_credential_secret"],
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TOKEN",
					"OS_AUTH_TOKEN",
				}, ""),
				Description: descriptions["token"],
			},

			"user_domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_NAME", ""),
				Description: descriptions["user_domain_name"],
			},

			"user_domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_DOMAIN_ID", ""),
				Description: descriptions["user_domain_id"],
			},

			"project_domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_NAME", ""),
				Description: descriptions["project_domain_name"],
			},

			"project_domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_ID", ""),
				Description: descriptions["project_domain_id"],
			},

			"domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_ID", ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_NAME", ""),
				Description: descriptions["domain_name"],
			},

			"default_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DEFAULT_DOMAIN", "default"),
				Description: descriptions["default_domain"],
			},

			"system_scope": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SYSTEM_SCOPE", false),
				Description: descriptions["system_scope"],
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", nil),
				Description: descriptions["insecure"],
			},

			"endpoint_type": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"swauth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SWAUTH", false),
				Description: descriptions["swauth"],
			},

			"use_octavia": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USE_OCTAVIA", true),
				Description: descriptions["use_octavia"],
			},

			"delayed_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DELAYED_AUTH", true),
				Description: descriptions["delayed_auth"],
			},

			"allow_reauth": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ALLOW_REAUTH", true),
				Description: descriptions["allow_reauth"],
			},

			"cloud": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CLOUD", ""),
				Description: descriptions["cloud"],
			},

			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: descriptions["max_retries"],
			},

			"endpoint_overrides": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions["endpoint_overrides"],
			},

			"disable_no_cache_header": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["disable_no_cache_header"],
			},

			"enable_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["enable_logging"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"viettelidc_blockstorage_availability_zones_v3":       dataSourceBlockStorageAvailabilityZonesV3(),
			"viettelidc_blockstorage_snapshot_v2":                 dataSourceBlockStorageSnapshotV2(),
			"viettelidc_blockstorage_snapshot_v3":                 dataSourceBlockStorageSnapshotV3(),
			"viettelidc_blockstorage_volume_v2":                   dataSourceBlockStorageVolumeV2(),
			"viettelidc_blockstorage_volume_v3":                   dataSourceBlockStorageVolumeV3(),
			"viettelidc_blockstorage_quotaset_v3":                 dataSourceBlockStorageQuotasetV3(),
			"viettelidc_compute_aggregate_v2":                     dataSourceComputeAggregateV2(),
			"viettelidc_compute_availability_zones_v2":            dataSourceComputeAvailabilityZonesV2(),
			"viettelidc_compute_instance_v2":                      dataSourceComputeInstanceV2(),
			"viettelidc_compute_flavor_v2":                        dataSourceComputeFlavorV2(),
			"viettelidc_compute_hypervisor_v2":                    dataSourceComputeHypervisorV2(),
			"viettelidc_compute_keypair_v2":                       dataSourceComputeKeypairV2(),
			"viettelidc_compute_quotaset_v2":                      dataSourceComputeQuotasetV2(),
			"viettelidc_compute_limits_v2":                        dataSourceComputeLimitsV2(),
			"viettelidc_containerinfra_nodegroup_v1":              dataSourceContainerInfraNodeGroupV1(),
			"viettelidc_containerinfra_clustertemplate_v1":        dataSourceContainerInfraClusterTemplateV1(),
			"viettelidc_containerinfra_cluster_v1":                dataSourceContainerInfraCluster(),
			"viettelidc_dns_zone_v2":                              dataSourceDNSZoneV2(),
			"viettelidc_fw_group_v2":                              dataSourceFWGroupV2(),
			"viettelidc_fw_policy_v1":                             dataSourceFWPolicyV1(),
			"viettelidc_fw_policy_v2":                             dataSourceFWPolicyV2(),
			"viettelidc_fw_rule_v2":                               dataSourceFWRuleV2(),
			"viettelidc_identity_role_v3":                         dataSourceIdentityRoleV3(),
			"viettelidc_identity_project_v3":                      dataSourceIdentityProjectV3(),
			"viettelidc_identity_user_v3":                         dataSourceIdentityUserV3(),
			"viettelidc_identity_auth_scope_v3":                   dataSourceIdentityAuthScopeV3(),
			"viettelidc_identity_endpoint_v3":                     dataSourceIdentityEndpointV3(),
			"viettelidc_identity_service_v3":                      dataSourceIdentityServiceV3(),
			"viettelidc_identity_group_v3":                        dataSourceIdentityGroupV3(),
			"viettelidc_images_image_v2":                          dataSourceImagesImageV2(),
			"viettelidc_images_image_ids_v2":                      dataSourceImagesImageIDsV2(),
			"viettelidc_networking_addressscope_v2":               dataSourceNetworkingAddressScopeV2(),
			"viettelidc_networking_network_v2":                    dataSourceNetworkingNetworkV2(),
			"viettelidc_networking_qos_bandwidth_limit_rule_v2":   dataSourceNetworkingQoSBandwidthLimitRuleV2(),
			"viettelidc_networking_qos_dscp_marking_rule_v2":      dataSourceNetworkingQoSDSCPMarkingRuleV2(),
			"viettelidc_networking_qos_minimum_bandwidth_rule_v2": dataSourceNetworkingQoSMinimumBandwidthRuleV2(),
			"viettelidc_networking_qos_policy_v2":                 dataSourceNetworkingQoSPolicyV2(),
			"viettelidc_networking_quota_v2":                      dataSourceNetworkingQuotaV2(),
			"viettelidc_networking_subnet_v2":                     dataSourceNetworkingSubnetV2(),
			"viettelidc_networking_subnet_ids_v2":                 dataSourceNetworkingSubnetIDsV2(),
			"viettelidc_networking_secgroup_v2":                   dataSourceNetworkingSecGroupV2(),
			"viettelidc_networking_subnetpool_v2":                 dataSourceNetworkingSubnetPoolV2(),
			"viettelidc_networking_floatingip_v2":                 dataSourceNetworkingFloatingIPV2(),
			"viettelidc_networking_router_v2":                     dataSourceNetworkingRouterV2(),
			"viettelidc_networking_port_v2":                       dataSourceNetworkingPortV2(),
			"viettelidc_networking_port_ids_v2":                   dataSourceNetworkingPortIDsV2(),
			"viettelidc_networking_trunk_v2":                      dataSourceNetworkingTrunkV2(),
			"viettelidc_sharedfilesystem_availability_zones_v2":   dataSourceSharedFilesystemAvailabilityZonesV2(),
			"viettelidc_sharedfilesystem_sharenetwork_v2":         dataSourceSharedFilesystemShareNetworkV2(),
			"viettelidc_sharedfilesystem_share_v2":                dataSourceSharedFilesystemShareV2(),
			"viettelidc_sharedfilesystem_snapshot_v2":             dataSourceSharedFilesystemSnapshotV2(),
			"viettelidc_keymanager_secret_v1":                     dataSourceKeyManagerSecretV1(),
			"viettelidc_keymanager_container_v1":                  dataSourceKeyManagerContainerV1(),
			"viettelidc_kubernetes_v1":                            dataSourceKubernetesV1(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"viettelidc_blockstorage_qos_association_v3":          resourceBlockStorageQosAssociationV3(),
			"viettelidc_blockstorage_qos_v3":                      resourceBlockStorageQosV3(),
			"viettelidc_blockstorage_quotaset_v2":                 resourceBlockStorageQuotasetV2(),
			"viettelidc_blockstorage_quotaset_v3":                 resourceBlockStorageQuotasetV3(),
			"viettelidc_blockstorage_volume_v1":                   resourceBlockStorageVolumeV1(),
			"viettelidc_blockstorage_volume_v2":                   resourceBlockStorageVolumeV2(),
			"viettelidc_blockstorage_volume_v3":                   resourceBlockStorageVolumeV3(),
			"viettelidc_blockstorage_volume_attach_v2":            resourceBlockStorageVolumeAttachV2(),
			"viettelidc_blockstorage_volume_attach_v3":            resourceBlockStorageVolumeAttachV3(),
			"viettelidc_blockstorage_volume_type_access_v3":       resourceBlockstorageVolumeTypeAccessV3(),
			"viettelidc_blockstorage_volume_type_v3":              resourceBlockStorageVolumeTypeV3(),
			"viettelidc_compute_aggregate_v2":                     resourceComputeAggregateV2(),
			"viettelidc_compute_flavor_v2":                        resourceComputeFlavorV2(),
			"viettelidc_compute_flavor_access_v2":                 resourceComputeFlavorAccessV2(),
			"viettelidc_compute_instance_v2":                      resourceComputeInstanceV2(),
			"viettelidc_compute_interface_attach_v2":              resourceComputeInterfaceAttachV2(),
			"viettelidc_compute_keypair_v2":                       resourceComputeKeypairV2(),
			"viettelidc_compute_secgroup_v2":                      resourceComputeSecGroupV2(),
			"viettelidc_compute_servergroup_v2":                   resourceComputeServerGroupV2(),
			"viettelidc_compute_quotaset_v2":                      resourceComputeQuotasetV2(),
			"viettelidc_compute_floatingip_v2":                    resourceComputeFloatingIPV2(),
			"viettelidc_compute_floatingip_associate_v2":          resourceComputeFloatingIPAssociateV2(),
			"viettelidc_compute_volume_attach_v2":                 resourceComputeVolumeAttachV2(),
			"viettelidc_containerinfra_nodegroup_v1":              resourceContainerInfraNodeGroupV1(),
			"viettelidc_containerinfra_clustertemplate_v1":        resourceContainerInfraClusterTemplateV1(),
			"viettelidc_containerinfra_cluster_v1":                resourceContainerInfraClusterV1(),
			"viettelidc_db_instance_v1":                           resourceDatabaseInstanceV1(),
			"viettelidc_db_user_v1":                               resourceDatabaseUserV1(),
			"viettelidc_db_configuration_v1":                      resourceDatabaseConfigurationV1(),
			"viettelidc_db_database_v1":                           resourceDatabaseDatabaseV1(),
			"viettelidc_dns_recordset_v2":                         resourceDNSRecordSetV2(),
			"viettelidc_dns_zone_v2":                              resourceDNSZoneV2(),
			"viettelidc_dns_transfer_request_v2":                  resourceDNSTransferRequestV2(),
			"viettelidc_dns_transfer_accept_v2":                   resourceDNSTransferAcceptV2(),
			"viettelidc_fw_firewall_v1":                           resourceFWFirewallV1(),
			"viettelidc_fw_group_v2":                              resourceFWGroupV2(),
			"viettelidc_fw_policy_v1":                             resourceFWPolicyV1(),
			"viettelidc_fw_policy_v2":                             resourceFWPolicyV2(),
			"viettelidc_fw_rule_v1":                               resourceFWRuleV1(),
			"viettelidc_fw_rule_v2":                               resourceFWRuleV2(),
			"viettelidc_identity_endpoint_v3":                     resourceIdentityEndpointV3(),
			"viettelidc_identity_project_v3":                      resourceIdentityProjectV3(),
			"viettelidc_identity_role_v3":                         resourceIdentityRoleV3(),
			"viettelidc_identity_role_assignment_v3":              resourceIdentityRoleAssignmentV3(),
			"viettelidc_identity_inherit_role_assignment_v3":      resourceIdentityInheritRoleAssignmentV3(),
			"viettelidc_identity_service_v3":                      resourceIdentityServiceV3(),
			"viettelidc_identity_user_v3":                         resourceIdentityUserV3(),
			"viettelidc_identity_user_membership_v3":              resourceIdentityUserMembershipV3(),
			"viettelidc_identity_group_v3":                        resourceIdentityGroupV3(),
			"viettelidc_identity_application_credential_v3":       resourceIdentityApplicationCredentialV3(),
			"viettelidc_identity_ec2_credential_v3":               resourceIdentityEc2CredentialV3(),
			"viettelidc_images_image_v2":                          resourceImagesImageV2(),
			"viettelidc_images_image_access_v2":                   resourceImagesImageAccessV2(),
			"viettelidc_images_image_access_accept_v2":            resourceImagesImageAccessAcceptV2(),
			"viettelidc_lb_member_v1":                             resourceLBMemberV1(),
			"viettelidc_lb_monitor_v1":                            resourceLBMonitorV1(),
			"viettelidc_lb_pool_v1":                               resourceLBPoolV1(),
			"viettelidc_lb_vip_v1":                                resourceLBVipV1(),
			"viettelidc_lb_loadbalancer_v2":                       resourceLoadBalancerV2(),
			"viettelidc_lb_listener_v2":                           resourceListenerV2(),
			"viettelidc_lb_pool_v2":                               resourcePoolV2(),
			"viettelidc_lb_member_v2":                             resourceMemberV2(),
			"viettelidc_lb_members_v2":                            resourceMembersV2(),
			"viettelidc_lb_monitor_v2":                            resourceMonitorV2(),
			"viettelidc_lb_l7policy_v2":                           resourceL7PolicyV2(),
			"viettelidc_lb_l7rule_v2":                             resourceL7RuleV2(),
			"viettelidc_lb_quota_v2":                              resourceLoadBalancerQuotaV2(),
			"viettelidc_networking_floatingip_v2":                 resourceNetworkingFloatingIPV2(),
			"viettelidc_networking_floatingip_associate_v2":       resourceNetworkingFloatingIPAssociateV2(),
			"viettelidc_networking_network_v2":                    resourceNetworkingNetworkV2(),
			"viettelidc_networking_port_v2":                       resourceNetworkingPortV2(),
			"viettelidc_networking_rbac_policy_v2":                resourceNetworkingRBACPolicyV2(),
			"viettelidc_networking_port_secgroup_associate_v2":    resourceNetworkingPortSecGroupAssociateV2(),
			"viettelidc_networking_qos_bandwidth_limit_rule_v2":   resourceNetworkingQoSBandwidthLimitRuleV2(),
			"viettelidc_networking_qos_dscp_marking_rule_v2":      resourceNetworkingQoSDSCPMarkingRuleV2(),
			"viettelidc_networking_qos_minimum_bandwidth_rule_v2": resourceNetworkingQoSMinimumBandwidthRuleV2(),
			"viettelidc_networking_qos_policy_v2":                 resourceNetworkingQoSPolicyV2(),
			"viettelidc_networking_quota_v2":                      resourceNetworkingQuotaV2(),
			"viettelidc_networking_router_v2":                     resourceNetworkingRouterV2(),
			"viettelidc_networking_router_interface_v2":           resourceNetworkingRouterInterfaceV2(),
			"viettelidc_networking_router_route_v2":               resourceNetworkingRouterRouteV2(),
			"viettelidc_networking_secgroup_v2":                   resourceNetworkingSecGroupV2(),
			"viettelidc_networking_secgroup_rule_v2":              resourceNetworkingSecGroupRuleV2(),
			"viettelidc_networking_subnet_v2":                     resourceNetworkingSubnetV2(),
			"viettelidc_networking_subnet_route_v2":               resourceNetworkingSubnetRouteV2(),
			"viettelidc_networking_subnetpool_v2":                 resourceNetworkingSubnetPoolV2(),
			"viettelidc_networking_addressscope_v2":               resourceNetworkingAddressScopeV2(),
			"viettelidc_networking_trunk_v2":                      resourceNetworkingTrunkV2(),
			"viettelidc_networking_portforwarding_v2":             resourceNetworkingPortForwardingV2(),
			"viettelidc_objectstorage_container_v1":               resourceObjectStorageContainerV1(),
			"viettelidc_objectstorage_object_v1":                  resourceObjectStorageObjectV1(),
			"viettelidc_objectstorage_tempurl_v1":                 resourceObjectstorageTempurlV1(),
			"viettelidc_orchestration_stack_v1":                   resourceOrchestrationStackV1(),
			"viettelidc_vpnaas_ipsec_policy_v2":                   resourceIPSecPolicyV2(),
			"viettelidc_vpnaas_service_v2":                        resourceServiceV2(),
			"viettelidc_vpnaas_ike_policy_v2":                     resourceIKEPolicyV2(),
			"viettelidc_vpnaas_endpoint_group_v2":                 resourceEndpointGroupV2(),
			"viettelidc_vpnaas_site_connection_v2":                resourceSiteConnectionV2(),
			"viettelidc_sharedfilesystem_securityservice_v2":      resourceSharedFilesystemSecurityServiceV2(),
			"viettelidc_sharedfilesystem_sharenetwork_v2":         resourceSharedFilesystemShareNetworkV2(),
			"viettelidc_sharedfilesystem_share_v2":                resourceSharedFilesystemShareV2(),
			"viettelidc_sharedfilesystem_share_access_v2":         resourceSharedFilesystemShareAccessV2(),
			"viettelidc_keymanager_secret_v1":                     resourceKeyManagerSecretV1(),
			"viettelidc_keymanager_container_v1":                  resourceKeyManagerContainerV1(),
			"viettelidc_keymanager_order_v1":                      resourceKeyManagerOrderV1(),
			"viettelidc_kubernetes_v1":                            resourceKubernetesV1(),
		},
	}

	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return configureProvider(d, terraformVersion)
	}

	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"cloud": "An entry in a `clouds.yaml` file to use.",

		"region": "The OpenStack region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"application_credential_id": "Application Credential ID to login with.",

		"application_credential_name": "Application Credential name to login with.",

		"application_credential_secret": "Application Credential secret to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"user_domain_name": "The name of the domain where the user resides (Identity v3).",

		"user_domain_id": "The ID of the domain where the user resides (Identity v3).",

		"project_domain_name": "The name of the domain where the project resides (Identity v3).",

		"project_domain_id": "The ID of the domain where the proejct resides (Identity v3).",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"default_domain": "The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).",

		"system_scope": "If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"endpoint_type": "The catalog endpoint type to use.",

		"endpoint_overrides": "A map of services with an endpoint to override what was\n" +
			"from the Keystone catalog",

		"swauth": "Use Swift's authentication system instead of Keystone. Only used for\n" +
			"interaction with Swift.",

		"use_octavia": "If set to `true`, API requests will go the Load Balancer\n" +
			"service (Octavia) instead of the Networking service (Neutron).",

		"disable_no_cache_header": "If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.",

		"delayed_auth": "If set to `false`, OpenStack authorization will be perfomed,\n" +
			"every time the service provider client is called. Defaults to `true`.",

		"allow_reauth": "If set to `false`, OpenStack authorization won't be perfomed\n" +
			"automatically, if the initial auth token get expired. Defaults to `true`",

		"max_retries": "How many times HTTP connection should be retried until giving up.",

		"enable_logging": "Outputs very verbose logs with all calls made to and responses from OpenStack",
	}
}

func configureProvider(d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	enableLogging := d.Get("enable_logging").(bool)
	if !enableLogging {
		// enforce logging (similar to OS_DEBUG) when TF_LOG is 'DEBUG' or 'TRACE'
		if logLevel := logging.LogLevel(); logLevel != "" && os.Getenv("OS_DEBUG") == "" {
			if logLevel == "DEBUG" || logLevel == "TRACE" {
				enableLogging = true
			}
		}
	}

	authOpts := &gophercloud.AuthOptions{
		Scope: &gophercloud.AuthScope{System: d.Get("system_scope").(bool)},
	}

	config := Config{
		auth.Config{
			CACertFile:                  d.Get("cacert_file").(string),
			ClientCertFile:              d.Get("cert").(string),
			ClientKeyFile:               d.Get("key").(string),
			Cloud:                       d.Get("cloud").(string),
			DefaultDomain:               d.Get("default_domain").(string),
			DomainID:                    d.Get("domain_id").(string),
			DomainName:                  d.Get("domain_name").(string),
			EndpointOverrides:           d.Get("endpoint_overrides").(map[string]interface{}),
			EndpointType:                d.Get("endpoint_type").(string),
			IdentityEndpoint:            d.Get("auth_url").(string),
			Password:                    d.Get("password").(string),
			ProjectDomainID:             d.Get("project_domain_id").(string),
			ProjectDomainName:           d.Get("project_domain_name").(string),
			Region:                      d.Get("region").(string),
			Swauth:                      d.Get("swauth").(bool),
			Token:                       d.Get("token").(string),
			TenantID:                    d.Get("tenant_id").(string),
			TenantName:                  d.Get("tenant_name").(string),
			UserDomainID:                d.Get("user_domain_id").(string),
			UserDomainName:              d.Get("user_domain_name").(string),
			Username:                    d.Get("user_name").(string),
			UserID:                      d.Get("user_id").(string),
			ApplicationCredentialID:     d.Get("application_credential_id").(string),
			ApplicationCredentialName:   d.Get("application_credential_name").(string),
			ApplicationCredentialSecret: d.Get("application_credential_secret").(string),
			UseOctavia:                  d.Get("use_octavia").(bool),
			DelayedAuth:                 d.Get("delayed_auth").(bool),
			AllowReauth:                 d.Get("allow_reauth").(bool),
			AuthOpts:                    authOpts,
			MaxRetries:                  d.Get("max_retries").(int),
			DisableNoCacheHeader:        d.Get("disable_no_cache_header").(bool),
			TerraformVersion:            terraformVersion,
			SDKVersion:                  meta.SDKVersionString(),
			MutexKV:                     mutexkv.NewMutexKV(),
			EnableLogger:                enableLogging,
		},
	}

	v, ok := d.GetOkExists("insecure")
	if ok {
		insecure := v.(bool)
		config.Insecure = &insecure
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, diag.FromErr(err)
	}

	return &config, nil
}
