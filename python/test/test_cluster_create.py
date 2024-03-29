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
from openapi_client.model.cluster_create_cloud import ClusterCreateCloud
from openapi_client.model.cluster_create_group import ClusterCreateGroup
from openapi_client.model.cluster_create_layout import ClusterCreateLayout
from openapi_client.model.cluster_create_type import ClusterCreateType
from openapi_client.model.cluster_server_create import ClusterServerCreate
globals()['ClusterCreateCloud'] = ClusterCreateCloud
globals()['ClusterCreateGroup'] = ClusterCreateGroup
globals()['ClusterCreateLayout'] = ClusterCreateLayout
globals()['ClusterCreateType'] = ClusterCreateType
globals()['ClusterServerCreate'] = ClusterServerCreate
from openapi_client.model.cluster_create import ClusterCreate


class TestClusterCreate(unittest.TestCase):
    """ClusterCreate unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testClusterCreate(self):
        """Test ClusterCreate"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ClusterCreate()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
