"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.prices_api import PricesApi  # noqa: E501


class TestPricesApi(unittest.TestCase):
    """PricesApi unit test stubs"""

    def setUp(self):
        self.api = PricesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_prices(self):
        """Test case for add_prices

        Creates a Price  # noqa: E501
        """
        pass

    def test_deactivate_prices(self):
        """Test case for deactivate_prices

        Deactivates a Price  # noqa: E501
        """
        pass

    def test_get_prices(self):
        """Test case for get_prices

        Retrieves a Specific Price  # noqa: E501
        """
        pass

    def test_list_prices(self):
        """Test case for list_prices

        Retrieves all Prices  # noqa: E501
        """
        pass

    def test_update_prices(self):
        """Test case for update_prices

        Updates a Price  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()