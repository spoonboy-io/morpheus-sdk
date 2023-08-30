# openapi_client.NetworksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_network_dhcp_relay**](NetworksApi.md#create_network_dhcp_relay) | **POST** /api/networks/servers/{serverId}/dhcp-relays | Create a Network DHCP Relay
[**create_network_dhcp_server**](NetworksApi.md#create_network_dhcp_server) | **POST** /api/networks/servers/{serverId}/dhcp-servers | Create a Network DHCP Server
[**create_network_domain**](NetworksApi.md#create_network_domain) | **POST** /api/networks/domains | Create a Network Domain
[**create_network_firewall_rule**](NetworksApi.md#create_network_firewall_rule) | **POST** /api/networks/servers/{serverId}/firewall-rules | Create a Network Firewall Rule
[**create_network_firewall_rule_group**](NetworksApi.md#create_network_firewall_rule_group) | **POST** /api/networks/servers/{serverId}/firewall-rule-groups | Create a Network Firewall Rule Group
[**create_network_group**](NetworksApi.md#create_network_group) | **POST** /api/networks/groups | Create a Network Group
[**create_network_pool**](NetworksApi.md#create_network_pool) | **POST** /api/networks/pools | Create a Network Pool
[**create_network_pool_ip**](NetworksApi.md#create_network_pool_ip) | **POST** /api/networks/pools/{id}/ips | Create a Network Pool IP Address
[**create_network_pool_server**](NetworksApi.md#create_network_pool_server) | **POST** /api/networks/pool-servers | Create a Network Pool Server
[**create_network_proxy**](NetworksApi.md#create_network_proxy) | **POST** /api/networks/proxies | Create a Network Proxy
[**create_network_router**](NetworksApi.md#create_network_router) | **POST** /api/networks/routers | Create a Network Router
[**create_network_router_bgp_neighbor**](NetworksApi.md#create_network_router_bgp_neighbor) | **POST** /api/networks/routers/{routerId}/bgp-neighbors | Create a Network Router BGP Neighbor
[**create_network_router_firewall_rule**](NetworksApi.md#create_network_router_firewall_rule) | **POST** /api/networks/routers/{routerId}/firewall-rules | Create a Network Router Firewall Rule
[**create_network_router_firewall_rule_group**](NetworksApi.md#create_network_router_firewall_rule_group) | **POST** /api/networks/routers/{routerId}/firewall-rule-groups | Create a Network Router Firewall Rule Group
[**create_network_router_nat**](NetworksApi.md#create_network_router_nat) | **POST** /api/networks/routers/{routerId}/nats | Create a Network Router NAT
[**create_network_router_route**](NetworksApi.md#create_network_router_route) | **POST** /api/networks/routers/{routerId}/routes | Create a Network Router Route
[**create_network_server_group**](NetworksApi.md#create_network_server_group) | **POST** /api/networks/servers/{serverId}/groups | Create a Network Server Group
[**create_network_transport_zone**](NetworksApi.md#create_network_transport_zone) | **POST** /api/networks/servers/{serverId}/scopes | Create a Network Transport Zone
[**create_networks**](NetworksApi.md#create_networks) | **POST** /api/networks | Create a Network
[**create_static_route**](NetworksApi.md#create_static_route) | **PUT** /api/networks/{id}/routes | Create a Network Static Route
[**create_subnet**](NetworksApi.md#create_subnet) | **POST** /api/subnets | Create a Subnet
[**delete_network**](NetworksApi.md#delete_network) | **DELETE** /api/networks/{id} | Delete a Network
[**delete_network_dhcp_relay**](NetworksApi.md#delete_network_dhcp_relay) | **DELETE** /api/networks/servers/{serverId}/dhcp-relays/{id} | Delete a Network DHCP Relay
[**delete_network_dhcp_server**](NetworksApi.md#delete_network_dhcp_server) | **DELETE** /api/networks/servers/{serverId}/dhcp-servers/{id} | Delete a Network DHCP Server
[**delete_network_domain**](NetworksApi.md#delete_network_domain) | **DELETE** /api/networks/domains/{id} | Delete a Network Domain
[**delete_network_firewall_rule**](NetworksApi.md#delete_network_firewall_rule) | **DELETE** /api/networks/servers/{serverId}/firewall-rules/{id} | Delete a Network Firewall Rule
[**delete_network_firewall_rule_group**](NetworksApi.md#delete_network_firewall_rule_group) | **DELETE** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Delete a Network firewall rule group
[**delete_network_group**](NetworksApi.md#delete_network_group) | **DELETE** /api/networks/groups/{id} | Delete a Network Group
[**delete_network_pool**](NetworksApi.md#delete_network_pool) | **DELETE** /api/networks/pools/{id} | Delete a Network Pool
[**delete_network_pool_ip**](NetworksApi.md#delete_network_pool_ip) | **DELETE** /api/networks/pools/{networkPoolId}/ips/{id} | Delete a host record associated with an IP Address for a Specific Network Pool
[**delete_network_pool_server**](NetworksApi.md#delete_network_pool_server) | **DELETE** /api/networks/pool-servers/{id} | Delete a Network Pool Server
[**delete_network_proxy**](NetworksApi.md#delete_network_proxy) | **DELETE** /api/networks/proxies/{id} | Delete a Network Proxy
[**delete_network_router**](NetworksApi.md#delete_network_router) | **DELETE** /api/networks/routers/{id} | Delete a Network Router
[**delete_network_router_bgp_neighbor**](NetworksApi.md#delete_network_router_bgp_neighbor) | **DELETE** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Delete a Network Router BGP Neighbor
[**delete_network_router_firewall_rule**](NetworksApi.md#delete_network_router_firewall_rule) | **DELETE** /api/networks/routers/{routerId}/firewall-rules/{id} | Delete a Network Router Firewall Rule
[**delete_network_router_firewall_rule_group**](NetworksApi.md#delete_network_router_firewall_rule_group) | **DELETE** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Delete a Network Router firewall rule group
[**delete_network_router_nat**](NetworksApi.md#delete_network_router_nat) | **DELETE** /api/networks/routers/{routerId}/nats/{id} | Delete a Network Router NAT
[**delete_network_router_route**](NetworksApi.md#delete_network_router_route) | **DELETE** /api/networks/routers/{routerId}/routes/{id} | Delete a Network Router Route
[**delete_network_server_group**](NetworksApi.md#delete_network_server_group) | **DELETE** /api/networks/servers/{serverId}/groups/{id} | Delete a Network Server Group
[**delete_network_transport_zone**](NetworksApi.md#delete_network_transport_zone) | **DELETE** /api/networks/servers/{serverId}/scopes/{id} | Delete a Network Transport Zone
[**delete_static_route**](NetworksApi.md#delete_static_route) | **DELETE** /api/networks/{id}/routes/{routeId} | Delete a Network Static Route
[**delete_subnet**](NetworksApi.md#delete_subnet) | **DELETE** /api/subnets/{id} | Delete a Subnet
[**get_all_network_floating_ips**](NetworksApi.md#get_all_network_floating_ips) | **GET** /api/networks/floating-ips | Get All Floating IPs
[**get_network**](NetworksApi.md#get_network) | **GET** /api/networks/{id} | Get a Specific Network
[**get_network_dhcp_relay**](NetworksApi.md#get_network_dhcp_relay) | **GET** /api/networks/servers/{serverId}/dhcp-relays/{id} | Get a Specific Network DHCP Relay
[**get_network_dhcp_relays**](NetworksApi.md#get_network_dhcp_relays) | **GET** /api/networks/servers/{serverId}/dhcp-relays | Get all Network DHCP Relays for Network Relay
[**get_network_dhcp_server**](NetworksApi.md#get_network_dhcp_server) | **GET** /api/networks/servers/{serverId}/dhcp-servers/{id} | Get a Specific Network DHCP Server
[**get_network_dhcp_servers**](NetworksApi.md#get_network_dhcp_servers) | **GET** /api/networks/servers/{serverId}/dhcp-servers | Get all Network DHCP Servers for Network Server
[**get_network_domain**](NetworksApi.md#get_network_domain) | **GET** /api/networks/domains/{id} | Get a Specific Network Domain
[**get_network_domains**](NetworksApi.md#get_network_domains) | **GET** /api/networks/domains | Get all Network Domains
[**get_network_edge_cluster**](NetworksApi.md#get_network_edge_cluster) | **GET** /api/networks/servers/{serverId}/edge-clusters/{id} | Get a Specific Network Edge Cluster
[**get_network_edge_clusters**](NetworksApi.md#get_network_edge_clusters) | **GET** /api/networks/servers/{serverId}/edge-clusters | Get all Network Edge Clusters for Network Server
[**get_network_firewall_rule**](NetworksApi.md#get_network_firewall_rule) | **GET** /api/networks/servers/{serverId}/firewall-rules/{id} | Get a Specific Network Firewall Rule
[**get_network_firewall_rule_group**](NetworksApi.md#get_network_firewall_rule_group) | **GET** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Get a Specific Network Firewall Rule Group
[**get_network_firewall_rule_groups**](NetworksApi.md#get_network_firewall_rule_groups) | **GET** /api/networks/servers/{serverId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Server
[**get_network_firewall_rules**](NetworksApi.md#get_network_firewall_rules) | **GET** /api/networks/servers/{serverId}/firewall-rules | Get all Network Firewall Rules for Network Server
[**get_network_floating_ip**](NetworksApi.md#get_network_floating_ip) | **GET** /api/networks/floating-ips/{id} | Get a Specific Floating IP
[**get_network_group**](NetworksApi.md#get_network_group) | **GET** /api/networks/groups/{id} | Get a Specific Network Group
[**get_network_groups**](NetworksApi.md#get_network_groups) | **GET** /api/networks/groups | Get all Network Groups
[**get_network_pool**](NetworksApi.md#get_network_pool) | **GET** /api/networks/pools/{id} | Get a Specific Network Pool
[**get_network_pool_ip**](NetworksApi.md#get_network_pool_ip) | **GET** /api/networks/pools/{networkPoolId}/ips/{id} | Get a Specific IP Address for a Specific Network Pool
[**get_network_pool_ips**](NetworksApi.md#get_network_pool_ips) | **GET** /api/networks/pools/{id}/ips | Get all IP Addresses for a Specific Network Pool
[**get_network_pool_server**](NetworksApi.md#get_network_pool_server) | **GET** /api/networks/pool-servers/{id} | Get a Specific Network Pool Server
[**get_network_pool_server_type**](NetworksApi.md#get_network_pool_server_type) | **GET** /api/networks/pool-server-types/{id} | Retrieves a Specific Network Pool Server Type
[**get_network_pools**](NetworksApi.md#get_network_pools) | **GET** /api/networks/pools | Get all Network Pools
[**get_network_proxies**](NetworksApi.md#get_network_proxies) | **GET** /api/networks/proxies | Get all Network Proxies
[**get_network_proxy**](NetworksApi.md#get_network_proxy) | **GET** /api/networks/proxies/{id} | Get a Specific Network Proxy
[**get_network_router**](NetworksApi.md#get_network_router) | **GET** /api/networks/routers/{id} | Get a Specific Network Router
[**get_network_router_bgp_neighbor**](NetworksApi.md#get_network_router_bgp_neighbor) | **GET** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Get a Network Router BGP Neighbor
[**get_network_router_firewall_rule**](NetworksApi.md#get_network_router_firewall_rule) | **GET** /api/networks/routers/{routerId}/firewall-rules/{id} | Get a Firewall Rule for Network Router
[**get_network_router_firewall_rule_group**](NetworksApi.md#get_network_router_firewall_rule_group) | **GET** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Get a Firewall Rule Group for Network Router
[**get_network_router_firewall_rule_groups**](NetworksApi.md#get_network_router_firewall_rule_groups) | **GET** /api/networks/routers/{routerId}/firewall-rule-groups | Get all Network Firewall Rule Groups for Network Router
[**get_network_router_nat**](NetworksApi.md#get_network_router_nat) | **GET** /api/networks/routers/{routerId}/nats/{id} | Get a Network Router NAT
[**get_network_router_route**](NetworksApi.md#get_network_router_route) | **GET** /api/networks/routers/{routerId}/routes/{id} | Get a Route for Network Router
[**get_network_router_type**](NetworksApi.md#get_network_router_type) | **GET** /api/network-router-types/{id} | Get a Specific Network Router Type
[**get_network_routers**](NetworksApi.md#get_network_routers) | **GET** /api/networks/routers | Get all Network Routers
[**get_network_routers_bgp_neighbors**](NetworksApi.md#get_network_routers_bgp_neighbors) | **GET** /api/networks/routers/{routerId}/bgp-neighbors | Get all BGP Neighbors for Network Router
[**get_network_routers_firewall_rules**](NetworksApi.md#get_network_routers_firewall_rules) | **GET** /api/networks/routers/{routerId}/firewall-rules | Get all Firewall Rules for Network Router
[**get_network_routers_nats**](NetworksApi.md#get_network_routers_nats) | **GET** /api/networks/routers/{routerId}/nats | Get all Network Router NATs for Network Router
[**get_network_routers_routes**](NetworksApi.md#get_network_routers_routes) | **GET** /api/networks/routers/{routerId}/routes | Get all Routes for Network Router
[**get_network_server_group**](NetworksApi.md#get_network_server_group) | **GET** /api/networks/servers/{serverId}/groups/{id} | Get a specific Network Server Group
[**get_network_subnets**](NetworksApi.md#get_network_subnets) | **GET** /api/networks/{id}/subnets | Get Subnets for a Network
[**get_network_transport_zone**](NetworksApi.md#get_network_transport_zone) | **GET** /api/networks/servers/{serverId}/scopes/{id} | Get a Specific Network Transport Zone
[**get_network_transport_zones**](NetworksApi.md#get_network_transport_zones) | **GET** /api/networks/servers/{serverId}/scopes | Get all Network Transport Zones for Network Server
[**get_network_type**](NetworksApi.md#get_network_type) | **GET** /api/network-types/{id} | Get a Specific Network Type
[**get_static_route**](NetworksApi.md#get_static_route) | **GET** /api/networks/{id}/routes/{routeId} | Get a Static Route for a Network
[**get_static_routes**](NetworksApi.md#get_static_routes) | **GET** /api/networks/{id}/routes | Get all Static Routes for a Network
[**get_subnet**](NetworksApi.md#get_subnet) | **GET** /api/subnets/{id} | Get a Specific Subnet
[**list_network_pool_server_types**](NetworksApi.md#list_network_pool_server_types) | **GET** /api/networks/pool-server-types | Get All Network Pool Server Types
[**list_network_pool_servers**](NetworksApi.md#list_network_pool_servers) | **GET** /api/networks/pool-servers | Get All Network Pool Servers
[**list_network_router_types**](NetworksApi.md#list_network_router_types) | **GET** /api/network-router-types | Get all Network Router Types
[**list_network_server_groups**](NetworksApi.md#list_network_server_groups) | **GET** /api/networks/servers/{serverId}/groups | Get all Network Server Groups for Network Server
[**list_network_services**](NetworksApi.md#list_network_services) | **GET** /api/networks/services | Get All Network Services
[**list_network_types**](NetworksApi.md#list_network_types) | **GET** /api/network-types | Network Types
[**list_networks**](NetworksApi.md#list_networks) | **GET** /api/networks | Get All Networks
[**list_subnet_types**](NetworksApi.md#list_subnet_types) | **GET** /api/subnet-types | Get All Subnet Types
[**list_subnets**](NetworksApi.md#list_subnets) | **GET** /api/subnets | Get All Subnets
[**refresh_network_server**](NetworksApi.md#refresh_network_server) | **POST** /api/networks/servers/{id}/refresh | Refresh a Network Server/Integration
[**release_network_floating_ip**](NetworksApi.md#release_network_floating_ip) | **PUT** /api/networks/floating-ips/{id}/release | Release a Floating IP
[**update_network**](NetworksApi.md#update_network) | **PUT** /api/networks/{id} | Update a Network
[**update_network_dhcp_relay**](NetworksApi.md#update_network_dhcp_relay) | **PUT** /api/networks/servers/{serverId}/dhcp-relays/{id} | Update a Network DHCP Relay
[**update_network_dhcp_server**](NetworksApi.md#update_network_dhcp_server) | **PUT** /api/networks/servers/{serverId}/dhcp-servers/{id} | Update a Network DHCP Server
[**update_network_domain**](NetworksApi.md#update_network_domain) | **PUT** /api/networks/domains/{id} | Update a Network Domain
[**update_network_edge_cluster**](NetworksApi.md#update_network_edge_cluster) | **PUT** /api/networks/servers/{serverId}/edge-clusters/{id} | Update a Network Edge Cluster
[**update_network_firewall_rule**](NetworksApi.md#update_network_firewall_rule) | **PUT** /api/networks/servers/{serverId}/firewall-rules/{id} | Update a Network Firewall Rule
[**update_network_firewall_rule_group**](NetworksApi.md#update_network_firewall_rule_group) | **PUT** /api/networks/servers/{serverId}/firewall-rule-groups/{id} | Update a Network Firewall Rule Group
[**update_network_group**](NetworksApi.md#update_network_group) | **PUT** /api/networks/groups/{id} | Update a Network Group
[**update_network_pool**](NetworksApi.md#update_network_pool) | **PUT** /api/networks/pools/{id} | Update a Network Pool
[**update_network_pool_server**](NetworksApi.md#update_network_pool_server) | **PUT** /api/networks/pool-servers/{id} | Update a Network Pool Server
[**update_network_proxy**](NetworksApi.md#update_network_proxy) | **PUT** /api/networks/proxies/{id} | Update a Network Proxy
[**update_network_router**](NetworksApi.md#update_network_router) | **PUT** /api/networks/routers/{id} | Update a Network Router
[**update_network_router_bgp_neighbor**](NetworksApi.md#update_network_router_bgp_neighbor) | **PUT** /api/networks/routers/{routerId}/bgp-neighbors/{id} | Update Network Router BGP Neighbor
[**update_network_router_firewall_rule**](NetworksApi.md#update_network_router_firewall_rule) | **PUT** /api/networks/routers/{routerId}/firewall-rules/{id} | Update a Network Router Firewall Rule
[**update_network_router_firewall_rule_group**](NetworksApi.md#update_network_router_firewall_rule_group) | **PUT** /api/networks/routers/{routerId}/firewall-rule-groups/{id} | Update a Network Router Firewall Rule Group
[**update_network_router_nat**](NetworksApi.md#update_network_router_nat) | **PUT** /api/networks/routers/{routerId}/nats/{id} | Update Network Router NAT
[**update_network_router_permissions**](NetworksApi.md#update_network_router_permissions) | **PUT** /api/networks/routers/{routerId}/permissions | Update Network Router Permissions
[**update_network_server_group**](NetworksApi.md#update_network_server_group) | **PUT** /api/networks/servers/{serverId}/groups/{id} | Update a Network Server Group
[**update_network_transport_zone**](NetworksApi.md#update_network_transport_zone) | **PUT** /api/networks/servers/{serverId}/scopes/{id} | Update a Network Transport Zone
[**update_static_route**](NetworksApi.md#update_static_route) | **PUT** /api/networks/{id}/routes/{routeId} | Update a Network Static Route
[**update_subnet**](NetworksApi.md#update_subnet) | **PUT** /api/subnets/{id} | Update a Subnet


# **create_network_dhcp_relay**
> SuccessId create_network_dhcp_relay(server_id)

Create a Network DHCP Relay

Create a Network DHCP Relay. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_dhcp_relay_request import CreateNetworkDhcpRelayRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    create_network_dhcp_relay_request = CreateNetworkDhcpRelayRequest(
        network_dhcp_relay=NetworkDhcpRelayCreate(
            name="name_example",
            server_ip_addresses=[
                "server_ip_addresses_example",
            ],
        ),
    ) # CreateNetworkDhcpRelayRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network DHCP Relay
        api_response = api_instance.create_network_dhcp_relay(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_dhcp_relay: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network DHCP Relay
        api_response = api_instance.create_network_dhcp_relay(server_id, create_network_dhcp_relay_request=create_network_dhcp_relay_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_dhcp_relay: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **create_network_dhcp_relay_request** | [**CreateNetworkDhcpRelayRequest**](CreateNetworkDhcpRelayRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_dhcp_server**
> SuccessId create_network_dhcp_server(server_id)

Create a Network DHCP Server

Create a Network DHCP Server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_dhcp_server_request import CreateNetworkDhcpServerRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    create_network_dhcp_server_request = CreateNetworkDhcpServerRequest(
        network_dhcp_server=NetworkDhcpServerCreate(
            server_ip_address="server_ip_address_example",
            lease_time=86400,
            name="name_example",
            config=NetworkDhcpServerCreateConfig(
                edge_cluster="edge_cluster_example",
                preferred_edge_node1="preferred_edge_node1_example",
                preferred_edge_node2="preferred_edge_node2_example",
            ),
        ),
    ) # CreateNetworkDhcpServerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network DHCP Server
        api_response = api_instance.create_network_dhcp_server(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_dhcp_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network DHCP Server
        api_response = api_instance.create_network_dhcp_server(server_id, create_network_dhcp_server_request=create_network_dhcp_server_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_dhcp_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **create_network_dhcp_server_request** | [**CreateNetworkDhcpServerRequest**](CreateNetworkDhcpServerRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_domain**
> CreateNetworkDomain200Response create_network_domain()

Create a Network Domain

Create a Network Domain. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_domain200_response import CreateNetworkDomain200Response
from openapi_client.model.create_network_domain_request import CreateNetworkDomainRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_network_domain_request = CreateNetworkDomainRequest(
        network_domain=NetworkDomainCreate(
            name="name_example",
            description="description_example",
            display_name="display_name_example",
            public_zone=False,
            task_set_id=1,
            active=True,
            domain_controller=True,
            domain_username="domain_username_example",
            domain_password="domain_password_example",
            dc_server="dc_server_example",
            ou_path="ou_path_example",
            guest_username="guest_username_example",
            guest_password="guest_password_example",
        ),
    ) # CreateNetworkDomainRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Domain
        api_response = api_instance.create_network_domain(create_network_domain_request=create_network_domain_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_domain: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_domain_request** | [**CreateNetworkDomainRequest**](CreateNetworkDomainRequest.md)|  | [optional]

### Return type

[**CreateNetworkDomain200Response**](CreateNetworkDomain200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_firewall_rule**
> SuccessId create_network_firewall_rule(server_id)

Create a Network Firewall Rule

Use this command to create a network firewall rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_firewall_rule_request import CreateNetworkFirewallRuleRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    create_network_firewall_rule_request = CreateNetworkFirewallRuleRequest(
        rule=NetworkFirewallRuleCreate(
            rule_group=NetworkFirewallRuleCreateRuleGroup(
                id=1,
            ),
            name="name_example",
            description="description_example",
            enabled=True,
            priority="priority_example",
            direction="direction_example",
            sources=NetworkFirewallRuleCreateSources(
                id=[
                    "id_example",
                ],
            ),
            destinations=NetworkFirewallRuleCreateSources(
                id=[
                    "id_example",
                ],
            ),
            config=NetworkFirewallRuleCreateConfig(
                application=[
                    "application_example",
                ],
                profile=[
                    "profile_example",
                ],
            ),
            scopes=NetworkFirewallRuleCreateSources(
                id=[
                    "id_example",
                ],
            ),
            policy="policy_example",
        ),
    ) # CreateNetworkFirewallRuleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Firewall Rule
        api_response = api_instance.create_network_firewall_rule(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_firewall_rule: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Firewall Rule
        api_response = api_instance.create_network_firewall_rule(server_id, create_network_firewall_rule_request=create_network_firewall_rule_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **create_network_firewall_rule_request** | [**CreateNetworkFirewallRuleRequest**](CreateNetworkFirewallRuleRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_firewall_rule_group**
> SuccessId create_network_firewall_rule_group(server_id)

Create a Network Firewall Rule Group

Use this command to create a network firewall rule group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_router_firewall_rule_group_request import CreateNetworkRouterFirewallRuleGroupRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    create_network_router_firewall_rule_group_request = CreateNetworkRouterFirewallRuleGroupRequest(
        rule_group=NetworkFirewallRuleGroupCreate(
            name="name_example",
            description="description_example",
            priority=1,
            external_type="external_type_example",
            group_layer="group_layer_example",
        ),
    ) # CreateNetworkRouterFirewallRuleGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Firewall Rule Group
        api_response = api_instance.create_network_firewall_rule_group(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_firewall_rule_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Firewall Rule Group
        api_response = api_instance.create_network_firewall_rule_group(server_id, create_network_router_firewall_rule_group_request=create_network_router_firewall_rule_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **create_network_router_firewall_rule_group_request** | [**CreateNetworkRouterFirewallRuleGroupRequest**](CreateNetworkRouterFirewallRuleGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_group**
> SuccessId create_network_group()

Create a Network Group

Use this command to create a network group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_group_request import CreateNetworkGroupRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_network_group_request = CreateNetworkGroupRequest(
        network_group=NetworkGroupsCreate(
            name="name_example",
            description="description_example",
            networks=[
                1,
            ],
            subnets=[
                {},
            ],
        ),
    ) # CreateNetworkGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Group
        api_response = api_instance.create_network_group(create_network_group_request=create_network_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_group_request** | [**CreateNetworkGroupRequest**](CreateNetworkGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_pool**
> CreateNetworkPool200Response create_network_pool()

Create a Network Pool

Create a Network Pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_pool200_response import CreateNetworkPool200Response
from openapi_client.model.create_network_pool_request import CreateNetworkPoolRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_network_pool_request = CreateNetworkPoolRequest(
        network_pool=NetworkPoolCreate(
            name="name_example",
            type=NetworkPoolCreateType(
                code=None,
            ),
            ip_ranges=[
                NetworkPoolCreateIpRangesInner(
                    start_address="start_address_example",
                    end_address="end_address_example",
                    cidr_ipv6="cidr_ipv6_example",
                ),
            ],
            config={},
        ),
    ) # CreateNetworkPoolRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Pool
        api_response = api_instance.create_network_pool(create_network_pool_request=create_network_pool_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_pool_request** | [**CreateNetworkPoolRequest**](CreateNetworkPoolRequest.md)|  | [optional]

### Return type

[**CreateNetworkPool200Response**](CreateNetworkPool200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_pool_ip**
> CreateNetworkPoolIp200Response create_network_pool_ip(id)

Create a Network Pool IP Address

Create an IP Address for a Specific Network Pool 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model406_error import Model406Error
from openapi_client.model.model404_error import Model404Error
from openapi_client.model.model429_error import Model429Error
from openapi_client.model.model500_error import Model500Error
from openapi_client.model.create_network_pool_ip200_response import CreateNetworkPoolIp200Response
from openapi_client.model.model400_error import Model400Error
from openapi_client.model.model405_error import Model405Error
from openapi_client.model.model403_error import Model403Error
from openapi_client.model.model503_error import Model503Error
from openapi_client.model.model401_error import Model401Error
from openapi_client.model.create_network_pool_ip_request import CreateNetworkPoolIpRequest
from openapi_client.model.model410_error import Model410Error
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_network_pool_ip_request = CreateNetworkPoolIpRequest(
        network_pool_ip=NetworkPoolIpCreate(
            ip_address="ip_address_example",
            hostname="hostname_example",
        ),
    ) # CreateNetworkPoolIpRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Pool IP Address
        api_response = api_instance.create_network_pool_ip(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_pool_ip: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Pool IP Address
        api_response = api_instance.create_network_pool_ip(id, create_network_pool_ip_request=create_network_pool_ip_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_pool_ip: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_network_pool_ip_request** | [**CreateNetworkPoolIpRequest**](CreateNetworkPoolIpRequest.md)|  | [optional]

### Return type

[**CreateNetworkPoolIp200Response**](CreateNetworkPoolIp200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_pool_server**
> CreateNetworkPoolServer200Response create_network_pool_server()

Create a Network Pool Server

This endpoint allows creating a Network Pool Server. Only certain types of integrations support creating and deleting network pool servers, such as Bluecat, Infoblox, phpIPAM, and Solar Winds. Configuration options vary by type. Note that creating a pool server will automatically create and associate the corresponding network integration object, but management is done via the network pool server object.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_pool_server200_response import CreateNetworkPoolServer200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_pool_server_request import CreateNetworkPoolServerRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_network_pool_server_request = CreateNetworkPoolServerRequest(
        network_pool_server=CreateNetworkPoolServerRequestNetworkPoolServer(),
    ) # CreateNetworkPoolServerRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Pool Server
        api_response = api_instance.create_network_pool_server(create_network_pool_server_request=create_network_pool_server_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_pool_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_pool_server_request** | [**CreateNetworkPoolServerRequest**](CreateNetworkPoolServerRequest.md)|  | [optional]

### Return type

[**CreateNetworkPoolServer200Response**](CreateNetworkPoolServer200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_proxy**
> CreateNetworkProxy200Response create_network_proxy()

Create a Network Proxy

Create a Network Proxy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_proxy200_response import CreateNetworkProxy200Response
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_proxy_request import CreateNetworkProxyRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_network_proxy_request = CreateNetworkProxyRequest(
        network_proxy=CreateNetworkProxyRequestNetworkProxy(
            name="name_example",
            proxy_host="proxy_host_example",
            proxy_port="proxy_port_example",
            proxy_user="proxy_user_example",
            proxy_password="proxy_password_example",
            proxy_domain="proxy_domain_example",
            proxy_workstation="proxy_workstation_example",
            visibility="private",
            account=CreateNetworkProxyRequestNetworkProxyAccount(
                id=1,
            ),
        ),
    ) # CreateNetworkProxyRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Proxy
        api_response = api_instance.create_network_proxy(create_network_proxy_request=create_network_proxy_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_proxy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_proxy_request** | [**CreateNetworkProxyRequest**](CreateNetworkProxyRequest.md)|  | [optional]

### Return type

[**CreateNetworkProxy200Response**](CreateNetworkProxy200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_router**
> SuccessId create_network_router()

Create a Network Router

Use this command to create a network router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_router_request import CreateNetworkRouterRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_network_router_request = CreateNetworkRouterRequest(
        network_router=NetworkRoutersCreate(
            name="name_example",
            type=NetworkRoutersCreateType(
                id=1,
            ),
            site=NetworkRoutersCreateSite(
                id=NetworkRoutersCreateSiteId(None),
            ),
            enabled=True,
            zone=NetworkRoutersCreateZone(
                id=1,
            ),
            network_server=NetworkRoutersCreateNetworkServer(
                id=1,
            ),
        ),
    ) # CreateNetworkRouterRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Router
        api_response = api_instance.create_network_router(create_network_router_request=create_network_router_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_network_router_request** | [**CreateNetworkRouterRequest**](CreateNetworkRouterRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_router_bgp_neighbor**
> SuccessId create_network_router_bgp_neighbor(router_id)

Create a Network Router BGP Neighbor

Use this command to create a BGP Neighbor for an existing network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_router_bgp_neighbor_request import CreateNetworkRouterBgpNeighborRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID
    create_network_router_bgp_neighbor_request = CreateNetworkRouterBgpNeighborRequest(
        network_router_bgp_neighbor={},
    ) # CreateNetworkRouterBgpNeighborRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Router BGP Neighbor
        api_response = api_instance.create_network_router_bgp_neighbor(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_bgp_neighbor: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Router BGP Neighbor
        api_response = api_instance.create_network_router_bgp_neighbor(router_id, create_network_router_bgp_neighbor_request=create_network_router_bgp_neighbor_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_bgp_neighbor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **create_network_router_bgp_neighbor_request** | [**CreateNetworkRouterBgpNeighborRequest**](CreateNetworkRouterBgpNeighborRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_router_firewall_rule**
> SuccessId create_network_router_firewall_rule(router_id)

Create a Network Router Firewall Rule

Use this command to create a firewall rule for an existing network router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_router_firewall_rule_request import CreateNetworkRouterFirewallRuleRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID
    create_network_router_firewall_rule_request = CreateNetworkRouterFirewallRuleRequest(
        rule=NetworkRouterFirewallRuleCreate(
            name="name_example",
            enabled=True,
            priority=1,
        ),
    ) # CreateNetworkRouterFirewallRuleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Router Firewall Rule
        api_response = api_instance.create_network_router_firewall_rule(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_firewall_rule: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Router Firewall Rule
        api_response = api_instance.create_network_router_firewall_rule(router_id, create_network_router_firewall_rule_request=create_network_router_firewall_rule_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **create_network_router_firewall_rule_request** | [**CreateNetworkRouterFirewallRuleRequest**](CreateNetworkRouterFirewallRuleRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_router_firewall_rule_group**
> SuccessId create_network_router_firewall_rule_group(router_id)

Create a Network Router Firewall Rule Group

Use this command to create a network firewall rule group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_router_firewall_rule_group_request import CreateNetworkRouterFirewallRuleGroupRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID
    create_network_router_firewall_rule_group_request = CreateNetworkRouterFirewallRuleGroupRequest(
        rule_group=NetworkFirewallRuleGroupCreate(
            name="name_example",
            description="description_example",
            priority=1,
            external_type="external_type_example",
            group_layer="group_layer_example",
        ),
    ) # CreateNetworkRouterFirewallRuleGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Router Firewall Rule Group
        api_response = api_instance.create_network_router_firewall_rule_group(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_firewall_rule_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Router Firewall Rule Group
        api_response = api_instance.create_network_router_firewall_rule_group(router_id, create_network_router_firewall_rule_group_request=create_network_router_firewall_rule_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **create_network_router_firewall_rule_group_request** | [**CreateNetworkRouterFirewallRuleGroupRequest**](CreateNetworkRouterFirewallRuleGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_router_nat**
> SuccessId create_network_router_nat(router_id)

Create a Network Router NAT

Use this command to create a NAT for an existing network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_router_nat_request import CreateNetworkRouterNatRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID
    create_network_router_nat_request = CreateNetworkRouterNatRequest(
        network_router_nat=CreateNetworkRouterNatRequestNetworkRouterNAT(
            name=None,
        ),
    ) # CreateNetworkRouterNatRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Router NAT
        api_response = api_instance.create_network_router_nat(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_nat: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Router NAT
        api_response = api_instance.create_network_router_nat(router_id, create_network_router_nat_request=create_network_router_nat_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_nat: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **create_network_router_nat_request** | [**CreateNetworkRouterNatRequest**](CreateNetworkRouterNatRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_router_route**
> SuccessId create_network_router_route(router_id)

Create a Network Router Route

Use this command to create a route for an existing network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_router_route_request import CreateNetworkRouterRouteRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID
    create_network_router_route_request = CreateNetworkRouterRouteRequest(
        network_route=NetworkRouterRouteCreate(
            name="name_example",
            description="description_example",
            enabled=False,
            default_route=False,
            source="source_example",
            destination="destination_example",
            network_mtu="network_mtu_example",
        ),
    ) # CreateNetworkRouterRouteRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Router Route
        api_response = api_instance.create_network_router_route(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_route: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Router Route
        api_response = api_instance.create_network_router_route(router_id, create_network_router_route_request=create_network_router_route_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_router_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **create_network_router_route_request** | [**CreateNetworkRouterRouteRequest**](CreateNetworkRouterRouteRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_server_group**
> SuccessId create_network_server_group(server_id)

Create a Network Server Group

Use this command to create a network server group. Note: Only available for NSX-T network integrations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_server_group_request import CreateNetworkServerGroupRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    create_network_server_group_request = CreateNetworkServerGroupRequest(
        group=NetworkServerGroupCreate(
            id=1,
            name="name_example",
            description="description_example",
            internal_id="internal_id_example",
            external_id="external_id_example",
            visibility="visibility_example",
            account=UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                id=1,
            ),
            owner=UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                id=1,
            ),
            network_server=UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                id=1,
            ),
            permissions=Permissions(
                resource_pool=ClusterPermissionsResourcePool(
                    id=1,
                    name="name_example",
                    visibility="visibility_example",
                ),
                resource_permissions=ResourcePermissions(
                    all=True,
                    sites=[
                        ResourcePermissionsSitesInner(
                            id=1,
                            name="name_example",
                            default=True,
                        ),
                    ],
                    all_plans=True,
                    plans=[
                        ResourcePermissionsSitesInner(
                            id=1,
                            name="name_example",
                            default=True,
                        ),
                    ],
                ),
                tenant_permissions=UpdateCloudFoldersRequestFolderTenantPermissionsInner(
                    accounts=[1,3],
                ),
            ),
            tags=[
                Tag(
                    id=1,
                    name="name_example",
                    value="value_example",
                ),
            ],
            members=[
                NetworkServerGroupMember(
                    id=1,
                    category="category_example",
                    type="type_example",
                    member_name="member_name_example",
                    member_type="member_type_example",
                    member_value="member_value_example",
                    member_expression="member_expression_example",
                    display_order=1,
                    internal_id="internal_id_example",
                    external_id="external_id_example",
                    members=[
                        {},
                    ],
                ),
            ],
        ),
    ) # CreateNetworkServerGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Server Group
        api_response = api_instance.create_network_server_group(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_server_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Server Group
        api_response = api_instance.create_network_server_group(server_id, create_network_server_group_request=create_network_server_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_server_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **create_network_server_group_request** | [**CreateNetworkServerGroupRequest**](CreateNetworkServerGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_network_transport_zone**
> SuccessId create_network_transport_zone(server_id)

Create a Network Transport Zone

Use this command to create a network transport zone.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_transport_zone_request import CreateNetworkTransportZoneRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    create_network_transport_zone_request = CreateNetworkTransportZoneRequest(
        network_scope=NetworkTransportZoneCreate(
            name="name_example",
            description="description_example",
            visibility="visibility_example",
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
        ),
    ) # CreateNetworkTransportZoneRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Transport Zone
        api_response = api_instance.create_network_transport_zone(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_transport_zone: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Transport Zone
        api_response = api_instance.create_network_transport_zone(server_id, create_network_transport_zone_request=create_network_transport_zone_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_network_transport_zone: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **create_network_transport_zone_request** | [**CreateNetworkTransportZoneRequest**](CreateNetworkTransportZoneRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_networks**
> CreateNetworks200Response create_networks()

Create a Network

This endpoint allows creating a Network. Only certain types of clouds support creating and deleting networks. Configuration options vary by Network Types.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_networks200_response import CreateNetworks200Response
from openapi_client.model.create_networks_request import CreateNetworksRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_networks_request = CreateNetworksRequest(
        network=NetworkCreate(
            name="name_example",
            display_name="display_name_example",
            labels=[
                "labels_example",
            ],
            description="description_example",
            site=NetworkCreateSite(
                id=1,
            ),
            zone=NetworkCreateZone(
                id=1,
            ),
            type=NetworkCreateType(
                id=1,
            ),
            ipv4_enabled=True,
            ipv6_enabled=True,
            cidr="cidr_example",
            gateway="gateway_example",
            dns_primary="dns_primary_example",
            dns_secondary="dns_secondary_example",
            gateway_ipv6="gateway_ipv6_example",
            netmask_ipv6="netmask_ipv6_example",
            dns_primary_ipv6="dns_primary_ipv6_example",
            dns_secondary_ipv6="dns_secondary_ipv6_example",
            cidr_ipv6="cidr_ipv6_example",
            vlan_id=1,
            pool=1,
            pool_ipv6=1,
            allow_static_override=True,
            assign_public_ip=True,
            active=True,
            dhcp_server=True,
            dhcp_server_ipv6=True,
            network_domain=NetworkCreateNetworkDomain(
                id=1,
            ),
            search_domains="search_domains_example",
            network_proxy=NetworkCreateNetworkProxy(
                id=1,
            ),
            appliance_url_proxy_bypass=True,
            no_proxy="no_proxy_example",
            visibility="private",
            config=NetworkCreateConfig(),
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            resource_permissions=NetworkCreateResourcePermissions(
                all=True,
                sites=[
                    1,
                ],
            ),
        ),
    ) # CreateNetworksRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network
        api_response = api_instance.create_networks(create_networks_request=create_networks_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_networks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_networks_request** | [**CreateNetworksRequest**](CreateNetworksRequest.md)|  | [optional]

### Return type

[**CreateNetworks200Response**](CreateNetworks200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_static_route**
> SuccessId create_static_route(id)

Create a Network Static Route

Use this command to create a route for an existing network. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_static_route_request import CreateStaticRouteRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_static_route_request = CreateStaticRouteRequest(
        network_route=NetworkStaticRouteCreate(
            source="source_example",
            destination="destination_example",
        ),
    ) # CreateStaticRouteRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Create a Network Static Route
        api_response = api_instance.create_static_route(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_static_route: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Network Static Route
        api_response = api_instance.create_static_route(id, create_static_route_request=create_static_route_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_static_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_static_route_request** | [**CreateStaticRouteRequest**](CreateStaticRouteRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_subnet**
> CreateSubnet200Response create_subnet()

Create a Subnet

This endpoint allows creating a Subnet. Only certain types of clouds support creating and deleting subnets. Configuration options vary for each Subnet Type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_subnet200_response import CreateSubnet200Response
from openapi_client.model.create_subnet_request import CreateSubnetRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    create_subnet_request = CreateSubnetRequest(
        subnet=CreateSubnetRequestSubnet(
            type=CreateSubnetRequestSubnetType(
                id=1,
            ),
            config={},
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            visibility="private",
            labels=[
                "labels_example",
            ],
        ),
        resource_permission=CreateSubnetRequestResourcePermission(
            all=True,
            sites=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            all_plans=True,
            plans=[
                {},
            ],
        ),
    ) # CreateSubnetRequest |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Create a Subnet
        api_response = api_instance.create_subnet(create_subnet_request=create_subnet_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->create_subnet: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_subnet_request** | [**CreateSubnetRequest**](CreateSubnetRequest.md)|  | [optional]

### Return type

[**CreateSubnet200Response**](CreateSubnet200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network**
> Model200Success delete_network(id)

Delete a Network

Will delete a Network from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network
        api_response = api_instance.delete_network(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_dhcp_relay**
> Model200Success delete_network_dhcp_relay(id, server_id)

Delete a Network DHCP Relay

Will delete a Network DHCP Relay from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network DHCP Relay
        api_response = api_instance.delete_network_dhcp_relay(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_dhcp_relay: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_dhcp_server**
> Model200Success delete_network_dhcp_server(id, server_id)

Delete a Network DHCP Server

Will delete a Network DHCP Server from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network DHCP Server
        api_response = api_instance.delete_network_dhcp_server(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_dhcp_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_domain**
> Model200Success delete_network_domain(id)

Delete a Network Domain

Will delete a Network Domain from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Domain
        api_response = api_instance.delete_network_domain(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_domain: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_firewall_rule**
> Model200Success delete_network_firewall_rule(id, server_id)

Delete a Network Firewall Rule

Will delete a Network Firewall Rule from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Firewall Rule
        api_response = api_instance.delete_network_firewall_rule(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_firewall_rule_group**
> Model200Success delete_network_firewall_rule_group(id, server_id)

Delete a Network firewall rule group

Will delete a network firewall rule group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network firewall rule group
        api_response = api_instance.delete_network_firewall_rule_group(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_group**
> Model200Success delete_network_group(id)

Delete a Network Group

Will delete a Network Group from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Group
        api_response = api_instance.delete_network_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_pool**
> Model200Success delete_network_pool(id)

Delete a Network Pool

Will delete a Network Pool from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Pool
        api_response = api_instance.delete_network_pool(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_pool_ip**
> Model200Success delete_network_pool_ip(network_pool_id, id)

Delete a host record associated with an IP Address for a Specific Network Pool

Will delete a host record associated with an IP address for a specific network pool and free up the address

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model406_error import Model406Error
from openapi_client.model.model404_error import Model404Error
from openapi_client.model.model429_error import Model429Error
from openapi_client.model.model500_error import Model500Error
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.model400_error import Model400Error
from openapi_client.model.model405_error import Model405Error
from openapi_client.model.model403_error import Model403Error
from openapi_client.model.model503_error import Model503Error
from openapi_client.model.model401_error import Model401Error
from openapi_client.model.model410_error import Model410Error
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    network_pool_id = 1 # int | Morpheus Network Pool ID of the Object being referenced
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a host record associated with an IP Address for a Specific Network Pool
        api_response = api_instance.delete_network_pool_ip(network_pool_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_pool_ip: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_pool_id** | **int**| Morpheus Network Pool ID of the Object being referenced |
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_pool_server**
> Model200Success delete_network_pool_server(id)

Delete a Network Pool Server

Will delete a Network Pool Server from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Pool Server
        api_response = api_instance.delete_network_pool_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_pool_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_proxy**
> Model200Success delete_network_proxy(id)

Delete a Network Proxy

Will delete a Network Proxy from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Proxy
        api_response = api_instance.delete_network_proxy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_proxy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_router**
> Model200Success delete_network_router(id)

Delete a Network Router

Will delete a Network Router from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Router
        api_response = api_instance.delete_network_router(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_router: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_router_bgp_neighbor**
> Model200Success delete_network_router_bgp_neighbor(id, router_id)

Delete a Network Router BGP Neighbor

Will delete a BGP Neighbor from a network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Router BGP Neighbor
        api_response = api_instance.delete_network_router_bgp_neighbor(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_router_bgp_neighbor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_router_firewall_rule**
> Model200Success delete_network_router_firewall_rule(id, router_id)

Delete a Network Router Firewall Rule

Will delete a firewall rule from a network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Router Firewall Rule
        api_response = api_instance.delete_network_router_firewall_rule(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_router_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_router_firewall_rule_group**
> Model200Success delete_network_router_firewall_rule_group(id, router_id)

Delete a Network Router firewall rule group

Will delete a network router firewall rule group.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Router firewall rule group
        api_response = api_instance.delete_network_router_firewall_rule_group(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_router_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_router_nat**
> Model200Success delete_network_router_nat(id, router_id)

Delete a Network Router NAT

Will delete a NAT from a network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Router NAT
        api_response = api_instance.delete_network_router_nat(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_router_nat: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_router_route**
> Model200Success delete_network_router_route(id, router_id)

Delete a Network Router Route

Will delete a Route from a network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Router Route
        api_response = api_instance.delete_network_router_route(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_router_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_server_group**
> Model200Success delete_network_server_group(id, server_id)

Delete a Network Server Group

Will delete a Network Server Group from the system and make it no longer usable. Note: Only available for NSX-T network integrations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Server Group
        api_response = api_instance.delete_network_server_group(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_server_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_network_transport_zone**
> Model200Success delete_network_transport_zone(id, server_id)

Delete a Network Transport Zone

Will delete a Network Transport Zone from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Transport Zone
        api_response = api_instance.delete_network_transport_zone(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_network_transport_zone: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_static_route**
> Model200Success delete_static_route(id, route_id)

Delete a Network Static Route

Will delete a route from a network.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    route_id = 4 # float | The ID of the route

    # example passing only required values which don't have defaults set
    try:
        # Delete a Network Static Route
        api_response = api_instance.delete_static_route(id, route_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_static_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **route_id** | **float**| The ID of the route |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_subnet**
> Model200Success delete_subnet(id)

Delete a Subnet

Will delete a Subnet from the system and make it no longer usable.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Delete a Subnet
        api_response = api_instance.delete_subnet(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->delete_subnet: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_all_network_floating_ips**
> GetAllNetworkFloatingIps200Response get_all_network_floating_ips()

Get All Floating IPs

This endpoint retrieves all network floating IPs associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_all_network_floating_ips200_response import GetAllNetworkFloatingIps200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    ip_address = "10.32.23.188" # str | Filter by IP Address (optional)
    ip_status = "assigned" # str | Filter by IP Status (optional)
    zone_id = 1 # int | Filter by Cloud ID (optional)
    server_id = 1 # int | Filter by Server ID (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Floating IPs
        api_response = api_instance.get_all_network_floating_ips(phrase=phrase, ip_address=ip_address, ip_status=ip_status, zone_id=zone_id, server_id=server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_all_network_floating_ips: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **ip_address** | **str**| Filter by IP Address | [optional]
 **ip_status** | **str**| Filter by IP Status | [optional]
 **zone_id** | **int**| Filter by Cloud ID | [optional]
 **server_id** | **int**| Filter by Server ID | [optional]

### Return type

[**GetAllNetworkFloatingIps200Response**](GetAllNetworkFloatingIps200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network**
> CreateNetworks200ResponseAllOf get_network(id)

Get a Specific Network

This endpoint retrieves a specific Network. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_networks200_response_all_of import CreateNetworks200ResponseAllOf
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network
        api_response = api_instance.get_network(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateNetworks200ResponseAllOf**](CreateNetworks200ResponseAllOf.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_dhcp_relay**
> GetNetworkDhcpRelay200Response get_network_dhcp_relay(id, server_id)

Get a Specific Network DHCP Relay

This endpoint retrieves a specific Network DHCP Relay. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_dhcp_relay200_response import GetNetworkDhcpRelay200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network DHCP Relay
        api_response = api_instance.get_network_dhcp_relay(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_dhcp_relay: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**GetNetworkDhcpRelay200Response**](GetNetworkDhcpRelay200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_dhcp_relays**
> GetNetworkDhcpRelays200Response get_network_dhcp_relays(server_id)

Get all Network DHCP Relays for Network Relay

This endpoint retrieves all Network DHCP Relays for a specified Network Service. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_dhcp_relays200_response import GetNetworkDhcpRelays200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network DHCP Relays for Network Relay
        api_response = api_instance.get_network_dhcp_relays(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_dhcp_relays: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network DHCP Relays for Network Relay
        api_response = api_instance.get_network_dhcp_relays(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_dhcp_relays: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkDhcpRelays200Response**](GetNetworkDhcpRelays200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_dhcp_server**
> GetNetworkDhcpServer200Response get_network_dhcp_server(id, server_id)

Get a Specific Network DHCP Server

This endpoint retrieves a specific Network DHCP Server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_dhcp_server200_response import GetNetworkDhcpServer200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network DHCP Server
        api_response = api_instance.get_network_dhcp_server(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_dhcp_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**GetNetworkDhcpServer200Response**](GetNetworkDhcpServer200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_dhcp_servers**
> GetNetworkDhcpServers200Response get_network_dhcp_servers(server_id)

Get all Network DHCP Servers for Network Server

This endpoint retrieves all Network DHCP Servers for a specified Network Service. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_dhcp_servers200_response import GetNetworkDhcpServers200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network DHCP Servers for Network Server
        api_response = api_instance.get_network_dhcp_servers(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_dhcp_servers: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network DHCP Servers for Network Server
        api_response = api_instance.get_network_dhcp_servers(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_dhcp_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkDhcpServers200Response**](GetNetworkDhcpServers200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_domain**
> CreateNetworkDomain200Response get_network_domain(id)

Get a Specific Network Domain

This endpoint retrieves a specific Network Domain. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_domain200_response import CreateNetworkDomain200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Domain
        api_response = api_instance.get_network_domain(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_domain: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateNetworkDomain200Response**](CreateNetworkDomain200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_domains**
> GetNetworkDomains200Response get_network_domains()

Get all Network Domains

This endpoint retrieves all Network Domains associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_domains200_response import GetNetworkDomains200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Domains
        api_response = api_instance.get_network_domains(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_domains: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkDomains200Response**](GetNetworkDomains200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_edge_cluster**
> GetNetworkEdgeCluster200Response get_network_edge_cluster(id, server_id)

Get a Specific Network Edge Cluster

This endpoint retrieves a specific network Edge Cluster. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_edge_cluster200_response import GetNetworkEdgeCluster200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Edge Cluster
        api_response = api_instance.get_network_edge_cluster(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_edge_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**GetNetworkEdgeCluster200Response**](GetNetworkEdgeCluster200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_edge_clusters**
> GetNetworkEdgeClusters200Response get_network_edge_clusters(server_id)

Get all Network Edge Clusters for Network Server

This endpoint retrieves all Network Edge Clusters for a specified Network Service. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_edge_clusters200_response import GetNetworkEdgeClusters200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Edge Clusters for Network Server
        api_response = api_instance.get_network_edge_clusters(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_edge_clusters: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Edge Clusters for Network Server
        api_response = api_instance.get_network_edge_clusters(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_edge_clusters: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkEdgeClusters200Response**](GetNetworkEdgeClusters200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_firewall_rule**
> GetNetworkFirewallRule200Response get_network_firewall_rule(id, server_id)

Get a Specific Network Firewall Rule

This endpoint retrieves a specific Network Firewall Rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_firewall_rule200_response import GetNetworkFirewallRule200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Firewall Rule
        api_response = api_instance.get_network_firewall_rule(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**GetNetworkFirewallRule200Response**](GetNetworkFirewallRule200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_firewall_rule_group**
> GetNetworkFirewallRuleGroup200Response get_network_firewall_rule_group(id, server_id)

Get a Specific Network Firewall Rule Group

This endpoint retrieves a specific Network Firewall Rule Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_firewall_rule_group200_response import GetNetworkFirewallRuleGroup200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Firewall Rule Group
        api_response = api_instance.get_network_firewall_rule_group(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**GetNetworkFirewallRuleGroup200Response**](GetNetworkFirewallRuleGroup200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_firewall_rule_groups**
> GetNetworkFirewallRuleGroups200Response get_network_firewall_rule_groups(server_id)

Get all Network Firewall Rule Groups for Network Server

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_firewall_rule_groups200_response import GetNetworkFirewallRuleGroups200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Firewall Rule Groups for Network Server
        api_response = api_instance.get_network_firewall_rule_groups(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_firewall_rule_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Firewall Rule Groups for Network Server
        api_response = api_instance.get_network_firewall_rule_groups(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_firewall_rule_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkFirewallRuleGroups200Response**](GetNetworkFirewallRuleGroups200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_firewall_rules**
> GetNetworkFirewallRules200Response get_network_firewall_rules(server_id)

Get all Network Firewall Rules for Network Server

This endpoint retrieves all Network Firewall Rules for a specified Network Service. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_firewall_rules200_response import GetNetworkFirewallRules200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Firewall Rules for Network Server
        api_response = api_instance.get_network_firewall_rules(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_firewall_rules: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Firewall Rules for Network Server
        api_response = api_instance.get_network_firewall_rules(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_firewall_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkFirewallRules200Response**](GetNetworkFirewallRules200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_floating_ip**
> GetNetworkFloatingIp200Response get_network_floating_ip(id)

Get a Specific Floating IP

This endpoint retrieves a specific Floating IP. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_floating_ip200_response import GetNetworkFloatingIp200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Floating IP
        api_response = api_instance.get_network_floating_ip(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_floating_ip: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkFloatingIp200Response**](GetNetworkFloatingIp200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_group**
> GetNetworkGroup200Response get_network_group(id)

Get a Specific Network Group

This endpoint retrieves a specific Network Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_group200_response import GetNetworkGroup200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Group
        api_response = api_instance.get_network_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkGroup200Response**](GetNetworkGroup200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_groups**
> GetNetworkGroups200Response get_network_groups()

Get all Network Groups

This endpoint retrieves all Network Groups associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_groups200_response import GetNetworkGroups200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get all Network Groups
        api_response = api_instance.get_network_groups()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_groups: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetNetworkGroups200Response**](GetNetworkGroups200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_pool**
> CreateNetworkPool200Response get_network_pool(id)

Get a Specific Network Pool

This endpoint retrieves a specific Network Pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_pool200_response import CreateNetworkPool200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Pool
        api_response = api_instance.get_network_pool(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateNetworkPool200Response**](CreateNetworkPool200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_pool_ip**
> GetNetworkPoolIps200Response get_network_pool_ip(network_pool_id, id)

Get a Specific IP Address for a Specific Network Pool

This endpoint retrieves a specific IP address for a specific Network Pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model406_error import Model406Error
from openapi_client.model.model404_error import Model404Error
from openapi_client.model.model429_error import Model429Error
from openapi_client.model.model500_error import Model500Error
from openapi_client.model.model400_error import Model400Error
from openapi_client.model.model405_error import Model405Error
from openapi_client.model.model403_error import Model403Error
from openapi_client.model.model503_error import Model503Error
from openapi_client.model.get_network_pool_ips200_response import GetNetworkPoolIps200Response
from openapi_client.model.model401_error import Model401Error
from openapi_client.model.model410_error import Model410Error
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    network_pool_id = 1 # int | Morpheus Network Pool ID of the Object being referenced
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific IP Address for a Specific Network Pool
        api_response = api_instance.get_network_pool_ip(network_pool_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_pool_ip: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **network_pool_id** | **int**| Morpheus Network Pool ID of the Object being referenced |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkPoolIps200Response**](GetNetworkPoolIps200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_pool_ips**
> GetNetworkPoolIps200Response get_network_pool_ips(id)

Get all IP Addresses for a Specific Network Pool

This endpoint retrieves a list of IP addresses for a specific Network Pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model406_error import Model406Error
from openapi_client.model.model404_error import Model404Error
from openapi_client.model.model429_error import Model429Error
from openapi_client.model.model500_error import Model500Error
from openapi_client.model.model400_error import Model400Error
from openapi_client.model.model405_error import Model405Error
from openapi_client.model.model403_error import Model403Error
from openapi_client.model.model503_error import Model503Error
from openapi_client.model.get_network_pool_ips200_response import GetNetworkPoolIps200Response
from openapi_client.model.model401_error import Model401Error
from openapi_client.model.model410_error import Model410Error
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get all IP Addresses for a Specific Network Pool
        api_response = api_instance.get_network_pool_ips(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_pool_ips: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkPoolIps200Response**](GetNetworkPoolIps200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_pool_server**
> CreateNetworkPoolServer200ResponseAllOf get_network_pool_server(id)

Get a Specific Network Pool Server

This endpoint retrieves a specific Network Pool Server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_pool_server200_response_all_of import CreateNetworkPoolServer200ResponseAllOf
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Pool Server
        api_response = api_instance.get_network_pool_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_pool_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateNetworkPoolServer200ResponseAllOf**](CreateNetworkPoolServer200ResponseAllOf.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_pool_server_type**
> GetNetworkPoolServerType200Response get_network_pool_server_type(id)

Retrieves a Specific Network Pool Server Type

Retrieves a specific Network Pool Server Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_pool_server_type200_response import GetNetworkPoolServerType200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Retrieves a Specific Network Pool Server Type
        api_response = api_instance.get_network_pool_server_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_pool_server_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkPoolServerType200Response**](GetNetworkPoolServerType200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_pools**
> GetNetworkPools200Response get_network_pools()

Get all Network Pools

This endpoint retrieves all Network Pools associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_pools200_response import GetNetworkPools200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Pools
        api_response = api_instance.get_network_pools(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_pools: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkPools200Response**](GetNetworkPools200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_proxies**
> GetNetworkProxies200Response get_network_proxies()

Get all Network Proxies

This endpoint retrieves all Network Proxies associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_proxies200_response import GetNetworkProxies200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Proxies
        api_response = api_instance.get_network_proxies(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_proxies: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkProxies200Response**](GetNetworkProxies200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_proxy**
> GetNetworkProxy200Response get_network_proxy(id)

Get a Specific Network Proxy

This endpoint retrieves a specific Network Proxy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_proxy200_response import GetNetworkProxy200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Proxy
        api_response = api_instance.get_network_proxy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_proxy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkProxy200Response**](GetNetworkProxy200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router**
> GetNetworkRouter200Response get_network_router(id)

Get a Specific Network Router

This endpoint retrieves a specific Network Router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_router200_response import GetNetworkRouter200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Router
        api_response = api_instance.get_network_router(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkRouter200Response**](GetNetworkRouter200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_bgp_neighbor**
> GetNetworkRouterBgpNeighbor200Response get_network_router_bgp_neighbor(id, router_id)

Get a Network Router BGP Neighbor

This endpoint retrieves a network router BGP Neighbor for specified network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_router_bgp_neighbor200_response import GetNetworkRouterBgpNeighbor200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Network Router BGP Neighbor
        api_response = api_instance.get_network_router_bgp_neighbor(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_bgp_neighbor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRouterBgpNeighbor200Response**](GetNetworkRouterBgpNeighbor200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_firewall_rule**
> GetNetworkRouterFirewallRule200Response get_network_router_firewall_rule(id, router_id)

Get a Firewall Rule for Network Router

This endpoint retrieves a firewall rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_router_firewall_rule200_response import GetNetworkRouterFirewallRule200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Firewall Rule for Network Router
        api_response = api_instance.get_network_router_firewall_rule(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRouterFirewallRule200Response**](GetNetworkRouterFirewallRule200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_firewall_rule_group**
> GetNetworkRouterFirewallRuleGroup200Response get_network_router_firewall_rule_group(id, router_id)

Get a Firewall Rule Group for Network Router

This endpoint retrieves a firewall rule group for specified network router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_router_firewall_rule_group200_response import GetNetworkRouterFirewallRuleGroup200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Firewall Rule Group for Network Router
        api_response = api_instance.get_network_router_firewall_rule_group(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRouterFirewallRuleGroup200Response**](GetNetworkRouterFirewallRuleGroup200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_firewall_rule_groups**
> GetNetworkRouterFirewallRuleGroups200Response get_network_router_firewall_rule_groups(router_id)

Get all Network Firewall Rule Groups for Network Router

This endpoint retrieves all Network Firewall Rule Groups for a specified Network Service. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_router_firewall_rule_groups200_response import GetNetworkRouterFirewallRuleGroups200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Firewall Rule Groups for Network Router
        api_response = api_instance.get_network_router_firewall_rule_groups(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_firewall_rule_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRouterFirewallRuleGroups200Response**](GetNetworkRouterFirewallRuleGroups200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_nat**
> GetNetworkRouterNat200Response get_network_router_nat(id, router_id)

Get a Network Router NAT

This endpoint retrieves a network router NAT for specified network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_router_nat200_response import GetNetworkRouterNat200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Network Router NAT
        api_response = api_instance.get_network_router_nat(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_nat: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRouterNat200Response**](GetNetworkRouterNat200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_route**
> GetNetworkRouterRoute200Response get_network_router_route(id, router_id)

Get a Route for Network Router

This endpoint retrieves a Route. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_router_route200_response import GetNetworkRouterRoute200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Route for Network Router
        api_response = api_instance.get_network_router_route(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRouterRoute200Response**](GetNetworkRouterRoute200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_router_type**
> GetNetworkRouterType200Response get_network_router_type(id)

Get a Specific Network Router Type

This endpoint retrieves a specific network router type. Use this API to retrieve list of available option types for a specific network router type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_router_type200_response import GetNetworkRouterType200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Router Type
        api_response = api_instance.get_network_router_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_router_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkRouterType200Response**](GetNetworkRouterType200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_routers**
> GetNetworkRouters200Response get_network_routers()

Get all Network Routers

This endpoint retrieves all Network Routers 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_routers200_response import GetNetworkRouters200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get all Network Routers
        api_response = api_instance.get_network_routers()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_routers: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**GetNetworkRouters200Response**](GetNetworkRouters200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_routers_bgp_neighbors**
> GetNetworkRoutersBgpNeighbors200Response get_network_routers_bgp_neighbors(router_id)

Get all BGP Neighbors for Network Router

This endpoint retrieves all BGP Neighbors for specified network router.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_routers_bgp_neighbors200_response import GetNetworkRoutersBgpNeighbors200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get all BGP Neighbors for Network Router
        api_response = api_instance.get_network_routers_bgp_neighbors(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_routers_bgp_neighbors: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRoutersBgpNeighbors200Response**](GetNetworkRoutersBgpNeighbors200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_routers_firewall_rules**
> GetNetworkRoutersFirewallRules200Response get_network_routers_firewall_rules(router_id)

Get all Firewall Rules for Network Router

This endpoint retrieves all firewall rules for specified network router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_routers_firewall_rules200_response import GetNetworkRoutersFirewallRules200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get all Firewall Rules for Network Router
        api_response = api_instance.get_network_routers_firewall_rules(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_routers_firewall_rules: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRoutersFirewallRules200Response**](GetNetworkRoutersFirewallRules200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_routers_nats**
> GetNetworkRoutersNats200Response get_network_routers_nats(router_id)

Get all Network Router NATs for Network Router

This endpoint retrieves all NATs for specified network router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_routers_nats200_response import GetNetworkRoutersNats200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Router NATs for Network Router
        api_response = api_instance.get_network_routers_nats(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_routers_nats: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRoutersNats200Response**](GetNetworkRoutersNats200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_routers_routes**
> GetNetworkRoutersRoutes200Response get_network_routers_routes(router_id)

Get all Routes for Network Router

This endpoint retrieves all Routes for specified network router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_routers_routes200_response import GetNetworkRoutersRoutes200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID

    # example passing only required values which don't have defaults set
    try:
        # Get all Routes for Network Router
        api_response = api_instance.get_network_routers_routes(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_routers_routes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |

### Return type

[**GetNetworkRoutersRoutes200Response**](GetNetworkRoutersRoutes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_server_group**
> GetNetworkServerGroup200Response get_network_server_group(server_id, id)

Get a specific Network Server Group

This endpoint retrieves a specific Network Server Group for a Network Server. Note: Only available for NSX-T network integrations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_network_server_group200_response import GetNetworkServerGroup200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a specific Network Server Group
        api_response = api_instance.get_network_server_group(server_id, id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_server_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkServerGroup200Response**](GetNetworkServerGroup200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_subnets**
> GetNetworkSubnets200Response get_network_subnets(id)

Get Subnets for a Network

This endpoint retrieves all Subnets under a specific network. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_subnets200_response import GetNetworkSubnets200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get Subnets for a Network
        api_response = api_instance.get_network_subnets(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_subnets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkSubnets200Response**](GetNetworkSubnets200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_transport_zone**
> GetNetworkTransportZone200Response get_network_transport_zone(id, server_id)

Get a Specific Network Transport Zone

This endpoint retrieves a specific Network Transport Zone. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_transport_zone200_response import GetNetworkTransportZone200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Transport Zone
        api_response = api_instance.get_network_transport_zone(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_transport_zone: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |

### Return type

[**GetNetworkTransportZone200Response**](GetNetworkTransportZone200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_transport_zones**
> GetNetworkTransportZones200Response get_network_transport_zones(server_id)

Get all Network Transport Zones for Network Server

This endpoint retrieves all Network Transport Zones for a specified Network Service.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_transport_zones200_response import GetNetworkTransportZones200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Transport Zones for Network Server
        api_response = api_instance.get_network_transport_zones(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_transport_zones: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Transport Zones for Network Server
        api_response = api_instance.get_network_transport_zones(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_transport_zones: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**GetNetworkTransportZones200Response**](GetNetworkTransportZones200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_network_type**
> GetNetworkType200Response get_network_type(id)

Get a Specific Network Type

This endpoint retrieves a specific Network Type.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_type200_response import GetNetworkType200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Network Type
        api_response = api_instance.get_network_type(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_network_type: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetNetworkType200Response**](GetNetworkType200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_static_route**
> GetStaticRoute200Response get_static_route(id, route_id)

Get a Static Route for a Network

This endpoint retrieves a network static route for specified network. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_static_route200_response import GetStaticRoute200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    route_id = 4 # float | The ID of the route

    # example passing only required values which don't have defaults set
    try:
        # Get a Static Route for a Network
        api_response = api_instance.get_static_route(id, route_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_static_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **route_id** | **float**| The ID of the route |

### Return type

[**GetStaticRoute200Response**](GetStaticRoute200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_static_routes**
> GetStaticRoutes200Response get_static_routes(id)

Get all Static Routes for a Network

This endpoint retrieves all routes for specified network. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.get_static_routes200_response import GetStaticRoutes200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get all Static Routes for a Network
        api_response = api_instance.get_static_routes(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_static_routes: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**GetStaticRoutes200Response**](GetStaticRoutes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_subnet**
> CreateSubnet200Response get_subnet(id)

Get a Specific Subnet

This endpoint retrieves a specific Subnet. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_subnet200_response import CreateSubnet200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Get a Specific Subnet
        api_response = api_instance.get_subnet(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->get_subnet: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**CreateSubnet200Response**](CreateSubnet200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_network_pool_server_types**
> ListNetworkPoolServerTypes200Response list_network_pool_server_types()

Get All Network Pool Server Types

This endpoint retrieves all Network Pool Server Types 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_network_pool_server_types200_response import ListNetworkPoolServerTypes200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Network Pool Server Types
        api_response = api_instance.list_network_pool_server_types(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name, code=code)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_pool_server_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]

### Return type

[**ListNetworkPoolServerTypes200Response**](ListNetworkPoolServerTypes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_network_pool_servers**
> ListNetworkPoolServers200Response list_network_pool_servers()

Get All Network Pool Servers

This endpoint retrieves all Network Pool Servers associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.list_network_pool_servers200_response import ListNetworkPoolServers200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Network Pool Servers
        api_response = api_instance.list_network_pool_servers(max=max, offset=offset, sort=sort, direction=direction, phrase=phrase, name=name)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_pool_servers: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

### Return type

[**ListNetworkPoolServers200Response**](ListNetworkPoolServers200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_network_router_types**
> ListNetworkRouterTypes200Response list_network_router_types()

Get all Network Router Types

Get all Network Router Types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.list_network_router_types200_response import ListNetworkRouterTypes200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get all Network Router Types
        api_response = api_instance.list_network_router_types()
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_router_types: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**ListNetworkRouterTypes200Response**](ListNetworkRouterTypes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_network_server_groups**
> ListNetworkServerGroups200Response list_network_server_groups(server_id)

Get all Network Server Groups for Network Server

This endpoint retrieves all Network Server Groups for a specified Network Service. Note: Only available for NSX-T network integrations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.list_network_server_groups200_response import ListNetworkServerGroups200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    server_id = 4 # float | Server ID
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get all Network Server Groups for Network Server
        api_response = api_instance.list_network_server_groups(server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_server_groups: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get all Network Server Groups for Network Server
        api_response = api_instance.list_network_server_groups(server_id, max=max, offset=offset, sort=sort, direction=direction, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_server_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server_id** | **float**| Server ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListNetworkServerGroups200Response**](ListNetworkServerGroups200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_network_services**
> ListNetworkServices200Response list_network_services()

Get All Network Services

This endpoint retrieves all Network Services associated with the account.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.list_network_services200_response import ListNetworkServices200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Network Services
        api_response = api_instance.list_network_services(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_services: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListNetworkServices200Response**](ListNetworkServices200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_network_types**
> ListNetworkTypes200Response list_network_types()

Network Types

Provides API for viewing Network Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.list_network_types200_response import ListNetworkTypes200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code = "azr" # str | If specified will return an exact match on code (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Network Types
        api_response = api_instance.list_network_types(name=name, code=code, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_network_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **str**| If specified will return an exact match on code | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListNetworkTypes200Response**](ListNetworkTypes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_networks**
> ListNetworks200Response list_networks()

Get All Networks

This endpoint retrieves all Networks associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.list_networks200_response import ListNetworks200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Networks
        api_response = api_instance.list_networks(name=name, phrase=phrase, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_networks: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**ListNetworks200Response**](ListNetworks200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_subnet_types**
> ListSubnetTypes200Response list_subnet_types()

Get All Subnet Types

Provides API for viewing Network Subnet Types and their configuration options.  This endpoint retrieves all Network Types. The sample response has been abbreviated. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.list_subnet_types200_response import ListSubnetTypes200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Subnet Types
        api_response = api_instance.list_subnet_types(name=name, phrase=phrase)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_subnet_types: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]

### Return type

[**ListSubnetTypes200Response**](ListSubnetTypes200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_subnets**
> GetNetworkSubnets200Response list_subnets()

Get All Subnets

This endpoint retrieves all Subnets associated with the account. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.get_network_subnets200_response import GetNetworkSubnets200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    name = "example-%" # str | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    labels = "labels_example" # str | Filter by label(s), matches records that contain any of the specified labels (optional)
    all_labels = "allLabels_example" # str | Filter by label(s), matches records that contain all of the specified labels (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get All Subnets
        api_response = api_instance.list_subnets(name=name, phrase=phrase, labels=labels, all_labels=all_labels)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->list_subnets: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **labels** | **str**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **str**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

[**GetNetworkSubnets200Response**](GetNetworkSubnets200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **refresh_network_server**
> SuccessId refresh_network_server(id)

Refresh a Network Server/Integration

Refreshes a network server/integration. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Refresh a Network Server/Integration
        api_response = api_instance.refresh_network_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->refresh_network_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **release_network_floating_ip**
> Model200Success release_network_floating_ip(id)

Release a Floating IP

Release a floating IP detaching it from the associated node/VM. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced

    # example passing only required values which don't have defaults set
    try:
        # Release a Floating IP
        api_response = api_instance.release_network_floating_ip(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->release_network_floating_ip: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network**
> CreateNetworks200Response update_network(id)

Update a Network

This endpoint allows updating a Network. Configuration options vary by Network Types. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.update_network_request import UpdateNetworkRequest
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_networks200_response import CreateNetworks200Response
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_network_request = UpdateNetworkRequest(
        network=NetworkUpdate(
            display_name="display_name_example",
            labels=[
                "labels_example",
            ],
            description="description_example",
            cidr="cidr_example",
            gateway="gateway_example",
            dns_primary="dns_primary_example",
            dns_secondary="dns_secondary_example",
            vlan_id=1,
            pool=1,
            allow_static_override=True,
            assign_public_ip=True,
            active=True,
            dhcp_server=True,
            network_domain=NetworkCreateNetworkDomain(
                id=1,
            ),
            search_domains="search_domains_example",
            network_proxy=NetworkCreateNetworkProxy(
                id=1,
            ),
            appliance_url_proxy_bypass=True,
            no_proxy="no_proxy_example",
            visibility="private",
            config={},
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            resource_permissions=NetworkCreateResourcePermissions(
                all=True,
                sites=[
                    1,
                ],
            ),
        ),
    ) # UpdateNetworkRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network
        api_response = api_instance.update_network(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network
        api_response = api_instance.update_network(id, update_network_request=update_network_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_network_request** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md)|  | [optional]

### Return type

[**CreateNetworks200Response**](CreateNetworks200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_dhcp_relay**
> Model200Success update_network_dhcp_relay(id, server_id)

Update a Network DHCP Relay

Use this command to update an existing Network DHCP Relay. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.update_network_dhcp_relay_request import UpdateNetworkDhcpRelayRequest
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_dhcp_relay_request = UpdateNetworkDhcpRelayRequest(
        network_dhcp_relay={},
    ) # UpdateNetworkDhcpRelayRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network DHCP Relay
        api_response = api_instance.update_network_dhcp_relay(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_dhcp_relay: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network DHCP Relay
        api_response = api_instance.update_network_dhcp_relay(id, server_id, update_network_dhcp_relay_request=update_network_dhcp_relay_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_dhcp_relay: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_dhcp_relay_request** | [**UpdateNetworkDhcpRelayRequest**](UpdateNetworkDhcpRelayRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_dhcp_server**
> Model200Success update_network_dhcp_server(id, server_id)

Update a Network DHCP Server

Use this command to update an existing Network DHCP Server. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_network_dhcp_server_request import UpdateNetworkDhcpServerRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_dhcp_server_request = UpdateNetworkDhcpServerRequest(
        network_dhcp_server={},
    ) # UpdateNetworkDhcpServerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network DHCP Server
        api_response = api_instance.update_network_dhcp_server(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_dhcp_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network DHCP Server
        api_response = api_instance.update_network_dhcp_server(id, server_id, update_network_dhcp_server_request=update_network_dhcp_server_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_dhcp_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_dhcp_server_request** | [**UpdateNetworkDhcpServerRequest**](UpdateNetworkDhcpServerRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_domain**
> CreateNetworkDomain200Response update_network_domain(id)

Update a Network Domain

Update a Network Domain. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_domain200_response import CreateNetworkDomain200Response
from openapi_client.model.create_network_domain_request import CreateNetworkDomainRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_network_domain_request = CreateNetworkDomainRequest(
        network_domain=NetworkDomainCreate(
            name="name_example",
            description="description_example",
            display_name="display_name_example",
            public_zone=False,
            task_set_id=1,
            active=True,
            domain_controller=True,
            domain_username="domain_username_example",
            domain_password="domain_password_example",
            dc_server="dc_server_example",
            ou_path="ou_path_example",
            guest_username="guest_username_example",
            guest_password="guest_password_example",
        ),
    ) # CreateNetworkDomainRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Domain
        api_response = api_instance.update_network_domain(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_domain: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Domain
        api_response = api_instance.update_network_domain(id, create_network_domain_request=create_network_domain_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_domain: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_network_domain_request** | [**CreateNetworkDomainRequest**](CreateNetworkDomainRequest.md)|  | [optional]

### Return type

[**CreateNetworkDomain200Response**](CreateNetworkDomain200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_edge_cluster**
> Model200Success update_network_edge_cluster(id, server_id)

Update a Network Edge Cluster

Use this command to update an existing network Edge Cluster. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_network_edge_cluster_request import UpdateNetworkEdgeClusterRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_edge_cluster_request = UpdateNetworkEdgeClusterRequest(
        network_edge_cluster={},
    ) # UpdateNetworkEdgeClusterRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Edge Cluster
        api_response = api_instance.update_network_edge_cluster(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_edge_cluster: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Edge Cluster
        api_response = api_instance.update_network_edge_cluster(id, server_id, update_network_edge_cluster_request=update_network_edge_cluster_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_edge_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_edge_cluster_request** | [**UpdateNetworkEdgeClusterRequest**](UpdateNetworkEdgeClusterRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_firewall_rule**
> Model200Success update_network_firewall_rule(id, server_id)

Update a Network Firewall Rule

Use this command to update an existing network firewall Rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_network_firewall_rule_request import UpdateNetworkFirewallRuleRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_firewall_rule_request = UpdateNetworkFirewallRuleRequest(
        rule={},
    ) # UpdateNetworkFirewallRuleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Firewall Rule
        api_response = api_instance.update_network_firewall_rule(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_firewall_rule: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Firewall Rule
        api_response = api_instance.update_network_firewall_rule(id, server_id, update_network_firewall_rule_request=update_network_firewall_rule_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_firewall_rule_request** | [**UpdateNetworkFirewallRuleRequest**](UpdateNetworkFirewallRuleRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_firewall_rule_group**
> Model200Success update_network_firewall_rule_group(id, server_id)

Update a Network Firewall Rule Group

Use this command to update an existing Network Firewall Rule Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_network_firewall_rule_group_request import UpdateNetworkFirewallRuleGroupRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_firewall_rule_group_request = UpdateNetworkFirewallRuleGroupRequest(
        rule_group={},
    ) # UpdateNetworkFirewallRuleGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Firewall Rule Group
        api_response = api_instance.update_network_firewall_rule_group(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_firewall_rule_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Firewall Rule Group
        api_response = api_instance.update_network_firewall_rule_group(id, server_id, update_network_firewall_rule_group_request=update_network_firewall_rule_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_firewall_rule_group_request** | [**UpdateNetworkFirewallRuleGroupRequest**](UpdateNetworkFirewallRuleGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_group**
> Model200Success update_network_group(id)

Update a Network Group

Use this command to update an existing network Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_group_request import CreateNetworkGroupRequest
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_network_group_request = CreateNetworkGroupRequest(
        network_group=NetworkGroupsCreate(
            name="name_example",
            description="description_example",
            networks=[
                1,
            ],
            subnets=[
                {},
            ],
        ),
    ) # CreateNetworkGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Group
        api_response = api_instance.update_network_group(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Group
        api_response = api_instance.update_network_group(id, create_network_group_request=create_network_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_network_group_request** | [**CreateNetworkGroupRequest**](CreateNetworkGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_pool**
> CreateNetworkPool200Response update_network_pool(id)

Update a Network Pool

Update a Network Pool. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_network_pool200_response import CreateNetworkPool200Response
from openapi_client.model.create_network_pool_request import CreateNetworkPoolRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_network_pool_request = CreateNetworkPoolRequest(
        network_pool=NetworkPoolCreate(
            name="name_example",
            type=NetworkPoolCreateType(
                code=None,
            ),
            ip_ranges=[
                NetworkPoolCreateIpRangesInner(
                    start_address="start_address_example",
                    end_address="end_address_example",
                    cidr_ipv6="cidr_ipv6_example",
                ),
            ],
            config={},
        ),
    ) # CreateNetworkPoolRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Pool
        api_response = api_instance.update_network_pool(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_pool: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Pool
        api_response = api_instance.update_network_pool(id, create_network_pool_request=create_network_pool_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_pool: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_network_pool_request** | [**CreateNetworkPoolRequest**](CreateNetworkPoolRequest.md)|  | [optional]

### Return type

[**CreateNetworkPool200Response**](CreateNetworkPool200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_pool_server**
> CreateNetworkPoolServer200Response update_network_pool_server(id)

Update a Network Pool Server

This endpoint allows updating a Network Pool Server. Configuration options vary by type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.update_network_pool_server_request import UpdateNetworkPoolServerRequest
from openapi_client.model.create_network_pool_server200_response import CreateNetworkPoolServer200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_network_pool_server_request = UpdateNetworkPoolServerRequest(
        network_pool_server=UpdateNetworkPoolServerRequestNetworkPoolServer(),
    ) # UpdateNetworkPoolServerRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Pool Server
        api_response = api_instance.update_network_pool_server(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_pool_server: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Pool Server
        api_response = api_instance.update_network_pool_server(id, update_network_pool_server_request=update_network_pool_server_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_pool_server: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_network_pool_server_request** | [**UpdateNetworkPoolServerRequest**](UpdateNetworkPoolServerRequest.md)|  | [optional]

### Return type

[**CreateNetworkPoolServer200Response**](CreateNetworkPoolServer200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_proxy**
> Model200Success update_network_proxy(id)

Update a Network Proxy

Use this command to update an existing network Proxy. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.create_network_proxy_request import CreateNetworkProxyRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_network_proxy_request = CreateNetworkProxyRequest(
        network_proxy=CreateNetworkProxyRequestNetworkProxy(
            name="name_example",
            proxy_host="proxy_host_example",
            proxy_port="proxy_port_example",
            proxy_user="proxy_user_example",
            proxy_password="proxy_password_example",
            proxy_domain="proxy_domain_example",
            proxy_workstation="proxy_workstation_example",
            visibility="private",
            account=CreateNetworkProxyRequestNetworkProxyAccount(
                id=1,
            ),
        ),
    ) # CreateNetworkProxyRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Proxy
        api_response = api_instance.update_network_proxy(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_proxy: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Proxy
        api_response = api_instance.update_network_proxy(id, create_network_proxy_request=create_network_proxy_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_proxy: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_network_proxy_request** | [**CreateNetworkProxyRequest**](CreateNetworkProxyRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_router**
> Model200Success update_network_router(id)

Update a Network Router

Use this command to update an existing network Router. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.update_network_router_request import UpdateNetworkRouterRequest
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    update_network_router_request = UpdateNetworkRouterRequest(
        network_router=NetworkRoutersUpdate(
            name="name_example",
            type=NetworkRoutersUpdateType(
                id=1,
            ),
            site=NetworkRoutersUpdateSite(
                id=NetworkRoutersCreateSiteId(None),
            ),
            enabled=True,
            zone=NetworkRoutersUpdateZone(
                id=1,
            ),
            network_server=NetworkRoutersUpdateNetworkServer(
                id=1,
            ),
        ),
    ) # UpdateNetworkRouterRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Router
        api_response = api_instance.update_network_router(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Router
        api_response = api_instance.update_network_router(id, update_network_router_request=update_network_router_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **update_network_router_request** | [**UpdateNetworkRouterRequest**](UpdateNetworkRouterRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_router_bgp_neighbor**
> Model200Success update_network_router_bgp_neighbor(id, router_id)

Update Network Router BGP Neighbor

Use this command to update an existing Network Router BGP Neighbor. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_network_router_bgp_neighbor_request import UpdateNetworkRouterBgpNeighborRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID
    update_network_router_bgp_neighbor_request = UpdateNetworkRouterBgpNeighborRequest(
        network_router_bgp_neighbor={},
    ) # UpdateNetworkRouterBgpNeighborRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Network Router BGP Neighbor
        api_response = api_instance.update_network_router_bgp_neighbor(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_bgp_neighbor: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Network Router BGP Neighbor
        api_response = api_instance.update_network_router_bgp_neighbor(id, router_id, update_network_router_bgp_neighbor_request=update_network_router_bgp_neighbor_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_bgp_neighbor: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **update_network_router_bgp_neighbor_request** | [**UpdateNetworkRouterBgpNeighborRequest**](UpdateNetworkRouterBgpNeighborRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_router_firewall_rule**
> Model200Success update_network_router_firewall_rule(id, router_id)

Update a Network Router Firewall Rule

Use this command to update an existing network router firewall rule. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.update_network_router_firewall_rule_request import UpdateNetworkRouterFirewallRuleRequest
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID
    update_network_router_firewall_rule_request = UpdateNetworkRouterFirewallRuleRequest(
        rule=NetworkRouterFirewallRuleCreate(
            name="name_example",
            enabled=True,
            priority=1,
        ),
    ) # UpdateNetworkRouterFirewallRuleRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Router Firewall Rule
        api_response = api_instance.update_network_router_firewall_rule(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_firewall_rule: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Router Firewall Rule
        api_response = api_instance.update_network_router_firewall_rule(id, router_id, update_network_router_firewall_rule_request=update_network_router_firewall_rule_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_firewall_rule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **update_network_router_firewall_rule_request** | [**UpdateNetworkRouterFirewallRuleRequest**](UpdateNetworkRouterFirewallRuleRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_router_firewall_rule_group**
> Model200Success update_network_router_firewall_rule_group(id, router_id)

Update a Network Router Firewall Rule Group

Use this command to update an existing Network Router Firewall Rule Group. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from openapi_client.model.update_network_router_firewall_rule_group_request import UpdateNetworkRouterFirewallRuleGroupRequest
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID
    update_network_router_firewall_rule_group_request = UpdateNetworkRouterFirewallRuleGroupRequest(
        rule_group={},
    ) # UpdateNetworkRouterFirewallRuleGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Router Firewall Rule Group
        api_response = api_instance.update_network_router_firewall_rule_group(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_firewall_rule_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Router Firewall Rule Group
        api_response = api_instance.update_network_router_firewall_rule_group(id, router_id, update_network_router_firewall_rule_group_request=update_network_router_firewall_rule_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_firewall_rule_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **update_network_router_firewall_rule_group_request** | [**UpdateNetworkRouterFirewallRuleGroupRequest**](UpdateNetworkRouterFirewallRuleGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_router_nat**
> Model200Success update_network_router_nat(id, router_id)

Update Network Router NAT

Use this command to update an existing Network Router NAT. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_network_router_nat_request import UpdateNetworkRouterNatRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    router_id = 4 # float | Router ID
    update_network_router_nat_request = UpdateNetworkRouterNatRequest(
        network_router_nat=UpdateNetworkRouterNatRequestNetworkRouterNAT(
            name=None,
        ),
    ) # UpdateNetworkRouterNatRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Network Router NAT
        api_response = api_instance.update_network_router_nat(id, router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_nat: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Network Router NAT
        api_response = api_instance.update_network_router_nat(id, router_id, update_network_router_nat_request=update_network_router_nat_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_nat: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **router_id** | **float**| Router ID |
 **update_network_router_nat_request** | [**UpdateNetworkRouterNatRequest**](UpdateNetworkRouterNatRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_router_permissions**
> SuccessId update_network_router_permissions(router_id)

Update Network Router Permissions

Update Network Router Permissions 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.update_network_router_permissions_request import UpdateNetworkRouterPermissionsRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    router_id = 4 # float | Router ID
    update_network_router_permissions_request = UpdateNetworkRouterPermissionsRequest(
        permissions=NetworkRouterPermissionsUpdate(
            visibility="visibility_example",
            tenant_permissions=NetworkRouterPermissionsUpdateTenantPermissions(
                accounts=[
                    1,
                ],
            ),
        ),
    ) # UpdateNetworkRouterPermissionsRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update Network Router Permissions
        api_response = api_instance.update_network_router_permissions(router_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_permissions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update Network Router Permissions
        api_response = api_instance.update_network_router_permissions(router_id, update_network_router_permissions_request=update_network_router_permissions_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_router_permissions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **router_id** | **float**| Router ID |
 **update_network_router_permissions_request** | [**UpdateNetworkRouterPermissionsRequest**](UpdateNetworkRouterPermissionsRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_server_group**
> Model200Success update_network_server_group(id, server_id)

Update a Network Server Group

Use this command to update an existing network server group. Note: Only available for NSX-T network integrations. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_network_server_group_request import UpdateNetworkServerGroupRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_server_group_request = UpdateNetworkServerGroupRequest(
        group={},
    ) # UpdateNetworkServerGroupRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Server Group
        api_response = api_instance.update_network_server_group(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_server_group: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Server Group
        api_response = api_instance.update_network_server_group(id, server_id, update_network_server_group_request=update_network_server_group_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_server_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_server_group_request** | [**UpdateNetworkServerGroupRequest**](UpdateNetworkServerGroupRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_network_transport_zone**
> Model200Success update_network_transport_zone(id, server_id)

Update a Network Transport Zone

Use this command to update an existing network Transport Zone. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.update_network_transport_zone_request import UpdateNetworkTransportZoneRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    server_id = 4 # float | Server ID
    update_network_transport_zone_request = UpdateNetworkTransportZoneRequest(
        network_scope=NetworkTransportZoneCreate(
            name="name_example",
            description="description_example",
            visibility="visibility_example",
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
        ),
    ) # UpdateNetworkTransportZoneRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Transport Zone
        api_response = api_instance.update_network_transport_zone(id, server_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_transport_zone: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Transport Zone
        api_response = api_instance.update_network_transport_zone(id, server_id, update_network_transport_zone_request=update_network_transport_zone_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_network_transport_zone: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **server_id** | **float**| Server ID |
 **update_network_transport_zone_request** | [**UpdateNetworkTransportZoneRequest**](UpdateNetworkTransportZoneRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_static_route**
> SuccessId update_static_route(id, route_id)

Update a Network Static Route

Use this command to update a route. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_static_route_request import CreateStaticRouteRequest
from openapi_client.model.success_id import SuccessId
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    route_id = 4 # float | The ID of the route
    create_static_route_request = CreateStaticRouteRequest(
        network_route=NetworkStaticRouteCreate(
            source="source_example",
            destination="destination_example",
        ),
    ) # CreateStaticRouteRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Network Static Route
        api_response = api_instance.update_static_route(id, route_id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_static_route: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Network Static Route
        api_response = api_instance.update_static_route(id, route_id, create_static_route_request=create_static_route_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_static_route: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **route_id** | **float**| The ID of the route |
 **create_static_route_request** | [**CreateStaticRouteRequest**](CreateStaticRouteRequest.md)|  | [optional]

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_subnet**
> CreateSubnet200Response update_subnet(id)

Update a Subnet

This endpoint allows updating a Subnet. Only certain types of clouds support this action. Configuration options vary for each Subnet Type. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import networks_api
from openapi_client.model.create_subnet200_response import CreateSubnet200Response
from openapi_client.model.create_subnet_request import CreateSubnetRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = networks_api.NetworksApi(api_client)
    id = 1 # int | Morpheus ID of the Object being referenced
    create_subnet_request = CreateSubnetRequest(
        subnet=CreateSubnetRequestSubnet(
            type=CreateSubnetRequestSubnetType(
                id=1,
            ),
            config={},
            tenants=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            visibility="private",
            labels=[
                "labels_example",
            ],
        ),
        resource_permission=CreateSubnetRequestResourcePermission(
            all=True,
            sites=[
                UpdateBlueprintPermissionsRequestResourcePermissionSitesInner(
                    id=1,
                ),
            ],
            all_plans=True,
            plans=[
                {},
            ],
        ),
    ) # CreateSubnetRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Update a Subnet
        api_response = api_instance.update_subnet(id)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_subnet: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Update a Subnet
        api_response = api_instance.update_subnet(id, create_subnet_request=create_subnet_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling NetworksApi->update_subnet: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **create_subnet_request** | [**CreateSubnetRequest**](CreateSubnetRequest.md)|  | [optional]

### Return type

[**CreateSubnet200Response**](CreateSubnet200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

