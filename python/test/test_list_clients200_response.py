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
from openapi_client.model.client import Client
from openapi_client.model.list_clients200_response_all_of import ListClients200ResponseAllOf
from openapi_client.model.meta import Meta
from openapi_client.model.meta_meta import MetaMeta
globals()['Client'] = Client
globals()['ListClients200ResponseAllOf'] = ListClients200ResponseAllOf
globals()['Meta'] = Meta
globals()['MetaMeta'] = MetaMeta
from openapi_client.model.list_clients200_response import ListClients200Response


class TestListClients200Response(unittest.TestCase):
    """ListClients200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testListClients200Response(self):
        """Test ListClients200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ListClients200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
