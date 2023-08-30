"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.cluster_layout_compute_servers_inner_compute_server_type import ClusterLayoutComputeServersInnerComputeServerType
from openapi_client.model.get_network_pool_ips200_response_all_of_network_pool_ips_inner_created_by import GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy
from openapi_client.model.guidance_vmware_sizing_resource_controllers_inner import GuidanceVmwareSizingResourceControllersInner
from openapi_client.model.guidance_vmware_sizing_resource_interfaces_inner import GuidanceVmwareSizingResourceInterfacesInner
from openapi_client.model.guidance_vmware_sizing_resource_server_os import GuidanceVmwareSizingResourceServerOs
from openapi_client.model.guidance_vmware_sizing_resource_stats import GuidanceVmwareSizingResourceStats
from openapi_client.model.guidance_vmware_sizing_resource_volumes_inner import GuidanceVmwareSizingResourceVolumesInner
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.list_load_balancer_monitors200_response_all_of_load_balancer_monitors_inner_load_balancer_type import ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
globals()['ClusterLayoutComputeServersInnerComputeServerType'] = ClusterLayoutComputeServersInnerComputeServerType
globals()['GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy'] = GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy
globals()['GuidanceVmwareSizingResourceControllersInner'] = GuidanceVmwareSizingResourceControllersInner
globals()['GuidanceVmwareSizingResourceInterfacesInner'] = GuidanceVmwareSizingResourceInterfacesInner
globals()['GuidanceVmwareSizingResourceServerOs'] = GuidanceVmwareSizingResourceServerOs
globals()['GuidanceVmwareSizingResourceStats'] = GuidanceVmwareSizingResourceStats
globals()['GuidanceVmwareSizingResourceVolumesInner'] = GuidanceVmwareSizingResourceVolumesInner
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType'] = ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
from openapi_client.model.guidance_vmware_sizing_resource import GuidanceVmwareSizingResource


class TestGuidanceVmwareSizingResource(unittest.TestCase):
    """GuidanceVmwareSizingResource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testGuidanceVmwareSizingResource(self):
        """Test GuidanceVmwareSizingResource"""
        # FIXME: construct object with mandatory attributes with example values
        # model = GuidanceVmwareSizingResource()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
