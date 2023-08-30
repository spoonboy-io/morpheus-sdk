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
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.list_load_balancer_monitors200_response_all_of_load_balancer_monitors_inner_load_balancer_type import ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType'] = ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType
from openapi_client.model.user_settings_user import UserSettingsUser


class TestUserSettingsUser(unittest.TestCase):
    """UserSettingsUser unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testUserSettingsUser(self):
        """Test UserSettingsUser"""
        # FIXME: construct object with mandatory attributes with example values
        # model = UserSettingsUser()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
