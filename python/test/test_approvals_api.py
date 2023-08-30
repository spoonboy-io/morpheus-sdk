"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.approvals_api import ApprovalsApi  # noqa: E501


class TestApprovalsApi(unittest.TestCase):
    """ApprovalsApi unit test stubs"""

    def setUp(self):
        self.api = ApprovalsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_approvals(self):
        """Test case for get_approvals

        Retrieves a Specific Approval  # noqa: E501
        """
        pass

    def test_get_approvals_item(self):
        """Test case for get_approvals_item

        Retrieves a Specific Approval Item  # noqa: E501
        """
        pass

    def test_list_approvals(self):
        """Test case for list_approvals

        Retrieves all Approvals  # noqa: E501
        """
        pass

    def test_update_approvals_item(self):
        """Test case for update_approvals_item

        Updates a Specific Approval Item  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
