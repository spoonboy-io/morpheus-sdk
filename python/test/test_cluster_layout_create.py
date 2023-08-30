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
from openapi_client.model.add_service_plans_request_service_plan_provision_type_inner import AddServicePlansRequestServicePlanProvisionTypeInner
from openapi_client.model.cluster_layout_create_environment_variables_inner import ClusterLayoutCreateEnvironmentVariablesInner
from openapi_client.model.cluster_layout_create_group_type import ClusterLayoutCreateGroupType
from openapi_client.model.cluster_layout_create_masters_inner import ClusterLayoutCreateMastersInner
from openapi_client.model.update_blueprint_permissions_request_resource_permission_sites_inner import UpdateBlueprintPermissionsRequestResourcePermissionSitesInner
globals()['AddServicePlansRequestServicePlanProvisionTypeInner'] = AddServicePlansRequestServicePlanProvisionTypeInner
globals()['ClusterLayoutCreateEnvironmentVariablesInner'] = ClusterLayoutCreateEnvironmentVariablesInner
globals()['ClusterLayoutCreateGroupType'] = ClusterLayoutCreateGroupType
globals()['ClusterLayoutCreateMastersInner'] = ClusterLayoutCreateMastersInner
globals()['UpdateBlueprintPermissionsRequestResourcePermissionSitesInner'] = UpdateBlueprintPermissionsRequestResourcePermissionSitesInner
from openapi_client.model.cluster_layout_create import ClusterLayoutCreate


class TestClusterLayoutCreate(unittest.TestCase):
    """ClusterLayoutCreate unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testClusterLayoutCreate(self):
        """Test ClusterLayoutCreate"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ClusterLayoutCreate()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
