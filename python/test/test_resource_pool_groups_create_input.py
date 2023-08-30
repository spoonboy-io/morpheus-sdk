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
from openapi_client.model.get_resource_pool_groups200_response_resource_pool_groups_inner_resource_permission import GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission'] = GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.resource_pool_groups_create_input import ResourcePoolGroupsCreateInput


class TestResourcePoolGroupsCreateInput(unittest.TestCase):
    """ResourcePoolGroupsCreateInput unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testResourcePoolGroupsCreateInput(self):
        """Test ResourcePoolGroupsCreateInput"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ResourcePoolGroupsCreateInput()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()