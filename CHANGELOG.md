# Terraform Provider MSO - Changelog

All notable changes to this project will be documented in this file.

## 1.0.0 (April 09, 2026)

BREAKING CHANGES:
- Removed site template contract service graph logic from mso_schema_template_contract_service_graph resource
- Add warning for upgrading ndo to version 4.2.2 which breaks terraform state (dcne-731) (#469)

* [major_change] add warning for upgrading ndo to version 4.2.2 which breaks terraform state

* update index.html.markdown

---------

co-authored-by: matt tarkington <mtarking@cisco.com>

DEPRECATIONS:
- Ability to add microsoft and redhat domains in mso_schema_site_anp_epg_domain (deprecate dn attribute) (#130)
- Fix mso_schema resource example template reference and remove deprecated usage of mso_schema resource (#202)
- Made deprecated argument available.
- Deprecated mso_schema_site_service_graph_node resource and added new resource and datasource for mso_schema_site_service_graph which is supported in ndo version 4.1.1i and above.

IMPROVEMENTS:
- Set subnet ip as forcenew
- Resource eepg subnet ip forcenew
- Fix path of terraform import for schema_template_external_epg_contract
- Ability to add multiple templates using the resource mso_schema (#135)
- Readme fix to point to correct tf mso repo
- Import fix with correct example in docs
- Go version change for m1 support
- Add name attribute for the subnet in resource_mso_schema_site_vrf_region (#140)
- Description and tcp_session_rules fix
- Add support for multiple dhcp label policies in mso_schema_template_bd (#161)
- Add mso_schema_template_deploy_ndo resource to support ndo4.1+ deploy api (#165)
- Add mso_schema_site_anp_epg_bulk_staticport resource and datasource (#173)
- Add ip_data_plane_learning and preferred_group arguments to mso_schema_template_vrf resource (#177)
- Add mso_rest data source (#184)
- Ndo4.0+ version conditional to execute replace instead of add patch operation during create in mso_schema_site_anp, mso_schema_site_anp_epg, mso_schema_site_bd, mso_schema_site_external_epg and mso_schema_site_vrf (#188)
- Added gcp specific attributes to mso_tenant resource and datasource
- Added gcp specific attributes to mso_site_vrf_region resource and datasource
- Added gcp specific attributes to mso_site_vrf_region_cidr_subnet resource and datasource
- Added missing primary, description and no_default_gateway attributes to mso_site_anp_epg_subnet and mso_template_anp_epg_subnet resource and data source
- Add svi_mac attribute to  mso_schema_site_bd and make template_name attribute mandatory (#214)
- Added description attribute to schema_template_bd and schema_template_anp_epg resources

* [minor_change] added description attribute to schema_template_bd and schema_template_anp_epg resources

* [ignore] added resourcemsoschematemplateanpepgsetattr to the mso_schema_template_anp_epg resource

* [ignore] fixed document changes
- Add resource and data source mso_schema_site_vrf_route_leak for site local vrf route leaking support
- Add template type attribute to mso_schema and mso_schema_template resources and datasources
- Allow schemas to be created without templates in ndo4.2 releases and above
- Add resource and datasource for mso_remote_location
- Fix for mso_site to detect changes to a site, after manually changing it to unmanaged and fix errorforobjectnotfound function in utils to error out when new "error" payload is found. (#231)
- Add primary and virtual attribute to schema_site_bd_subnet and schema_template_bd_subnet resource and datasource
- Add mso_system_config resource and datasource (#241)
- Add attributes filter_type, directives, action and priority to filter_relationship attribute and add target_dscp and priority attributes to mso_schema_template_contract and add input validation for attributes and align datasource with resource
- Add action and priority attributes to template_contract_filter
- Set id in correct format to state for mso_schema_template_contract and mso_schema_template_contract_filter
- Separating service graph terraform provider into site level and template level.
- [minor changes] removed site association from template service graph resource and datasource. service graph association to template without site properties made possible in v4.1 and above.
- Updated the documentation according to the changes in the resource and datasource of service graph.
- Made changes to site service graph.
- Add support for site epg vmm domain attributes in version 4.2
- Added site_template_contract_service_graph resource and datasource to manage the site template contract service graph
- [minor_changes] adding missing cloud apic site parameters for site_service_graph module.
- Removed firewall_provider_connector_type parameter. added customizediff function to verify the inpurts for provider_connector_type when service node type is 'other' or 'firewall'.
- Add mso_schema_site_contract_service_graph_listener resource and data source to manage azure cnc contract service graph load balancer - listeners and rules (#256)
- Addition of the attribute 'description' to a few applicable resources
- Add support for site_aware_policy_enforcement in mso_schema_template_vrf (#274)
- Added additional check to avoid error generation for checks only required for cloud servicegraphs.
- Add l3out_schema_id, l3out_template and l3out_on_apic attributes to mso_schema_site_external_epg (#291)
- Bump required go version to go v1.21
- Bump required go version to go v1.22
- Add support for vpc connected to fex in resource mso_schema_site_anp_epg_bulk_staticport
- Add resource and datasource for mso_template
- Add support for endpoint move detection mode in schema_template_bd
- Add support for dpc path_type input with fex for resource_mso_schema_site_anp_epg_staticport and resource_mso_schema_site_anp_epg_bulk_staticport
- Addition of new resource and data source for mso_tenant_policies_ipsla_monitoring_policy
- Add resource and data source for multicast route map policy.
- Change rp_ip to rendezvous_point_ip in route_map_entries_multicast. add compatible ndo versions in documentation for resource and data source. add extra test step to check out-of-range order value.
- Add rendezvous_points as a new attribute for mso_schema_template_vrf resource and datasource.
- Added tenant_policies_dhcp_relay_policy resource and introduced uuid to the schema_template_anp_epg and schema_template_external_epg resources
- Add mso_fabric_policies_vlan_pool resource and data source.
- Added mso_fabric_policies_physical_domain resource and datasource.
- Add retries to provider schema to customise to amount of retries for failed api requests
- Addition of logic to display error (or) retry deploy task in resource_ndo_schema_template_deploy
- Fix bd configuration in example of schema_site_anp_bulk_staticport
- Allow update operation for mso_rest without payload, path or method changes by adding a boolean based retrigger mechanism
- Added ability to deploy tenant_policies in mso_schema_template_deploy_ndo
- Addition of resource and datasource for mso_service_device_cluster
- Add resource and data source for synce interface policy.
- Add resource and data source for macsec policy under fabric policy template.
- Addition of resource and data source for service chaining under contract in template
- Addition of resource and data source for bgp peer prefix policy under tenant policies
- Addition of resource and data source for mso_fabric_policies_l3_domain
- Addition of resource and data source for custom qos policy under tenant policies
- Addition of resource and data source for mso_tenant_policies_dhcp_option_policy
- Added ability to un-deploy non-application template types in resource_ndo_schema_template_deploy
- Addition of resource and data source for mso_tenant_policies_mld_snooping_policy
- Add mso_fabric_policies_mcp_global_policy resource and datasource.
- Added mso_tenant_policies_ipsla_track_list resource
- Addition of new resource and data source for mso_tenant_policies_l3out_interface_routing_policy
- Add mso_fabric_policies_interface_setting resource and datasource.
- [minor] upgrade mso-go-client to increase connection pool and enable tls session ticket caching

BUG FIXES:
- Fixes import in mso_schema_template_anp_epg which caused configured values to be null (#129)
- Set subnet name as empty string in read
- Resource eepg subnet set name as empty string in read func
- Update example for schema_template_external_epg_subnet
- Add detail for name attr in schema_template_external_epg_subnet doc
- Bump mso-go-client to 1.4.3
- Rm old version of mso-go-client
- Fix mso_schema_template_anp_epg if only bd or vrf is provided
- In eepg contract resource, add check for relation type in order to get correct contractrelationship
- Update example for schema_template_external_epg_contract
- Update example for schema template external epgcontract
- Fix mso_schema_template_bd import issues (#136)
- Fix bug issue for mso_site import when platform is nd to keep consistency as mso platform (#137)
- Add cloudepg parameters only when epg_type is service
- Add option in mso_tenant to decide if deleting tenant from mso/ndo only or not (#162)
- Ensure correct path is used for mso_schema_template_deploy_ndo when provider platform is set to nd
- Fix enhanced_lag_policy for vmm domain attribute in mso_schema_site_anp_epg_domain (#180)
- Skip delete if method set explicitly in mso_rest resource (#191)
- Removed an unused variable tempvar from the resource_mso_rest.go
- Fix mso_schema read issue when the object is not present in the mso/ndo
- Fix mso resources state file refresh issue when the terraform managed objects were missing from the mso/ndo (#216)
- Fix path calculation for vpc and fex type ports in resource mso_schema_site_anp_epg_bulk_staticport (#218)
- Fix import and read functionality  and add documentation for mso_schema_site_external_epg (#204)
- Ensure correct information is retrieved with template/site combination for mso_schema_site_vrf_region and mso_schema_site_vrf_region_cidr_subnet datasource
- Template selection fix for import statements of site_vrf_region and site_vrf_region_cidr_subnet
- Allow multiple user association to be set in statefile from mso_tenant datatsource
- Fix conditional to set password when store in statefile is false
- Fix creation without aggregate or scope in mso_schema_template_external_epg_subnet for ndo4.1 (#228)
- Fix dhcp policies issue in mso_schema_template_bd on ndo 4.1+ (#230)
- Fix description setting for datasource_remote_location and resource_mso_remote_location because response changed to exclude empty descriptions
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_role
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_label
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and set rbac and email correctly in statefile for datasource_mso_user
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and set cloud site attributes correctly in statefile for datasource_mso_tenant
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add missing attributes type, group_id, version, status, reprovision, proxy, sr_l3out, template_count for datasource_mso_site
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_service_node_type
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and fix schema description setting to statefile for datasource_mso_schema
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_vrf
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_vrf_contract
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and writing multiple site_nodes to statefile for datasource_mso_schema_template_service_graph
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_l3out
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_filter_entry
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_external_epg and fix sellist setting in read and import of external epg resource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_external_epg_subnet
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_external_epg_selector
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_external_epg_contract
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and added directives, action to filter and removed directives from contract for datasource_mso_schema_template_contract
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and added bd_schema_id, bd_template_name to service_graph for datasource_mso_schema_template_contract_service_graph
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_contract_filter
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add primary attribute for datasource_mso_schema_template_bd_subnet
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_anp
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and prevent {} is written to statefile for datasource_mso_schema_template_anp_epg_useg_attr
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_anp_epg_subnet
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_anp_epg_selector
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly for datasource_mso_schema_template_anp_epg_contract
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_vrf
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_vrf_routeleak and change the id setting for resource/datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_vrf_region and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_vrf_region_cidr and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_vrf_region_cidr_subnet and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_external_epg and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check for datasource_mso_schema_site_external_epg_selector and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_bd and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_bd_l3out and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_bd_subnet and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg_subnet and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg_static_port and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg_static_leaf and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg_selector and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg_domain and change the id setting for datasource
- Fix schema attributes and documentation to reflect computed (read-only) attributes correctly and add template_name check and for datasource_mso_schema_site_anp_epg_bulk_staticport and change the id setting for datasource
- Fix for correct schema id and template name display for sub-objects l3out and contract in site_bd_l3out, site_external_epg, template_anp_epg_contract, template_contract_service_graph, template_external_epg_contract, template_vrf_contract datasources
- Fix to raise error when bd is not found and change the id set to include schema_id, template_name and bd_name for datasource_schema_template_bd
- Fix to change the id set to include schema_id, template_name and other object unique identifiers
- Fix for template_type attribute change check in mso_schema resource when not template_type is not specified and other attributes like description are changed
- Modified code to incorporate changes made in ndo version >= 4.2's api response for 'regions'
- Modified code in datasource 'mso_schema_site_vrf_region' to incorporate changes made in ndo version >= 4.2's api response for 'regions'
- Modified code in 'mso_tenant' to overcome index out of range error for aws sites
- Added an if condition to check the healthcheck list length before accessing the list elements
- Updated tls version on mso-go-client and upsted the changes on provider.
- Fix update and delete of mso_schema_template_anp_epg
- Prevent destroy operation for static ports in bulk list when updating the list
- Fix to prevent l3out reference is removed when host based routing for bd is updated
- Fix to prevent panic when labels are null in mso_site datasource
- Fix to prevent subnet from template bd is removed when template bd attributes are updated
- Fixed idempotency on site service graph. (#286)
- Fix undeployment on destroy by checking first if template is deployed for resource mso_schema_site.
- Fix importing for mso_schema_site_contract_service_graph redirect policy and cluster interface when they include hyphens in naming
- Optimization through an additional api call at list-identity schema endpoint to avoid retrieval of all schemas in order to get the id of a schema (#302)
- Fix read and import for mso_schema_site_anp resource to select anp in correct site and template combination
- Fix importing for mso_schema_template_anp_epg_contract to select the contract with the correct type
- Fix import for mso_schema_template_anp_epg_static_leaf to include template (dcne-214) (#306)
- Fix customise diff validation in mso_schema_site_service_graph to avoid retrieving all schemas when the schema id is unknown during plan
- Fix import order in mso_schema_site_bd_subnet and docs in mso_schema_template_bd_subnet (#305)
- Fix importing for mso_schema_site to support multiple templates in the same schema with the same site
- Fix fex and micro_seg_vlan attributes in resource_mso_schema_site_anp_epg_bulk_staticport to be correctly set when index shift occur in the static_ports list
- Add ability to retry status code 500 error when the error string is matching a proxy request error (#366)
- Fix mso_schema_site to add a retry mechanism on a pending undeploy operation and display error message when failing to undeploy
- Add re-login and retry mechanism when login session expire and 401 is returned
- Fix error handling message for non 200 status responses
- Avoid panic error on redirect policy when string is the empty or incorrectly split in mso_schema_site_contract_service_graph
- Sanitize url inputs to accepts wrongly appended / by updating mso-go-client to v1.33.4. (#383)
- Ensure service graph configuration in schema template contract is not deleted when there are updates to attributes of the contract
- Changed terraform resource id format from just schema_id to {schema_id}/template/{template_name} to distinguish deploy resources using the same schema_id
- Fix to update epg attributes which should not delete child objects for resource_mso_schema_template_anp_epg
- Fix to update vrf and bd without destruction of epg resource with resource_mso_schema_template_anp_epg to avoid removal of child objects
- Fix to update epg site attributes which should not delete child objects for resource_mso_schema_site_anp_epg
- Fix to update template attributes which should not delete child objects for resource_mso_schema_template
- Fix to update anp attributes which should not delete child objects for resource_mso_schema_template_anp
- Fix to update external epg attributes which should not delete child objects for resource_mso_schema_template_external_epg (dcne-719) (#462)

* [bugfix] fix to update external epg attributes which should not delete child objects for resource_mso_schema_template_external_epg

* [ignore] fix small errors in inline documentation sentences
- Fix to update vrf attributes which should not delete child objects for resource_mso_schema_template_vrf (dcne-717) (#464)

* [bugfix] fix to update vrf attributes which should not delete child objects for resource_mso_schema_template_vrf

* [ignore] fix small errors in inline documentation sentences

* [ignore] simplify logic of rendezvous_points and add tests

OTHER:
- Add missing provider environment variables
- Merge pull request #465 from akinross/epg_site_update_bug_fix

[bugfix] fix to update epg site attributes which should not delete child objects for resource_mso_schema_site_anp_epg (dcne-721)
- Merge pull request #463 from akinross/template_update_bug_fix

[bugfix] fix to update template attributes which should not delete child objects for resource_mso_schema_template (dcne-720)
- Merge pull request #461 from akinross/anp_update_bug_fix

[bugfix] fix to update anp attributes which should not delete child objects for resource_mso_schema_template_anp (dcne-718)

## 1.7.0 (December 08, 2025)

IMPROVEMENTS:
- Add mso_fabric_policies_macsec_policy resource and datasource.
- Add mso_schema_template_contract_service_chaining resource and datasource.
- Add mso_tenant_policies_bgp_peer_prefix_policy resource and datasource.

BUG FIXES:
- Fix mso_schema_template_contract to prevent deletion when attributes change.

## 1.6.0 (November 5, 2025)

IMPROVEMENTS:
- Add retrigger attribute to mso_rest to enable updates without payload, path or method changes.
- Add template_type and template_id attributes to mso_schema_template_deploy_ndo to enable template deployment of any type.
- Add mso_service_device_cluster resource and datasource.
- Add mso_fabric_policies_synce_interface_policy resource and datasource.
- Add uuid read-only attribute to mso_schema_template_bd resource and datasource.

BUG FIXES:
- Fix error when consumer_connector_redirect_policy or provider_connector_redirect_policy attributes are empty strings in mso_schema_site_contract_service_graph.
- Fix url attribute on the mso provider to accept appended slash characters.

## 1.5.3 (August 21, 2025)
BUG FIXES:
- Fix error handling issue that ignore error messages returned by the API for non 200 status responses introduced in v1.5.1

## 1.5.2 (August 8, 2025)
BUG FIXES:
- Add re-login and retry mechanism to fix issue when login session expire and 401 is returned

## 1.5.1 (July 24, 2025)
BUG FIXES:
- Fix mso_schema_site to add a retry mechanism on a pending undeploy operation and display error message when failing to undeploy
- Add ability to retry status code 500 error when the error string is matching a proxy request error

## 1.5.0 (July 17, 2025)
IMPROVEMENTS:
- Add mso_fabric_policies_physical_domain resource and datasource.
- Add mso_fabric_policies_vlan_pool resource and datasource.
- Add mso_tenant_policies_dhcp_relay_policy resource and datasource.
- Add mso_tenant_policies_ipsla_monitoring_policy resource and datasource. 
- Add mso_tenant_policies_route_map_policy_multicast resource and datasource.
- Add ability to wait for deploy task to finish and display error message if present in resource_ndo_schema_template_deploy.
- Add retry mechanism to provider for failed API requests due to network or capacity issues.
- Add uuid attribute to mso_schema_template_anp_epg and mso_schema_template_external_epg resources and datasources.
- Add rendezvous_points attribute in mso_schema_template_vrf resource and datasource.

## 1.4.0 (January 22, 2025)
IMPROVEMENTS:
- Add support for dpc path_type input with fex for mso_schema_site_anp_epg_staticport and mso_schema_site_anp_epg_bulk_staticport resources
- Add support for endpoint move detection mode in schema_template_bd

## 1.3.0 (December 2, 2024)
BUG FIXES:
- Fix fex and micro_seg_vlan attributes in resource_mso_schema_site_anp_epg_bulk_staticport to be correctly set when index shift occur in the static_ports list
- Fix importing for mso_schema_site to support multiple templates in the same schema with the same site
- Fix import order in mso_schema_site_bd_subnet and docs in mso_schema_template_bd_subnet (#305)
- Fix customize diff validation in mso_schema_site_service_graph to avoid retrieving all schemas when the schema id is unknown during plan
- Fix import for mso_schema_template_anp_epg_static_leaf to include template (DCNE-214) (#306)
- Fix importing for mso_schema_template_anp_epg_contract to select the contract with the correct type
- Fix read and import for mso_schema_site_anp resource to select anp in correct site and template combination
- Optimization through an additional api call at list-identity schema endpoint to avoid retrieval of all schemas in order to get the id of a schema (#302)
- Fix importing for mso_schema_site_contract_service_graph redirect policy and cluster interface when they include hyphens in naming
- Fix undeployment on destroy by checking first if template is deployed for resource mso_schema_site.

IMPROVEMENTS:
- Add l3out_schema_id, l3out_template and l3out_on_apic attributes to mso_schema_site_external_epg (#291)
- Add resource and datasource for mso_template
- Add support for vpc connected to fex in resource mso_schema_site_anp_epg_bulk_staticport

## 1.2.2 (August 6, 2024)
BUG FIXES:
- Fix idempotency issues in mso_schema_template_contract, mso_schema_site_contract_service_graph and mso_schema_site_service_graph resources.

## 1.2.1 (July 12, 2024)
BUG FIXES:
- Add check to avoid error in plan when mso_schema_site_service_graph resource is used when template does not exist yet.

## 1.2.0 (July 2, 2024)
BUG FIXES:
- Prevent destroy operation for static ports in bulk list when updating the list
- Fix update and delete of mso_schema_template_anp_epg
- Fix to prevent subnets from mso_schema_template_bd to be removed when the template bd attributes are updated
- Fix to prevent panic when labels are null in mso_site datasource
- Fix to prevent the L3out reference from being removed when host based routing for bd is updated in mso_schema_site_bd

IMPROVEMENTS:
- Add support for site_aware_policy_enforcement in mso_schema_template_vrf (#274)

## 1.1.0 (April 5, 2024)
BUG FIXES:
- Fix maximum TLS version to 1.3.

IMPROVEMENTS:
- Add attribute 'description' to a few applicable resources and data sources
- Add mso_schema_site_contract_service_graph_listener resource and data source to manage Azure CNC Contract Service Graph Load Balancer - Listeners and Rules (#256)
- Add missing Cloud APIC / CNC site parameters for mso_schema_site_service_graph module.

## 1.0.0 (December 6, 2023)
BREAKING CHANGE:
- Separating site and template level service graph provider from mso_schema_template_service_graph (#240)
- Removed Site Contract Service Graph logic from mso_schema_template_contract_service_graph resource (#244)

BUG FIXES:
- Fix schema attributes and documentation to reflect computed (Read-Only) attributes for all data-sources (#235)
- Fix template name selections for all data-sources (#235)
- Fix consistency of id attribute for all data-sources (#235)
- Fix to prevent {} to be written to statefile in datasource mso_schema_template_anp_epg_useg_attr (#235)
- Fix for setting correct schema id and template name for schema attributes l3out and contract in mso_schema_site_bd_l3out, mso_schema_site_external_epg, mso_schema_template_anp_epg_contract, mso_schema_template_contract_service_graph, mso_schema_template_external_epg_contract, and template_vrf_contract datasources (#235)
- Fix to raise error when bd is not found in datasource mso_schema_template_bd (#235)
- Fix for changes when template_type is not specified in mso_schema resource (#237)
- Fix to overcome index out of range error for aws sites in mso_tenant (#254)
- Fix for regions API changes in NDO version >= 4.2's in mso_schema_site_vrf_region (#254)

IMPROVEMENTS:
- Add support for mso_schema_site_contract_service_graph (#248)
- Add mso_system_config resource and datasource (#241)
- Add missing attributes type, group_id, version, status, reprovision, proxy, sr_l3out, template_count for datasource mso_site (#235)
- Add primary and virtual attribute to mso_schema_site_bd_subnet and mso_schema_template_bd_subnet (#235)
- Add bd_schema_id, bd_template_name to service_graph for datasource mso_schema_template_contract_service_graph (#235)
- Add filter_type, directives, action and priority attributes to filter_relationship to mso_schema_template_contract (#238)
- Add target_dscp and priority attributes to mso_schema_template_contract (#238)
- Add directives, action and priority attributes to mso_schema_template_contract_filter (#239)
- Add support for attributes binding_type, delimiter, num_ports, port_allocation, netflow, allow_promiscuous, forged_transmits, mac_changes, and custom_epg_name in mso_schema_site_epg_domain (#247)


## 0.11.1 (July 31, 2023)
BUG FIXES:
- Fix for mso_site to detect changes to a site, after manually changing it to unmanaged (#231)
- Fix errorForObjectNotFound function in utils to error out when new "error" payload is found. (#231)
- Fix DHCP Policies issue in mso_schema_template_bd on NDO 4.1+ (#230)
- Fix creation without aggregate or scope in mso_schema_template_external_epg_subnet for NDO4.1 (#228)

## 0.11.0 (July 1, 2023)
BUG FIXES:
- Fix conditional to set password when store in statefile is false
- Allow multiple user association to be set in statefile from mso_tenant data source
- Template selection fix for import statements of site_vrf_region and site_vrf_region_cidr_subnet
- Ensure correct information is retrieved with template/site combination for mso_schema_site_vrf_region and mso_schema_site_vrf_region_cidr_subnet datasource
- Fix import and read functionality  and add documentation for mso_schema_site_external_epg (#204)
- Fix path calculation for VPC and FEX type ports in resource mso_schema_site_anp_epg_bulk_staticport (#218)
- Fix MSO resources state file refresh issue when the terraform managed objects were missing from the MSO/NDO (#216)

IMPROVEMENTS:
- Add resource and datasource for mso_remote_location
- Add resource and data source mso_schema_site_vrf_route_leak for site local vrf route leaking support
- Allow schemas to be created without templates in ndo4.2 releases and above
- Add template type attribute to mso_schema and mso_schema_template resources and datasources
- Add description attribute to schema_template_bd and schema_template_anp_epg resources
- Add svi_mac attribute to  mso_schema_site_bd and make template_name attribute mandatory (#214)
- Add missing primary, description and no_default_gateway attributes to mso_site_anp_epg_subnet and mso_template_anp_epg_subnet resource and data source
- Add gcp specific attributes to mso_site_vrf_region_cidr_subnet resource and datasource
- Add gcp specific attributes to mso_site_vrf_region resource and datasource
- Add gcp specific attributes to mso_tenant resource and datasource

## 0.10.0 (May 24, 2023)
BUG FIXES:
- Fix mso_schema read issue when the object is not present in the MSO/NDO
- Skip delete if method set explicitely in mso_rest resource (#191)
- Fix mso_schema resource example template reference and remove deprecated usage of mso_schema resource (#202)
- Fix enhanced_lag_policy for vmm domain attribute in mso_schema_site_anp_epg_domain (#180)

IMPROVEMENTS:
- Add support for NDO4.1+ to mso_schema_site_anp, mso_schema_site_anp_epg, mso_schema_site_bd, mso_schema_site_external_epg and mso_schema_site_vrf (#188)
- Add mso_rest data source (#184)
- Add ip_data_plane_learning and preferred_group arguments to mso_schema_template_vrf resource (#177)
- Add missing provider level environment variables

## 0.9.0 (April 3, 2023)
IMPROVEMENTS:
- Add mso_schema_site_anp_epg_bulk_staticport resource and data source

## 0.8.1 (February 2, 2023)
BUG FIXES:
- Fix issue with platform set to nd and require platform set to nd when using mso_schema_template_deploy_ndo

## 0.8.0 (January 31, 2023)
BUG FIXES:
- Fix concurrency issues by implementing a mutex in the MSO/NDO Golang client

IMPROVEMENTS:
- Add mso_schema_template_deploy_ndo resource to support NDO4.1+ deploy API (#165)
- Add support for multiple DHCP Label policies in mso_schema_template_bd (#161)
- Add option in mso_tenant to decide if deleting tenant from mso/ndo only or not (#162)

## 0.7.1 (October 14, 2022)
BUG FIXES:
- Fix Cloud EPG default attribute issue in mso_schema_template_anp_epg
- Fix mso_schema_template_filter_entry crash when tcp_session_rules attribute is not provided

## 0.7.0 (August 26, 2022)
IMPROVEMENTS:
- Ability to add Microsoft and Redhat domains in mso_schema_site_anp_epg_domain (deprecate dn attribute) (#130)
- Ability to add multiple templates using the resource mso_schema (#135)
- Add import support for mso_schema_template_vrf_contract
- Apple M1 support
- Add name attribute to the subnet in resource_mso_schema_site_vrf_region (#140)

BUG FIXES:
- Fix import in mso_schema_template_anp_epg which caused configured values to be null (#129)
- Fix subnet name idempotency issue in mso_schema_template_external_epg_subnet
- Update example and update documentation for mso_schema_template_external_epg_subnet
- Fix mso_schema_template_anp_epg if only bd or vrf is provided
- Fix terraform import of schema_template_external_epg_contract
- Fix mso_schema_template_bd import issues (#136)
- Fix mso_schema_template_external_epg_subnet recreation issue when IP was changed
- Fix bug issue for mso_site import when platform is nd to keep consistency as mso platform (#137)
- Fix mso_schema_template_external_epg_contract relation type and update example

## 0.6.0 (March 15, 2022)
IMPROVEMENTS:
- Add Service EPG support for cloud sites (#113)

BUG FIXES:
- Update mso_schema_site_anp_epg_static_port documentation and example for path_type = vpc/dpc (#118)

## 0.5.0 (February 28, 2022)
IMPROVEMENTS:
- Add arp_flooding, virtual_mac_address, unicast_routing, ipv6_unknown_multicast_flooding, multi_destination_flooding and unknown_multicast_flooding in mso_schema_template_bd

BUG FIXES:
- Fix import documentation for mso_schema_template_external_epg_subnet, mso_schema_site_bd_subnet and mso_schema_template_filter_entry.
- Fix aci_site import and idempotency issues with NDO
- Add check to ignore error code 141 Resource Not Present to all Delete methods
- Fix idempotency issue of mso_schema_template_bd

## 0.4.1 (December 17, 2021)
BUG FIXES:
- Fix documentation for mso_schema_template_external_epg, mso_schema_template_external_epg_subnet and mso_schema_template_contract.

## 0.4.0 (December 16, 2021)
BEHAVIOR CHANGE:
- Modify mso_schema_template_deploy so destroy undeploy all sites

BUG FIXES:
- Fix idemptotency issue with site_id in template_external_epg

## 0.3.4 (December 16, 2021)

BUG FIXES:
- Fix login payload issue due to NDO API change (mso-go-client v1.2.6 update)
- Fix API change in l3outRef and add some error catching conditions

## 0.3.3 (December 15, 2021)

BUG FIXES:
- Fix login domain issue with NDO (mso-go-client v1.2.3 update)
- Fix mso_user data source when running on NDO
- Fix mso_schema_template_external_epg and mso_schema_site_external_epg when running on NDO

## 0.3.2 (October 29, 2021)

BUG FIXES:
- Fix issue in mso_tenant with Cloud Accounts

## 0.3.1 (September 30, 2021)

BUG FIXES:
- Fix DeletebyId crash issue with nil pointer (mso-go-client v1.2.2 update)
- Make zone attribute not mandatory in mso_schema_site_vrf_region and mso_schema_site_vrf_region_cidr_subnet (to support Azure sites)

## 0.3.0 (September 24, 2021)

BUG FIXES:
- Fix mso_schema_site_anp_epg Import when VRF or BD is not configured

IMPROVEMENTS:
- Added new resource mso_schema_site_external_epg
- Updated mso_user to support MSO/ND platform
- Updated resource and  datasource mso_site to work with ND-based MSO / NDO
- Updated mso_schema_template_contract to support multiple filter_relationship and deprecate filter_relationships attribute.

## 0.2.1 (July 20, 2021)

BUG FIXES:
- Added examples and documentations for Multi-Site associations
- Fix ressources creation and update when PATCH request return 204 No Content

## 0.2.0 (June 23, 2021)

BUG FIXES:
- Fix mso_schema_site_anp_epg_static_leaf documentation example
- Fix mso_site documentation for location attribute (locations -> location)

IMPROVEMENTS:
- Support of ND authentication (platform = "nd") for MSO v3.2+
- Updated mso_schema_site and mso_schema_template_vrf resource docs name
- Updated all terraform examples with required_provider section added in recent Terraform versions
- Added example of how to use for_each to support multiple static ports
- Added import capability on all resources
- Add support for fex ports in mso_schema_site_anp_epg_static_port
- Cleanup of go.mod and vendor directory

## 0.1.5 (January 25, 2021)

BUG FIXES:
- Fixed an issue with mso_tenant resource, where users were not able to attach the On-Prem site to the tenants due to panic error.

## 0.1.3 (October 28, 2020)

IMPROVEMENTS:

- Enabled the Auth-token sharing between API calls.

BUG FIXES:
- Fixed an issue with mso_schema_template_contract resource, where wrong filter was being read when there are multiple filters added to the same contract.

## 0.1.2 (September 3, 2020)

IMPROVEMENTS:

- Renamed resources for naming consistency.
- First Terraform Registry release.

## 0.1.1 (July 22, 2020)

IMPROVEMENTS:

- Added new resources to manage selectors on various levels.
- Added support for login_domains to site resource.
- Added Azure and AWS account support.
- Renamed resources for better formatting.

## 0.1.0 (June 17, 2020)

- Initial Release
