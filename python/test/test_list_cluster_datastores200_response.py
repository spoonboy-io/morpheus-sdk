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
from openapi_client.model.cluster_datastores import ClusterDatastores
from openapi_client.model.list_cluster_datastores200_response_all_of import ListClusterDatastores200ResponseAllOf
from openapi_client.model.meta import Meta
from openapi_client.model.meta_meta import MetaMeta
globals()['ClusterDatastores'] = ClusterDatastores
globals()['ListClusterDatastores200ResponseAllOf'] = ListClusterDatastores200ResponseAllOf
globals()['Meta'] = Meta
globals()['MetaMeta'] = MetaMeta
from openapi_client.model.list_cluster_datastores200_response import ListClusterDatastores200Response


class TestListClusterDatastores200Response(unittest.TestCase):
    """ListClusterDatastores200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testListClusterDatastores200Response(self):
        """Test ListClusterDatastores200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ListClusterDatastores200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
