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
from openapi_client.model.check_type import CheckType
from openapi_client.model.list_check_types200_response_all_of import ListCheckTypes200ResponseAllOf
from openapi_client.model.meta import Meta
from openapi_client.model.meta_meta import MetaMeta
globals()['CheckType'] = CheckType
globals()['ListCheckTypes200ResponseAllOf'] = ListCheckTypes200ResponseAllOf
globals()['Meta'] = Meta
globals()['MetaMeta'] = MetaMeta
from openapi_client.model.list_check_types200_response import ListCheckTypes200Response


class TestListCheckTypes200Response(unittest.TestCase):
    """ListCheckTypes200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testListCheckTypes200Response(self):
        """Test ListCheckTypes200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ListCheckTypes200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
