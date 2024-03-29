"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.cypher_api import CypherApi  # noqa: E501


class TestCypherApi(unittest.TestCase):
    """CypherApi unit test stubs"""

    def setUp(self):
        self.api = CypherApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_cypher_key(self):
        """Test case for add_cypher_key

        Write a Cypher  # noqa: E501
        """
        pass

    def test_get_cypher_key(self):
        """Test case for get_cypher_key

        Read or Create a Cypher Key  # noqa: E501
        """
        pass

    def test_list_cypher_keys(self):
        """Test case for list_cypher_keys

        List Cypher Keys  # noqa: E501
        """
        pass

    def test_remove_cypher(self):
        """Test case for remove_cypher

        Delete a Cypher  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
