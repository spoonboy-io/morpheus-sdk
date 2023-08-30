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
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.role import Role
from openapi_client.model.role_app_template_permissions_inner import RoleAppTemplatePermissionsInner
from openapi_client.model.role_feature_permissions_inner import RoleFeaturePermissionsInner
from openapi_client.model.role_instance_type_permissions_inner import RoleInstanceTypePermissionsInner
from openapi_client.model.role_role import RoleRole
from openapi_client.model.role_sites_inner import RoleSitesInner
globals()['Model200Success'] = Model200Success
globals()['Role'] = Role
globals()['RoleAppTemplatePermissionsInner'] = RoleAppTemplatePermissionsInner
globals()['RoleFeaturePermissionsInner'] = RoleFeaturePermissionsInner
globals()['RoleInstanceTypePermissionsInner'] = RoleInstanceTypePermissionsInner
globals()['RoleRole'] = RoleRole
globals()['RoleSitesInner'] = RoleSitesInner
from openapi_client.model.add_roles200_response import AddRoles200Response


class TestAddRoles200Response(unittest.TestCase):
    """AddRoles200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAddRoles200Response(self):
        """Test AddRoles200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AddRoles200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()