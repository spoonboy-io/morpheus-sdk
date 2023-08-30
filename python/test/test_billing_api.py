"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.billing_api import BillingApi  # noqa: E501


class TestBillingApi(unittest.TestCase):
    """BillingApi unit test stubs"""

    def setUp(self):
        self.api = BillingApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_billing_account(self):
        """Test case for get_billing_account

        This endpoint will retrieve a specific account by id if the user has permission to access it  # noqa: E501
        """
        pass

    def test_get_billing_instances_identifier(self):
        """Test case for get_billing_instances_identifier

        Retrieves billing information for an instance in the requestor's account. Use instanceUUID whenever possible.  # noqa: E501
        """
        pass

    def test_get_billing_servers_identifier(self):
        """Test case for get_billing_servers_identifier

        Retrieves billing information for a specific server (container host) in the requestor's account. Use refUUID whenever possible.  # noqa: E501
        """
        pass

    def test_get_billing_zone_identifier(self):
        """Test case for get_billing_zone_identifier

        Retrieves billing information for a specific zone in the requestor's account. Use zoneUUID whenever possible.  # noqa: E501
        """
        pass

    def test_list_billing_account(self):
        """Test case for list_billing_account

        Retrieves billing information for the requesting user's account.  # noqa: E501
        """
        pass

    def test_list_billing_instances(self):
        """Test case for list_billing_instances

        Retrieves billing information for all instances on the requestor's account.  # noqa: E501
        """
        pass

    def test_list_billing_servers(self):
        """Test case for list_billing_servers

        Retrieves billing information for all servers (container hosts) on the requestor's account.  # noqa: E501
        """
        pass

    def test_list_billing_zone(self):
        """Test case for list_billing_zone

        Retrieves billing information for all zones on the requestor's account.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
