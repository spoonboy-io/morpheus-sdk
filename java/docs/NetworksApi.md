# NetworksApi

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


<a name="createNetworkDhcpRelay"></a>
# **createNetworkDhcpRelay**
> SuccessId createNetworkDhcpRelay(serverId, inlineObject167)

Create a Network DHCP Relay

Create a Network DHCP Relay. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject167 inlineObject167 = new InlineObject167(); // InlineObject167 | 
    try {
      SuccessId result = apiInstance.createNetworkDhcpRelay(serverId, inlineObject167);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkDhcpRelay");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject167** | [**InlineObject167**](InlineObject167.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkDhcpServer"></a>
# **createNetworkDhcpServer**
> SuccessId createNetworkDhcpServer(serverId, inlineObject169)

Create a Network DHCP Server

Create a Network DHCP Server. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject169 inlineObject169 = new InlineObject169(); // InlineObject169 | 
    try {
      SuccessId result = apiInstance.createNetworkDhcpServer(serverId, inlineObject169);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkDhcpServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject169** | [**InlineObject169**](InlineObject169.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkDomain"></a>
# **createNetworkDomain**
> InlineResponse200109 createNetworkDomain(inlineObject163)

Create a Network Domain

Create a Network Domain. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject163 inlineObject163 = new InlineObject163(); // InlineObject163 | 
    try {
      InlineResponse200109 result = apiInstance.createNetworkDomain(inlineObject163);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkDomain");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkFirewallRule"></a>
# **createNetworkFirewallRule**
> SuccessId createNetworkFirewallRule(serverId, inlineObject172)

Create a Network Firewall Rule

Use this command to create a network firewall rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject172 inlineObject172 = new InlineObject172(); // InlineObject172 | 
    try {
      SuccessId result = apiInstance.createNetworkFirewallRule(serverId, inlineObject172);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject172** | [**InlineObject172**](InlineObject172.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkFirewallRuleGroup"></a>
# **createNetworkFirewallRuleGroup**
> SuccessId createNetworkFirewallRuleGroup(serverId, inlineObject174)

Create a Network Firewall Rule Group

Use this command to create a network firewall rule group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject174 inlineObject174 = new InlineObject174(); // InlineObject174 | 
    try {
      SuccessId result = apiInstance.createNetworkFirewallRuleGroup(serverId, inlineObject174);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject174** | [**InlineObject174**](InlineObject174.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkGroup"></a>
# **createNetworkGroup**
> SuccessId createNetworkGroup(inlineObject146)

Create a Network Group

Use this command to create a network group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject146 inlineObject146 = new InlineObject146(); // InlineObject146 | 
    try {
      SuccessId result = apiInstance.createNetworkGroup(inlineObject146);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkPool"></a>
# **createNetworkPool**
> InlineResponse200106 createNetworkPool(inlineObject160)

Create a Network Pool

Create a Network Pool. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject160 inlineObject160 = new InlineObject160(); // InlineObject160 | 
    try {
      InlineResponse200106 result = apiInstance.createNetworkPool(inlineObject160);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkPool");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkPoolIp"></a>
# **createNetworkPoolIp**
> InlineResponse200107 createNetworkPoolIp(id, inlineObject162)

Create a Network Pool IP Address

Create an IP Address for a Specific Network Pool 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject162 inlineObject162 = new InlineObject162(); // InlineObject162 | 
    try {
      InlineResponse200107 result = apiInstance.createNetworkPoolIp(id, inlineObject162);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkPoolIp");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject162** | [**InlineObject162**](InlineObject162.md)|  | [optional]

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**400** | Bad Request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Object Not Found. |  -  |
**405** | Method Not Allowed. |  -  |
**406** | Not Acceptable. |  -  |
**410** | Gone. |  -  |
**429** | Too Many Requests. |  -  |
**500** | Internal Server Error. |  -  |
**503** | Service Unavailable. |  -  |

<a name="createNetworkPoolServer"></a>
# **createNetworkPoolServer**
> Model200Success createNetworkPoolServer(inlineObject180)

Create a Network Pool Server

This endpoint allows creating a Network Pool Server. Only certain types of integrations support creating and deleting network pool servers, such as Bluecat, Infoblox, phpIPAM, and Solar Winds. Configuration options vary by type. Note that creating a pool server will automatically create and associate the corresponding network integration object, but management is done via the network pool server object.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject180 inlineObject180 = new InlineObject180(); // InlineObject180 | 
    try {
      Model200Success result = apiInstance.createNetworkPoolServer(inlineObject180);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkPoolServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkProxy"></a>
# **createNetworkProxy**
> InlineResponse200110 createNetworkProxy(inlineObject165)

Create a Network Proxy

Create a Network Proxy. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject165 inlineObject165 = new InlineObject165(); // InlineObject165 | 
    try {
      InlineResponse200110 result = apiInstance.createNetworkProxy(inlineObject165);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkProxy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkRouter"></a>
# **createNetworkRouter**
> SuccessId createNetworkRouter(inlineObject148)

Create a Network Router

Use this command to create a network router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject148 inlineObject148 = new InlineObject148(); // InlineObject148 | 
    try {
      SuccessId result = apiInstance.createNetworkRouter(inlineObject148);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkRouter");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkRouterBgpNeighbor"></a>
# **createNetworkRouterBgpNeighbor**
> SuccessId createNetworkRouterBgpNeighbor(routerId, inlineObject150)

Create a Network Router BGP Neighbor

Use this command to create a BGP Neighbor for an existing network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject150 inlineObject150 = new InlineObject150(); // InlineObject150 | 
    try {
      SuccessId result = apiInstance.createNetworkRouterBgpNeighbor(routerId, inlineObject150);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkRouterBgpNeighbor");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject150** | [**InlineObject150**](InlineObject150.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkRouterFirewallRule"></a>
# **createNetworkRouterFirewallRule**
> SuccessId createNetworkRouterFirewallRule(routerId, inlineObject152)

Create a Network Router Firewall Rule

Use this command to create a firewall rule for an existing network router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject152 inlineObject152 = new InlineObject152(); // InlineObject152 | 
    try {
      SuccessId result = apiInstance.createNetworkRouterFirewallRule(routerId, inlineObject152);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkRouterFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject152** | [**InlineObject152**](InlineObject152.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkRouterFirewallRuleGroup"></a>
# **createNetworkRouterFirewallRuleGroup**
> SuccessId createNetworkRouterFirewallRuleGroup(routerId, inlineObject154)

Create a Network Router Firewall Rule Group

Use this command to create a network firewall rule group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject154 inlineObject154 = new InlineObject154(); // InlineObject154 | 
    try {
      SuccessId result = apiInstance.createNetworkRouterFirewallRuleGroup(routerId, inlineObject154);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkRouterFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject154** | [**InlineObject154**](InlineObject154.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkRouterNat"></a>
# **createNetworkRouterNat**
> SuccessId createNetworkRouterNat(routerId, inlineObject156)

Create a Network Router NAT

Use this command to create a NAT for an existing network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject156 inlineObject156 = new InlineObject156(); // InlineObject156 | 
    try {
      SuccessId result = apiInstance.createNetworkRouterNat(routerId, inlineObject156);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkRouterNat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject156** | [**InlineObject156**](InlineObject156.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkRouterRoute"></a>
# **createNetworkRouterRoute**
> SuccessId createNetworkRouterRoute(routerId, inlineObject159)

Create a Network Router Route

Use this command to create a route for an existing network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject159 inlineObject159 = new InlineObject159(); // InlineObject159 | 
    try {
      SuccessId result = apiInstance.createNetworkRouterRoute(routerId, inlineObject159);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkRouterRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject159** | [**InlineObject159**](InlineObject159.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkServerGroup"></a>
# **createNetworkServerGroup**
> SuccessId createNetworkServerGroup(serverId, inlineObject176)

Create a Network Server Group

Use this command to create a network server group. Note: Only available for NSX-T network integrations. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject176 inlineObject176 = new InlineObject176(); // InlineObject176 | 
    try {
      SuccessId result = apiInstance.createNetworkServerGroup(serverId, inlineObject176);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkServerGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject176** | [**InlineObject176**](InlineObject176.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworkTransportZone"></a>
# **createNetworkTransportZone**
> SuccessId createNetworkTransportZone(serverId, inlineObject178)

Create a Network Transport Zone

Use this command to create a network transport zone.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject178 inlineObject178 = new InlineObject178(); // InlineObject178 | 
    try {
      SuccessId result = apiInstance.createNetworkTransportZone(serverId, inlineObject178);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworkTransportZone");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject178** | [**InlineObject178**](InlineObject178.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createNetworks"></a>
# **createNetworks**
> Object createNetworks(inlineObject142)

Create a Network

This endpoint allows creating a Network. Only certain types of clouds support creating and deleting networks. Configuration options vary by Network Types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject142 inlineObject142 = new InlineObject142(); // InlineObject142 | 
    try {
      Object result = apiInstance.createNetworks(inlineObject142);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createNetworks");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createStaticRoute"></a>
# **createStaticRoute**
> SuccessId createStaticRoute(id, inlineObject144)

Create a Network Static Route

Use this command to create a route for an existing network. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject144 inlineObject144 = new InlineObject144(); // InlineObject144 | 
    try {
      SuccessId result = apiInstance.createStaticRoute(id, inlineObject144);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createStaticRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject144** | [**InlineObject144**](InlineObject144.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createSubnet"></a>
# **createSubnet**
> InlineResponse200154 createSubnet(inlineObject244)

Create a Subnet

This endpoint allows creating a Subnet. Only certain types of clouds support creating and deleting subnets. Configuration options vary for each Subnet Type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    InlineObject244 inlineObject244 = new InlineObject244(); // InlineObject244 | 
    try {
      InlineResponse200154 result = apiInstance.createSubnet(inlineObject244);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#createSubnet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetwork"></a>
# **deleteNetwork**
> Model200Success deleteNetwork(id)

Delete a Network

Will delete a Network from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetwork(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetwork");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkDhcpRelay"></a>
# **deleteNetworkDhcpRelay**
> Model200Success deleteNetworkDhcpRelay(id, serverId)

Delete a Network DHCP Relay

Will delete a Network DHCP Relay from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      Model200Success result = apiInstance.deleteNetworkDhcpRelay(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkDhcpRelay");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkDhcpServer"></a>
# **deleteNetworkDhcpServer**
> Model200Success deleteNetworkDhcpServer(id, serverId)

Delete a Network DHCP Server

Will delete a Network DHCP Server from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      Model200Success result = apiInstance.deleteNetworkDhcpServer(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkDhcpServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkDomain"></a>
# **deleteNetworkDomain**
> Model200Success deleteNetworkDomain(id)

Delete a Network Domain

Will delete a Network Domain from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkDomain(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkDomain");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkFirewallRule"></a>
# **deleteNetworkFirewallRule**
> Model200Success deleteNetworkFirewallRule(id, serverId)

Delete a Network Firewall Rule

Will delete a Network Firewall Rule from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      Model200Success result = apiInstance.deleteNetworkFirewallRule(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkFirewallRuleGroup"></a>
# **deleteNetworkFirewallRuleGroup**
> Model200Success deleteNetworkFirewallRuleGroup(id, serverId)

Delete a Network firewall rule group

Will delete a network firewall rule group.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      Model200Success result = apiInstance.deleteNetworkFirewallRuleGroup(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkGroup"></a>
# **deleteNetworkGroup**
> Model200Success deleteNetworkGroup(id)

Delete a Network Group

Will delete a Network Group from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkGroup(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkPool"></a>
# **deleteNetworkPool**
> Model200Success deleteNetworkPool(id)

Delete a Network Pool

Will delete a Network Pool from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkPool(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkPool");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkPoolIp"></a>
# **deleteNetworkPoolIp**
> Model200Success deleteNetworkPoolIp(networkPoolId, id)

Delete a host record associated with an IP Address for a Specific Network Pool

Will delete a host record associated with an IP address for a specific network pool and free up the address

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long networkPoolId = 1L; // Long | Morpheus Network Pool ID of the Object being referenced
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkPoolIp(networkPoolId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkPoolIp");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkPoolId** | **Long**| Morpheus Network Pool ID of the Object being referenced |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**400** | Bad Request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Object Not Found. |  -  |
**405** | Method Not Allowed. |  -  |
**406** | Not Acceptable. |  -  |
**410** | Gone. |  -  |
**429** | Too Many Requests. |  -  |
**500** | Internal Server Error. |  -  |
**503** | Service Unavailable. |  -  |

<a name="deleteNetworkPoolServer"></a>
# **deleteNetworkPoolServer**
> Model200Success deleteNetworkPoolServer(id)

Delete a Network Pool Server

Will delete a Network Pool Server from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkPoolServer(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkPoolServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkProxy"></a>
# **deleteNetworkProxy**
> Model200Success deleteNetworkProxy(id)

Delete a Network Proxy

Will delete a Network Proxy from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkProxy(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkProxy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkRouter"></a>
# **deleteNetworkRouter**
> Model200Success deleteNetworkRouter(id)

Delete a Network Router

Will delete a Network Router from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteNetworkRouter(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkRouter");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkRouterBgpNeighbor"></a>
# **deleteNetworkRouterBgpNeighbor**
> Model200Success deleteNetworkRouterBgpNeighbor(id, routerId)

Delete a Network Router BGP Neighbor

Will delete a BGP Neighbor from a network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      Model200Success result = apiInstance.deleteNetworkRouterBgpNeighbor(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkRouterBgpNeighbor");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkRouterFirewallRule"></a>
# **deleteNetworkRouterFirewallRule**
> Model200Success deleteNetworkRouterFirewallRule(id, routerId)

Delete a Network Router Firewall Rule

Will delete a firewall rule from a network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      Model200Success result = apiInstance.deleteNetworkRouterFirewallRule(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkRouterFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkRouterFirewallRuleGroup"></a>
# **deleteNetworkRouterFirewallRuleGroup**
> Model200Success deleteNetworkRouterFirewallRuleGroup(id, routerId)

Delete a Network Router firewall rule group

Will delete a network router firewall rule group.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      Model200Success result = apiInstance.deleteNetworkRouterFirewallRuleGroup(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkRouterFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkRouterNat"></a>
# **deleteNetworkRouterNat**
> Model200Success deleteNetworkRouterNat(id, routerId)

Delete a Network Router NAT

Will delete a NAT from a network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      Model200Success result = apiInstance.deleteNetworkRouterNat(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkRouterNat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkRouterRoute"></a>
# **deleteNetworkRouterRoute**
> Model200Success deleteNetworkRouterRoute(id, routerId)

Delete a Network Router Route

Will delete a Route from a network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      Model200Success result = apiInstance.deleteNetworkRouterRoute(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkRouterRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkServerGroup"></a>
# **deleteNetworkServerGroup**
> Model200Success deleteNetworkServerGroup(id, serverId)

Delete a Network Server Group

Will delete a Network Server Group from the system and make it no longer usable. Note: Only available for NSX-T network integrations. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      Model200Success result = apiInstance.deleteNetworkServerGroup(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkServerGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteNetworkTransportZone"></a>
# **deleteNetworkTransportZone**
> Model200Success deleteNetworkTransportZone(id, serverId)

Delete a Network Transport Zone

Will delete a Network Transport Zone from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      Model200Success result = apiInstance.deleteNetworkTransportZone(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteNetworkTransportZone");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteStaticRoute"></a>
# **deleteStaticRoute**
> Model200Success deleteStaticRoute(id, routeId)

Delete a Network Static Route

Will delete a route from a network.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routeId = new BigDecimal(78); // BigDecimal | The ID of the route
    try {
      Model200Success result = apiInstance.deleteStaticRoute(id, routeId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteStaticRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routeId** | **BigDecimal**| The ID of the route |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteSubnet"></a>
# **deleteSubnet**
> Model200Success deleteSubnet(id)

Delete a Subnet

Will delete a Subnet from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteSubnet(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#deleteSubnet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getAllNetworkFloatingIps"></a>
# **getAllNetworkFloatingIps**
> Object getAllNetworkFloatingIps(phrase, ipAddress, ipStatus, zoneId, serverId)

Get All Floating IPs

This endpoint retrieves all network floating IPs associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String ipAddress = "10.32.23.188"; // String | Filter by IP Address
    String ipStatus = "ipStatus_example"; // String | Filter by IP Status
    Long zoneId = 56L; // Long | Filter by Cloud ID
    Long serverId = 56L; // Long | Filter by Server ID
    try {
      Object result = apiInstance.getAllNetworkFloatingIps(phrase, ipAddress, ipStatus, zoneId, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getAllNetworkFloatingIps");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **ipAddress** | **String**| Filter by IP Address | [optional]
 **ipStatus** | **String**| Filter by IP Status | [optional] [enum: assigned, free, pending]
 **zoneId** | **Long**| Filter by Cloud ID | [optional]
 **serverId** | **Long**| Filter by Server ID | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetwork"></a>
# **getNetwork**
> InlineResponse20087 getNetwork(id)

Get a Specific Network

This endpoint retrieves a specific Network. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20087 result = apiInstance.getNetwork(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetwork");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkDhcpRelay"></a>
# **getNetworkDhcpRelay**
> InlineResponse200112 getNetworkDhcpRelay(id, serverId)

Get a Specific Network DHCP Relay

This endpoint retrieves a specific Network DHCP Relay. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      InlineResponse200112 result = apiInstance.getNetworkDhcpRelay(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkDhcpRelay");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**InlineResponse200112**](InlineResponse200112.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkDhcpRelays"></a>
# **getNetworkDhcpRelays**
> Object getNetworkDhcpRelays(serverId, max, offset, sort, direction, phrase)

Get all Network DHCP Relays for Network Relay

This endpoint retrieves all Network DHCP Relays for a specified Network Service. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkDhcpRelays(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkDhcpRelays");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkDhcpServer"></a>
# **getNetworkDhcpServer**
> InlineResponse200113 getNetworkDhcpServer(id, serverId)

Get a Specific Network DHCP Server

This endpoint retrieves a specific Network DHCP Server. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      InlineResponse200113 result = apiInstance.getNetworkDhcpServer(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkDhcpServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**InlineResponse200113**](InlineResponse200113.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkDhcpServers"></a>
# **getNetworkDhcpServers**
> Object getNetworkDhcpServers(serverId, max, offset, sort, direction, phrase)

Get all Network DHCP Servers for Network Server

This endpoint retrieves all Network DHCP Servers for a specified Network Service. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkDhcpServers(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkDhcpServers");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkDomain"></a>
# **getNetworkDomain**
> InlineResponse200109 getNetworkDomain(id)

Get a Specific Network Domain

This endpoint retrieves a specific Network Domain. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200109 result = apiInstance.getNetworkDomain(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkDomain");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkDomains"></a>
# **getNetworkDomains**
> Object getNetworkDomains(name, phrase)

Get all Network Domains

This endpoint retrieves all Network Domains associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkDomains(name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkDomains");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkEdgeCluster"></a>
# **getNetworkEdgeCluster**
> InlineResponse200114 getNetworkEdgeCluster(id, serverId)

Get a Specific Network Edge Cluster

This endpoint retrieves a specific network Edge Cluster. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      InlineResponse200114 result = apiInstance.getNetworkEdgeCluster(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkEdgeCluster");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**InlineResponse200114**](InlineResponse200114.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkEdgeClusters"></a>
# **getNetworkEdgeClusters**
> Object getNetworkEdgeClusters(serverId, max, offset, sort, direction, phrase)

Get all Network Edge Clusters for Network Server

This endpoint retrieves all Network Edge Clusters for a specified Network Service. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkEdgeClusters(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkEdgeClusters");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkFirewallRule"></a>
# **getNetworkFirewallRule**
> InlineResponse200115 getNetworkFirewallRule(id, serverId)

Get a Specific Network Firewall Rule

This endpoint retrieves a specific Network Firewall Rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      InlineResponse200115 result = apiInstance.getNetworkFirewallRule(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**InlineResponse200115**](InlineResponse200115.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkFirewallRuleGroup"></a>
# **getNetworkFirewallRuleGroup**
> InlineResponse200116 getNetworkFirewallRuleGroup(id, serverId)

Get a Specific Network Firewall Rule Group

This endpoint retrieves a specific Network Firewall Rule Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      InlineResponse200116 result = apiInstance.getNetworkFirewallRuleGroup(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**InlineResponse200116**](InlineResponse200116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkFirewallRuleGroups"></a>
# **getNetworkFirewallRuleGroups**
> Object getNetworkFirewallRuleGroups(serverId, max, offset, sort, direction, phrase)

Get all Network Firewall Rule Groups for Network Server

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkFirewallRuleGroups(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkFirewallRuleGroups");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkFirewallRules"></a>
# **getNetworkFirewallRules**
> Object getNetworkFirewallRules(serverId, max, offset, sort, direction, phrase)

Get all Network Firewall Rules for Network Server

This endpoint retrieves all Network Firewall Rules for a specified Network Service. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkFirewallRules(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkFirewallRules");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkFloatingIp"></a>
# **getNetworkFloatingIp**
> InlineResponse200108 getNetworkFloatingIp(id)

Get a Specific Floating IP

This endpoint retrieves a specific Floating IP. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200108 result = apiInstance.getNetworkFloatingIp(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkFloatingIp");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkGroup"></a>
# **getNetworkGroup**
> InlineResponse20091 getNetworkGroup(id)

Get a Specific Network Group

This endpoint retrieves a specific Network Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20091 result = apiInstance.getNetworkGroup(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkGroups"></a>
# **getNetworkGroups**
> InlineResponse20090 getNetworkGroups()

Get all Network Groups

This endpoint retrieves all Network Groups associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    try {
      InlineResponse20090 result = apiInstance.getNetworkGroups();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkGroups");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkPool"></a>
# **getNetworkPool**
> InlineResponse200106 getNetworkPool(id)

Get a Specific Network Pool

This endpoint retrieves a specific Network Pool. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200106 result = apiInstance.getNetworkPool(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkPool");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200106**](InlineResponse200106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkPoolIp"></a>
# **getNetworkPoolIp**
> Object getNetworkPoolIp(networkPoolId, id)

Get a Specific IP Address for a Specific Network Pool

This endpoint retrieves a specific IP address for a specific Network Pool. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long networkPoolId = 1L; // Long | Morpheus Network Pool ID of the Object being referenced
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getNetworkPoolIp(networkPoolId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkPoolIp");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkPoolId** | **Long**| Morpheus Network Pool ID of the Object being referenced |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**400** | Bad Request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Object Not Found. |  -  |
**405** | Method Not Allowed. |  -  |
**406** | Not Acceptable. |  -  |
**410** | Gone. |  -  |
**429** | Too Many Requests. |  -  |
**500** | Internal Server Error. |  -  |
**503** | Service Unavailable. |  -  |

<a name="getNetworkPoolIps"></a>
# **getNetworkPoolIps**
> Object getNetworkPoolIps(id)

Get all IP Addresses for a Specific Network Pool

This endpoint retrieves a list of IP addresses for a specific Network Pool. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getNetworkPoolIps(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkPoolIps");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**400** | Bad Request. |  -  |
**401** | Unauthorized. |  -  |
**403** | Forbidden. |  -  |
**404** | Object Not Found. |  -  |
**405** | Method Not Allowed. |  -  |
**406** | Not Acceptable. |  -  |
**410** | Gone. |  -  |
**429** | Too Many Requests. |  -  |
**500** | Internal Server Error. |  -  |
**503** | Service Unavailable. |  -  |

<a name="getNetworkPoolServer"></a>
# **getNetworkPoolServer**
> InlineResponse200120 getNetworkPoolServer(id)

Get a Specific Network Pool Server

This endpoint retrieves a specific Network Pool Server. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200120 result = apiInstance.getNetworkPoolServer(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkPoolServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200120**](InlineResponse200120.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkPoolServerType"></a>
# **getNetworkPoolServerType**
> InlineResponse200121 getNetworkPoolServerType(id)

Retrieves a Specific Network Pool Server Type

Retrieves a specific Network Pool Server Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200121 result = apiInstance.getNetworkPoolServerType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkPoolServerType");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200121**](InlineResponse200121.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkPools"></a>
# **getNetworkPools**
> Object getNetworkPools(name, phrase)

Get all Network Pools

This endpoint retrieves all Network Pools associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkPools(name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkPools");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkProxies"></a>
# **getNetworkProxies**
> Object getNetworkProxies(name, phrase)

Get all Network Proxies

This endpoint retrieves all Network Proxies associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkProxies(name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkProxies");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkProxy"></a>
# **getNetworkProxy**
> InlineResponse200111 getNetworkProxy(id)

Get a Specific Network Proxy

This endpoint retrieves a specific Network Proxy. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200111 result = apiInstance.getNetworkProxy(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkProxy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouter"></a>
# **getNetworkRouter**
> InlineResponse20095 getNetworkRouter(id)

Get a Specific Network Router

This endpoint retrieves a specific Network Router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20095 result = apiInstance.getNetworkRouter(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouter");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterBgpNeighbor"></a>
# **getNetworkRouterBgpNeighbor**
> InlineResponse20097 getNetworkRouterBgpNeighbor(id, routerId)

Get a Network Router BGP Neighbor

This endpoint retrieves a network router BGP Neighbor for specified network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse20097 result = apiInstance.getNetworkRouterBgpNeighbor(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterBgpNeighbor");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterFirewallRule"></a>
# **getNetworkRouterFirewallRule**
> InlineResponse20099 getNetworkRouterFirewallRule(id, routerId)

Get a Firewall Rule for Network Router

This endpoint retrieves a firewall rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse20099 result = apiInstance.getNetworkRouterFirewallRule(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse20099**](InlineResponse20099.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterFirewallRuleGroup"></a>
# **getNetworkRouterFirewallRuleGroup**
> InlineResponse200101 getNetworkRouterFirewallRuleGroup(id, routerId)

Get a Firewall Rule Group for Network Router

This endpoint retrieves a firewall rule group for specified network router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse200101 result = apiInstance.getNetworkRouterFirewallRuleGroup(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterFirewallRuleGroups"></a>
# **getNetworkRouterFirewallRuleGroups**
> InlineResponse200100 getNetworkRouterFirewallRuleGroups(routerId)

Get all Network Firewall Rule Groups for Network Router

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse200100 result = apiInstance.getNetworkRouterFirewallRuleGroups(routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterFirewallRuleGroups");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterNat"></a>
# **getNetworkRouterNat**
> InlineResponse200103 getNetworkRouterNat(id, routerId)

Get a Network Router NAT

This endpoint retrieves a network router NAT for specified network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse200103 result = apiInstance.getNetworkRouterNat(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterNat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse200103**](InlineResponse200103.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterRoute"></a>
# **getNetworkRouterRoute**
> InlineResponse200105 getNetworkRouterRoute(id, routerId)

Get a Route for Network Router

This endpoint retrieves a Route. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse200105 result = apiInstance.getNetworkRouterRoute(id, routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouterType"></a>
# **getNetworkRouterType**
> InlineResponse20093 getNetworkRouterType(id)

Get a Specific Network Router Type

This endpoint retrieves a specific network router type. Use this API to retrieve list of available option types for a specific network router type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20093 result = apiInstance.getNetworkRouterType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouterType");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRouters"></a>
# **getNetworkRouters**
> InlineResponse20094 getNetworkRouters()

Get all Network Routers

This endpoint retrieves all Network Routers 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    try {
      InlineResponse20094 result = apiInstance.getNetworkRouters();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRouters");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRoutersBgpNeighbors"></a>
# **getNetworkRoutersBgpNeighbors**
> InlineResponse20096 getNetworkRoutersBgpNeighbors(routerId)

Get all BGP Neighbors for Network Router

This endpoint retrieves all BGP Neighbors for specified network router.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse20096 result = apiInstance.getNetworkRoutersBgpNeighbors(routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRoutersBgpNeighbors");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRoutersFirewallRules"></a>
# **getNetworkRoutersFirewallRules**
> InlineResponse20098 getNetworkRoutersFirewallRules(routerId)

Get all Firewall Rules for Network Router

This endpoint retrieves all firewall rules for specified network router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse20098 result = apiInstance.getNetworkRoutersFirewallRules(routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRoutersFirewallRules");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRoutersNats"></a>
# **getNetworkRoutersNats**
> InlineResponse200102 getNetworkRoutersNats(routerId)

Get all Network Router NATs for Network Router

This endpoint retrieves all NATs for specified network router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse200102 result = apiInstance.getNetworkRoutersNats(routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRoutersNats");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkRoutersRoutes"></a>
# **getNetworkRoutersRoutes**
> InlineResponse200104 getNetworkRoutersRoutes(routerId)

Get all Routes for Network Router

This endpoint retrieves all Routes for specified network router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    try {
      InlineResponse200104 result = apiInstance.getNetworkRoutersRoutes(routerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkRoutersRoutes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |

### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkServerGroup"></a>
# **getNetworkServerGroup**
> InlineResponse200117 getNetworkServerGroup(serverId, id)

Get a specific Network Server Group

This endpoint retrieves a specific Network Server Group for a Network Server. Note: Only available for NSX-T network integrations. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200117 result = apiInstance.getNetworkServerGroup(serverId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkServerGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200117**](InlineResponse200117.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkSubnets"></a>
# **getNetworkSubnets**
> Object getNetworkSubnets(id)

Get Subnets for a Network

This endpoint retrieves all Subnets under a specific network. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getNetworkSubnets(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkSubnets");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkTransportZone"></a>
# **getNetworkTransportZone**
> InlineResponse200118 getNetworkTransportZone(id, serverId)

Get a Specific Network Transport Zone

This endpoint retrieves a specific Network Transport Zone. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    try {
      InlineResponse200118 result = apiInstance.getNetworkTransportZone(id, serverId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkTransportZone");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |

### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkTransportZones"></a>
# **getNetworkTransportZones**
> Object getNetworkTransportZones(serverId, max, offset, sort, direction, phrase)

Get all Network Transport Zones for Network Server

This endpoint retrieves all Network Transport Zones for a specified Network Service.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.getNetworkTransportZones(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkTransportZones");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getNetworkType"></a>
# **getNetworkType**
> InlineResponse20086 getNetworkType(id)

Get a Specific Network Type

This endpoint retrieves a specific Network Type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20086 result = apiInstance.getNetworkType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getNetworkType");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getStaticRoute"></a>
# **getStaticRoute**
> InlineResponse20089 getStaticRoute(id, routeId)

Get a Static Route for a Network

This endpoint retrieves a network static route for specified network. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routeId = new BigDecimal(78); // BigDecimal | The ID of the route
    try {
      InlineResponse20089 result = apiInstance.getStaticRoute(id, routeId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getStaticRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routeId** | **BigDecimal**| The ID of the route |

### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getStaticRoutes"></a>
# **getStaticRoutes**
> InlineResponse20088 getStaticRoutes(id)

Get all Static Routes for a Network

This endpoint retrieves all routes for specified network. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20088 result = apiInstance.getStaticRoutes(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getStaticRoutes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20088**](InlineResponse20088.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getSubnet"></a>
# **getSubnet**
> InlineResponse200154 getSubnet(id)

Get a Specific Subnet

This endpoint retrieves a specific Subnet. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200154 result = apiInstance.getSubnet(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#getSubnet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200154**](InlineResponse200154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworkPoolServerTypes"></a>
# **listNetworkPoolServerTypes**
> Object listNetworkPoolServerTypes(max, offset, sort, direction, phrase, name, code)

Get All Network Pool Server Types

This endpoint retrieves all Network Pool Server Types 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    try {
      Object result = apiInstance.listNetworkPoolServerTypes(max, offset, sort, direction, phrase, name, code);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworkPoolServerTypes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworkPoolServers"></a>
# **listNetworkPoolServers**
> Object listNetworkPoolServers(max, offset, sort, direction, phrase, name)

Get All Network Pool Servers

This endpoint retrieves all Network Pool Servers associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      Object result = apiInstance.listNetworkPoolServers(max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworkPoolServers");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworkRouterTypes"></a>
# **listNetworkRouterTypes**
> InlineResponse20092 listNetworkRouterTypes()

Get all Network Router Types

Get all Network Router Types. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    try {
      InlineResponse20092 result = apiInstance.listNetworkRouterTypes();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworkRouterTypes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworkServerGroups"></a>
# **listNetworkServerGroups**
> Object listNetworkServerGroups(serverId, max, offset, sort, direction, phrase)

Get all Network Server Groups for Network Server

This endpoint retrieves all Network Server Groups for a specified Network Service. Note: Only available for NSX-T network integrations. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listNetworkServerGroups(serverId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworkServerGroups");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **BigDecimal**| Server ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworkServices"></a>
# **listNetworkServices**
> InlineResponse200119 listNetworkServices(name, phrase)

Get All Network Services

This endpoint retrieves all Network Services associated with the account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      InlineResponse200119 result = apiInstance.listNetworkServices(name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworkServices");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworkTypes"></a>
# **listNetworkTypes**
> Object listNetworkTypes(name, code, phrase)

Network Types

Provides API for viewing Network Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listNetworkTypes(name, code, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworkTypes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listNetworks"></a>
# **listNetworks**
> Object listNetworks(name, phrase, labels, allLabels)

Get All Networks

This endpoint retrieves all Networks associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      Object result = apiInstance.listNetworks(name, phrase, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listNetworks");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listSubnetTypes"></a>
# **listSubnetTypes**
> Object listSubnetTypes(name, phrase)

Get All Subnet Types

Provides API for viewing Network Subnet Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listSubnetTypes(name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listSubnetTypes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listSubnets"></a>
# **listSubnets**
> Object listSubnets(name, phrase, labels, allLabels)

Get All Subnets

This endpoint retrieves all Subnets associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      Object result = apiInstance.listSubnets(name, phrase, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#listSubnets");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="refreshNetworkServer"></a>
# **refreshNetworkServer**
> SuccessId refreshNetworkServer(id)

Refresh a Network Server/Integration

Refreshes a network server/integration. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessId result = apiInstance.refreshNetworkServer(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#refreshNetworkServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="releaseNetworkFloatingIp"></a>
# **releaseNetworkFloatingIp**
> Model200Success releaseNetworkFloatingIp(id)

Release a Floating IP

Release a floating IP detaching it from the associated node/VM. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.releaseNetworkFloatingIp(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#releaseNetworkFloatingIp");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetwork"></a>
# **updateNetwork**
> Object updateNetwork(id, inlineObject143)

Update a Network

This endpoint allows updating a Network. Configuration options vary by Network Types. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject143 inlineObject143 = new InlineObject143(); // InlineObject143 | 
    try {
      Object result = apiInstance.updateNetwork(id, inlineObject143);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetwork");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject143** | [**InlineObject143**](InlineObject143.md)|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkDhcpRelay"></a>
# **updateNetworkDhcpRelay**
> Model200Success updateNetworkDhcpRelay(id, serverId, inlineObject168)

Update a Network DHCP Relay

Use this command to update an existing Network DHCP Relay. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject168 inlineObject168 = new InlineObject168(); // InlineObject168 | 
    try {
      Model200Success result = apiInstance.updateNetworkDhcpRelay(id, serverId, inlineObject168);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkDhcpRelay");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject168** | [**InlineObject168**](InlineObject168.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkDhcpServer"></a>
# **updateNetworkDhcpServer**
> Model200Success updateNetworkDhcpServer(id, serverId, inlineObject170)

Update a Network DHCP Server

Use this command to update an existing Network DHCP Server. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject170 inlineObject170 = new InlineObject170(); // InlineObject170 | 
    try {
      Model200Success result = apiInstance.updateNetworkDhcpServer(id, serverId, inlineObject170);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkDhcpServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject170** | [**InlineObject170**](InlineObject170.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkDomain"></a>
# **updateNetworkDomain**
> InlineResponse200109 updateNetworkDomain(id, inlineObject164)

Update a Network Domain

Update a Network Domain. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject164 inlineObject164 = new InlineObject164(); // InlineObject164 | 
    try {
      InlineResponse200109 result = apiInstance.updateNetworkDomain(id, inlineObject164);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkDomain");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject164** | [**InlineObject164**](InlineObject164.md)|  | [optional]

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkEdgeCluster"></a>
# **updateNetworkEdgeCluster**
> Model200Success updateNetworkEdgeCluster(id, serverId, inlineObject171)

Update a Network Edge Cluster

Use this command to update an existing network Edge Cluster. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject171 inlineObject171 = new InlineObject171(); // InlineObject171 | 
    try {
      Model200Success result = apiInstance.updateNetworkEdgeCluster(id, serverId, inlineObject171);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkEdgeCluster");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject171** | [**InlineObject171**](InlineObject171.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkFirewallRule"></a>
# **updateNetworkFirewallRule**
> Model200Success updateNetworkFirewallRule(id, serverId, inlineObject173)

Update a Network Firewall Rule

Use this command to update an existing network firewall Rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject173 inlineObject173 = new InlineObject173(); // InlineObject173 | 
    try {
      Model200Success result = apiInstance.updateNetworkFirewallRule(id, serverId, inlineObject173);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject173** | [**InlineObject173**](InlineObject173.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkFirewallRuleGroup"></a>
# **updateNetworkFirewallRuleGroup**
> Model200Success updateNetworkFirewallRuleGroup(id, serverId, inlineObject175)

Update a Network Firewall Rule Group

Use this command to update an existing Network Firewall Rule Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject175 inlineObject175 = new InlineObject175(); // InlineObject175 | 
    try {
      Model200Success result = apiInstance.updateNetworkFirewallRuleGroup(id, serverId, inlineObject175);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject175** | [**InlineObject175**](InlineObject175.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkGroup"></a>
# **updateNetworkGroup**
> Model200Success updateNetworkGroup(id, inlineObject147)

Update a Network Group

Use this command to update an existing network Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject147 inlineObject147 = new InlineObject147(); // InlineObject147 | 
    try {
      Model200Success result = apiInstance.updateNetworkGroup(id, inlineObject147);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject147** | [**InlineObject147**](InlineObject147.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkPool"></a>
# **updateNetworkPool**
> InlineResponse200106 updateNetworkPool(id, inlineObject161)

Update a Network Pool

Update a Network Pool. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject161 inlineObject161 = new InlineObject161(); // InlineObject161 | 
    try {
      InlineResponse200106 result = apiInstance.updateNetworkPool(id, inlineObject161);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkPool");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject161** | [**InlineObject161**](InlineObject161.md)|  | [optional]

### Return type

[**InlineResponse200106**](InlineResponse200106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkPoolServer"></a>
# **updateNetworkPoolServer**
> Model200Success updateNetworkPoolServer(id, inlineObject181)

Update a Network Pool Server

This endpoint allows updating a Network Pool Server. Configuration options vary by type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject181 inlineObject181 = new InlineObject181(); // InlineObject181 | 
    try {
      Model200Success result = apiInstance.updateNetworkPoolServer(id, inlineObject181);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkPoolServer");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject181** | [**InlineObject181**](InlineObject181.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkProxy"></a>
# **updateNetworkProxy**
> Model200Success updateNetworkProxy(id, inlineObject166)

Update a Network Proxy

Use this command to update an existing network Proxy. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject166 inlineObject166 = new InlineObject166(); // InlineObject166 | 
    try {
      Model200Success result = apiInstance.updateNetworkProxy(id, inlineObject166);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkProxy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject166** | [**InlineObject166**](InlineObject166.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkRouter"></a>
# **updateNetworkRouter**
> Model200Success updateNetworkRouter(id, inlineObject149)

Update a Network Router

Use this command to update an existing network Router. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject149 inlineObject149 = new InlineObject149(); // InlineObject149 | 
    try {
      Model200Success result = apiInstance.updateNetworkRouter(id, inlineObject149);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkRouter");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject149** | [**InlineObject149**](InlineObject149.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkRouterBgpNeighbor"></a>
# **updateNetworkRouterBgpNeighbor**
> Model200Success updateNetworkRouterBgpNeighbor(id, routerId, inlineObject151)

Update Network Router BGP Neighbor

Use this command to update an existing Network Router BGP Neighbor. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject151 inlineObject151 = new InlineObject151(); // InlineObject151 | 
    try {
      Model200Success result = apiInstance.updateNetworkRouterBgpNeighbor(id, routerId, inlineObject151);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkRouterBgpNeighbor");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject151** | [**InlineObject151**](InlineObject151.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkRouterFirewallRule"></a>
# **updateNetworkRouterFirewallRule**
> Model200Success updateNetworkRouterFirewallRule(id, routerId, inlineObject153)

Update a Network Router Firewall Rule

Use this command to update an existing network router firewall rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject153 inlineObject153 = new InlineObject153(); // InlineObject153 | 
    try {
      Model200Success result = apiInstance.updateNetworkRouterFirewallRule(id, routerId, inlineObject153);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkRouterFirewallRule");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject153** | [**InlineObject153**](InlineObject153.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkRouterFirewallRuleGroup"></a>
# **updateNetworkRouterFirewallRuleGroup**
> Model200Success updateNetworkRouterFirewallRuleGroup(id, routerId, inlineObject155)

Update a Network Router Firewall Rule Group

Use this command to update an existing Network Router Firewall Rule Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject155 inlineObject155 = new InlineObject155(); // InlineObject155 | 
    try {
      Model200Success result = apiInstance.updateNetworkRouterFirewallRuleGroup(id, routerId, inlineObject155);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkRouterFirewallRuleGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject155** | [**InlineObject155**](InlineObject155.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkRouterNat"></a>
# **updateNetworkRouterNat**
> Model200Success updateNetworkRouterNat(id, routerId, inlineObject157)

Update Network Router NAT

Use this command to update an existing Network Router NAT. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject157 inlineObject157 = new InlineObject157(); // InlineObject157 | 
    try {
      Model200Success result = apiInstance.updateNetworkRouterNat(id, routerId, inlineObject157);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkRouterNat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject157** | [**InlineObject157**](InlineObject157.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkRouterPermissions"></a>
# **updateNetworkRouterPermissions**
> SuccessId updateNetworkRouterPermissions(routerId, inlineObject158)

Update Network Router Permissions

Update Network Router Permissions 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    BigDecimal routerId = new BigDecimal(78); // BigDecimal | Router ID
    InlineObject158 inlineObject158 = new InlineObject158(); // InlineObject158 | 
    try {
      SuccessId result = apiInstance.updateNetworkRouterPermissions(routerId, inlineObject158);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkRouterPermissions");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routerId** | **BigDecimal**| Router ID |
 **inlineObject158** | [**InlineObject158**](InlineObject158.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkServerGroup"></a>
# **updateNetworkServerGroup**
> Model200Success updateNetworkServerGroup(id, serverId, inlineObject177)

Update a Network Server Group

Use this command to update an existing network server group. Note: Only available for NSX-T network integrations. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject177 inlineObject177 = new InlineObject177(); // InlineObject177 | 
    try {
      Model200Success result = apiInstance.updateNetworkServerGroup(id, serverId, inlineObject177);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkServerGroup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject177** | [**InlineObject177**](InlineObject177.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateNetworkTransportZone"></a>
# **updateNetworkTransportZone**
> Model200Success updateNetworkTransportZone(id, serverId, inlineObject179)

Update a Network Transport Zone

Use this command to update an existing network Transport Zone. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal serverId = new BigDecimal(78); // BigDecimal | Server ID
    InlineObject179 inlineObject179 = new InlineObject179(); // InlineObject179 | 
    try {
      Model200Success result = apiInstance.updateNetworkTransportZone(id, serverId, inlineObject179);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateNetworkTransportZone");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **serverId** | **BigDecimal**| Server ID |
 **inlineObject179** | [**InlineObject179**](InlineObject179.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateStaticRoute"></a>
# **updateStaticRoute**
> SuccessId updateStaticRoute(id, routeId, inlineObject145)

Update a Network Static Route

Use this command to update a route. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal routeId = new BigDecimal(78); // BigDecimal | The ID of the route
    InlineObject145 inlineObject145 = new InlineObject145(); // InlineObject145 | 
    try {
      SuccessId result = apiInstance.updateStaticRoute(id, routeId, inlineObject145);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateStaticRoute");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **routeId** | **BigDecimal**| The ID of the route |
 **inlineObject145** | [**InlineObject145**](InlineObject145.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateSubnet"></a>
# **updateSubnet**
> InlineResponse200154 updateSubnet(id, inlineObject245)

Update a Subnet

This endpoint allows updating a Subnet. Only certain types of clouds support this action. Configuration options vary for each Subnet Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.NetworksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    NetworksApi apiInstance = new NetworksApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject245 inlineObject245 = new InlineObject245(); // InlineObject245 | 
    try {
      InlineResponse200154 result = apiInstance.updateSubnet(id, inlineObject245);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NetworksApi#updateSubnet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject245** | [**InlineObject245**](InlineObject245.md)|  | [optional]

### Return type

[**InlineResponse200154**](InlineResponse200154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

