# \NetworksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkDhcpRelay**](NetworksApi.md#CreateNetworkDhcpRelay) | **Post** /api/networks/servers/{serverId}/dhcp-relays | Create a Network DHCP Relay
[**CreateNetworkDhcpServer**](NetworksApi.md#CreateNetworkDhcpServer) | **Post** /api/networks/servers/{serverId}/dhcp-servers | Create a Network DHCP Server
[**CreateNetworkDomain**](NetworksApi.md#CreateNetworkDomain) | **Post** /api/networks/domains | Create a Network Domain
[**CreateNetworkFirewallRule**](NetworksApi.md#CreateNetworkFirewallRule) | **Post** /api/networks/servers/{serverId}/firewall-rules | Create a Network Firewall Rule
[**CreateNetworkFirewallRuleGroup**](NetworksApi.md#CreateNetworkFirewallRuleGroup) | **Post** /api/networks/servers/{serverId}/firewall-rule-groups | Create a Network Firewall Rule Group
[**CreateNetworkGroup**](NetworksApi.md#CreateNetworkGroup) | **Post** /api/networks/groups | Create a Network Group
[**CreateNetworkPool**](NetworksApi.md#CreateNetworkPool) | **Post** /api/networks/pools | Create a Network Pool
[**CreateNetworkPoolIp**](NetworksApi.md#CreateNetworkPoolIp) | **Post** /api/networks/pools/{id}/ips | Create a Network Pool IP Address
[**CreateNetworkPoolServer**](NetworksApi.md#CreateNetworkPoolServer) | **Post** /api/networks/pool-servers | Create a Network Pool Server
[**CreateNetworkProxy**](NetworksApi.md#CreateNetworkProxy) | **Post** /api/networks/proxies | Create a Network Proxy
[**CreateNetworkRouter**](NetworksApi.md#CreateNetworkRouter) | **Post** /api/networks/routers | Create a Network Router
[**CreateNetworkRouterBgpNeighbor**](NetworksApi.md#CreateNetworkRouterBgpNeighbor) | **Post** /api/networks/routers/{routerId}/bgp-neighbors | Create a Network Router BGP Neighbor
[**CreateNetworkRouterFirewallRule**](NetworksApi.md#CreateNetworkRouterFirewallRule) | **Post** /api/networks/routers/{routerId}/firewall-rules | Create a Network Router Firewall Rule
[**CreateNetworkRouterFirewallRuleGroup**](NetworksApi.md#CreateNetworkRouterFirewallRuleGroup) | **Post** /api/networks/routers/{routerId}/firewall-rule-groups | Create a Network Router Firewall Rule Group
[**CreateNetworkRouterNat**](NetworksApi.md#CreateNetworkRouterNat) | **Post** /api/networks/routers/{routerId}/nats | Create a Network Router NAT
[**CreateNetworkRouterRoute**](NetworksApi.md#CreateNetworkRouterRoute) | **Post** /api/networks/routers/{routerId}/routes | Create a Network Router Route
[**CreateNetworkServerGroup**](NetworksApi.md#CreateNetworkServerGroup) | **Post** /api/networks/servers/{serverId}/groups | Create a Network Server Group
[**CreateNetworkTransportZone**](NetworksApi.md#CreateNetworkTransportZone) | **Post** /api/networks/servers/{serverId}/scopes | Create a Network Transport Zone
[**CreateNetworks**](NetworksApi.md#CreateNetworks) | **Post** /api/networks | Create a Network
[**CreateStaticRoute**](NetworksApi.md#CreateStaticRoute) | **Put** /api/networks/{id}/routes | Create a Network Static Route
[**CreateSubnet**](NetworksApi.md#CreateSubnet) | **Post** /api/subnets | Create a Subnet
[**DeleteNetwork**](NetworksApi.md#DeleteNetwork) | **Delete** /api/networks/{id} | Delete a Network
[**DeleteNetworkDhcpRelay**](NetworksApi.md#DeleteNetworkDhcpRelay) | **Delete** /api/networks/servers/{serverId}/dhcp-relays/{id} | Delete a Network DHCP Relay
[**DeleteNetworkDhcpServer**](NetworksApi.md#DeleteNetworkDhcpServer) | **Delete** /api/networks/servers/{serverId}/dhcp-servers/{id} | Delete a Network DHCP Server
[**DeleteNetworkDomain**](NetworksApi.md#DeleteNetworkDomain) | **Delete** /api/networks/domains/{id} | Delete a Network Domain
[**DeleteNetworkFirewallRule**](NetworksApi.md#DeleteNetworkFirewallRule) | **Delete** /api/networks/servers/{serverId}/firewall-rules/{id} | Delete a Network Firewall Rule
[**DeleteNetworkFirewallRuleGroup**](NetworksApi.md#DeleteNetworkFirewallRuleGroup) | **Delete** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Delete a Network firewall rule group
[**DeleteNetworkGroup**](NetworksApi.md#DeleteNetworkGroup) | **Delete** /api/networks/groups/{id} | Delete a Network Group
[**DeleteNetworkPool**](NetworksApi.md#DeleteNetworkPool) | **Delete** /api/networks/pools/{id} | Delete a Network Pool
[**DeleteNetworkPoolIp**](NetworksApi.md#DeleteNetworkPoolIp) | **Delete** /api/networks/pools/{networkPoolId}/ips/{id} | Delete a host record associated with an IP Address for a Specific Network Pool
[**DeleteNetworkPoolServer**](NetworksApi.md#DeleteNetworkPoolServer) | **Delete** /api/networks/pool-servers/{id} | Delete a Network Pool Server
[**DeleteNetworkProxy**](NetworksApi.md#DeleteNetworkProxy) | **Delete** /api/networks/proxies/{id} | Delete a Network Proxy
[**DeleteNetworkRouter**](NetworksApi.md#DeleteNetworkRouter) | **Delete** /api/networks/routers/{id} | Delete a Network Router
[**DeleteNetworkRouterBgpNeighbor**](NetworksApi.md#DeleteNetworkRouterBgpNeighbor) | **Delete** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Delete a Network Router BGP Neighbor
[**DeleteNetworkRouterFirewallRule**](NetworksApi.md#DeleteNetworkRouterFirewallRule) | **Delete** /api/networks/routers/{routerId}/firewall-rules/{id} | Delete a Network Router Firewall Rule
[**DeleteNetworkRouterFirewallRuleGroup**](NetworksApi.md#DeleteNetworkRouterFirewallRuleGroup) | **Delete** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Delete a Network Router firewall rule group
[**DeleteNetworkRouterNat**](NetworksApi.md#DeleteNetworkRouterNat) | **Delete** /api/networks/routers/{routerId}/nats/{id} | Delete a Network Router NAT
[**DeleteNetworkRouterRoute**](NetworksApi.md#DeleteNetworkRouterRoute) | **Delete** /api/networks/routers/{routerId}/routes/{id} | Delete a Network Router Route
[**DeleteNetworkServerGroup**](NetworksApi.md#DeleteNetworkServerGroup) | **Delete** /api/networks/servers/{serverId}/groups/{id} | Delete a Network Server Group
[**DeleteNetworkTransportZone**](NetworksApi.md#DeleteNetworkTransportZone) | **Delete** /api/networks/servers/{serverId}/scopes/{id} | Delete a Network Transport Zone
[**DeleteStaticRoute**](NetworksApi.md#DeleteStaticRoute) | **Delete** /api/networks/{id}/routes/{routeId} | Delete a Network Static Route
[**DeleteSubnet**](NetworksApi.md#DeleteSubnet) | **Delete** /api/subnets/{id} | Delete a Subnet
[**GetAllNetworkFloatingIps**](NetworksApi.md#GetAllNetworkFloatingIps) | **Get** /api/networks/floating-ips | Get All Floating IPs
[**GetNetwork**](NetworksApi.md#GetNetwork) | **Get** /api/networks/{id} | Get a Specific Network
[**GetNetworkDhcpRelay**](NetworksApi.md#GetNetworkDhcpRelay) | **Get** /api/networks/servers/{serverId}/dhcp-relays/{id} | Get a Specific Network DHCP Relay
[**GetNetworkDhcpRelays**](NetworksApi.md#GetNetworkDhcpRelays) | **Get** /api/networks/servers/{serverId}/dhcp-relays | Get all Network DHCP Relays for Network Relay
[**GetNetworkDhcpServer**](NetworksApi.md#GetNetworkDhcpServer) | **Get** /api/networks/servers/{serverId}/dhcp-servers/{id} | Get a Specific Network DHCP Server
[**GetNetworkDhcpServers**](NetworksApi.md#GetNetworkDhcpServers) | **Get** /api/networks/servers/{serverId}/dhcp-servers | Get all Network DHCP Servers for Network Server
[**GetNetworkDomain**](NetworksApi.md#GetNetworkDomain) | **Get** /api/networks/domains/{id} | Get a Specific Network Domain
[**GetNetworkDomains**](NetworksApi.md#GetNetworkDomains) | **Get** /api/networks/domains | Get all Network Domains
[**GetNetworkEdgeCluster**](NetworksApi.md#GetNetworkEdgeCluster) | **Get** /api/networks/servers/{serverId}/edge-clusters/{id} | Get a Specific Network Edge Cluster
[**GetNetworkEdgeClusters**](NetworksApi.md#GetNetworkEdgeClusters) | **Get** /api/networks/servers/{serverId}/edge-clusters | Get all Network Edge Clusters for Network Server
[**GetNetworkFirewallRule**](NetworksApi.md#GetNetworkFirewallRule) | **Get** /api/networks/servers/{serverId}/firewall-rules/{id} | Get a Specific Network Firewall Rule
[**GetNetworkFirewallRuleGroup**](NetworksApi.md#GetNetworkFirewallRuleGroup) | **Get** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Get a Specific Network Firewall Rule Group
[**GetNetworkFirewallRuleGroups**](NetworksApi.md#GetNetworkFirewallRuleGroups) | **Get** /api/networks/servers/{serverId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Server
[**GetNetworkFirewallRules**](NetworksApi.md#GetNetworkFirewallRules) | **Get** /api/networks/servers/{serverId}/firewall-rules | Get all Network Firewall Rules for Network Server
[**GetNetworkFloatingIp**](NetworksApi.md#GetNetworkFloatingIp) | **Get** /api/networks/floating-ips/{id} | Get a Specific Floating IP
[**GetNetworkGroup**](NetworksApi.md#GetNetworkGroup) | **Get** /api/networks/groups/{id} | Get a Specific Network Group
[**GetNetworkGroups**](NetworksApi.md#GetNetworkGroups) | **Get** /api/networks/groups | Get all Network Groups
[**GetNetworkPool**](NetworksApi.md#GetNetworkPool) | **Get** /api/networks/pools/{id} | Get a Specific Network Pool
[**GetNetworkPoolIp**](NetworksApi.md#GetNetworkPoolIp) | **Get** /api/networks/pools/{networkPoolId}/ips/{id} | Get a Specific IP Address for a Specific Network Pool
[**GetNetworkPoolIps**](NetworksApi.md#GetNetworkPoolIps) | **Get** /api/networks/pools/{id}/ips | Get all IP Addresses for a Specific Network Pool
[**GetNetworkPoolServer**](NetworksApi.md#GetNetworkPoolServer) | **Get** /api/networks/pool-servers/{id} | Get a Specific Network Pool Server
[**GetNetworkPoolServerType**](NetworksApi.md#GetNetworkPoolServerType) | **Get** /api/networks/pool-server-types/{id} | Retrieves a Specific Network Pool Server Type
[**GetNetworkPools**](NetworksApi.md#GetNetworkPools) | **Get** /api/networks/pools | Get all Network Pools
[**GetNetworkProxies**](NetworksApi.md#GetNetworkProxies) | **Get** /api/networks/proxies | Get all Network Proxies
[**GetNetworkProxy**](NetworksApi.md#GetNetworkProxy) | **Get** /api/networks/proxies/{id} | Get a Specific Network Proxy
[**GetNetworkRouter**](NetworksApi.md#GetNetworkRouter) | **Get** /api/networks/routers/{id} | Get a Specific Network Router
[**GetNetworkRouterBgpNeighbor**](NetworksApi.md#GetNetworkRouterBgpNeighbor) | **Get** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Get a Network Router BGP Neighbor
[**GetNetworkRouterFirewallRule**](NetworksApi.md#GetNetworkRouterFirewallRule) | **Get** /api/networks/routers/{routerId}/firewall-rules/{id} | Get a Firewall Rule for Network Router
[**GetNetworkRouterFirewallRuleGroup**](NetworksApi.md#GetNetworkRouterFirewallRuleGroup) | **Get** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Get a Firewall Rule Group for Network Router
[**GetNetworkRouterFirewallRuleGroups**](NetworksApi.md#GetNetworkRouterFirewallRuleGroups) | **Get** /api/networks/routers/{routerId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Router
[**GetNetworkRouterNat**](NetworksApi.md#GetNetworkRouterNat) | **Get** /api/networks/routers/{routerId}/nats/{id} | Get a Network Router NAT
[**GetNetworkRouterRoute**](NetworksApi.md#GetNetworkRouterRoute) | **Get** /api/networks/routers/{routerId}/routes/{id} | Get a Route for Network Router
[**GetNetworkRouterType**](NetworksApi.md#GetNetworkRouterType) | **Get** /api/network-router-types/{id} | Get a Specific Network Router Type
[**GetNetworkRouters**](NetworksApi.md#GetNetworkRouters) | **Get** /api/networks/routers | Get all Network Routers
[**GetNetworkRoutersBgpNeighbors**](NetworksApi.md#GetNetworkRoutersBgpNeighbors) | **Get** /api/networks/routers/{routerId}/bgp-neighbors | Get all BGP Neighbors for Network Router
[**GetNetworkRoutersFirewallRules**](NetworksApi.md#GetNetworkRoutersFirewallRules) | **Get** /api/networks/routers/{routerId}/firewall-rules | Get all Firewall Rules for Network Router
[**GetNetworkRoutersNats**](NetworksApi.md#GetNetworkRoutersNats) | **Get** /api/networks/routers/{routerId}/nats | Get all Network Router NATs for Network Router
[**GetNetworkRoutersRoutes**](NetworksApi.md#GetNetworkRoutersRoutes) | **Get** /api/networks/routers/{routerId}/routes | Get all Routes for Network Router
[**GetNetworkServerGroup**](NetworksApi.md#GetNetworkServerGroup) | **Get** /api/networks/servers/{serverId}/groups/{id} | Get a specific Network Server Group
[**GetNetworkSubnets**](NetworksApi.md#GetNetworkSubnets) | **Get** /api/networks/{id}/subnets | Get Subnets for a Network
[**GetNetworkTransportZone**](NetworksApi.md#GetNetworkTransportZone) | **Get** /api/networks/servers/{serverId}/scopes/{id} | Get a Specific Network Transport Zone
[**GetNetworkTransportZones**](NetworksApi.md#GetNetworkTransportZones) | **Get** /api/networks/servers/{serverId}/scopes | Get all Network Transport Zones for Network Server
[**GetNetworkType**](NetworksApi.md#GetNetworkType) | **Get** /api/network-types/{id} | Get a Specific Network Type
[**GetStaticRoute**](NetworksApi.md#GetStaticRoute) | **Get** /api/networks/{id}/routes/{routeId} | Get a Static Route for a Network
[**GetStaticRoutes**](NetworksApi.md#GetStaticRoutes) | **Get** /api/networks/{id}/routes | Get all Static Routes for a Network
[**GetSubnet**](NetworksApi.md#GetSubnet) | **Get** /api/subnets/{id} | Get a Specific Subnet
[**ListNetworkPoolServerTypes**](NetworksApi.md#ListNetworkPoolServerTypes) | **Get** /api/networks/pool-server-types | Get All Network Pool Server Types
[**ListNetworkPoolServers**](NetworksApi.md#ListNetworkPoolServers) | **Get** /api/networks/pool-servers | Get All Network Pool Servers
[**ListNetworkRouterTypes**](NetworksApi.md#ListNetworkRouterTypes) | **Get** /api/network-router-types | Get all Network Router Types
[**ListNetworkServerGroups**](NetworksApi.md#ListNetworkServerGroups) | **Get** /api/networks/servers/{serverId}/groups | Get all Network Server Groups for Network Server
[**ListNetworkServices**](NetworksApi.md#ListNetworkServices) | **Get** /api/networks/services | Get All Network Services
[**ListNetworkTypes**](NetworksApi.md#ListNetworkTypes) | **Get** /api/network-types | Network Types
[**ListNetworks**](NetworksApi.md#ListNetworks) | **Get** /api/networks | Get All Networks
[**ListSubnetTypes**](NetworksApi.md#ListSubnetTypes) | **Get** /api/subnet-types | Get All Subnet Types
[**ListSubnets**](NetworksApi.md#ListSubnets) | **Get** /api/subnets | Get All Subnets
[**RefreshNetworkServer**](NetworksApi.md#RefreshNetworkServer) | **Post** /api/networks/servers/{id}/refresh | Refresh a Network Server/Integration
[**ReleaseNetworkFloatingIp**](NetworksApi.md#ReleaseNetworkFloatingIp) | **Put** /api/networks/floating-ips/{id}/release | Release a Floating IP
[**UpdateNetwork**](NetworksApi.md#UpdateNetwork) | **Put** /api/networks/{id} | Update a Network
[**UpdateNetworkDhcpRelay**](NetworksApi.md#UpdateNetworkDhcpRelay) | **Put** /api/networks/servers/{serverId}/dhcp-relays/{id} | Update a Network DHCP Relay
[**UpdateNetworkDhcpServer**](NetworksApi.md#UpdateNetworkDhcpServer) | **Put** /api/networks/servers/{serverId}/dhcp-servers/{id} | Update a Network DHCP Server
[**UpdateNetworkDomain**](NetworksApi.md#UpdateNetworkDomain) | **Put** /api/networks/domains/{id} | Update a Network Domain
[**UpdateNetworkEdgeCluster**](NetworksApi.md#UpdateNetworkEdgeCluster) | **Put** /api/networks/servers/{serverId}/edge-clusters/{id} | Update a Network Edge Cluster
[**UpdateNetworkFirewallRule**](NetworksApi.md#UpdateNetworkFirewallRule) | **Put** /api/networks/servers/{serverId}/firewall-rules/{id} | Update a Network Firewall Rule
[**UpdateNetworkFirewallRuleGroup**](NetworksApi.md#UpdateNetworkFirewallRuleGroup) | **Put** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Update a Network Firewall Rule Group
[**UpdateNetworkGroup**](NetworksApi.md#UpdateNetworkGroup) | **Put** /api/networks/groups/{id} | Update a Network Group
[**UpdateNetworkPool**](NetworksApi.md#UpdateNetworkPool) | **Put** /api/networks/pools/{id} | Update a Network Pool
[**UpdateNetworkPoolServer**](NetworksApi.md#UpdateNetworkPoolServer) | **Put** /api/networks/pool-servers/{id} | Update a Network Pool Server
[**UpdateNetworkProxy**](NetworksApi.md#UpdateNetworkProxy) | **Put** /api/networks/proxies/{id} | Update a Network Proxy
[**UpdateNetworkRouter**](NetworksApi.md#UpdateNetworkRouter) | **Put** /api/networks/routers/{id} | Update a Network Router
[**UpdateNetworkRouterBgpNeighbor**](NetworksApi.md#UpdateNetworkRouterBgpNeighbor) | **Put** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Update Network Router BGP Neighbor
[**UpdateNetworkRouterFirewallRule**](NetworksApi.md#UpdateNetworkRouterFirewallRule) | **Put** /api/networks/routers/{routerId}/firewall-rules/{id} | Update a Network Router Firewall Rule
[**UpdateNetworkRouterFirewallRuleGroup**](NetworksApi.md#UpdateNetworkRouterFirewallRuleGroup) | **Put** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Update a Network Router Firewall Rule Group
[**UpdateNetworkRouterNat**](NetworksApi.md#UpdateNetworkRouterNat) | **Put** /api/networks/routers/{routerId}/nats/{id} | Update Network Router NAT
[**UpdateNetworkRouterPermissions**](NetworksApi.md#UpdateNetworkRouterPermissions) | **Put** /api/networks/routers/{routerId}/permissions | Update Network Router Permissions
[**UpdateNetworkServerGroup**](NetworksApi.md#UpdateNetworkServerGroup) | **Put** /api/networks/servers/{serverId}/groups/{id} | Update a Network Server Group
[**UpdateNetworkTransportZone**](NetworksApi.md#UpdateNetworkTransportZone) | **Put** /api/networks/servers/{serverId}/scopes/{id} | Update a Network Transport Zone
[**UpdateStaticRoute**](NetworksApi.md#UpdateStaticRoute) | **Put** /api/networks/{id}/routes/{routeId} | Update a Network Static Route
[**UpdateSubnet**](NetworksApi.md#UpdateSubnet) | **Put** /api/subnets/{id} | Update a Subnet



## CreateNetworkDhcpRelay

> SuccessId CreateNetworkDhcpRelay(ctx, serverId).InlineObject167(inlineObject167).Execute()

Create a Network DHCP Relay



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    inlineObject167 := *openapiclient.NewInlineObject167() // InlineObject167 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkDhcpRelay(context.Background(), serverId).InlineObject167(inlineObject167).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkDhcpRelay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkDhcpRelay`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject167** | [**InlineObject167**](InlineObject167.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkDhcpServer

> SuccessId CreateNetworkDhcpServer(ctx, serverId).InlineObject169(inlineObject169).Execute()

Create a Network DHCP Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    inlineObject169 := *openapiclient.NewInlineObject169() // InlineObject169 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkDhcpServer(context.Background(), serverId).InlineObject169(inlineObject169).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkDhcpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkDhcpServer`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject169** | [**InlineObject169**](InlineObject169.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkDomain

> InlineResponse200109 CreateNetworkDomain(ctx).InlineObject163(inlineObject163).Execute()

Create a Network Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject163 := *openapiclient.NewInlineObject163() // InlineObject163 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkDomain(context.Background()).InlineObject163(inlineObject163).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkDomain`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject163** | [**InlineObject163**](InlineObject163.md) |  | 

### Return type

[**InlineResponse200109**](inline_response_200_109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirewallRule

> SuccessId CreateNetworkFirewallRule(ctx, serverId).InlineObject172(inlineObject172).Execute()

Create a Network Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    inlineObject172 := *openapiclient.NewInlineObject172() // InlineObject172 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkFirewallRule(context.Background(), serverId).InlineObject172(inlineObject172).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirewallRule`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject172** | [**InlineObject172**](InlineObject172.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirewallRuleGroup

> SuccessId CreateNetworkFirewallRuleGroup(ctx, serverId).InlineObject174(inlineObject174).Execute()

Create a Network Firewall Rule Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    inlineObject174 := *openapiclient.NewInlineObject174() // InlineObject174 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkFirewallRuleGroup(context.Background(), serverId).InlineObject174(inlineObject174).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirewallRuleGroup`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject174** | [**InlineObject174**](InlineObject174.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkGroup

> SuccessId CreateNetworkGroup(ctx).InlineObject146(inlineObject146).Execute()

Create a Network Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject146 := *openapiclient.NewInlineObject146() // InlineObject146 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkGroup(context.Background()).InlineObject146(inlineObject146).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkGroup`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject146** | [**InlineObject146**](InlineObject146.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPool

> InlineResponse200106 CreateNetworkPool(ctx).InlineObject160(inlineObject160).Execute()

Create a Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject160 := *openapiclient.NewInlineObject160() // InlineObject160 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkPool(context.Background()).InlineObject160(inlineObject160).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPool`: InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject160** | [**InlineObject160**](InlineObject160.md) |  | 

### Return type

[**InlineResponse200106**](inline_response_200_106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPoolIp

> InlineResponse200107 CreateNetworkPoolIp(ctx, id).InlineObject162(inlineObject162).Execute()

Create a Network Pool IP Address



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject162 := *openapiclient.NewInlineObject162() // InlineObject162 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkPoolIp(context.Background(), id).InlineObject162(inlineObject162).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkPoolIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPoolIp`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkPoolIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject162** | [**InlineObject162**](InlineObject162.md) |  | 

### Return type

[**InlineResponse200107**](inline_response_200_107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPoolServer

> Model200Success CreateNetworkPoolServer(ctx).InlineObject180(inlineObject180).Execute()

Create a Network Pool Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject180 := *openapiclient.NewInlineObject180() // InlineObject180 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkPoolServer(context.Background()).InlineObject180(inlineObject180).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkPoolServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPoolServer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject180** | [**InlineObject180**](InlineObject180.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkProxy

> InlineResponse200110 CreateNetworkProxy(ctx).InlineObject165(inlineObject165).Execute()

Create a Network Proxy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject165 := *openapiclient.NewInlineObject165() // InlineObject165 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkProxy(context.Background()).InlineObject165(inlineObject165).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkProxy`: InlineResponse200110
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject165** | [**InlineObject165**](InlineObject165.md) |  | 

### Return type

[**InlineResponse200110**](inline_response_200_110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouter

> SuccessId CreateNetworkRouter(ctx).InlineObject148(inlineObject148).Execute()

Create a Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject148 := *openapiclient.NewInlineObject148() // InlineObject148 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkRouter(context.Background()).InlineObject148(inlineObject148).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkRouter`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkRouter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject148** | [**InlineObject148**](InlineObject148.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterBgpNeighbor

> SuccessId CreateNetworkRouterBgpNeighbor(ctx, routerId).InlineObject150(inlineObject150).Execute()

Create a Network Router BGP Neighbor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID
    inlineObject150 := *openapiclient.NewInlineObject150() // InlineObject150 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkRouterBgpNeighbor(context.Background(), routerId).InlineObject150(inlineObject150).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkRouterBgpNeighbor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkRouterBgpNeighbor`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject150** | [**InlineObject150**](InlineObject150.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterFirewallRule

> SuccessId CreateNetworkRouterFirewallRule(ctx, routerId).InlineObject152(inlineObject152).Execute()

Create a Network Router Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID
    inlineObject152 := *openapiclient.NewInlineObject152() // InlineObject152 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkRouterFirewallRule(context.Background(), routerId).InlineObject152(inlineObject152).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkRouterFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkRouterFirewallRule`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject152** | [**InlineObject152**](InlineObject152.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterFirewallRuleGroup

> SuccessId CreateNetworkRouterFirewallRuleGroup(ctx, routerId).InlineObject154(inlineObject154).Execute()

Create a Network Router Firewall Rule Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID
    inlineObject154 := *openapiclient.NewInlineObject154() // InlineObject154 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkRouterFirewallRuleGroup(context.Background(), routerId).InlineObject154(inlineObject154).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkRouterFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkRouterFirewallRuleGroup`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject154** | [**InlineObject154**](InlineObject154.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterNat

> SuccessId CreateNetworkRouterNat(ctx, routerId).InlineObject156(inlineObject156).Execute()

Create a Network Router NAT



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID
    inlineObject156 := *openapiclient.NewInlineObject156() // InlineObject156 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkRouterNat(context.Background(), routerId).InlineObject156(inlineObject156).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkRouterNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkRouterNat`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject156** | [**InlineObject156**](InlineObject156.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkRouterRoute

> SuccessId CreateNetworkRouterRoute(ctx, routerId).InlineObject159(inlineObject159).Execute()

Create a Network Router Route



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID
    inlineObject159 := *openapiclient.NewInlineObject159() // InlineObject159 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkRouterRoute(context.Background(), routerId).InlineObject159(inlineObject159).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkRouterRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkRouterRoute`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkRouterRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRouterRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject159** | [**InlineObject159**](InlineObject159.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkServerGroup

> SuccessId CreateNetworkServerGroup(ctx, serverId).InlineObject176(inlineObject176).Execute()

Create a Network Server Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    inlineObject176 := *openapiclient.NewInlineObject176() // InlineObject176 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkServerGroup(context.Background(), serverId).InlineObject176(inlineObject176).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkServerGroup`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject176** | [**InlineObject176**](InlineObject176.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkTransportZone

> SuccessId CreateNetworkTransportZone(ctx, serverId).InlineObject178(inlineObject178).Execute()

Create a Network Transport Zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    inlineObject178 := *openapiclient.NewInlineObject178() // InlineObject178 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworkTransportZone(context.Background(), serverId).InlineObject178(inlineObject178).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworkTransportZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkTransportZone`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject178** | [**InlineObject178**](InlineObject178.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworks

> map[string]interface{} CreateNetworks(ctx).InlineObject142(inlineObject142).Execute()

Create a Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject142 := *openapiclient.NewInlineObject142() // InlineObject142 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateNetworks(context.Background()).InlineObject142(inlineObject142).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject142** | [**InlineObject142**](InlineObject142.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStaticRoute

> SuccessId CreateStaticRoute(ctx, id).InlineObject144(inlineObject144).Execute()

Create a Network Static Route



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject144 := *openapiclient.NewInlineObject144() // InlineObject144 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateStaticRoute(context.Background(), id).InlineObject144(inlineObject144).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStaticRoute`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject144** | [**InlineObject144**](InlineObject144.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubnet

> InlineResponse200154 CreateSubnet(ctx).InlineObject244(inlineObject244).Execute()

Create a Subnet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject244 := *openapiclient.NewInlineObject244() // InlineObject244 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.CreateSubnet(context.Background()).InlineObject244(inlineObject244).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.CreateSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubnet`: InlineResponse200154
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject244** | [**InlineObject244**](InlineObject244.md) |  | 

### Return type

[**InlineResponse200154**](inline_response_200_154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> Model200Success DeleteNetwork(ctx, id).Execute()

Delete a Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetwork(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetwork`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDhcpRelay

> Model200Success DeleteNetworkDhcpRelay(ctx, id, serverId).Execute()

Delete a Network DHCP Relay



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkDhcpRelay(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkDhcpRelay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkDhcpRelay`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDhcpServer

> Model200Success DeleteNetworkDhcpServer(ctx, id, serverId).Execute()

Delete a Network DHCP Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkDhcpServer(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkDhcpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkDhcpServer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkDomain

> Model200Success DeleteNetworkDomain(ctx, id).Execute()

Delete a Network Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkDomain(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkDomain`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkFirewallRule

> Model200Success DeleteNetworkFirewallRule(ctx, id, serverId).Execute()

Delete a Network Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkFirewallRule(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkFirewallRule`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkFirewallRuleGroup

> Model200Success DeleteNetworkFirewallRuleGroup(ctx, id, serverId).Execute()

Delete a Network firewall rule group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkFirewallRuleGroup(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkFirewallRuleGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkGroup

> Model200Success DeleteNetworkGroup(ctx, id).Execute()

Delete a Network Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkGroup(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPool

> Model200Success DeleteNetworkPool(ctx, id).Execute()

Delete a Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkPool(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkPool`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPoolIp

> Model200Success DeleteNetworkPoolIp(ctx, networkPoolId, id).Execute()

Delete a host record associated with an IP Address for a Specific Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkPoolId := int64(1) // int64 | Morpheus Network Pool ID of the Object being referenced
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkPoolIp(context.Background(), networkPoolId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkPoolIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkPoolIp`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkPoolIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkPoolId** | **int64** | Morpheus Network Pool ID of the Object being referenced | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPoolServer

> Model200Success DeleteNetworkPoolServer(ctx, id).Execute()

Delete a Network Pool Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkPoolServer(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkPoolServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkPoolServer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkProxy

> Model200Success DeleteNetworkProxy(ctx, id).Execute()

Delete a Network Proxy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkProxy(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkProxy`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkRouter

> Model200Success DeleteNetworkRouter(ctx, id).Execute()

Delete a Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkRouter(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkRouter`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkRouterBgpNeighbor

> Model200Success DeleteNetworkRouterBgpNeighbor(ctx, id, routerId).Execute()

Delete a Network Router BGP Neighbor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkRouterBgpNeighbor(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkRouterBgpNeighbor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkRouterBgpNeighbor`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkRouterFirewallRule

> Model200Success DeleteNetworkRouterFirewallRule(ctx, id, routerId).Execute()

Delete a Network Router Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkRouterFirewallRule(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkRouterFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkRouterFirewallRule`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkRouterFirewallRuleGroup

> Model200Success DeleteNetworkRouterFirewallRuleGroup(ctx, id, routerId).Execute()

Delete a Network Router firewall rule group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkRouterFirewallRuleGroup(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkRouterFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkRouterFirewallRuleGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkRouterNat

> Model200Success DeleteNetworkRouterNat(ctx, id, routerId).Execute()

Delete a Network Router NAT



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkRouterNat(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkRouterNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkRouterNat`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkRouterRoute

> Model200Success DeleteNetworkRouterRoute(ctx, id, routerId).Execute()

Delete a Network Router Route



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkRouterRoute(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkRouterRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkRouterRoute`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkRouterRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRouterRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkServerGroup

> Model200Success DeleteNetworkServerGroup(ctx, id, serverId).Execute()

Delete a Network Server Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkServerGroup(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkServerGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkTransportZone

> Model200Success DeleteNetworkTransportZone(ctx, id, serverId).Execute()

Delete a Network Transport Zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteNetworkTransportZone(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteNetworkTransportZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkTransportZone`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStaticRoute

> Model200Success DeleteStaticRoute(ctx, id, routeId).Execute()

Delete a Network Static Route



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routeId := float32(4) // float32 | The ID of the route

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteStaticRoute(context.Background(), id, routeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStaticRoute`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routeId** | **float32** | The ID of the route | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnet

> Model200Success DeleteSubnet(ctx, id).Execute()

Delete a Subnet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.DeleteSubnet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.DeleteSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubnet`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.DeleteSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNetworkFloatingIps

> map[string]interface{} GetAllNetworkFloatingIps(ctx).Phrase(phrase).IpAddress(ipAddress).IpStatus(ipStatus).ZoneId(zoneId).ServerId(serverId).Execute()

Get All Floating IPs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    ipAddress := "10.32.23.188" // string | Filter by IP Address (optional)
    ipStatus := "ipStatus_example" // string | Filter by IP Status (optional)
    zoneId := int64(789) // int64 | Filter by Cloud ID (optional)
    serverId := int64(789) // int64 | Filter by Server ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetAllNetworkFloatingIps(context.Background()).Phrase(phrase).IpAddress(ipAddress).IpStatus(ipStatus).ZoneId(zoneId).ServerId(serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetAllNetworkFloatingIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNetworkFloatingIps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetAllNetworkFloatingIps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworkFloatingIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **ipAddress** | **string** | Filter by IP Address | 
 **ipStatus** | **string** | Filter by IP Status | 
 **zoneId** | **int64** | Filter by Cloud ID | 
 **serverId** | **int64** | Filter by Server ID | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetwork

> InlineResponse20087 GetNetwork(ctx, id).Execute()

Get a Specific Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetwork(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetwork`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20087**](inline_response_200_87.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpRelay

> InlineResponse200112 GetNetworkDhcpRelay(ctx, id, serverId).Execute()

Get a Specific Network DHCP Relay



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkDhcpRelay(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDhcpRelay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDhcpRelay`: InlineResponse200112
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200112**](inline_response_200_112.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpRelays

> map[string]interface{} GetNetworkDhcpRelays(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network DHCP Relays for Network Relay



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkDhcpRelays(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDhcpRelays``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDhcpRelays`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDhcpRelays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpRelaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpServer

> InlineResponse200113 GetNetworkDhcpServer(ctx, id, serverId).Execute()

Get a Specific Network DHCP Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkDhcpServer(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDhcpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDhcpServer`: InlineResponse200113
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDhcpServers

> map[string]interface{} GetNetworkDhcpServers(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network DHCP Servers for Network Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkDhcpServers(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDhcpServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDhcpServers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDhcpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDhcpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDomain

> InlineResponse200109 GetNetworkDomain(ctx, id).Execute()

Get a Specific Network Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkDomain(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDomain`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200109**](inline_response_200_109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDomains

> map[string]interface{} GetNetworkDomains(ctx).Name(name).Phrase(phrase).Execute()

Get all Network Domains



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkDomains(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkDomains`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEdgeCluster

> InlineResponse200114 GetNetworkEdgeCluster(ctx, id, serverId).Execute()

Get a Specific Network Edge Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkEdgeCluster(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkEdgeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkEdgeCluster`: InlineResponse200114
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200114**](inline_response_200_114.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEdgeClusters

> map[string]interface{} GetNetworkEdgeClusters(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Edge Clusters for Network Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkEdgeClusters(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkEdgeClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkEdgeClusters`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkEdgeClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEdgeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRule

> InlineResponse200115 GetNetworkFirewallRule(ctx, id, serverId).Execute()

Get a Specific Network Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkFirewallRule(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirewallRule`: InlineResponse200115
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRuleGroup

> InlineResponse200116 GetNetworkFirewallRuleGroup(ctx, id, serverId).Execute()

Get a Specific Network Firewall Rule Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkFirewallRuleGroup(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirewallRuleGroup`: InlineResponse200116
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRuleGroups

> map[string]interface{} GetNetworkFirewallRuleGroups(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Firewall Rule Groups for Network Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkFirewallRuleGroups(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirewallRuleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirewallRuleGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirewallRuleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRuleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirewallRules

> map[string]interface{} GetNetworkFirewallRules(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Firewall Rules for Network Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkFirewallRules(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloatingIp

> InlineResponse200108 GetNetworkFloatingIp(ctx, id).Execute()

Get a Specific Floating IP



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkFloatingIp(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkFloatingIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloatingIp`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkFloatingIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloatingIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200108**](inline_response_200_108.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroup

> InlineResponse20091 GetNetworkGroup(ctx, id).Execute()

Get a Specific Network Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkGroup(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroup`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20091**](inline_response_200_91.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroups

> InlineResponse20090 GetNetworkGroups(ctx).Execute()

Get all Network Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkGroups(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroups`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupsRequest struct via the builder pattern


### Return type

[**InlineResponse20090**](inline_response_200_90.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPool

> InlineResponse200106 GetNetworkPool(ctx, id).Execute()

Get a Specific Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkPool(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPool`: InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200106**](inline_response_200_106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolIp

> map[string]interface{} GetNetworkPoolIp(ctx, networkPoolId, id).Execute()

Get a Specific IP Address for a Specific Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkPoolId := int64(1) // int64 | Morpheus Network Pool ID of the Object being referenced
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkPoolIp(context.Background(), networkPoolId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPoolIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPoolIp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPoolIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkPoolId** | **int64** | Morpheus Network Pool ID of the Object being referenced | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolIps

> map[string]interface{} GetNetworkPoolIps(ctx, id).Execute()

Get all IP Addresses for a Specific Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkPoolIps(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPoolIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPoolIps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPoolIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolServer

> InlineResponse200120 GetNetworkPoolServer(ctx, id).Execute()

Get a Specific Network Pool Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkPoolServer(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPoolServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPoolServer`: InlineResponse200120
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200120**](inline_response_200_120.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolServerType

> InlineResponse200121 GetNetworkPoolServerType(ctx, id).Execute()

Retrieves a Specific Network Pool Server Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkPoolServerType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPoolServerType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPoolServerType`: InlineResponse200121
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPoolServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200121**](inline_response_200_121.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPools

> map[string]interface{} GetNetworkPools(ctx).Name(name).Phrase(phrase).Execute()

Get all Network Pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkPools(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPools`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkProxies

> map[string]interface{} GetNetworkProxies(ctx).Name(name).Phrase(phrase).Execute()

Get all Network Proxies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkProxies(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkProxies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProxies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkProxy

> InlineResponse200111 GetNetworkProxy(ctx, id).Execute()

Get a Specific Network Proxy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkProxy(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkProxy`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouter

> InlineResponse20095 GetNetworkRouter(ctx, id).Execute()

Get a Specific Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouter(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouter`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20095**](inline_response_200_95.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterBgpNeighbor

> InlineResponse20097 GetNetworkRouterBgpNeighbor(ctx, id, routerId).Execute()

Get a Network Router BGP Neighbor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterBgpNeighbor(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterBgpNeighbor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterBgpNeighbor`: InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20097**](inline_response_200_97.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterFirewallRule

> InlineResponse20099 GetNetworkRouterFirewallRule(ctx, id, routerId).Execute()

Get a Firewall Rule for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterFirewallRule(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterFirewallRule`: InlineResponse20099
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20099**](inline_response_200_99.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterFirewallRuleGroup

> InlineResponse200101 GetNetworkRouterFirewallRuleGroup(ctx, id, routerId).Execute()

Get a Firewall Rule Group for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterFirewallRuleGroup(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterFirewallRuleGroup`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200101**](inline_response_200_101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterFirewallRuleGroups

> InlineResponse200100 GetNetworkRouterFirewallRuleGroups(ctx, routerId).Execute()

Get all Network Firewall Rule Groups for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterFirewallRuleGroups(context.Background(), routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterFirewallRuleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterFirewallRuleGroups`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterFirewallRuleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterFirewallRuleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200100**](inline_response_200_100.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterNat

> InlineResponse200103 GetNetworkRouterNat(ctx, id, routerId).Execute()

Get a Network Router NAT



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterNat(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterNat`: InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200103**](inline_response_200_103.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterRoute

> InlineResponse200105 GetNetworkRouterRoute(ctx, id, routerId).Execute()

Get a Route for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterRoute(context.Background(), id, routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterRoute`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200105**](inline_response_200_105.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouterType

> InlineResponse20093 GetNetworkRouterType(ctx, id).Execute()

Get a Specific Network Router Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouterType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouterType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouterType`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouterType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRouterTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20093**](inline_response_200_93.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRouters

> InlineResponse20094 GetNetworkRouters(ctx).Execute()

Get all Network Routers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRouters(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRouters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRouters`: InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRouters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersRequest struct via the builder pattern


### Return type

[**InlineResponse20094**](inline_response_200_94.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersBgpNeighbors

> InlineResponse20096 GetNetworkRoutersBgpNeighbors(ctx, routerId).Execute()

Get all BGP Neighbors for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRoutersBgpNeighbors(context.Background(), routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRoutersBgpNeighbors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRoutersBgpNeighbors`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRoutersBgpNeighbors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersBgpNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20096**](inline_response_200_96.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersFirewallRules

> InlineResponse20098 GetNetworkRoutersFirewallRules(ctx, routerId).Execute()

Get all Firewall Rules for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRoutersFirewallRules(context.Background(), routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRoutersFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRoutersFirewallRules`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRoutersFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20098**](inline_response_200_98.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersNats

> InlineResponse200102 GetNetworkRoutersNats(ctx, routerId).Execute()

Get all Network Router NATs for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRoutersNats(context.Background(), routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRoutersNats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRoutersNats`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRoutersNats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersNatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200102**](inline_response_200_102.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkRoutersRoutes

> InlineResponse200104 GetNetworkRoutersRoutes(ctx, routerId).Execute()

Get all Routes for Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkRoutersRoutes(context.Background(), routerId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkRoutersRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkRoutersRoutes`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkRoutersRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRoutersRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200104**](inline_response_200_104.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkServerGroup

> InlineResponse200117 GetNetworkServerGroup(ctx, serverId, id).Execute()

Get a specific Network Server Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkServerGroup(context.Background(), serverId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkServerGroup`: InlineResponse200117
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSubnets

> map[string]interface{} GetNetworkSubnets(ctx, id).Execute()

Get Subnets for a Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkSubnets(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSubnets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTransportZone

> InlineResponse200118 GetNetworkTransportZone(ctx, id, serverId).Execute()

Get a Specific Network Transport Zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkTransportZone(context.Background(), id, serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTransportZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTransportZone`: InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTransportZones

> map[string]interface{} GetNetworkTransportZones(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Transport Zones for Network Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkTransportZones(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkTransportZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTransportZones`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkTransportZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTransportZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkType

> InlineResponse20086 GetNetworkType(ctx, id).Execute()

Get a Specific Network Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetNetworkType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetNetworkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkType`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetNetworkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20086**](inline_response_200_86.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticRoute

> InlineResponse20089 GetStaticRoute(ctx, id, routeId).Execute()

Get a Static Route for a Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routeId := float32(4) // float32 | The ID of the route

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetStaticRoute(context.Background(), id, routeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaticRoute`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routeId** | **float32** | The ID of the route | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20089**](inline_response_200_89.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticRoutes

> InlineResponse20088 GetStaticRoutes(ctx, id).Execute()

Get all Static Routes for a Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetStaticRoutes(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaticRoutes`: InlineResponse20088
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20088**](inline_response_200_88.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubnet

> InlineResponse200154 GetSubnet(ctx, id).Execute()

Get a Specific Subnet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.GetSubnet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.GetSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubnet`: InlineResponse200154
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.GetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200154**](inline_response_200_154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkPoolServerTypes

> map[string]interface{} ListNetworkPoolServerTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Execute()

Get All Network Pool Server Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworkPoolServerTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Code(code).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworkPoolServerTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkPoolServerTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworkPoolServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkPoolServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkPoolServers

> map[string]interface{} ListNetworkPoolServers(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()

Get All Network Pool Servers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworkPoolServers(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworkPoolServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkPoolServers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworkPoolServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkPoolServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkRouterTypes

> InlineResponse20092 ListNetworkRouterTypes(ctx).Execute()

Get all Network Router Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworkRouterTypes(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworkRouterTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkRouterTypes`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworkRouterTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkRouterTypesRequest struct via the builder pattern


### Return type

[**InlineResponse20092**](inline_response_200_92.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkServerGroups

> map[string]interface{} ListNetworkServerGroups(ctx, serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()

Get all Network Server Groups for Network Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serverId := float32(4) // float32 | Server ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworkServerGroups(context.Background(), serverId).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworkServerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkServerGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworkServerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkServices

> InlineResponse200119 ListNetworkServices(ctx).Name(name).Phrase(phrase).Execute()

Get All Network Services



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworkServices(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworkServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkServices`: InlineResponse200119
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworkServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

[**InlineResponse200119**](inline_response_200_119.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkTypes

> map[string]interface{} ListNetworkTypes(ctx).Name(name).Code(code).Phrase(phrase).Execute()

Network Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworkTypes(context.Background()).Name(name).Code(code).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworkTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworkTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworks

> map[string]interface{} ListNetworks(ctx).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).Execute()

Get All Networks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListNetworks(context.Background()).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnetTypes

> map[string]interface{} ListSubnetTypes(ctx).Name(name).Phrase(phrase).Execute()

Get All Subnet Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListSubnetTypes(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListSubnetTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubnetTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListSubnetTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnets

> map[string]interface{} ListSubnets(ctx).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).Execute()

Get All Subnets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ListSubnets(context.Background()).Name(name).Phrase(phrase).Labels(labels).AllLabels(allLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ListSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubnets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ListSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshNetworkServer

> SuccessId RefreshNetworkServer(ctx, id).Execute()

Refresh a Network Server/Integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.RefreshNetworkServer(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.RefreshNetworkServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshNetworkServer`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.RefreshNetworkServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshNetworkServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseNetworkFloatingIp

> Model200Success ReleaseNetworkFloatingIp(ctx, id).Execute()

Release a Floating IP



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.ReleaseNetworkFloatingIp(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.ReleaseNetworkFloatingIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseNetworkFloatingIp`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.ReleaseNetworkFloatingIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseNetworkFloatingIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> map[string]interface{} UpdateNetwork(ctx, id).InlineObject143(inlineObject143).Execute()

Update a Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject143 := *openapiclient.NewInlineObject143() // InlineObject143 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetwork(context.Background(), id).InlineObject143(inlineObject143).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetwork`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject143** | [**InlineObject143**](InlineObject143.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDhcpRelay

> Model200Success UpdateNetworkDhcpRelay(ctx, id, serverId).InlineObject168(inlineObject168).Execute()

Update a Network DHCP Relay



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject168 := *openapiclient.NewInlineObject168() // InlineObject168 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkDhcpRelay(context.Background(), id, serverId).InlineObject168(inlineObject168).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkDhcpRelay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkDhcpRelay`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkDhcpRelay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDhcpRelayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject168** | [**InlineObject168**](InlineObject168.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDhcpServer

> Model200Success UpdateNetworkDhcpServer(ctx, id, serverId).InlineObject170(inlineObject170).Execute()

Update a Network DHCP Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject170 := *openapiclient.NewInlineObject170() // InlineObject170 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkDhcpServer(context.Background(), id, serverId).InlineObject170(inlineObject170).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkDhcpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkDhcpServer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkDhcpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDhcpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject170** | [**InlineObject170**](InlineObject170.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDomain

> InlineResponse200109 UpdateNetworkDomain(ctx, id).InlineObject164(inlineObject164).Execute()

Update a Network Domain



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject164 := *openapiclient.NewInlineObject164() // InlineObject164 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkDomain(context.Background(), id).InlineObject164(inlineObject164).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkDomain`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject164** | [**InlineObject164**](InlineObject164.md) |  | 

### Return type

[**InlineResponse200109**](inline_response_200_109.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkEdgeCluster

> Model200Success UpdateNetworkEdgeCluster(ctx, id, serverId).InlineObject171(inlineObject171).Execute()

Update a Network Edge Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject171 := *openapiclient.NewInlineObject171() // InlineObject171 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkEdgeCluster(context.Background(), id, serverId).InlineObject171(inlineObject171).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkEdgeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkEdgeCluster`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject171** | [**InlineObject171**](InlineObject171.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirewallRule

> Model200Success UpdateNetworkFirewallRule(ctx, id, serverId).InlineObject173(inlineObject173).Execute()

Update a Network Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject173 := *openapiclient.NewInlineObject173() // InlineObject173 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkFirewallRule(context.Background(), id, serverId).InlineObject173(inlineObject173).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirewallRule`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject173** | [**InlineObject173**](InlineObject173.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirewallRuleGroup

> Model200Success UpdateNetworkFirewallRuleGroup(ctx, id, serverId).InlineObject175(inlineObject175).Execute()

Update a Network Firewall Rule Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject175 := *openapiclient.NewInlineObject175() // InlineObject175 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkFirewallRuleGroup(context.Background(), id, serverId).InlineObject175(inlineObject175).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirewallRuleGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject175** | [**InlineObject175**](InlineObject175.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkGroup

> Model200Success UpdateNetworkGroup(ctx, id).InlineObject147(inlineObject147).Execute()

Update a Network Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject147 := *openapiclient.NewInlineObject147() // InlineObject147 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkGroup(context.Background(), id).InlineObject147(inlineObject147).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject147** | [**InlineObject147**](InlineObject147.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkPool

> InlineResponse200106 UpdateNetworkPool(ctx, id).InlineObject161(inlineObject161).Execute()

Update a Network Pool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject161 := *openapiclient.NewInlineObject161() // InlineObject161 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkPool(context.Background(), id).InlineObject161(inlineObject161).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkPool`: InlineResponse200106
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject161** | [**InlineObject161**](InlineObject161.md) |  | 

### Return type

[**InlineResponse200106**](inline_response_200_106.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkPoolServer

> Model200Success UpdateNetworkPoolServer(ctx, id).InlineObject181(inlineObject181).Execute()

Update a Network Pool Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject181 := *openapiclient.NewInlineObject181() // InlineObject181 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkPoolServer(context.Background(), id).InlineObject181(inlineObject181).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkPoolServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkPoolServer`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkPoolServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkPoolServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject181** | [**InlineObject181**](InlineObject181.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkProxy

> Model200Success UpdateNetworkProxy(ctx, id).InlineObject166(inlineObject166).Execute()

Update a Network Proxy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject166 := *openapiclient.NewInlineObject166() // InlineObject166 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkProxy(context.Background(), id).InlineObject166(inlineObject166).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkProxy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkProxy`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject166** | [**InlineObject166**](InlineObject166.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkRouter

> Model200Success UpdateNetworkRouter(ctx, id).InlineObject149(inlineObject149).Execute()

Update a Network Router



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject149 := *openapiclient.NewInlineObject149() // InlineObject149 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkRouter(context.Background(), id).InlineObject149(inlineObject149).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkRouter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkRouter`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkRouter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject149** | [**InlineObject149**](InlineObject149.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkRouterBgpNeighbor

> Model200Success UpdateNetworkRouterBgpNeighbor(ctx, id, routerId).InlineObject151(inlineObject151).Execute()

Update Network Router BGP Neighbor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID
    inlineObject151 := *openapiclient.NewInlineObject151() // InlineObject151 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkRouterBgpNeighbor(context.Background(), id, routerId).InlineObject151(inlineObject151).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkRouterBgpNeighbor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkRouterBgpNeighbor`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkRouterBgpNeighbor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterBgpNeighborRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject151** | [**InlineObject151**](InlineObject151.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkRouterFirewallRule

> Model200Success UpdateNetworkRouterFirewallRule(ctx, id, routerId).InlineObject153(inlineObject153).Execute()

Update a Network Router Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID
    inlineObject153 := *openapiclient.NewInlineObject153() // InlineObject153 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkRouterFirewallRule(context.Background(), id, routerId).InlineObject153(inlineObject153).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkRouterFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkRouterFirewallRule`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkRouterFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject153** | [**InlineObject153**](InlineObject153.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkRouterFirewallRuleGroup

> Model200Success UpdateNetworkRouterFirewallRuleGroup(ctx, id, routerId).InlineObject155(inlineObject155).Execute()

Update a Network Router Firewall Rule Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID
    inlineObject155 := *openapiclient.NewInlineObject155() // InlineObject155 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkRouterFirewallRuleGroup(context.Background(), id, routerId).InlineObject155(inlineObject155).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkRouterFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkRouterFirewallRuleGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkRouterFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject155** | [**InlineObject155**](InlineObject155.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkRouterNat

> Model200Success UpdateNetworkRouterNat(ctx, id, routerId).InlineObject157(inlineObject157).Execute()

Update Network Router NAT



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routerId := float32(4) // float32 | Router ID
    inlineObject157 := *openapiclient.NewInlineObject157() // InlineObject157 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkRouterNat(context.Background(), id, routerId).InlineObject157(inlineObject157).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkRouterNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkRouterNat`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkRouterNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject157** | [**InlineObject157**](InlineObject157.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkRouterPermissions

> SuccessId UpdateNetworkRouterPermissions(ctx, routerId).InlineObject158(inlineObject158).Execute()

Update Network Router Permissions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    routerId := float32(4) // float32 | Router ID
    inlineObject158 := *openapiclient.NewInlineObject158() // InlineObject158 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkRouterPermissions(context.Background(), routerId).InlineObject158(inlineObject158).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkRouterPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkRouterPermissions`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkRouterPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerId** | **float32** | Router ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRouterPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject158** | [**InlineObject158**](InlineObject158.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkServerGroup

> Model200Success UpdateNetworkServerGroup(ctx, id, serverId).InlineObject177(inlineObject177).Execute()

Update a Network Server Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject177 := *openapiclient.NewInlineObject177() // InlineObject177 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkServerGroup(context.Background(), id, serverId).InlineObject177(inlineObject177).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkServerGroup`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject177** | [**InlineObject177**](InlineObject177.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkTransportZone

> Model200Success UpdateNetworkTransportZone(ctx, id, serverId).InlineObject179(inlineObject179).Execute()

Update a Network Transport Zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    serverId := float32(4) // float32 | Server ID
    inlineObject179 := *openapiclient.NewInlineObject179() // InlineObject179 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateNetworkTransportZone(context.Background(), id, serverId).InlineObject179(inlineObject179).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateNetworkTransportZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkTransportZone`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateNetworkTransportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**serverId** | **float32** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkTransportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject179** | [**InlineObject179**](InlineObject179.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStaticRoute

> SuccessId UpdateStaticRoute(ctx, id, routeId).InlineObject145(inlineObject145).Execute()

Update a Network Static Route



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    routeId := float32(4) // float32 | The ID of the route
    inlineObject145 := *openapiclient.NewInlineObject145() // InlineObject145 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateStaticRoute(context.Background(), id, routeId).InlineObject145(inlineObject145).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStaticRoute`: SuccessId
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**routeId** | **float32** | The ID of the route | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject145** | [**InlineObject145**](InlineObject145.md) |  | 

### Return type

[**SuccessId**](successId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubnet

> InlineResponse200154 UpdateSubnet(ctx, id).InlineObject245(inlineObject245).Execute()

Update a Subnet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject245 := *openapiclient.NewInlineObject245() // InlineObject245 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworksApi.UpdateSubnet(context.Background(), id).InlineObject245(inlineObject245).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksApi.UpdateSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubnet`: InlineResponse200154
    fmt.Fprintf(os.Stdout, "Response from `NetworksApi.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject245** | [**InlineObject245**](InlineObject245.md) |  | 

### Return type

[**InlineResponse200154**](inline_response_200_154.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

