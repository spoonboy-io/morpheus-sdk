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
from openapi_client.model.get_network_pool_ips200_response_all_of_network_pool_ips_inner_created_by import GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy
from openapi_client.model.guidance_vmware_sizing_resource_controllers_inner import GuidanceVmwareSizingResourceControllersInner
from openapi_client.model.instance_config import InstanceConfig
from openapi_client.model.instance_connection_info_inner import InstanceConnectionInfoInner
from openapi_client.model.instance_instance_price import InstanceInstancePrice
from openapi_client.model.instance_instance_type import InstanceInstanceType
from openapi_client.model.instance_interfaces_inner import InstanceInterfacesInner
from openapi_client.model.instance_layout import InstanceLayout
from openapi_client.model.instance_stats import InstanceStats
from openapi_client.model.instance_volumes_inner import InstanceVolumesInner
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.list_load_balancer_monitors200_response_all_of_load_balancer_monitors_inner_load_balancer_type import ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
globals()['GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy'] = GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy
globals()['GuidanceVmwareSizingResourceControllersInner'] = GuidanceVmwareSizingResourceControllersInner
globals()['InstanceConfig'] = InstanceConfig
globals()['InstanceConnectionInfoInner'] = InstanceConnectionInfoInner
globals()['InstanceInstancePrice'] = InstanceInstancePrice
globals()['InstanceInstanceType'] = InstanceInstanceType
globals()['InstanceInterfacesInner'] = InstanceInterfacesInner
globals()['InstanceLayout'] = InstanceLayout
globals()['InstanceStats'] = InstanceStats
globals()['InstanceVolumesInner'] = InstanceVolumesInner
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType'] = ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
from openapi_client.model.instance import Instance


class TestInstance(unittest.TestCase):
    """Instance unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testInstance(self):
        """Test Instance"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Instance()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()