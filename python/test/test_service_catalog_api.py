"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.service_catalog_api import ServiceCatalogApi  # noqa: E501


class TestServiceCatalogApi(unittest.TestCase):
    """ServiceCatalogApi unit test stubs"""

    def setUp(self):
        self.api = ServiceCatalogApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_catalog_cart(self):
        """Test case for add_catalog_cart

        Checkout Catalog Cart  # noqa: E501
        """
        pass

    def test_add_catalog_cart_item(self):
        """Test case for add_catalog_cart_item

        Add Catalog Item to Cart  # noqa: E501
        """
        pass

    def test_add_catalog_order(self):
        """Test case for add_catalog_order

        Place Catalog Order  # noqa: E501
        """
        pass

    def test_delete_catalog_cart(self):
        """Test case for delete_catalog_cart

        Clear Catalog Cart  # noqa: E501
        """
        pass

    def test_delete_catalog_cart_item(self):
        """Test case for delete_catalog_cart_item

        Remove a Catalog Item From Cart  # noqa: E501
        """
        pass

    def test_delete_catalog_item(self):
        """Test case for delete_catalog_item

        Delete a Catalog Inventory Item  # noqa: E501
        """
        pass

    def test_get_catalog_item(self):
        """Test case for get_catalog_item

        Get a Specific Catalog Inventory Item  # noqa: E501
        """
        pass

    def test_get_catalog_type(self):
        """Test case for get_catalog_type

        Get a Specific Catalog Type  # noqa: E501
        """
        pass

    def test_list_catalog_cart(self):
        """Test case for list_catalog_cart

        Get Catalog Cart  # noqa: E501
        """
        pass

    def test_list_catalog_items(self):
        """Test case for list_catalog_items

        List Catalog Inventory Items  # noqa: E501
        """
        pass

    def test_list_catalog_types(self):
        """Test case for list_catalog_types

        List Catalog Types  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()