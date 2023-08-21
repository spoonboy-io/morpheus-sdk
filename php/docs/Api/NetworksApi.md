# OpenAPI\Client\NetworksApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**createNetworkDhcpRelay()**](NetworksApi.md#createNetworkDhcpRelay) | **POST** /api/networks/servers/{serverId}/dhcp-relays | Create a Network DHCP Relay
[**createNetworkDhcpServer()**](NetworksApi.md#createNetworkDhcpServer) | **POST** /api/networks/servers/{serverId}/dhcp-servers | Create a Network DHCP Server
[**createNetworkDomain()**](NetworksApi.md#createNetworkDomain) | **POST** /api/networks/domains | Create a Network Domain
[**createNetworkFirewallRule()**](NetworksApi.md#createNetworkFirewallRule) | **POST** /api/networks/servers/{serverId}/firewall-rules | Create a Network Firewall Rule
[**createNetworkFirewallRuleGroup()**](NetworksApi.md#createNetworkFirewallRuleGroup) | **POST** /api/networks/servers/{serverId}/firewall-rule-groups | Create a Network Firewall Rule Group
[**createNetworkGroup()**](NetworksApi.md#createNetworkGroup) | **POST** /api/networks/groups | Create a Network Group
[**createNetworkPool()**](NetworksApi.md#createNetworkPool) | **POST** /api/networks/pools | Create a Network Pool
[**createNetworkPoolIp()**](NetworksApi.md#createNetworkPoolIp) | **POST** /api/networks/pools/{id}/ips | Create a Network Pool IP Address
[**createNetworkPoolServer()**](NetworksApi.md#createNetworkPoolServer) | **POST** /api/networks/pool-servers | Create a Network Pool Server
[**createNetworkProxy()**](NetworksApi.md#createNetworkProxy) | **POST** /api/networks/proxies | Create a Network Proxy
[**createNetworkRouter()**](NetworksApi.md#createNetworkRouter) | **POST** /api/networks/routers | Create a Network Router
[**createNetworkRouterBgpNeighbor()**](NetworksApi.md#createNetworkRouterBgpNeighbor) | **POST** /api/networks/routers/{routerId}/bgp-neighbors | Create a Network Router BGP Neighbor
[**createNetworkRouterFirewallRule()**](NetworksApi.md#createNetworkRouterFirewallRule) | **POST** /api/networks/routers/{routerId}/firewall-rules | Create a Network Router Firewall Rule
[**createNetworkRouterFirewallRuleGroup()**](NetworksApi.md#createNetworkRouterFirewallRuleGroup) | **POST** /api/networks/routers/{routerId}/firewall-rule-groups | Create a Network Router Firewall Rule Group
[**createNetworkRouterNat()**](NetworksApi.md#createNetworkRouterNat) | **POST** /api/networks/routers/{routerId}/nats | Create a Network Router NAT
[**createNetworkRouterRoute()**](NetworksApi.md#createNetworkRouterRoute) | **POST** /api/networks/routers/{routerId}/routes | Create a Network Router Route
[**createNetworkServerGroup()**](NetworksApi.md#createNetworkServerGroup) | **POST** /api/networks/servers/{serverId}/groups | Create a Network Server Group
[**createNetworkTransportZone()**](NetworksApi.md#createNetworkTransportZone) | **POST** /api/networks/servers/{serverId}/scopes | Create a Network Transport Zone
[**createNetworks()**](NetworksApi.md#createNetworks) | **POST** /api/networks | Create a Network
[**createStaticRoute()**](NetworksApi.md#createStaticRoute) | **PUT** /api/networks/{id}/routes | Create a Network Static Route
[**createSubnet()**](NetworksApi.md#createSubnet) | **POST** /api/subnets | Create a Subnet
[**deleteNetwork()**](NetworksApi.md#deleteNetwork) | **DELETE** /api/networks/{id} | Delete a Network
[**deleteNetworkDhcpRelay()**](NetworksApi.md#deleteNetworkDhcpRelay) | **DELETE** /api/networks/servers/{serverId}/dhcp-relays/{id} | Delete a Network DHCP Relay
[**deleteNetworkDhcpServer()**](NetworksApi.md#deleteNetworkDhcpServer) | **DELETE** /api/networks/servers/{serverId}/dhcp-servers/{id} | Delete a Network DHCP Server
[**deleteNetworkDomain()**](NetworksApi.md#deleteNetworkDomain) | **DELETE** /api/networks/domains/{id} | Delete a Network Domain
[**deleteNetworkFirewallRule()**](NetworksApi.md#deleteNetworkFirewallRule) | **DELETE** /api/networks/servers/{serverId}/firewall-rules/{id} | Delete a Network Firewall Rule
[**deleteNetworkFirewallRuleGroup()**](NetworksApi.md#deleteNetworkFirewallRuleGroup) | **DELETE** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Delete a Network firewall rule group
[**deleteNetworkGroup()**](NetworksApi.md#deleteNetworkGroup) | **DELETE** /api/networks/groups/{id} | Delete a Network Group
[**deleteNetworkPool()**](NetworksApi.md#deleteNetworkPool) | **DELETE** /api/networks/pools/{id} | Delete a Network Pool
[**deleteNetworkPoolIp()**](NetworksApi.md#deleteNetworkPoolIp) | **DELETE** /api/networks/pools/{networkPoolId}/ips/{id} | Delete a host record associated with an IP Address for a Specific Network Pool
[**deleteNetworkPoolServer()**](NetworksApi.md#deleteNetworkPoolServer) | **DELETE** /api/networks/pool-servers/{id} | Delete a Network Pool Server
[**deleteNetworkProxy()**](NetworksApi.md#deleteNetworkProxy) | **DELETE** /api/networks/proxies/{id} | Delete a Network Proxy
[**deleteNetworkRouter()**](NetworksApi.md#deleteNetworkRouter) | **DELETE** /api/networks/routers/{id} | Delete a Network Router
[**deleteNetworkRouterBgpNeighbor()**](NetworksApi.md#deleteNetworkRouterBgpNeighbor) | **DELETE** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Delete a Network Router BGP Neighbor
[**deleteNetworkRouterFirewallRule()**](NetworksApi.md#deleteNetworkRouterFirewallRule) | **DELETE** /api/networks/routers/{routerId}/firewall-rules/{id} | Delete a Network Router Firewall Rule
[**deleteNetworkRouterFirewallRuleGroup()**](NetworksApi.md#deleteNetworkRouterFirewallRuleGroup) | **DELETE** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Delete a Network Router firewall rule group
[**deleteNetworkRouterNat()**](NetworksApi.md#deleteNetworkRouterNat) | **DELETE** /api/networks/routers/{routerId}/nats/{id} | Delete a Network Router NAT
[**deleteNetworkRouterRoute()**](NetworksApi.md#deleteNetworkRouterRoute) | **DELETE** /api/networks/routers/{routerId}/routes/{id} | Delete a Network Router Route
[**deleteNetworkServerGroup()**](NetworksApi.md#deleteNetworkServerGroup) | **DELETE** /api/networks/servers/{serverId}/groups/{id} | Delete a Network Server Group
[**deleteNetworkTransportZone()**](NetworksApi.md#deleteNetworkTransportZone) | **DELETE** /api/networks/servers/{serverId}/scopes/{id} | Delete a Network Transport Zone
[**deleteStaticRoute()**](NetworksApi.md#deleteStaticRoute) | **DELETE** /api/networks/{id}/routes/{routeId} | Delete a Network Static Route
[**deleteSubnet()**](NetworksApi.md#deleteSubnet) | **DELETE** /api/subnets/{id} | Delete a Subnet
[**getAllNetworkFloatingIps()**](NetworksApi.md#getAllNetworkFloatingIps) | **GET** /api/networks/floating-ips | Get All Floating IPs
[**getNetwork()**](NetworksApi.md#getNetwork) | **GET** /api/networks/{id} | Get a Specific Network
[**getNetworkDhcpRelay()**](NetworksApi.md#getNetworkDhcpRelay) | **GET** /api/networks/servers/{serverId}/dhcp-relays/{id} | Get a Specific Network DHCP Relay
[**getNetworkDhcpRelays()**](NetworksApi.md#getNetworkDhcpRelays) | **GET** /api/networks/servers/{serverId}/dhcp-relays | Get all Network DHCP Relays for Network Relay
[**getNetworkDhcpServer()**](NetworksApi.md#getNetworkDhcpServer) | **GET** /api/networks/servers/{serverId}/dhcp-servers/{id} | Get a Specific Network DHCP Server
[**getNetworkDhcpServers()**](NetworksApi.md#getNetworkDhcpServers) | **GET** /api/networks/servers/{serverId}/dhcp-servers | Get all Network DHCP Servers for Network Server
[**getNetworkDomain()**](NetworksApi.md#getNetworkDomain) | **GET** /api/networks/domains/{id} | Get a Specific Network Domain
[**getNetworkDomains()**](NetworksApi.md#getNetworkDomains) | **GET** /api/networks/domains | Get all Network Domains
[**getNetworkEdgeCluster()**](NetworksApi.md#getNetworkEdgeCluster) | **GET** /api/networks/servers/{serverId}/edge-clusters/{id} | Get a Specific Network Edge Cluster
[**getNetworkEdgeClusters()**](NetworksApi.md#getNetworkEdgeClusters) | **GET** /api/networks/servers/{serverId}/edge-clusters | Get all Network Edge Clusters for Network Server
[**getNetworkFirewallRule()**](NetworksApi.md#getNetworkFirewallRule) | **GET** /api/networks/servers/{serverId}/firewall-rules/{id} | Get a Specific Network Firewall Rule
[**getNetworkFirewallRuleGroup()**](NetworksApi.md#getNetworkFirewallRuleGroup) | **GET** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Get a Specific Network Firewall Rule Group
[**getNetworkFirewallRuleGroups()**](NetworksApi.md#getNetworkFirewallRuleGroups) | **GET** /api/networks/servers/{serverId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Server
[**getNetworkFirewallRules()**](NetworksApi.md#getNetworkFirewallRules) | **GET** /api/networks/servers/{serverId}/firewall-rules | Get all Network Firewall Rules for Network Server
[**getNetworkFloatingIp()**](NetworksApi.md#getNetworkFloatingIp) | **GET** /api/networks/floating-ips/{id} | Get a Specific Floating IP
[**getNetworkGroup()**](NetworksApi.md#getNetworkGroup) | **GET** /api/networks/groups/{id} | Get a Specific Network Group
[**getNetworkGroups()**](NetworksApi.md#getNetworkGroups) | **GET** /api/networks/groups | Get all Network Groups
[**getNetworkPool()**](NetworksApi.md#getNetworkPool) | **GET** /api/networks/pools/{id} | Get a Specific Network Pool
[**getNetworkPoolIp()**](NetworksApi.md#getNetworkPoolIp) | **GET** /api/networks/pools/{networkPoolId}/ips/{id} | Get a Specific IP Address for a Specific Network Pool
[**getNetworkPoolIps()**](NetworksApi.md#getNetworkPoolIps) | **GET** /api/networks/pools/{id}/ips | Get all IP Addresses for a Specific Network Pool
[**getNetworkPoolServer()**](NetworksApi.md#getNetworkPoolServer) | **GET** /api/networks/pool-servers/{id} | Get a Specific Network Pool Server
[**getNetworkPoolServerType()**](NetworksApi.md#getNetworkPoolServerType) | **GET** /api/networks/pool-server-types/{id} | Retrieves a Specific Network Pool Server Type
[**getNetworkPools()**](NetworksApi.md#getNetworkPools) | **GET** /api/networks/pools | Get all Network Pools
[**getNetworkProxies()**](NetworksApi.md#getNetworkProxies) | **GET** /api/networks/proxies | Get all Network Proxies
[**getNetworkProxy()**](NetworksApi.md#getNetworkProxy) | **GET** /api/networks/proxies/{id} | Get a Specific Network Proxy
[**getNetworkRouter()**](NetworksApi.md#getNetworkRouter) | **GET** /api/networks/routers/{id} | Get a Specific Network Router
[**getNetworkRouterBgpNeighbor()**](NetworksApi.md#getNetworkRouterBgpNeighbor) | **GET** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Get a Network Router BGP Neighbor
[**getNetworkRouterFirewallRule()**](NetworksApi.md#getNetworkRouterFirewallRule) | **GET** /api/networks/routers/{routerId}/firewall-rules/{id} | Get a Firewall Rule for Network Router
[**getNetworkRouterFirewallRuleGroup()**](NetworksApi.md#getNetworkRouterFirewallRuleGroup) | **GET** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Get a Firewall Rule Group for Network Router
[**getNetworkRouterFirewallRuleGroups()**](NetworksApi.md#getNetworkRouterFirewallRuleGroups) | **GET** /api/networks/routers/{routerId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Router
[**getNetworkRouterNat()**](NetworksApi.md#getNetworkRouterNat) | **GET** /api/networks/routers/{routerId}/nats/{id} | Get a Network Router NAT
[**getNetworkRouterRoute()**](NetworksApi.md#getNetworkRouterRoute) | **GET** /api/networks/routers/{routerId}/routes/{id} | Get a Route for Network Router
[**getNetworkRouterType()**](NetworksApi.md#getNetworkRouterType) | **GET** /api/network-router-types/{id} | Get a Specific Network Router Type
[**getNetworkRouters()**](NetworksApi.md#getNetworkRouters) | **GET** /api/networks/routers | Get all Network Routers
[**getNetworkRoutersBgpNeighbors()**](NetworksApi.md#getNetworkRoutersBgpNeighbors) | **GET** /api/networks/routers/{routerId}/bgp-neighbors | Get all BGP Neighbors for Network Router
[**getNetworkRoutersFirewallRules()**](NetworksApi.md#getNetworkRoutersFirewallRules) | **GET** /api/networks/routers/{routerId}/firewall-rules | Get all Firewall Rules for Network Router
[**getNetworkRoutersNats()**](NetworksApi.md#getNetworkRoutersNats) | **GET** /api/networks/routers/{routerId}/nats | Get all Network Router NATs for Network Router
[**getNetworkRoutersRoutes()**](NetworksApi.md#getNetworkRoutersRoutes) | **GET** /api/networks/routers/{routerId}/routes | Get all Routes for Network Router
[**getNetworkServerGroup()**](NetworksApi.md#getNetworkServerGroup) | **GET** /api/networks/servers/{serverId}/groups/{id} | Get a specific Network Server Group
[**getNetworkSubnets()**](NetworksApi.md#getNetworkSubnets) | **GET** /api/networks/{id}/subnets | Get Subnets for a Network
[**getNetworkTransportZone()**](NetworksApi.md#getNetworkTransportZone) | **GET** /api/networks/servers/{serverId}/scopes/{id} | Get a Specific Network Transport Zone
[**getNetworkTransportZones()**](NetworksApi.md#getNetworkTransportZones) | **GET** /api/networks/servers/{serverId}/scopes | Get all Network Transport Zones for Network Server
[**getNetworkType()**](NetworksApi.md#getNetworkType) | **GET** /api/network-types/{id} | Get a Specific Network Type
[**getStaticRoute()**](NetworksApi.md#getStaticRoute) | **GET** /api/networks/{id}/routes/{routeId} | Get a Static Route for a Network
[**getStaticRoutes()**](NetworksApi.md#getStaticRoutes) | **GET** /api/networks/{id}/routes | Get all Static Routes for a Network
[**getSubnet()**](NetworksApi.md#getSubnet) | **GET** /api/subnets/{id} | Get a Specific Subnet
[**listNetworkPoolServerTypes()**](NetworksApi.md#listNetworkPoolServerTypes) | **GET** /api/networks/pool-server-types | Get All Network Pool Server Types
[**listNetworkPoolServers()**](NetworksApi.md#listNetworkPoolServers) | **GET** /api/networks/pool-servers | Get All Network Pool Servers
[**listNetworkRouterTypes()**](NetworksApi.md#listNetworkRouterTypes) | **GET** /api/network-router-types | Get all Network Router Types
[**listNetworkServerGroups()**](NetworksApi.md#listNetworkServerGroups) | **GET** /api/networks/servers/{serverId}/groups | Get all Network Server Groups for Network Server
[**listNetworkServices()**](NetworksApi.md#listNetworkServices) | **GET** /api/networks/services | Get All Network Services
[**listNetworkTypes()**](NetworksApi.md#listNetworkTypes) | **GET** /api/network-types | Network Types
[**listNetworks()**](NetworksApi.md#listNetworks) | **GET** /api/networks | Get All Networks
[**listSubnetTypes()**](NetworksApi.md#listSubnetTypes) | **GET** /api/subnet-types | Get All Subnet Types
[**listSubnets()**](NetworksApi.md#listSubnets) | **GET** /api/subnets | Get All Subnets
[**refreshNetworkServer()**](NetworksApi.md#refreshNetworkServer) | **POST** /api/networks/servers/{id}/refresh | Refresh a Network Server/Integration
[**releaseNetworkFloatingIp()**](NetworksApi.md#releaseNetworkFloatingIp) | **PUT** /api/networks/floating-ips/{id}/release | Release a Floating IP
[**updateNetwork()**](NetworksApi.md#updateNetwork) | **PUT** /api/networks/{id} | Update a Network
[**updateNetworkDhcpRelay()**](NetworksApi.md#updateNetworkDhcpRelay) | **PUT** /api/networks/servers/{serverId}/dhcp-relays/{id} | Update a Network DHCP Relay
[**updateNetworkDhcpServer()**](NetworksApi.md#updateNetworkDhcpServer) | **PUT** /api/networks/servers/{serverId}/dhcp-servers/{id} | Update a Network DHCP Server
[**updateNetworkDomain()**](NetworksApi.md#updateNetworkDomain) | **PUT** /api/networks/domains/{id} | Update a Network Domain
[**updateNetworkEdgeCluster()**](NetworksApi.md#updateNetworkEdgeCluster) | **PUT** /api/networks/servers/{serverId}/edge-clusters/{id} | Update a Network Edge Cluster
[**updateNetworkFirewallRule()**](NetworksApi.md#updateNetworkFirewallRule) | **PUT** /api/networks/servers/{serverId}/firewall-rules/{id} | Update a Network Firewall Rule
[**updateNetworkFirewallRuleGroup()**](NetworksApi.md#updateNetworkFirewallRuleGroup) | **PUT** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Update a Network Firewall Rule Group
[**updateNetworkGroup()**](NetworksApi.md#updateNetworkGroup) | **PUT** /api/networks/groups/{id} | Update a Network Group
[**updateNetworkPool()**](NetworksApi.md#updateNetworkPool) | **PUT** /api/networks/pools/{id} | Update a Network Pool
[**updateNetworkPoolServer()**](NetworksApi.md#updateNetworkPoolServer) | **PUT** /api/networks/pool-servers/{id} | Update a Network Pool Server
[**updateNetworkProxy()**](NetworksApi.md#updateNetworkProxy) | **PUT** /api/networks/proxies/{id} | Update a Network Proxy
[**updateNetworkRouter()**](NetworksApi.md#updateNetworkRouter) | **PUT** /api/networks/routers/{id} | Update a Network Router
[**updateNetworkRouterBgpNeighbor()**](NetworksApi.md#updateNetworkRouterBgpNeighbor) | **PUT** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Update Network Router BGP Neighbor
[**updateNetworkRouterFirewallRule()**](NetworksApi.md#updateNetworkRouterFirewallRule) | **PUT** /api/networks/routers/{routerId}/firewall-rules/{id} | Update a Network Router Firewall Rule
[**updateNetworkRouterFirewallRuleGroup()**](NetworksApi.md#updateNetworkRouterFirewallRuleGroup) | **PUT** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Update a Network Router Firewall Rule Group
[**updateNetworkRouterNat()**](NetworksApi.md#updateNetworkRouterNat) | **PUT** /api/networks/routers/{routerId}/nats/{id} | Update Network Router NAT
[**updateNetworkRouterPermissions()**](NetworksApi.md#updateNetworkRouterPermissions) | **PUT** /api/networks/routers/{routerId}/permissions | Update Network Router Permissions
[**updateNetworkServerGroup()**](NetworksApi.md#updateNetworkServerGroup) | **PUT** /api/networks/servers/{serverId}/groups/{id} | Update a Network Server Group
[**updateNetworkTransportZone()**](NetworksApi.md#updateNetworkTransportZone) | **PUT** /api/networks/servers/{serverId}/scopes/{id} | Update a Network Transport Zone
[**updateStaticRoute()**](NetworksApi.md#updateStaticRoute) | **PUT** /api/networks/{id}/routes/{routeId} | Update a Network Static Route
[**updateSubnet()**](NetworksApi.md#updateSubnet) | **PUT** /api/subnets/{id} | Update a Subnet


## `createNetworkDhcpRelay()`

```php
createNetworkDhcpRelay($server_id, $inline_object167): \OpenAPI\Client\Model\SuccessId
```

Create a Network DHCP Relay

Create a Network DHCP Relay.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$inline_object167 = new \OpenAPI\Client\Model\InlineObject167(); // \OpenAPI\Client\Model\InlineObject167

try {
    $result = $apiInstance->createNetworkDhcpRelay($server_id, $inline_object167);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkDhcpRelay: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **inline_object167** | [**\OpenAPI\Client\Model\InlineObject167**](../Model/InlineObject167.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkDhcpServer()`

```php
createNetworkDhcpServer($server_id, $inline_object169): \OpenAPI\Client\Model\SuccessId
```

Create a Network DHCP Server

Create a Network DHCP Server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$inline_object169 = new \OpenAPI\Client\Model\InlineObject169(); // \OpenAPI\Client\Model\InlineObject169

try {
    $result = $apiInstance->createNetworkDhcpServer($server_id, $inline_object169);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkDhcpServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **inline_object169** | [**\OpenAPI\Client\Model\InlineObject169**](../Model/InlineObject169.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkDomain()`

```php
createNetworkDomain($inline_object163): \OpenAPI\Client\Model\InlineResponse200109
```

Create a Network Domain

Create a Network Domain.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object163 = new \OpenAPI\Client\Model\InlineObject163(); // \OpenAPI\Client\Model\InlineObject163

try {
    $result = $apiInstance->createNetworkDomain($inline_object163);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkDomain: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object163** | [**\OpenAPI\Client\Model\InlineObject163**](../Model/InlineObject163.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200109**](../Model/InlineResponse200109.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkFirewallRule()`

```php
createNetworkFirewallRule($server_id, $inline_object172): \OpenAPI\Client\Model\SuccessId
```

Create a Network Firewall Rule

Use this command to create a network firewall rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$inline_object172 = new \OpenAPI\Client\Model\InlineObject172(); // \OpenAPI\Client\Model\InlineObject172

try {
    $result = $apiInstance->createNetworkFirewallRule($server_id, $inline_object172);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **inline_object172** | [**\OpenAPI\Client\Model\InlineObject172**](../Model/InlineObject172.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkFirewallRuleGroup()`

```php
createNetworkFirewallRuleGroup($server_id, $inline_object174): \OpenAPI\Client\Model\SuccessId
```

Create a Network Firewall Rule Group

Use this command to create a network firewall rule group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$inline_object174 = new \OpenAPI\Client\Model\InlineObject174(); // \OpenAPI\Client\Model\InlineObject174

try {
    $result = $apiInstance->createNetworkFirewallRuleGroup($server_id, $inline_object174);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **inline_object174** | [**\OpenAPI\Client\Model\InlineObject174**](../Model/InlineObject174.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkGroup()`

```php
createNetworkGroup($inline_object146): \OpenAPI\Client\Model\SuccessId
```

Create a Network Group

Use this command to create a network group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object146 = new \OpenAPI\Client\Model\InlineObject146(); // \OpenAPI\Client\Model\InlineObject146

try {
    $result = $apiInstance->createNetworkGroup($inline_object146);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object146** | [**\OpenAPI\Client\Model\InlineObject146**](../Model/InlineObject146.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkPool()`

```php
createNetworkPool($inline_object160): \OpenAPI\Client\Model\InlineResponse200106
```

Create a Network Pool

Create a Network Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object160 = new \OpenAPI\Client\Model\InlineObject160(); // \OpenAPI\Client\Model\InlineObject160

try {
    $result = $apiInstance->createNetworkPool($inline_object160);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object160** | [**\OpenAPI\Client\Model\InlineObject160**](../Model/InlineObject160.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200106**](../Model/InlineResponse200106.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkPoolIp()`

```php
createNetworkPoolIp($id, $inline_object162): \OpenAPI\Client\Model\InlineResponse200107
```

Create a Network Pool IP Address

Create an IP Address for a Specific Network Pool

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object162 = new \OpenAPI\Client\Model\InlineObject162(); // \OpenAPI\Client\Model\InlineObject162

try {
    $result = $apiInstance->createNetworkPoolIp($id, $inline_object162);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkPoolIp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object162** | [**\OpenAPI\Client\Model\InlineObject162**](../Model/InlineObject162.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200107**](../Model/InlineResponse200107.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkPoolServer()`

```php
createNetworkPoolServer($inline_object180): Model200Success
```

Create a Network Pool Server

This endpoint allows creating a Network Pool Server. Only certain types of integrations support creating and deleting network pool servers, such as Bluecat, Infoblox, phpIPAM, and Solar Winds. Configuration options vary by type. Note that creating a pool server will automatically create and associate the corresponding network integration object, but management is done via the network pool server object.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object180 = new \OpenAPI\Client\Model\InlineObject180(); // \OpenAPI\Client\Model\InlineObject180

try {
    $result = $apiInstance->createNetworkPoolServer($inline_object180);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkPoolServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object180** | [**\OpenAPI\Client\Model\InlineObject180**](../Model/InlineObject180.md)|  | [optional]

### Return type

[**Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkProxy()`

```php
createNetworkProxy($inline_object165): \OpenAPI\Client\Model\InlineResponse200110
```

Create a Network Proxy

Create a Network Proxy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object165 = new \OpenAPI\Client\Model\InlineObject165(); // \OpenAPI\Client\Model\InlineObject165

try {
    $result = $apiInstance->createNetworkProxy($inline_object165);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkProxy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object165** | [**\OpenAPI\Client\Model\InlineObject165**](../Model/InlineObject165.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200110**](../Model/InlineResponse200110.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkRouter()`

```php
createNetworkRouter($inline_object148): \OpenAPI\Client\Model\SuccessId
```

Create a Network Router

Use this command to create a network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object148 = new \OpenAPI\Client\Model\InlineObject148(); // \OpenAPI\Client\Model\InlineObject148

try {
    $result = $apiInstance->createNetworkRouter($inline_object148);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkRouter: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object148** | [**\OpenAPI\Client\Model\InlineObject148**](../Model/InlineObject148.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkRouterBgpNeighbor()`

```php
createNetworkRouterBgpNeighbor($router_id, $inline_object150): \OpenAPI\Client\Model\SuccessId
```

Create a Network Router BGP Neighbor

Use this command to create a BGP Neighbor for an existing network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID
$inline_object150 = new \OpenAPI\Client\Model\InlineObject150(); // \OpenAPI\Client\Model\InlineObject150

try {
    $result = $apiInstance->createNetworkRouterBgpNeighbor($router_id, $inline_object150);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkRouterBgpNeighbor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **inline_object150** | [**\OpenAPI\Client\Model\InlineObject150**](../Model/InlineObject150.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkRouterFirewallRule()`

```php
createNetworkRouterFirewallRule($router_id, $inline_object152): \OpenAPI\Client\Model\SuccessId
```

Create a Network Router Firewall Rule

Use this command to create a firewall rule for an existing network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID
$inline_object152 = new \OpenAPI\Client\Model\InlineObject152(); // \OpenAPI\Client\Model\InlineObject152

try {
    $result = $apiInstance->createNetworkRouterFirewallRule($router_id, $inline_object152);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkRouterFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **inline_object152** | [**\OpenAPI\Client\Model\InlineObject152**](../Model/InlineObject152.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkRouterFirewallRuleGroup()`

```php
createNetworkRouterFirewallRuleGroup($router_id, $inline_object154): \OpenAPI\Client\Model\SuccessId
```

Create a Network Router Firewall Rule Group

Use this command to create a network firewall rule group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID
$inline_object154 = new \OpenAPI\Client\Model\InlineObject154(); // \OpenAPI\Client\Model\InlineObject154

try {
    $result = $apiInstance->createNetworkRouterFirewallRuleGroup($router_id, $inline_object154);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkRouterFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **inline_object154** | [**\OpenAPI\Client\Model\InlineObject154**](../Model/InlineObject154.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkRouterNat()`

```php
createNetworkRouterNat($router_id, $inline_object156): \OpenAPI\Client\Model\SuccessId
```

Create a Network Router NAT

Use this command to create a NAT for an existing network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID
$inline_object156 = new \OpenAPI\Client\Model\InlineObject156(); // \OpenAPI\Client\Model\InlineObject156

try {
    $result = $apiInstance->createNetworkRouterNat($router_id, $inline_object156);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkRouterNat: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **inline_object156** | [**\OpenAPI\Client\Model\InlineObject156**](../Model/InlineObject156.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkRouterRoute()`

```php
createNetworkRouterRoute($router_id, $inline_object159): \OpenAPI\Client\Model\SuccessId
```

Create a Network Router Route

Use this command to create a route for an existing network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID
$inline_object159 = new \OpenAPI\Client\Model\InlineObject159(); // \OpenAPI\Client\Model\InlineObject159

try {
    $result = $apiInstance->createNetworkRouterRoute($router_id, $inline_object159);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkRouterRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **inline_object159** | [**\OpenAPI\Client\Model\InlineObject159**](../Model/InlineObject159.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkServerGroup()`

```php
createNetworkServerGroup($server_id, $inline_object176): \OpenAPI\Client\Model\SuccessId
```

Create a Network Server Group

Use this command to create a network server group. Note: Only available for NSX-T network integrations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$inline_object176 = new \OpenAPI\Client\Model\InlineObject176(); // \OpenAPI\Client\Model\InlineObject176

try {
    $result = $apiInstance->createNetworkServerGroup($server_id, $inline_object176);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkServerGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **inline_object176** | [**\OpenAPI\Client\Model\InlineObject176**](../Model/InlineObject176.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworkTransportZone()`

```php
createNetworkTransportZone($server_id, $inline_object178): \OpenAPI\Client\Model\SuccessId
```

Create a Network Transport Zone

Use this command to create a network transport zone.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$inline_object178 = new \OpenAPI\Client\Model\InlineObject178(); // \OpenAPI\Client\Model\InlineObject178

try {
    $result = $apiInstance->createNetworkTransportZone($server_id, $inline_object178);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworkTransportZone: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **inline_object178** | [**\OpenAPI\Client\Model\InlineObject178**](../Model/InlineObject178.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createNetworks()`

```php
createNetworks($inline_object142): object
```

Create a Network

This endpoint allows creating a Network. Only certain types of clouds support creating and deleting networks. Configuration options vary by Network Types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object142 = new \OpenAPI\Client\Model\InlineObject142(); // \OpenAPI\Client\Model\InlineObject142

try {
    $result = $apiInstance->createNetworks($inline_object142);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createNetworks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object142** | [**\OpenAPI\Client\Model\InlineObject142**](../Model/InlineObject142.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createStaticRoute()`

```php
createStaticRoute($id, $inline_object144): \OpenAPI\Client\Model\SuccessId
```

Create a Network Static Route

Use this command to create a route for an existing network.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object144 = new \OpenAPI\Client\Model\InlineObject144(); // \OpenAPI\Client\Model\InlineObject144

try {
    $result = $apiInstance->createStaticRoute($id, $inline_object144);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createStaticRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object144** | [**\OpenAPI\Client\Model\InlineObject144**](../Model/InlineObject144.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createSubnet()`

```php
createSubnet($inline_object244): \OpenAPI\Client\Model\InlineResponse200154
```

Create a Subnet

This endpoint allows creating a Subnet. Only certain types of clouds support creating and deleting subnets. Configuration options vary for each Subnet Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object244 = new \OpenAPI\Client\Model\InlineObject244(); // \OpenAPI\Client\Model\InlineObject244

try {
    $result = $apiInstance->createSubnet($inline_object244);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->createSubnet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object244** | [**\OpenAPI\Client\Model\InlineObject244**](../Model/InlineObject244.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200154**](../Model/InlineResponse200154.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetwork()`

```php
deleteNetwork($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network

Will delete a Network from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetwork($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetwork: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkDhcpRelay()`

```php
deleteNetworkDhcpRelay($id, $server_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network DHCP Relay

Will delete a Network DHCP Relay from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->deleteNetworkDhcpRelay($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkDhcpRelay: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkDhcpServer()`

```php
deleteNetworkDhcpServer($id, $server_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network DHCP Server

Will delete a Network DHCP Server from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->deleteNetworkDhcpServer($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkDhcpServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkDomain()`

```php
deleteNetworkDomain($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Domain

Will delete a Network Domain from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkDomain($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkDomain: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkFirewallRule()`

```php
deleteNetworkFirewallRule($id, $server_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Firewall Rule

Will delete a Network Firewall Rule from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->deleteNetworkFirewallRule($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkFirewallRuleGroup()`

```php
deleteNetworkFirewallRuleGroup($id, $server_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network firewall rule group

Will delete a network firewall rule group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->deleteNetworkFirewallRuleGroup($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkGroup()`

```php
deleteNetworkGroup($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Group

Will delete a Network Group from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkGroup($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkPool()`

```php
deleteNetworkPool($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Pool

Will delete a Network Pool from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkPool($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkPoolIp()`

```php
deleteNetworkPoolIp($network_pool_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a host record associated with an IP Address for a Specific Network Pool

Will delete a host record associated with an IP address for a specific network pool and free up the address

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$network_pool_id = 1; // int | Morpheus Network Pool ID of the Object being referenced
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkPoolIp($network_pool_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkPoolIp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_pool_id** | **int**| Morpheus Network Pool ID of the Object being referenced |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkPoolServer()`

```php
deleteNetworkPoolServer($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Pool Server

Will delete a Network Pool Server from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkPoolServer($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkPoolServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkProxy()`

```php
deleteNetworkProxy($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Proxy

Will delete a Network Proxy from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkProxy($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkProxy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkRouter()`

```php
deleteNetworkRouter($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Router

Will delete a Network Router from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNetworkRouter($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkRouter: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkRouterBgpNeighbor()`

```php
deleteNetworkRouterBgpNeighbor($id, $router_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Router BGP Neighbor

Will delete a BGP Neighbor from a network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->deleteNetworkRouterBgpNeighbor($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkRouterBgpNeighbor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkRouterFirewallRule()`

```php
deleteNetworkRouterFirewallRule($id, $router_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Router Firewall Rule

Will delete a firewall rule from a network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->deleteNetworkRouterFirewallRule($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkRouterFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkRouterFirewallRuleGroup()`

```php
deleteNetworkRouterFirewallRuleGroup($id, $router_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Router firewall rule group

Will delete a network router firewall rule group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->deleteNetworkRouterFirewallRuleGroup($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkRouterFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkRouterNat()`

```php
deleteNetworkRouterNat($id, $router_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Router NAT

Will delete a NAT from a network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->deleteNetworkRouterNat($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkRouterNat: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkRouterRoute()`

```php
deleteNetworkRouterRoute($id, $router_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Router Route

Will delete a Route from a network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->deleteNetworkRouterRoute($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkRouterRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkServerGroup()`

```php
deleteNetworkServerGroup($id, $server_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Server Group

Will delete a Network Server Group from the system and make it no longer usable. Note: Only available for NSX-T network integrations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->deleteNetworkServerGroup($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkServerGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteNetworkTransportZone()`

```php
deleteNetworkTransportZone($id, $server_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Transport Zone

Will delete a Network Transport Zone from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->deleteNetworkTransportZone($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteNetworkTransportZone: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteStaticRoute()`

```php
deleteStaticRoute($id, $route_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Network Static Route

Will delete a route from a network.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$route_id = 4; // float | The ID of the route

try {
    $result = $apiInstance->deleteStaticRoute($id, $route_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteStaticRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **route_id** | **float**| The ID of the route |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteSubnet()`

```php
deleteSubnet($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Subnet

Will delete a Subnet from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteSubnet($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->deleteSubnet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getAllNetworkFloatingIps()`

```php
getAllNetworkFloatingIps($phrase, $ip_address, $ip_status, $zone_id, $server_id): object
```

Get All Floating IPs

This endpoint retrieves all network floating IPs associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$ip_address = 10.32.23.188; // string | Filter by IP Address
$ip_status = 'ip_status_example'; // string | Filter by IP Status
$zone_id = 56; // int | Filter by Cloud ID
$server_id = 56; // int | Filter by Server ID

try {
    $result = $apiInstance->getAllNetworkFloatingIps($phrase, $ip_address, $ip_status, $zone_id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getAllNetworkFloatingIps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **ip_address** | **string**| Filter by IP Address | [optional]
 **ip_status** | **string**| Filter by IP Status | [optional]
 **zone_id** | **int**| Filter by Cloud ID | [optional]
 **server_id** | **int**| Filter by Server ID | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetwork()`

```php
getNetwork($id): \OpenAPI\Client\Model\InlineResponse20087
```

Get a Specific Network

This endpoint retrieves a specific Network.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetwork($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetwork: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20087**](../Model/InlineResponse20087.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkDhcpRelay()`

```php
getNetworkDhcpRelay($id, $server_id): \OpenAPI\Client\Model\InlineResponse200112
```

Get a Specific Network DHCP Relay

This endpoint retrieves a specific Network DHCP Relay.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->getNetworkDhcpRelay($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkDhcpRelay: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200112**](../Model/InlineResponse200112.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkDhcpRelays()`

```php
getNetworkDhcpRelays($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network DHCP Relays for Network Relay

This endpoint retrieves all Network DHCP Relays for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkDhcpRelays($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkDhcpRelays: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkDhcpServer()`

```php
getNetworkDhcpServer($id, $server_id): \OpenAPI\Client\Model\InlineResponse200113
```

Get a Specific Network DHCP Server

This endpoint retrieves a specific Network DHCP Server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->getNetworkDhcpServer($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkDhcpServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200113**](../Model/InlineResponse200113.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkDhcpServers()`

```php
getNetworkDhcpServers($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network DHCP Servers for Network Server

This endpoint retrieves all Network DHCP Servers for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkDhcpServers($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkDhcpServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkDomain()`

```php
getNetworkDomain($id): \OpenAPI\Client\Model\InlineResponse200109
```

Get a Specific Network Domain

This endpoint retrieves a specific Network Domain.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkDomain($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkDomain: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200109**](../Model/InlineResponse200109.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkDomains()`

```php
getNetworkDomains($name, $phrase): object
```

Get all Network Domains

This endpoint retrieves all Network Domains associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkDomains($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkDomains: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkEdgeCluster()`

```php
getNetworkEdgeCluster($id, $server_id): \OpenAPI\Client\Model\InlineResponse200114
```

Get a Specific Network Edge Cluster

This endpoint retrieves a specific network Edge Cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->getNetworkEdgeCluster($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkEdgeCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200114**](../Model/InlineResponse200114.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkEdgeClusters()`

```php
getNetworkEdgeClusters($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network Edge Clusters for Network Server

This endpoint retrieves all Network Edge Clusters for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkEdgeClusters($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkEdgeClusters: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkFirewallRule()`

```php
getNetworkFirewallRule($id, $server_id): \OpenAPI\Client\Model\InlineResponse200115
```

Get a Specific Network Firewall Rule

This endpoint retrieves a specific Network Firewall Rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->getNetworkFirewallRule($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200115**](../Model/InlineResponse200115.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkFirewallRuleGroup()`

```php
getNetworkFirewallRuleGroup($id, $server_id): \OpenAPI\Client\Model\InlineResponse200116
```

Get a Specific Network Firewall Rule Group

This endpoint retrieves a specific Network Firewall Rule Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->getNetworkFirewallRuleGroup($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200116**](../Model/InlineResponse200116.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkFirewallRuleGroups()`

```php
getNetworkFirewallRuleGroups($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network Firewall Rule Groups for Network Server

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkFirewallRuleGroups($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkFirewallRuleGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkFirewallRules()`

```php
getNetworkFirewallRules($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network Firewall Rules for Network Server

This endpoint retrieves all Network Firewall Rules for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkFirewallRules($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkFirewallRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkFloatingIp()`

```php
getNetworkFloatingIp($id): \OpenAPI\Client\Model\InlineResponse200108
```

Get a Specific Floating IP

This endpoint retrieves a specific Floating IP.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkFloatingIp($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkFloatingIp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200108**](../Model/InlineResponse200108.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkGroup()`

```php
getNetworkGroup($id): \OpenAPI\Client\Model\InlineResponse20091
```

Get a Specific Network Group

This endpoint retrieves a specific Network Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkGroup($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20091**](../Model/InlineResponse20091.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkGroups()`

```php
getNetworkGroups(): \OpenAPI\Client\Model\InlineResponse20090
```

Get all Network Groups

This endpoint retrieves all Network Groups associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getNetworkGroups();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20090**](../Model/InlineResponse20090.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkPool()`

```php
getNetworkPool($id): \OpenAPI\Client\Model\InlineResponse200106
```

Get a Specific Network Pool

This endpoint retrieves a specific Network Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkPool($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200106**](../Model/InlineResponse200106.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkPoolIp()`

```php
getNetworkPoolIp($network_pool_id, $id): object
```

Get a Specific IP Address for a Specific Network Pool

This endpoint retrieves a specific IP address for a specific Network Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$network_pool_id = 1; // int | Morpheus Network Pool ID of the Object being referenced
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkPoolIp($network_pool_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkPoolIp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_pool_id** | **int**| Morpheus Network Pool ID of the Object being referenced |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkPoolIps()`

```php
getNetworkPoolIps($id): object
```

Get all IP Addresses for a Specific Network Pool

This endpoint retrieves a list of IP addresses for a specific Network Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkPoolIps($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkPoolIps: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkPoolServer()`

```php
getNetworkPoolServer($id): \OpenAPI\Client\Model\InlineResponse200120
```

Get a Specific Network Pool Server

This endpoint retrieves a specific Network Pool Server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkPoolServer($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkPoolServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200120**](../Model/InlineResponse200120.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkPoolServerType()`

```php
getNetworkPoolServerType($id): \OpenAPI\Client\Model\InlineResponse200121
```

Retrieves a Specific Network Pool Server Type

Retrieves a specific Network Pool Server Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkPoolServerType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkPoolServerType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200121**](../Model/InlineResponse200121.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkPools()`

```php
getNetworkPools($name, $phrase): object
```

Get all Network Pools

This endpoint retrieves all Network Pools associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkPools($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkPools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkProxies()`

```php
getNetworkProxies($name, $phrase): object
```

Get all Network Proxies

This endpoint retrieves all Network Proxies associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkProxies($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkProxies: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkProxy()`

```php
getNetworkProxy($id): \OpenAPI\Client\Model\InlineResponse200111
```

Get a Specific Network Proxy

This endpoint retrieves a specific Network Proxy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkProxy($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkProxy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200111**](../Model/InlineResponse200111.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouter()`

```php
getNetworkRouter($id): \OpenAPI\Client\Model\InlineResponse20095
```

Get a Specific Network Router

This endpoint retrieves a specific Network Router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkRouter($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouter: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20095**](../Model/InlineResponse20095.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterBgpNeighbor()`

```php
getNetworkRouterBgpNeighbor($id, $router_id): \OpenAPI\Client\Model\InlineResponse20097
```

Get a Network Router BGP Neighbor

This endpoint retrieves a network router BGP Neighbor for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRouterBgpNeighbor($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterBgpNeighbor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20097**](../Model/InlineResponse20097.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterFirewallRule()`

```php
getNetworkRouterFirewallRule($id, $router_id): \OpenAPI\Client\Model\InlineResponse20099
```

Get a Firewall Rule for Network Router

This endpoint retrieves a firewall rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRouterFirewallRule($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20099**](../Model/InlineResponse20099.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterFirewallRuleGroup()`

```php
getNetworkRouterFirewallRuleGroup($id, $router_id): \OpenAPI\Client\Model\InlineResponse200101
```

Get a Firewall Rule Group for Network Router

This endpoint retrieves a firewall rule group for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRouterFirewallRuleGroup($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200101**](../Model/InlineResponse200101.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterFirewallRuleGroups()`

```php
getNetworkRouterFirewallRuleGroups($router_id): \OpenAPI\Client\Model\InlineResponse200100
```

Get all Network Firewall Rule Groups for Network Router

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRouterFirewallRuleGroups($router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterFirewallRuleGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200100**](../Model/InlineResponse200100.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterNat()`

```php
getNetworkRouterNat($id, $router_id): \OpenAPI\Client\Model\InlineResponse200103
```

Get a Network Router NAT

This endpoint retrieves a network router NAT for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRouterNat($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterNat: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200103**](../Model/InlineResponse200103.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterRoute()`

```php
getNetworkRouterRoute($id, $router_id): \OpenAPI\Client\Model\InlineResponse200105
```

Get a Route for Network Router

This endpoint retrieves a Route.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRouterRoute($id, $router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200105**](../Model/InlineResponse200105.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouterType()`

```php
getNetworkRouterType($id): \OpenAPI\Client\Model\InlineResponse20093
```

Get a Specific Network Router Type

This endpoint retrieves a specific network router type. Use this API to retrieve list of available option types for a specific network router type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkRouterType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouterType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20093**](../Model/InlineResponse20093.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRouters()`

```php
getNetworkRouters(): \OpenAPI\Client\Model\InlineResponse20094
```

Get all Network Routers

This endpoint retrieves all Network Routers

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getNetworkRouters();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRouters: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20094**](../Model/InlineResponse20094.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRoutersBgpNeighbors()`

```php
getNetworkRoutersBgpNeighbors($router_id): \OpenAPI\Client\Model\InlineResponse20096
```

Get all BGP Neighbors for Network Router

This endpoint retrieves all BGP Neighbors for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRoutersBgpNeighbors($router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRoutersBgpNeighbors: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20096**](../Model/InlineResponse20096.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRoutersFirewallRules()`

```php
getNetworkRoutersFirewallRules($router_id): \OpenAPI\Client\Model\InlineResponse20098
```

Get all Firewall Rules for Network Router

This endpoint retrieves all firewall rules for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRoutersFirewallRules($router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRoutersFirewallRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20098**](../Model/InlineResponse20098.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRoutersNats()`

```php
getNetworkRoutersNats($router_id): \OpenAPI\Client\Model\InlineResponse200102
```

Get all Network Router NATs for Network Router

This endpoint retrieves all NATs for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRoutersNats($router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRoutersNats: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200102**](../Model/InlineResponse200102.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkRoutersRoutes()`

```php
getNetworkRoutersRoutes($router_id): \OpenAPI\Client\Model\InlineResponse200104
```

Get all Routes for Network Router

This endpoint retrieves all Routes for specified network router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID

try {
    $result = $apiInstance->getNetworkRoutersRoutes($router_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkRoutersRoutes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200104**](../Model/InlineResponse200104.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkServerGroup()`

```php
getNetworkServerGroup($server_id, $id): \OpenAPI\Client\Model\InlineResponse200117
```

Get a specific Network Server Group

This endpoint retrieves a specific Network Server Group for a Network Server. Note: Only available for NSX-T network integrations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkServerGroup($server_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkServerGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200117**](../Model/InlineResponse200117.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkSubnets()`

```php
getNetworkSubnets($id): object
```

Get Subnets for a Network

This endpoint retrieves all Subnets under a specific network.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkSubnets($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkSubnets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkTransportZone()`

```php
getNetworkTransportZone($id, $server_id): \OpenAPI\Client\Model\InlineResponse200118
```

Get a Specific Network Transport Zone

This endpoint retrieves a specific Network Transport Zone.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID

try {
    $result = $apiInstance->getNetworkTransportZone($id, $server_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkTransportZone: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200118**](../Model/InlineResponse200118.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkTransportZones()`

```php
getNetworkTransportZones($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network Transport Zones for Network Server

This endpoint retrieves all Network Transport Zones for a specified Network Service.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->getNetworkTransportZones($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkTransportZones: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNetworkType()`

```php
getNetworkType($id): \OpenAPI\Client\Model\InlineResponse20086
```

Get a Specific Network Type

This endpoint retrieves a specific Network Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNetworkType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getNetworkType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20086**](../Model/InlineResponse20086.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStaticRoute()`

```php
getStaticRoute($id, $route_id): \OpenAPI\Client\Model\InlineResponse20089
```

Get a Static Route for a Network

This endpoint retrieves a network static route for specified network.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$route_id = 4; // float | The ID of the route

try {
    $result = $apiInstance->getStaticRoute($id, $route_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getStaticRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **route_id** | **float**| The ID of the route |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20089**](../Model/InlineResponse20089.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getStaticRoutes()`

```php
getStaticRoutes($id): \OpenAPI\Client\Model\InlineResponse20088
```

Get all Static Routes for a Network

This endpoint retrieves all routes for specified network.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getStaticRoutes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getStaticRoutes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20088**](../Model/InlineResponse20088.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getSubnet()`

```php
getSubnet($id): \OpenAPI\Client\Model\InlineResponse200154
```

Get a Specific Subnet

This endpoint retrieves a specific Subnet.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getSubnet($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->getSubnet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200154**](../Model/InlineResponse200154.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworkPoolServerTypes()`

```php
listNetworkPoolServerTypes($max, $offset, $sort, $direction, $phrase, $name, $code): object
```

Get All Network Pool Server Types

This endpoint retrieves all Network Pool Server Types

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code

try {
    $result = $apiInstance->listNetworkPoolServerTypes($max, $offset, $sort, $direction, $phrase, $name, $code);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworkPoolServerTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworkPoolServers()`

```php
listNetworkPoolServers($max, $offset, $sort, $direction, $phrase, $name): object
```

Get All Network Pool Servers

This endpoint retrieves all Network Pool Servers associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listNetworkPoolServers($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworkPoolServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworkRouterTypes()`

```php
listNetworkRouterTypes(): \OpenAPI\Client\Model\InlineResponse20092
```

Get all Network Router Types

Get all Network Router Types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listNetworkRouterTypes();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworkRouterTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20092**](../Model/InlineResponse20092.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworkServerGroups()`

```php
listNetworkServerGroups($server_id, $max, $offset, $sort, $direction, $phrase): object
```

Get all Network Server Groups for Network Server

This endpoint retrieves all Network Server Groups for a specified Network Service. Note: Only available for NSX-T network integrations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$server_id = 4; // float | Server ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listNetworkServerGroups($server_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworkServerGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworkServices()`

```php
listNetworkServices($name, $phrase): \OpenAPI\Client\Model\InlineResponse200119
```

Get All Network Services

This endpoint retrieves all Network Services associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listNetworkServices($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworkServices: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200119**](../Model/InlineResponse200119.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworkTypes()`

```php
listNetworkTypes($name, $code, $phrase): object
```

Network Types

Provides API for viewing Network Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listNetworkTypes($name, $code, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworkTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listNetworks()`

```php
listNetworks($name, $phrase, $labels, $all_labels): object
```

Get All Networks

This endpoint retrieves all Networks associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listNetworks($name, $phrase, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listNetworks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listSubnetTypes()`

```php
listSubnetTypes($name, $phrase): object
```

Get All Subnet Types

Provides API for viewing Network Subnet Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listSubnetTypes($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listSubnetTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listSubnets()`

```php
listSubnets($name, $phrase, $labels, $all_labels): object
```

Get All Subnets

This endpoint retrieves all Subnets associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listSubnets($name, $phrase, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->listSubnets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `refreshNetworkServer()`

```php
refreshNetworkServer($id): \OpenAPI\Client\Model\SuccessId
```

Refresh a Network Server/Integration

Refreshes a network server/integration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->refreshNetworkServer($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->refreshNetworkServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `releaseNetworkFloatingIp()`

```php
releaseNetworkFloatingIp($id): \OpenAPI\Client\Model\Model200Success
```

Release a Floating IP

Release a floating IP detaching it from the associated node/VM.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->releaseNetworkFloatingIp($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->releaseNetworkFloatingIp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetwork()`

```php
updateNetwork($id, $inline_object143): object
```

Update a Network

This endpoint allows updating a Network. Configuration options vary by Network Types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object143 = new \OpenAPI\Client\Model\InlineObject143(); // \OpenAPI\Client\Model\InlineObject143

try {
    $result = $apiInstance->updateNetwork($id, $inline_object143);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetwork: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object143** | [**\OpenAPI\Client\Model\InlineObject143**](../Model/InlineObject143.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkDhcpRelay()`

```php
updateNetworkDhcpRelay($id, $server_id, $inline_object168): \OpenAPI\Client\Model\Model200Success
```

Update a Network DHCP Relay

Use this command to update an existing Network DHCP Relay.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object168 = new \OpenAPI\Client\Model\InlineObject168(); // \OpenAPI\Client\Model\InlineObject168

try {
    $result = $apiInstance->updateNetworkDhcpRelay($id, $server_id, $inline_object168);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkDhcpRelay: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object168** | [**\OpenAPI\Client\Model\InlineObject168**](../Model/InlineObject168.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkDhcpServer()`

```php
updateNetworkDhcpServer($id, $server_id, $inline_object170): \OpenAPI\Client\Model\Model200Success
```

Update a Network DHCP Server

Use this command to update an existing Network DHCP Server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object170 = new \OpenAPI\Client\Model\InlineObject170(); // \OpenAPI\Client\Model\InlineObject170

try {
    $result = $apiInstance->updateNetworkDhcpServer($id, $server_id, $inline_object170);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkDhcpServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object170** | [**\OpenAPI\Client\Model\InlineObject170**](../Model/InlineObject170.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkDomain()`

```php
updateNetworkDomain($id, $inline_object164): \OpenAPI\Client\Model\InlineResponse200109
```

Update a Network Domain

Update a Network Domain.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object164 = new \OpenAPI\Client\Model\InlineObject164(); // \OpenAPI\Client\Model\InlineObject164

try {
    $result = $apiInstance->updateNetworkDomain($id, $inline_object164);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkDomain: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object164** | [**\OpenAPI\Client\Model\InlineObject164**](../Model/InlineObject164.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200109**](../Model/InlineResponse200109.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkEdgeCluster()`

```php
updateNetworkEdgeCluster($id, $server_id, $inline_object171): \OpenAPI\Client\Model\Model200Success
```

Update a Network Edge Cluster

Use this command to update an existing network Edge Cluster.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object171 = new \OpenAPI\Client\Model\InlineObject171(); // \OpenAPI\Client\Model\InlineObject171

try {
    $result = $apiInstance->updateNetworkEdgeCluster($id, $server_id, $inline_object171);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkEdgeCluster: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object171** | [**\OpenAPI\Client\Model\InlineObject171**](../Model/InlineObject171.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkFirewallRule()`

```php
updateNetworkFirewallRule($id, $server_id, $inline_object173): \OpenAPI\Client\Model\Model200Success
```

Update a Network Firewall Rule

Use this command to update an existing network firewall Rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object173 = new \OpenAPI\Client\Model\InlineObject173(); // \OpenAPI\Client\Model\InlineObject173

try {
    $result = $apiInstance->updateNetworkFirewallRule($id, $server_id, $inline_object173);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object173** | [**\OpenAPI\Client\Model\InlineObject173**](../Model/InlineObject173.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkFirewallRuleGroup()`

```php
updateNetworkFirewallRuleGroup($id, $server_id, $inline_object175): \OpenAPI\Client\Model\Model200Success
```

Update a Network Firewall Rule Group

Use this command to update an existing Network Firewall Rule Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object175 = new \OpenAPI\Client\Model\InlineObject175(); // \OpenAPI\Client\Model\InlineObject175

try {
    $result = $apiInstance->updateNetworkFirewallRuleGroup($id, $server_id, $inline_object175);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object175** | [**\OpenAPI\Client\Model\InlineObject175**](../Model/InlineObject175.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkGroup()`

```php
updateNetworkGroup($id, $inline_object147): \OpenAPI\Client\Model\Model200Success
```

Update a Network Group

Use this command to update an existing network Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object147 = new \OpenAPI\Client\Model\InlineObject147(); // \OpenAPI\Client\Model\InlineObject147

try {
    $result = $apiInstance->updateNetworkGroup($id, $inline_object147);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object147** | [**\OpenAPI\Client\Model\InlineObject147**](../Model/InlineObject147.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkPool()`

```php
updateNetworkPool($id, $inline_object161): \OpenAPI\Client\Model\InlineResponse200106
```

Update a Network Pool

Update a Network Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object161 = new \OpenAPI\Client\Model\InlineObject161(); // \OpenAPI\Client\Model\InlineObject161

try {
    $result = $apiInstance->updateNetworkPool($id, $inline_object161);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object161** | [**\OpenAPI\Client\Model\InlineObject161**](../Model/InlineObject161.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200106**](../Model/InlineResponse200106.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkPoolServer()`

```php
updateNetworkPoolServer($id, $inline_object181): Model200Success
```

Update a Network Pool Server

This endpoint allows updating a Network Pool Server. Configuration options vary by type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object181 = new \OpenAPI\Client\Model\InlineObject181(); // \OpenAPI\Client\Model\InlineObject181

try {
    $result = $apiInstance->updateNetworkPoolServer($id, $inline_object181);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkPoolServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object181** | [**\OpenAPI\Client\Model\InlineObject181**](../Model/InlineObject181.md)|  | [optional]

### Return type

[**Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkProxy()`

```php
updateNetworkProxy($id, $inline_object166): \OpenAPI\Client\Model\Model200Success
```

Update a Network Proxy

Use this command to update an existing network Proxy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object166 = new \OpenAPI\Client\Model\InlineObject166(); // \OpenAPI\Client\Model\InlineObject166

try {
    $result = $apiInstance->updateNetworkProxy($id, $inline_object166);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkProxy: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object166** | [**\OpenAPI\Client\Model\InlineObject166**](../Model/InlineObject166.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkRouter()`

```php
updateNetworkRouter($id, $inline_object149): \OpenAPI\Client\Model\Model200Success
```

Update a Network Router

Use this command to update an existing network Router.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object149 = new \OpenAPI\Client\Model\InlineObject149(); // \OpenAPI\Client\Model\InlineObject149

try {
    $result = $apiInstance->updateNetworkRouter($id, $inline_object149);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkRouter: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object149** | [**\OpenAPI\Client\Model\InlineObject149**](../Model/InlineObject149.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkRouterBgpNeighbor()`

```php
updateNetworkRouterBgpNeighbor($id, $router_id, $inline_object151): \OpenAPI\Client\Model\Model200Success
```

Update Network Router BGP Neighbor

Use this command to update an existing Network Router BGP Neighbor.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID
$inline_object151 = new \OpenAPI\Client\Model\InlineObject151(); // \OpenAPI\Client\Model\InlineObject151

try {
    $result = $apiInstance->updateNetworkRouterBgpNeighbor($id, $router_id, $inline_object151);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkRouterBgpNeighbor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **inline_object151** | [**\OpenAPI\Client\Model\InlineObject151**](../Model/InlineObject151.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkRouterFirewallRule()`

```php
updateNetworkRouterFirewallRule($id, $router_id, $inline_object153): \OpenAPI\Client\Model\Model200Success
```

Update a Network Router Firewall Rule

Use this command to update an existing network router firewall rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID
$inline_object153 = new \OpenAPI\Client\Model\InlineObject153(); // \OpenAPI\Client\Model\InlineObject153

try {
    $result = $apiInstance->updateNetworkRouterFirewallRule($id, $router_id, $inline_object153);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkRouterFirewallRule: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **inline_object153** | [**\OpenAPI\Client\Model\InlineObject153**](../Model/InlineObject153.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkRouterFirewallRuleGroup()`

```php
updateNetworkRouterFirewallRuleGroup($id, $router_id, $inline_object155): \OpenAPI\Client\Model\Model200Success
```

Update a Network Router Firewall Rule Group

Use this command to update an existing Network Router Firewall Rule Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID
$inline_object155 = new \OpenAPI\Client\Model\InlineObject155(); // \OpenAPI\Client\Model\InlineObject155

try {
    $result = $apiInstance->updateNetworkRouterFirewallRuleGroup($id, $router_id, $inline_object155);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkRouterFirewallRuleGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **inline_object155** | [**\OpenAPI\Client\Model\InlineObject155**](../Model/InlineObject155.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkRouterNat()`

```php
updateNetworkRouterNat($id, $router_id, $inline_object157): \OpenAPI\Client\Model\Model200Success
```

Update Network Router NAT

Use this command to update an existing Network Router NAT.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$router_id = 4; // float | Router ID
$inline_object157 = new \OpenAPI\Client\Model\InlineObject157(); // \OpenAPI\Client\Model\InlineObject157

try {
    $result = $apiInstance->updateNetworkRouterNat($id, $router_id, $inline_object157);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkRouterNat: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **inline_object157** | [**\OpenAPI\Client\Model\InlineObject157**](../Model/InlineObject157.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkRouterPermissions()`

```php
updateNetworkRouterPermissions($router_id, $inline_object158): \OpenAPI\Client\Model\SuccessId
```

Update Network Router Permissions

Update Network Router Permissions

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$router_id = 4; // float | Router ID
$inline_object158 = new \OpenAPI\Client\Model\InlineObject158(); // \OpenAPI\Client\Model\InlineObject158

try {
    $result = $apiInstance->updateNetworkRouterPermissions($router_id, $inline_object158);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkRouterPermissions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **inline_object158** | [**\OpenAPI\Client\Model\InlineObject158**](../Model/InlineObject158.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkServerGroup()`

```php
updateNetworkServerGroup($id, $server_id, $inline_object177): \OpenAPI\Client\Model\Model200Success
```

Update a Network Server Group

Use this command to update an existing network server group. Note: Only available for NSX-T network integrations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object177 = new \OpenAPI\Client\Model\InlineObject177(); // \OpenAPI\Client\Model\InlineObject177

try {
    $result = $apiInstance->updateNetworkServerGroup($id, $server_id, $inline_object177);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkServerGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object177** | [**\OpenAPI\Client\Model\InlineObject177**](../Model/InlineObject177.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateNetworkTransportZone()`

```php
updateNetworkTransportZone($id, $server_id, $inline_object179): \OpenAPI\Client\Model\Model200Success
```

Update a Network Transport Zone

Use this command to update an existing network Transport Zone.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$server_id = 4; // float | Server ID
$inline_object179 = new \OpenAPI\Client\Model\InlineObject179(); // \OpenAPI\Client\Model\InlineObject179

try {
    $result = $apiInstance->updateNetworkTransportZone($id, $server_id, $inline_object179);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateNetworkTransportZone: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **inline_object179** | [**\OpenAPI\Client\Model\InlineObject179**](../Model/InlineObject179.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateStaticRoute()`

```php
updateStaticRoute($id, $route_id, $inline_object145): \OpenAPI\Client\Model\SuccessId
```

Update a Network Static Route

Use this command to update a route.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$route_id = 4; // float | The ID of the route
$inline_object145 = new \OpenAPI\Client\Model\InlineObject145(); // \OpenAPI\Client\Model\InlineObject145

try {
    $result = $apiInstance->updateStaticRoute($id, $route_id, $inline_object145);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateStaticRoute: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **route_id** | **float**| The ID of the route |
 **inline_object145** | [**\OpenAPI\Client\Model\InlineObject145**](../Model/InlineObject145.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateSubnet()`

```php
updateSubnet($id, $inline_object245): \OpenAPI\Client\Model\InlineResponse200154
```

Update a Subnet

This endpoint allows updating a Subnet. Only certain types of clouds support this action. Configuration options vary for each Subnet Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\NetworksApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object245 = new \OpenAPI\Client\Model\InlineObject245(); // \OpenAPI\Client\Model\InlineObject245

try {
    $result = $apiInstance->updateSubnet($id, $inline_object245);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NetworksApi->updateSubnet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object245** | [**\OpenAPI\Client\Model\InlineObject245**](../Model/InlineObject245.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200154**](../Model/InlineResponse200154.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
