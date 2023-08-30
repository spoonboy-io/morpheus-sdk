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
from openapi_client.model.role_permission_cloud import RolePermissionCloud
from openapi_client.model.role_permission_cloud_all import RolePermissionCloudAll
globals()['RolePermissionCloud'] = RolePermissionCloud
globals()['RolePermissionCloudAll'] = RolePermissionCloudAll
from openapi_client.model.update_role_cloud_access_request import UpdateRoleCloudAccessRequest


class TestUpdateRoleCloudAccessRequest(unittest.TestCase):
    """UpdateRoleCloudAccessRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testUpdateRoleCloudAccessRequest(self):
        """Test UpdateRoleCloudAccessRequest"""
        # FIXME: construct object with mandatory attributes with example values
        # model = UpdateRoleCloudAccessRequest()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()