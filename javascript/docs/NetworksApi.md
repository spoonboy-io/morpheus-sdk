# MorpheusApi.NetworksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createNetworkDhcpRelay**](NetworksApi.md#createNetworkDhcpRelay) | **POST** /api/networks/servers/{serverId}/dhcp-relays | Create a Network DHCP Relay
[**createNetworkDhcpServer**](NetworksApi.md#createNetworkDhcpServer) | **POST** /api/networks/servers/{serverId}/dhcp-servers | Create a Network DHCP Server
[**createNetworkDomain**](NetworksApi.md#createNetworkDomain) | **POST** /api/networks/domains | Create a Network Domain
[**createNetworkFirewallRule**](NetworksApi.md#createNetworkFirewallRule) | **POST** /api/networks/servers/{serverId}/firewall-rules | Create a Network Firewall Rule
[**createNetworkFirewallRuleGroup**](NetworksApi.md#createNetworkFirewallRuleGroup) | **POST** /api/networks/servers/{serverId}/firewall-rule-groups | Create a Network Firewall Rule Group
[**createNetworkGroup**](NetworksApi.md#createNetworkGroup) | **POST** /api/networks/groups | Create a Network Group
[**createNetworkPool**](NetworksApi.md#createNetworkPool) | **POST** /api/networks/pools | Create a Network Pool
[**createNetworkPoolIp**](NetworksApi.md#createNetworkPoolIp) | **POST** /api/networks/pools/{id}/ips | Create a Network Pool IP Address
[**createNetworkPoolServer**](NetworksApi.md#createNetworkPoolServer) | **POST** /api/networks/pool-servers | Create a Network Pool Server
[**createNetworkProxy**](NetworksApi.md#createNetworkProxy) | **POST** /api/networks/proxies | Create a Network Proxy
[**createNetworkRouter**](NetworksApi.md#createNetworkRouter) | **POST** /api/networks/routers | Create a Network Router
[**createNetworkRouterBgpNeighbor**](NetworksApi.md#createNetworkRouterBgpNeighbor) | **POST** /api/networks/routers/{routerId}/bgp-neighbors | Create a Network Router BGP Neighbor
[**createNetworkRouterFirewallRule**](NetworksApi.md#createNetworkRouterFirewallRule) | **POST** /api/networks/routers/{routerId}/firewall-rules | Create a Network Router Firewall Rule
[**createNetworkRouterFirewallRuleGroup**](NetworksApi.md#createNetworkRouterFirewallRuleGroup) | **POST** /api/networks/routers/{routerId}/firewall-rule-groups | Create a Network Router Firewall Rule Group
[**createNetworkRouterNat**](NetworksApi.md#createNetworkRouterNat) | **POST** /api/networks/routers/{routerId}/nats | Create a Network Router NAT
[**createNetworkRouterRoute**](NetworksApi.md#createNetworkRouterRoute) | **POST** /api/networks/routers/{routerId}/routes | Create a Network Router Route
[**createNetworkServerGroup**](NetworksApi.md#createNetworkServerGroup) | **POST** /api/networks/servers/{serverId}/groups | Create a Network Server Group
[**createNetworkTransportZone**](NetworksApi.md#createNetworkTransportZone) | **POST** /api/networks/servers/{serverId}/scopes | Create a Network Transport Zone
[**createNetworks**](NetworksApi.md#createNetworks) | **POST** /api/networks | Create a Network
[**createStaticRoute**](NetworksApi.md#createStaticRoute) | **PUT** /api/networks/{id}/routes | Create a Network Static Route
[**createSubnet**](NetworksApi.md#createSubnet) | **POST** /api/subnets | Create a Subnet
[**deleteNetwork**](NetworksApi.md#deleteNetwork) | **DELETE** /api/networks/{id} | Delete a Network
[**deleteNetworkDhcpRelay**](NetworksApi.md#deleteNetworkDhcpRelay) | **DELETE** /api/networks/servers/{serverId}/dhcp-relays/{id} | Delete a Network DHCP Relay
[**deleteNetworkDhcpServer**](NetworksApi.md#deleteNetworkDhcpServer) | **DELETE** /api/networks/servers/{serverId}/dhcp-servers/{id} | Delete a Network DHCP Server
[**deleteNetworkDomain**](NetworksApi.md#deleteNetworkDomain) | **DELETE** /api/networks/domains/{id} | Delete a Network Domain
[**deleteNetworkFirewallRule**](NetworksApi.md#deleteNetworkFirewallRule) | **DELETE** /api/networks/servers/{serverId}/firewall-rules/{id} | Delete a Network Firewall Rule
[**deleteNetworkFirewallRuleGroup**](NetworksApi.md#deleteNetworkFirewallRuleGroup) | **DELETE** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Delete a Network firewall rule group
[**deleteNetworkGroup**](NetworksApi.md#deleteNetworkGroup) | **DELETE** /api/networks/groups/{id} | Delete a Network Group
[**deleteNetworkPool**](NetworksApi.md#deleteNetworkPool) | **DELETE** /api/networks/pools/{id} | Delete a Network Pool
[**deleteNetworkPoolIp**](NetworksApi.md#deleteNetworkPoolIp) | **DELETE** /api/networks/pools/{networkPoolId}/ips/{id} | Delete a host record associated with an IP Address for a Specific Network Pool
[**deleteNetworkPoolServer**](NetworksApi.md#deleteNetworkPoolServer) | **DELETE** /api/networks/pool-servers/{id} | Delete a Network Pool Server
[**deleteNetworkProxy**](NetworksApi.md#deleteNetworkProxy) | **DELETE** /api/networks/proxies/{id} | Delete a Network Proxy
[**deleteNetworkRouter**](NetworksApi.md#deleteNetworkRouter) | **DELETE** /api/networks/routers/{id} | Delete a Network Router
[**deleteNetworkRouterBgpNeighbor**](NetworksApi.md#deleteNetworkRouterBgpNeighbor) | **DELETE** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Delete a Network Router BGP Neighbor
[**deleteNetworkRouterFirewallRule**](NetworksApi.md#deleteNetworkRouterFirewallRule) | **DELETE** /api/networks/routers/{routerId}/firewall-rules/{id} | Delete a Network Router Firewall Rule
[**deleteNetworkRouterFirewallRuleGroup**](NetworksApi.md#deleteNetworkRouterFirewallRuleGroup) | **DELETE** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Delete a Network Router firewall rule group
[**deleteNetworkRouterNat**](NetworksApi.md#deleteNetworkRouterNat) | **DELETE** /api/networks/routers/{routerId}/nats/{id} | Delete a Network Router NAT
[**deleteNetworkRouterRoute**](NetworksApi.md#deleteNetworkRouterRoute) | **DELETE** /api/networks/routers/{routerId}/routes/{id} | Delete a Network Router Route
[**deleteNetworkServerGroup**](NetworksApi.md#deleteNetworkServerGroup) | **DELETE** /api/networks/servers/{serverId}/groups/{id} | Delete a Network Server Group
[**deleteNetworkTransportZone**](NetworksApi.md#deleteNetworkTransportZone) | **DELETE** /api/networks/servers/{serverId}/scopes/{id} | Delete a Network Transport Zone
[**deleteStaticRoute**](NetworksApi.md#deleteStaticRoute) | **DELETE** /api/networks/{id}/routes/{routeId} | Delete a Network Static Route
[**deleteSubnet**](NetworksApi.md#deleteSubnet) | **DELETE** /api/subnets/{id} | Delete a Subnet
[**getAllNetworkFloatingIps**](NetworksApi.md#getAllNetworkFloatingIps) | **GET** /api/networks/floating-ips | Get All Floating IPs
[**getNetwork**](NetworksApi.md#getNetwork) | **GET** /api/networks/{id} | Get a Specific Network
[**getNetworkDhcpRelay**](NetworksApi.md#getNetworkDhcpRelay) | **GET** /api/networks/servers/{serverId}/dhcp-relays/{id} | Get a Specific Network DHCP Relay
[**getNetworkDhcpRelays**](NetworksApi.md#getNetworkDhcpRelays) | **GET** /api/networks/servers/{serverId}/dhcp-relays | Get all Network DHCP Relays for Network Relay
[**getNetworkDhcpServer**](NetworksApi.md#getNetworkDhcpServer) | **GET** /api/networks/servers/{serverId}/dhcp-servers/{id} | Get a Specific Network DHCP Server
[**getNetworkDhcpServers**](NetworksApi.md#getNetworkDhcpServers) | **GET** /api/networks/servers/{serverId}/dhcp-servers | Get all Network DHCP Servers for Network Server
[**getNetworkDomain**](NetworksApi.md#getNetworkDomain) | **GET** /api/networks/domains/{id} | Get a Specific Network Domain
[**getNetworkDomains**](NetworksApi.md#getNetworkDomains) | **GET** /api/networks/domains | Get all Network Domains
[**getNetworkEdgeCluster**](NetworksApi.md#getNetworkEdgeCluster) | **GET** /api/networks/servers/{serverId}/edge-clusters/{id} | Get a Specific Network Edge Cluster
[**getNetworkEdgeClusters**](NetworksApi.md#getNetworkEdgeClusters) | **GET** /api/networks/servers/{serverId}/edge-clusters | Get all Network Edge Clusters for Network Server
[**getNetworkFirewallRule**](NetworksApi.md#getNetworkFirewallRule) | **GET** /api/networks/servers/{serverId}/firewall-rules/{id} | Get a Specific Network Firewall Rule
[**getNetworkFirewallRuleGroup**](NetworksApi.md#getNetworkFirewallRuleGroup) | **GET** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Get a Specific Network Firewall Rule Group
[**getNetworkFirewallRuleGroups**](NetworksApi.md#getNetworkFirewallRuleGroups) | **GET** /api/networks/servers/{serverId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Server
[**getNetworkFirewallRules**](NetworksApi.md#getNetworkFirewallRules) | **GET** /api/networks/servers/{serverId}/firewall-rules | Get all Network Firewall Rules for Network Server
[**getNetworkFloatingIp**](NetworksApi.md#getNetworkFloatingIp) | **GET** /api/networks/floating-ips/{id} | Get a Specific Floating IP
[**getNetworkGroup**](NetworksApi.md#getNetworkGroup) | **GET** /api/networks/groups/{id} | Get a Specific Network Group
[**getNetworkGroups**](NetworksApi.md#getNetworkGroups) | **GET** /api/networks/groups | Get all Network Groups
[**getNetworkPool**](NetworksApi.md#getNetworkPool) | **GET** /api/networks/pools/{id} | Get a Specific Network Pool
[**getNetworkPoolIp**](NetworksApi.md#getNetworkPoolIp) | **GET** /api/networks/pools/{networkPoolId}/ips/{id} | Get a Specific IP Address for a Specific Network Pool
[**getNetworkPoolIps**](NetworksApi.md#getNetworkPoolIps) | **GET** /api/networks/pools/{id}/ips | Get all IP Addresses for a Specific Network Pool
[**getNetworkPoolServer**](NetworksApi.md#getNetworkPoolServer) | **GET** /api/networks/pool-servers/{id} | Get a Specific Network Pool Server
[**getNetworkPoolServerType**](NetworksApi.md#getNetworkPoolServerType) | **GET** /api/networks/pool-server-types/{id} | Retrieves a Specific Network Pool Server Type
[**getNetworkPools**](NetworksApi.md#getNetworkPools) | **GET** /api/networks/pools | Get all Network Pools
[**getNetworkProxies**](NetworksApi.md#getNetworkProxies) | **GET** /api/networks/proxies | Get all Network Proxies
[**getNetworkProxy**](NetworksApi.md#getNetworkProxy) | **GET** /api/networks/proxies/{id} | Get a Specific Network Proxy
[**getNetworkRouter**](NetworksApi.md#getNetworkRouter) | **GET** /api/networks/routers/{id} | Get a Specific Network Router
[**getNetworkRouterBgpNeighbor**](NetworksApi.md#getNetworkRouterBgpNeighbor) | **GET** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Get a Network Router BGP Neighbor
[**getNetworkRouterFirewallRule**](NetworksApi.md#getNetworkRouterFirewallRule) | **GET** /api/networks/routers/{routerId}/firewall-rules/{id} | Get a Firewall Rule for Network Router
[**getNetworkRouterFirewallRuleGroup**](NetworksApi.md#getNetworkRouterFirewallRuleGroup) | **GET** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Get a Firewall Rule Group for Network Router
[**getNetworkRouterFirewallRuleGroups**](NetworksApi.md#getNetworkRouterFirewallRuleGroups) | **GET** /api/networks/routers/{routerId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Router
[**getNetworkRouterNat**](NetworksApi.md#getNetworkRouterNat) | **GET** /api/networks/routers/{routerId}/nats/{id} | Get a Network Router NAT
[**getNetworkRouterRoute**](NetworksApi.md#getNetworkRouterRoute) | **GET** /api/networks/routers/{routerId}/routes/{id} | Get a Route for Network Router
[**getNetworkRouterType**](NetworksApi.md#getNetworkRouterType) | **GET** /api/network-router-types/{id} | Get a Specific Network Router Type
[**getNetworkRouters**](NetworksApi.md#getNetworkRouters) | **GET** /api/networks/routers | Get all Network Routers
[**getNetworkRoutersBgpNeighbors**](NetworksApi.md#getNetworkRoutersBgpNeighbors) | **GET** /api/networks/routers/{routerId}/bgp-neighbors | Get all BGP Neighbors for Network Router
[**getNetworkRoutersFirewallRules**](NetworksApi.md#getNetworkRoutersFirewallRules) | **GET** /api/networks/routers/{routerId}/firewall-rules | Get all Firewall Rules for Network Router
[**getNetworkRoutersNats**](NetworksApi.md#getNetworkRoutersNats) | **GET** /api/networks/routers/{routerId}/nats | Get all Network Router NATs for Network Router
[**getNetworkRoutersRoutes**](NetworksApi.md#getNetworkRoutersRoutes) | **GET** /api/networks/routers/{routerId}/routes | Get all Routes for Network Router
[**getNetworkServerGroup**](NetworksApi.md#getNetworkServerGroup) | **GET** /api/networks/servers/{serverId}/groups/{id} | Get a specific Network Server Group
[**getNetworkSubnets**](NetworksApi.md#getNetworkSubnets) | **GET** /api/networks/{id}/subnets | Get Subnets for a Network
[**getNetworkTransportZone**](NetworksApi.md#getNetworkTransportZone) | **GET** /api/networks/servers/{serverId}/scopes/{id} | Get a Specific Network Transport Zone
[**getNetworkTransportZones**](NetworksApi.md#getNetworkTransportZones) | **GET** /api/networks/servers/{serverId}/scopes | Get all Network Transport Zones for Network Server
[**getNetworkType**](NetworksApi.md#getNetworkType) | **GET** /api/network-types/{id} | Get a Specific Network Type
[**getStaticRoute**](NetworksApi.md#getStaticRoute) | **GET** /api/networks/{id}/routes/{routeId} | Get a Static Route for a Network
[**getStaticRoutes**](NetworksApi.md#getStaticRoutes) | **GET** /api/networks/{id}/routes | Get all Static Routes for a Network
[**getSubnet**](NetworksApi.md#getSubnet) | **GET** /api/subnets/{id} | Get a Specific Subnet
[**listNetworkPoolServerTypes**](NetworksApi.md#listNetworkPoolServerTypes) | **GET** /api/networks/pool-server-types | Get All Network Pool Server Types
[**listNetworkPoolServers**](NetworksApi.md#listNetworkPoolServers) | **GET** /api/networks/pool-servers | Get All Network Pool Servers
[**listNetworkRouterTypes**](NetworksApi.md#listNetworkRouterTypes) | **GET** /api/network-router-types | Get all Network Router Types
[**listNetworkServerGroups**](NetworksApi.md#listNetworkServerGroups) | **GET** /api/networks/servers/{serverId}/groups | Get all Network Server Groups for Network Server
[**listNetworkServices**](NetworksApi.md#listNetworkServices) | **GET** /api/networks/services | Get All Network Services
[**listNetworkTypes**](NetworksApi.md#listNetworkTypes) | **GET** /api/network-types | Network Types
[**listNetworks**](NetworksApi.md#listNetworks) | **GET** /api/networks | Get All Networks
[**listSubnetTypes**](NetworksApi.md#listSubnetTypes) | **GET** /api/subnet-types | Get All Subnet Types
[**listSubnets**](NetworksApi.md#listSubnets) | **GET** /api/subnets | Get All Subnets
[**refreshNetworkServer**](NetworksApi.md#refreshNetworkServer) | **POST** /api/networks/servers/{id}/refresh | Refresh a Network Server/Integration
[**releaseNetworkFloatingIp**](NetworksApi.md#releaseNetworkFloatingIp) | **PUT** /api/networks/floating-ips/{id}/release | Release a Floating IP
[**updateNetwork**](NetworksApi.md#updateNetwork) | **PUT** /api/networks/{id} | Update a Network
[**updateNetworkDhcpRelay**](NetworksApi.md#updateNetworkDhcpRelay) | **PUT** /api/networks/servers/{serverId}/dhcp-relays/{id} | Update a Network DHCP Relay
[**updateNetworkDhcpServer**](NetworksApi.md#updateNetworkDhcpServer) | **PUT** /api/networks/servers/{serverId}/dhcp-servers/{id} | Update a Network DHCP Server
[**updateNetworkDomain**](NetworksApi.md#updateNetworkDomain) | **PUT** /api/networks/domains/{id} | Update a Network Domain
[**updateNetworkEdgeCluster**](NetworksApi.md#updateNetworkEdgeCluster) | **PUT** /api/networks/servers/{serverId}/edge-clusters/{id} | Update a Network Edge Cluster
[**updateNetworkFirewallRule**](NetworksApi.md#updateNetworkFirewallRule) | **PUT** /api/networks/servers/{serverId}/firewall-rules/{id} | Update a Network Firewall Rule
[**updateNetworkFirewallRuleGroup**](NetworksApi.md#updateNetworkFirewallRuleGroup) | **PUT** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Update a Network Firewall Rule Group
[**updateNetworkGroup**](NetworksApi.md#updateNetworkGroup) | **PUT** /api/networks/groups/{id} | Update a Network Group
[**updateNetworkPool**](NetworksApi.md#updateNetworkPool) | **PUT** /api/networks/pools/{id} | Update a Network Pool
[**updateNetworkPoolServer**](NetworksApi.md#updateNetworkPoolServer) | **PUT** /api/networks/pool-servers/{id} | Update a Network Pool Server
[**updateNetworkProxy**](NetworksApi.md#updateNetworkProxy) | **PUT** /api/networks/proxies/{id} | Update a Network Proxy
[**updateNetworkRouter**](NetworksApi.md#updateNetworkRouter) | **PUT** /api/networks/routers/{id} | Update a Network Router
[**updateNetworkRouterBgpNeighbor**](NetworksApi.md#updateNetworkRouterBgpNeighbor) | **PUT** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Update Network Router BGP Neighbor
[**updateNetworkRouterFirewallRule**](NetworksApi.md#updateNetworkRouterFirewallRule) | **PUT** /api/networks/routers/{routerId}/firewall-rules/{id} | Update a Network Router Firewall Rule
[**updateNetworkRouterFirewallRuleGroup**](NetworksApi.md#updateNetworkRouterFirewallRuleGroup) | **PUT** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Update a Network Router Firewall Rule Group
[**updateNetworkRouterNat**](NetworksApi.md#updateNetworkRouterNat) | **PUT** /api/networks/routers/{routerId}/nats/{id} | Update Network Router NAT
[**updateNetworkRouterPermissions**](NetworksApi.md#updateNetworkRouterPermissions) | **PUT** /api/networks/routers/{routerId}/permissions | Update Network Router Permissions
[**updateNetworkServerGroup**](NetworksApi.md#updateNetworkServerGroup) | **PUT** /api/networks/servers/{serverId}/groups/{id} | Update a Network Server Group
[**updateNetworkTransportZone**](NetworksApi.md#updateNetworkTransportZone) | **PUT** /api/networks/servers/{serverId}/scopes/{id} | Update a Network Transport Zone
[**updateStaticRoute**](NetworksApi.md#updateStaticRoute) | **PUT** /api/networks/{id}/routes/{routeId} | Update a Network Static Route
[**updateSubnet**](NetworksApi.md#updateSubnet) | **PUT** /api/subnets/{id} | Update a Subnet



## createNetworkDhcpRelay

> SuccessId createNetworkDhcpRelay(serverId, opts)

Create a Network DHCP Relay

Create a Network DHCP Relay. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject167': new MorpheusApi.InlineObject167() // InlineObject167 | 
};
apiInstance.createNetworkDhcpRelay(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **inlineObject167** | [**InlineObject167**](InlineObject167.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkDhcpServer

> SuccessId createNetworkDhcpServer(serverId, opts)

Create a Network DHCP Server

Create a Network DHCP Server. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject169': new MorpheusApi.InlineObject169() // InlineObject169 | 
};
apiInstance.createNetworkDhcpServer(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **inlineObject169** | [**InlineObject169**](InlineObject169.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkDomain

> InlineResponse200109 createNetworkDomain(opts)

Create a Network Domain

Create a Network Domain. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject163': new MorpheusApi.InlineObject163() // InlineObject163 | 
};
apiInstance.createNetworkDomain(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject163** | [**InlineObject163**](InlineObject163.md)|  | [optional] 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkFirewallRule

> SuccessId createNetworkFirewallRule(serverId, opts)

Create a Network Firewall Rule

Use this command to create a network firewall rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject172': new MorpheusApi.InlineObject172() // InlineObject172 | 
};
apiInstance.createNetworkFirewallRule(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **inlineObject172** | [**InlineObject172**](InlineObject172.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkFirewallRuleGroup

> SuccessId createNetworkFirewallRuleGroup(serverId, opts)

Create a Network Firewall Rule Group

Use this command to create a network firewall rule group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject174': new MorpheusApi.InlineObject174() // InlineObject174 | 
};
apiInstance.createNetworkFirewallRuleGroup(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **inlineObject174** | [**InlineObject174**](InlineObject174.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkGroup

> SuccessId createNetworkGroup(opts)

Create a Network Group

Use this command to create a network group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject146': new MorpheusApi.InlineObject146() // InlineObject146 | 
};
apiInstance.createNetworkGroup(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject146** | [**InlineObject146**](InlineObject146.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkPool

> InlineResponse200106 createNetworkPool(opts)

Create a Network Pool

Create a Network Pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject160': new MorpheusApi.InlineObject160() // InlineObject160 | 
};
apiInstance.createNetworkPool(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject160** | [**InlineObject160**](InlineObject160.md)|  | [optional] 

### Return type

[**InlineResponse200106**](InlineResponse200106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkPoolIp

> InlineResponse200107 createNetworkPoolIp(id, opts)

Create a Network Pool IP Address

Create an IP Address for a Specific Network Pool 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject162': new MorpheusApi.InlineObject162() // InlineObject162 | 
};
apiInstance.createNetworkPoolIp(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject162** | [**InlineObject162**](InlineObject162.md)|  | [optional] 

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkPoolServer

> Model200Success createNetworkPoolServer(opts)

Create a Network Pool Server

This endpoint allows creating a Network Pool Server. Only certain types of integrations support creating and deleting network pool servers, such as Bluecat, Infoblox, phpIPAM, and Solar Winds. Configuration options vary by type. Note that creating a pool server will automatically create and associate the corresponding network integration object, but management is done via the network pool server object.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject180': new MorpheusApi.InlineObject180() // InlineObject180 | 
};
apiInstance.createNetworkPoolServer(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject180** | [**InlineObject180**](InlineObject180.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkProxy

> InlineResponse200110 createNetworkProxy(opts)

Create a Network Proxy

Create a Network Proxy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject165': new MorpheusApi.InlineObject165() // InlineObject165 | 
};
apiInstance.createNetworkProxy(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject165** | [**InlineObject165**](InlineObject165.md)|  | [optional] 

### Return type

[**InlineResponse200110**](InlineResponse200110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkRouter

> SuccessId createNetworkRouter(opts)

Create a Network Router

Use this command to create a network router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject148': new MorpheusApi.InlineObject148() // InlineObject148 | 
};
apiInstance.createNetworkRouter(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject148** | [**InlineObject148**](InlineObject148.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkRouterBgpNeighbor

> SuccessId createNetworkRouterBgpNeighbor(routerId, opts)

Create a Network Router BGP Neighbor

Use this command to create a BGP Neighbor for an existing network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject150': new MorpheusApi.InlineObject150() // InlineObject150 | 
};
apiInstance.createNetworkRouterBgpNeighbor(routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 
 **inlineObject150** | [**InlineObject150**](InlineObject150.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkRouterFirewallRule

> SuccessId createNetworkRouterFirewallRule(routerId, opts)

Create a Network Router Firewall Rule

Use this command to create a firewall rule for an existing network router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject152': new MorpheusApi.InlineObject152() // InlineObject152 | 
};
apiInstance.createNetworkRouterFirewallRule(routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 
 **inlineObject152** | [**InlineObject152**](InlineObject152.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkRouterFirewallRuleGroup

> SuccessId createNetworkRouterFirewallRuleGroup(routerId, opts)

Create a Network Router Firewall Rule Group

Use this command to create a network firewall rule group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject154': new MorpheusApi.InlineObject154() // InlineObject154 | 
};
apiInstance.createNetworkRouterFirewallRuleGroup(routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 
 **inlineObject154** | [**InlineObject154**](InlineObject154.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkRouterNat

> SuccessId createNetworkRouterNat(routerId, opts)

Create a Network Router NAT

Use this command to create a NAT for an existing network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject156': new MorpheusApi.InlineObject156() // InlineObject156 | 
};
apiInstance.createNetworkRouterNat(routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 
 **inlineObject156** | [**InlineObject156**](InlineObject156.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkRouterRoute

> SuccessId createNetworkRouterRoute(routerId, opts)

Create a Network Router Route

Use this command to create a route for an existing network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject159': new MorpheusApi.InlineObject159() // InlineObject159 | 
};
apiInstance.createNetworkRouterRoute(routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 
 **inlineObject159** | [**InlineObject159**](InlineObject159.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkServerGroup

> SuccessId createNetworkServerGroup(serverId, opts)

Create a Network Server Group

Use this command to create a network server group. Note: Only available for NSX-T network integrations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject176': new MorpheusApi.InlineObject176() // InlineObject176 | 
};
apiInstance.createNetworkServerGroup(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **inlineObject176** | [**InlineObject176**](InlineObject176.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworkTransportZone

> SuccessId createNetworkTransportZone(serverId, opts)

Create a Network Transport Zone

Use this command to create a network transport zone.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject178': new MorpheusApi.InlineObject178() // InlineObject178 | 
};
apiInstance.createNetworkTransportZone(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **inlineObject178** | [**InlineObject178**](InlineObject178.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createNetworks

> Object createNetworks(opts)

Create a Network

This endpoint allows creating a Network. Only certain types of clouds support creating and deleting networks. Configuration options vary by Network Types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject142': new MorpheusApi.InlineObject142() // InlineObject142 | 
};
apiInstance.createNetworks(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject142** | [**InlineObject142**](InlineObject142.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createStaticRoute

> SuccessId createStaticRoute(id, opts)

Create a Network Static Route

Use this command to create a route for an existing network. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject144': new MorpheusApi.InlineObject144() // InlineObject144 | 
};
apiInstance.createStaticRoute(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject144** | [**InlineObject144**](InlineObject144.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createSubnet

> InlineResponse200154 createSubnet(opts)

Create a Subnet

This endpoint allows creating a Subnet. Only certain types of clouds support creating and deleting subnets. Configuration options vary for each Subnet Type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'inlineObject244': new MorpheusApi.InlineObject244() // InlineObject244 | 
};
apiInstance.createSubnet(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject244** | [**InlineObject244**](InlineObject244.md)|  | [optional] 

### Return type

[**InlineResponse200154**](InlineResponse200154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteNetwork

> Model200Success deleteNetwork(id)

Delete a Network

Will delete a Network from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetwork(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkDhcpRelay

> Model200Success deleteNetworkDhcpRelay(id, serverId)

Delete a Network DHCP Relay

Will delete a Network DHCP Relay from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.deleteNetworkDhcpRelay(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkDhcpServer

> Model200Success deleteNetworkDhcpServer(id, serverId)

Delete a Network DHCP Server

Will delete a Network DHCP Server from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.deleteNetworkDhcpServer(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkDomain

> Model200Success deleteNetworkDomain(id)

Delete a Network Domain

Will delete a Network Domain from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkDomain(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkFirewallRule

> Model200Success deleteNetworkFirewallRule(id, serverId)

Delete a Network Firewall Rule

Will delete a Network Firewall Rule from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.deleteNetworkFirewallRule(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkFirewallRuleGroup

> Model200Success deleteNetworkFirewallRuleGroup(id, serverId)

Delete a Network firewall rule group

Will delete a network firewall rule group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.deleteNetworkFirewallRuleGroup(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkGroup

> Model200Success deleteNetworkGroup(id)

Delete a Network Group

Will delete a Network Group from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkGroup(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkPool

> Model200Success deleteNetworkPool(id)

Delete a Network Pool

Will delete a Network Pool from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkPool(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkPoolIp

> Model200Success deleteNetworkPoolIp(networkPoolId, id)

Delete a host record associated with an IP Address for a Specific Network Pool

Will delete a host record associated with an IP address for a specific network pool and free up the address

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let networkPoolId = 1; // Number | Morpheus Network Pool ID of the Object being referenced
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkPoolIp(networkPoolId, id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkPoolId** | **Number**| Morpheus Network Pool ID of the Object being referenced | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkPoolServer

> Model200Success deleteNetworkPoolServer(id)

Delete a Network Pool Server

Will delete a Network Pool Server from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkPoolServer(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkProxy

> Model200Success deleteNetworkProxy(id)

Delete a Network Proxy

Will delete a Network Proxy from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkProxy(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkRouter

> Model200Success deleteNetworkRouter(id)

Delete a Network Router

Will delete a Network Router from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNetworkRouter(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkRouterBgpNeighbor

> Model200Success deleteNetworkRouterBgpNeighbor(id, routerId)

Delete a Network Router BGP Neighbor

Will delete a BGP Neighbor from a network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.deleteNetworkRouterBgpNeighbor(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkRouterFirewallRule

> Model200Success deleteNetworkRouterFirewallRule(id, routerId)

Delete a Network Router Firewall Rule

Will delete a firewall rule from a network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.deleteNetworkRouterFirewallRule(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkRouterFirewallRuleGroup

> Model200Success deleteNetworkRouterFirewallRuleGroup(id, routerId)

Delete a Network Router firewall rule group

Will delete a network router firewall rule group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.deleteNetworkRouterFirewallRuleGroup(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkRouterNat

> Model200Success deleteNetworkRouterNat(id, routerId)

Delete a Network Router NAT

Will delete a NAT from a network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.deleteNetworkRouterNat(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkRouterRoute

> Model200Success deleteNetworkRouterRoute(id, routerId)

Delete a Network Router Route

Will delete a Route from a network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.deleteNetworkRouterRoute(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkServerGroup

> Model200Success deleteNetworkServerGroup(id, serverId)

Delete a Network Server Group

Will delete a Network Server Group from the system and make it no longer usable. Note: Only available for NSX-T network integrations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.deleteNetworkServerGroup(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteNetworkTransportZone

> Model200Success deleteNetworkTransportZone(id, serverId)

Delete a Network Transport Zone

Will delete a Network Transport Zone from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.deleteNetworkTransportZone(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteStaticRoute

> Model200Success deleteStaticRoute(id, routeId)

Delete a Network Static Route

Will delete a route from a network.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routeId = 4; // Number | The ID of the route
apiInstance.deleteStaticRoute(id, routeId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routeId** | **Number**| The ID of the route | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteSubnet

> Model200Success deleteSubnet(id)

Delete a Subnet

Will delete a Subnet from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteSubnet(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getAllNetworkFloatingIps

> Object getAllNetworkFloatingIps(opts)

Get All Floating IPs

This endpoint retrieves all network floating IPs associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'ipAddress': 10.32.23.188, // String | Filter by IP Address
  'ipStatus': "ipStatus_example", // String | Filter by IP Status
  'zoneId': 789, // Number | Filter by Cloud ID
  'serverId': 789 // Number | Filter by Server ID
};
apiInstance.getAllNetworkFloatingIps(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **ipAddress** | **String**| Filter by IP Address | [optional] 
 **ipStatus** | **String**| Filter by IP Status | [optional] 
 **zoneId** | **Number**| Filter by Cloud ID | [optional] 
 **serverId** | **Number**| Filter by Server ID | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetwork

> InlineResponse20087 getNetwork(id)

Get a Specific Network

This endpoint retrieves a specific Network. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetwork(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkDhcpRelay

> InlineResponse200112 getNetworkDhcpRelay(id, serverId)

Get a Specific Network DHCP Relay

This endpoint retrieves a specific Network DHCP Relay. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.getNetworkDhcpRelay(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**InlineResponse200112**](InlineResponse200112.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkDhcpRelays

> Object getNetworkDhcpRelays(serverId, opts)

Get all Network DHCP Relays for Network Relay

This endpoint retrieves all Network DHCP Relays for a specified Network Service. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkDhcpRelays(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkDhcpServer

> InlineResponse200113 getNetworkDhcpServer(id, serverId)

Get a Specific Network DHCP Server

This endpoint retrieves a specific Network DHCP Server. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.getNetworkDhcpServer(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**InlineResponse200113**](InlineResponse200113.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkDhcpServers

> Object getNetworkDhcpServers(serverId, opts)

Get all Network DHCP Servers for Network Server

This endpoint retrieves all Network DHCP Servers for a specified Network Service. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkDhcpServers(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkDomain

> InlineResponse200109 getNetworkDomain(id)

Get a Specific Network Domain

This endpoint retrieves a specific Network Domain. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkDomain(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkDomains

> Object getNetworkDomains(opts)

Get all Network Domains

This endpoint retrieves all Network Domains associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkDomains(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkEdgeCluster

> InlineResponse200114 getNetworkEdgeCluster(id, serverId)

Get a Specific Network Edge Cluster

This endpoint retrieves a specific network Edge Cluster. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.getNetworkEdgeCluster(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**InlineResponse200114**](InlineResponse200114.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkEdgeClusters

> Object getNetworkEdgeClusters(serverId, opts)

Get all Network Edge Clusters for Network Server

This endpoint retrieves all Network Edge Clusters for a specified Network Service. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkEdgeClusters(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkFirewallRule

> InlineResponse200115 getNetworkFirewallRule(id, serverId)

Get a Specific Network Firewall Rule

This endpoint retrieves a specific Network Firewall Rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.getNetworkFirewallRule(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**InlineResponse200115**](InlineResponse200115.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkFirewallRuleGroup

> InlineResponse200116 getNetworkFirewallRuleGroup(id, serverId)

Get a Specific Network Firewall Rule Group

This endpoint retrieves a specific Network Firewall Rule Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.getNetworkFirewallRuleGroup(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**InlineResponse200116**](InlineResponse200116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkFirewallRuleGroups

> Object getNetworkFirewallRuleGroups(serverId, opts)

Get all Network Firewall Rule Groups for Network Server

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkFirewallRuleGroups(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkFirewallRules

> Object getNetworkFirewallRules(serverId, opts)

Get all Network Firewall Rules for Network Server

This endpoint retrieves all Network Firewall Rules for a specified Network Service. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkFirewallRules(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkFloatingIp

> InlineResponse200108 getNetworkFloatingIp(id)

Get a Specific Floating IP

This endpoint retrieves a specific Floating IP. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkFloatingIp(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkGroup

> InlineResponse20091 getNetworkGroup(id)

Get a Specific Network Group

This endpoint retrieves a specific Network Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkGroup(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkGroups

> InlineResponse20090 getNetworkGroups()

Get all Network Groups

This endpoint retrieves all Network Groups associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
apiInstance.getNetworkGroups((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkPool

> InlineResponse200106 getNetworkPool(id)

Get a Specific Network Pool

This endpoint retrieves a specific Network Pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkPool(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200106**](InlineResponse200106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkPoolIp

> Object getNetworkPoolIp(networkPoolId, id)

Get a Specific IP Address for a Specific Network Pool

This endpoint retrieves a specific IP address for a specific Network Pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let networkPoolId = 1; // Number | Morpheus Network Pool ID of the Object being referenced
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkPoolIp(networkPoolId, id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkPoolId** | **Number**| Morpheus Network Pool ID of the Object being referenced | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkPoolIps

> Object getNetworkPoolIps(id)

Get all IP Addresses for a Specific Network Pool

This endpoint retrieves a list of IP addresses for a specific Network Pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkPoolIps(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkPoolServer

> InlineResponse200120 getNetworkPoolServer(id)

Get a Specific Network Pool Server

This endpoint retrieves a specific Network Pool Server. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkPoolServer(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200120**](InlineResponse200120.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkPoolServerType

> InlineResponse200121 getNetworkPoolServerType(id)

Retrieves a Specific Network Pool Server Type

Retrieves a specific Network Pool Server Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkPoolServerType(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200121**](InlineResponse200121.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkPools

> Object getNetworkPools(opts)

Get all Network Pools

This endpoint retrieves all Network Pools associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkPools(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkProxies

> Object getNetworkProxies(opts)

Get all Network Proxies

This endpoint retrieves all Network Proxies associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkProxies(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkProxy

> InlineResponse200111 getNetworkProxy(id)

Get a Specific Network Proxy

This endpoint retrieves a specific Network Proxy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkProxy(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouter

> InlineResponse20095 getNetworkRouter(id)

Get a Specific Network Router

This endpoint retrieves a specific Network Router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkRouter(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterBgpNeighbor

> InlineResponse20097 getNetworkRouterBgpNeighbor(id, routerId)

Get a Network Router BGP Neighbor

This endpoint retrieves a network router BGP Neighbor for specified network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRouterBgpNeighbor(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterFirewallRule

> InlineResponse20099 getNetworkRouterFirewallRule(id, routerId)

Get a Firewall Rule for Network Router

This endpoint retrieves a firewall rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRouterFirewallRule(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterFirewallRuleGroup

> InlineResponse200101 getNetworkRouterFirewallRuleGroup(id, routerId)

Get a Firewall Rule Group for Network Router

This endpoint retrieves a firewall rule group for specified network router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRouterFirewallRuleGroup(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterFirewallRuleGroups

> InlineResponse200100 getNetworkRouterFirewallRuleGroups(routerId)

Get all Network Firewall Rule Groups for Network Router

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRouterFirewallRuleGroups(routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterNat

> InlineResponse200103 getNetworkRouterNat(id, routerId)

Get a Network Router NAT

This endpoint retrieves a network router NAT for specified network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRouterNat(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse200103**](InlineResponse200103.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterRoute

> InlineResponse200105 getNetworkRouterRoute(id, routerId)

Get a Route for Network Router

This endpoint retrieves a Route. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRouterRoute(id, routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouterType

> InlineResponse20093 getNetworkRouterType(id)

Get a Specific Network Router Type

This endpoint retrieves a specific network router type. Use this API to retrieve list of available option types for a specific network router type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkRouterType(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRouters

> InlineResponse20094 getNetworkRouters()

Get all Network Routers

This endpoint retrieves all Network Routers 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
apiInstance.getNetworkRouters((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20094**](InlineResponse20094.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRoutersBgpNeighbors

> InlineResponse20096 getNetworkRoutersBgpNeighbors(routerId)

Get all BGP Neighbors for Network Router

This endpoint retrieves all BGP Neighbors for specified network router.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRoutersBgpNeighbors(routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRoutersFirewallRules

> InlineResponse20098 getNetworkRoutersFirewallRules(routerId)

Get all Firewall Rules for Network Router

This endpoint retrieves all firewall rules for specified network router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRoutersFirewallRules(routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRoutersNats

> InlineResponse200102 getNetworkRoutersNats(routerId)

Get all Network Router NATs for Network Router

This endpoint retrieves all NATs for specified network router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRoutersNats(routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkRoutersRoutes

> InlineResponse200104 getNetworkRoutersRoutes(routerId)

Get all Routes for Network Router

This endpoint retrieves all Routes for specified network router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
apiInstance.getNetworkRoutersRoutes(routerId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 

### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkServerGroup

> InlineResponse200117 getNetworkServerGroup(serverId, id)

Get a specific Network Server Group

This endpoint retrieves a specific Network Server Group for a Network Server. Note: Only available for NSX-T network integrations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkServerGroup(serverId, id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200117**](InlineResponse200117.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkSubnets

> Object getNetworkSubnets(id)

Get Subnets for a Network

This endpoint retrieves all Subnets under a specific network. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkSubnets(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkTransportZone

> InlineResponse200118 getNetworkTransportZone(id, serverId)

Get a Specific Network Transport Zone

This endpoint retrieves a specific Network Transport Zone. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
apiInstance.getNetworkTransportZone(id, serverId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 

### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkTransportZones

> Object getNetworkTransportZones(serverId, opts)

Get all Network Transport Zones for Network Server

This endpoint retrieves all Network Transport Zones for a specified Network Service.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.getNetworkTransportZones(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNetworkType

> InlineResponse20086 getNetworkType(id)

Get a Specific Network Type

This endpoint retrieves a specific Network Type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNetworkType(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getStaticRoute

> InlineResponse20089 getStaticRoute(id, routeId)

Get a Static Route for a Network

This endpoint retrieves a network static route for specified network. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routeId = 4; // Number | The ID of the route
apiInstance.getStaticRoute(id, routeId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routeId** | **Number**| The ID of the route | 

### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getStaticRoutes

> InlineResponse20088 getStaticRoutes(id)

Get all Static Routes for a Network

This endpoint retrieves all routes for specified network. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getStaticRoutes(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20088**](InlineResponse20088.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getSubnet

> InlineResponse200154 getSubnet(id)

Get a Specific Subnet

This endpoint retrieves a specific Subnet. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getSubnet(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200154**](InlineResponse200154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworkPoolServerTypes

> Object listNetworkPoolServerTypes(opts)

Get All Network Pool Server Types

This endpoint retrieves all Network Pool Server Types 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr // String | If specified will return an exact match on code
};
apiInstance.listNetworkPoolServerTypes(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworkPoolServers

> Object listNetworkPoolServers(opts)

Get All Network Pool Servers

This endpoint retrieves all Network Pool Servers associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listNetworkPoolServers(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworkRouterTypes

> InlineResponse20092 listNetworkRouterTypes()

Get all Network Router Types

Get all Network Router Types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
apiInstance.listNetworkRouterTypes((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworkServerGroups

> Object listNetworkServerGroups(serverId, opts)

Get all Network Server Groups for Network Server

This endpoint retrieves all Network Server Groups for a specified Network Service. Note: Only available for NSX-T network integrations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let serverId = 4; // Number | Server ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listNetworkServerGroups(serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **Number**| Server ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworkServices

> InlineResponse200119 listNetworkServices(opts)

Get All Network Services

This endpoint retrieves all Network Services associated with the account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listNetworkServices(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

[**InlineResponse200119**](InlineResponse200119.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworkTypes

> Object listNetworkTypes(opts)

Network Types

Provides API for viewing Network Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listNetworkTypes(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNetworks

> Object listNetworks(opts)

Get All Networks

This endpoint retrieves all Networks associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listNetworks(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSubnetTypes

> Object listSubnetTypes(opts)

Get All Subnet Types

Provides API for viewing Network Subnet Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listSubnetTypes(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSubnets

> Object listSubnets(opts)

Get All Subnets

This endpoint retrieves all Subnets associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listSubnets(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## refreshNetworkServer

> SuccessId refreshNetworkServer(id)

Refresh a Network Server/Integration

Refreshes a network server/integration. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.refreshNetworkServer(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## releaseNetworkFloatingIp

> Model200Success releaseNetworkFloatingIp(id)

Release a Floating IP

Release a floating IP detaching it from the associated node/VM. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.releaseNetworkFloatingIp(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateNetwork

> Object updateNetwork(id, opts)

Update a Network

This endpoint allows updating a Network. Configuration options vary by Network Types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject143': new MorpheusApi.InlineObject143() // InlineObject143 | 
};
apiInstance.updateNetwork(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject143** | [**InlineObject143**](InlineObject143.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkDhcpRelay

> Model200Success updateNetworkDhcpRelay(id, serverId, opts)

Update a Network DHCP Relay

Use this command to update an existing Network DHCP Relay. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject168': new MorpheusApi.InlineObject168() // InlineObject168 | 
};
apiInstance.updateNetworkDhcpRelay(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject168** | [**InlineObject168**](InlineObject168.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkDhcpServer

> Model200Success updateNetworkDhcpServer(id, serverId, opts)

Update a Network DHCP Server

Use this command to update an existing Network DHCP Server. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject170': new MorpheusApi.InlineObject170() // InlineObject170 | 
};
apiInstance.updateNetworkDhcpServer(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject170** | [**InlineObject170**](InlineObject170.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkDomain

> InlineResponse200109 updateNetworkDomain(id, opts)

Update a Network Domain

Update a Network Domain. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject164': new MorpheusApi.InlineObject164() // InlineObject164 | 
};
apiInstance.updateNetworkDomain(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject164** | [**InlineObject164**](InlineObject164.md)|  | [optional] 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkEdgeCluster

> Model200Success updateNetworkEdgeCluster(id, serverId, opts)

Update a Network Edge Cluster

Use this command to update an existing network Edge Cluster. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject171': new MorpheusApi.InlineObject171() // InlineObject171 | 
};
apiInstance.updateNetworkEdgeCluster(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject171** | [**InlineObject171**](InlineObject171.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkFirewallRule

> Model200Success updateNetworkFirewallRule(id, serverId, opts)

Update a Network Firewall Rule

Use this command to update an existing network firewall Rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject173': new MorpheusApi.InlineObject173() // InlineObject173 | 
};
apiInstance.updateNetworkFirewallRule(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject173** | [**InlineObject173**](InlineObject173.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkFirewallRuleGroup

> Model200Success updateNetworkFirewallRuleGroup(id, serverId, opts)

Update a Network Firewall Rule Group

Use this command to update an existing Network Firewall Rule Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject175': new MorpheusApi.InlineObject175() // InlineObject175 | 
};
apiInstance.updateNetworkFirewallRuleGroup(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject175** | [**InlineObject175**](InlineObject175.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkGroup

> Model200Success updateNetworkGroup(id, opts)

Update a Network Group

Use this command to update an existing network Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject147': new MorpheusApi.InlineObject147() // InlineObject147 | 
};
apiInstance.updateNetworkGroup(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject147** | [**InlineObject147**](InlineObject147.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkPool

> InlineResponse200106 updateNetworkPool(id, opts)

Update a Network Pool

Update a Network Pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject161': new MorpheusApi.InlineObject161() // InlineObject161 | 
};
apiInstance.updateNetworkPool(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject161** | [**InlineObject161**](InlineObject161.md)|  | [optional] 

### Return type

[**InlineResponse200106**](InlineResponse200106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkPoolServer

> Model200Success updateNetworkPoolServer(id, opts)

Update a Network Pool Server

This endpoint allows updating a Network Pool Server. Configuration options vary by type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject181': new MorpheusApi.InlineObject181() // InlineObject181 | 
};
apiInstance.updateNetworkPoolServer(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject181** | [**InlineObject181**](InlineObject181.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkProxy

> Model200Success updateNetworkProxy(id, opts)

Update a Network Proxy

Use this command to update an existing network Proxy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject166': new MorpheusApi.InlineObject166() // InlineObject166 | 
};
apiInstance.updateNetworkProxy(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject166** | [**InlineObject166**](InlineObject166.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkRouter

> Model200Success updateNetworkRouter(id, opts)

Update a Network Router

Use this command to update an existing network Router. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject149': new MorpheusApi.InlineObject149() // InlineObject149 | 
};
apiInstance.updateNetworkRouter(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject149** | [**InlineObject149**](InlineObject149.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkRouterBgpNeighbor

> Model200Success updateNetworkRouterBgpNeighbor(id, routerId, opts)

Update Network Router BGP Neighbor

Use this command to update an existing Network Router BGP Neighbor. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject151': new MorpheusApi.InlineObject151() // InlineObject151 | 
};
apiInstance.updateNetworkRouterBgpNeighbor(id, routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 
 **inlineObject151** | [**InlineObject151**](InlineObject151.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkRouterFirewallRule

> Model200Success updateNetworkRouterFirewallRule(id, routerId, opts)

Update a Network Router Firewall Rule

Use this command to update an existing network router firewall rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject153': new MorpheusApi.InlineObject153() // InlineObject153 | 
};
apiInstance.updateNetworkRouterFirewallRule(id, routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 
 **inlineObject153** | [**InlineObject153**](InlineObject153.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkRouterFirewallRuleGroup

> Model200Success updateNetworkRouterFirewallRuleGroup(id, routerId, opts)

Update a Network Router Firewall Rule Group

Use this command to update an existing Network Router Firewall Rule Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject155': new MorpheusApi.InlineObject155() // InlineObject155 | 
};
apiInstance.updateNetworkRouterFirewallRuleGroup(id, routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 
 **inlineObject155** | [**InlineObject155**](InlineObject155.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkRouterNat

> Model200Success updateNetworkRouterNat(id, routerId, opts)

Update Network Router NAT

Use this command to update an existing Network Router NAT. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject157': new MorpheusApi.InlineObject157() // InlineObject157 | 
};
apiInstance.updateNetworkRouterNat(id, routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routerId** | **Number**| Router ID | 
 **inlineObject157** | [**InlineObject157**](InlineObject157.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkRouterPermissions

> SuccessId updateNetworkRouterPermissions(routerId, opts)

Update Network Router Permissions

Update Network Router Permissions 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let routerId = 4; // Number | Router ID
let opts = {
  'inlineObject158': new MorpheusApi.InlineObject158() // InlineObject158 | 
};
apiInstance.updateNetworkRouterPermissions(routerId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **Number**| Router ID | 
 **inlineObject158** | [**InlineObject158**](InlineObject158.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkServerGroup

> Model200Success updateNetworkServerGroup(id, serverId, opts)

Update a Network Server Group

Use this command to update an existing network server group. Note: Only available for NSX-T network integrations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject177': new MorpheusApi.InlineObject177() // InlineObject177 | 
};
apiInstance.updateNetworkServerGroup(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject177** | [**InlineObject177**](InlineObject177.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNetworkTransportZone

> Model200Success updateNetworkTransportZone(id, serverId, opts)

Update a Network Transport Zone

Use this command to update an existing network Transport Zone. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let serverId = 4; // Number | Server ID
let opts = {
  'inlineObject179': new MorpheusApi.InlineObject179() // InlineObject179 | 
};
apiInstance.updateNetworkTransportZone(id, serverId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **serverId** | **Number**| Server ID | 
 **inlineObject179** | [**InlineObject179**](InlineObject179.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateStaticRoute

> SuccessId updateStaticRoute(id, routeId, opts)

Update a Network Static Route

Use this command to update a route. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let routeId = 4; // Number | The ID of the route
let opts = {
  'inlineObject145': new MorpheusApi.InlineObject145() // InlineObject145 | 
};
apiInstance.updateStaticRoute(id, routeId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **routeId** | **Number**| The ID of the route | 
 **inlineObject145** | [**InlineObject145**](InlineObject145.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateSubnet

> InlineResponse200154 updateSubnet(id, opts)

Update a Subnet

This endpoint allows updating a Subnet. Only certain types of clouds support this action. Configuration options vary for each Subnet Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.NetworksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject245': new MorpheusApi.InlineObject245() // InlineObject245 | 
};
apiInstance.updateSubnet(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject245** | [**InlineObject245**](InlineObject245.md)|  | [optional] 

### Return type

[**InlineResponse200154**](InlineResponse200154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

