"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.groups_api import GroupsApi  # noqa: E501


class TestGroupsApi(unittest.TestCase):
    """GroupsApi unit test stubs"""

    def setUp(self):
        self.api = GroupsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_groups(self):
        """Test case for add_groups

        Creates a Group  # noqa: E501
        """
        pass

    def test_get_groups(self):
        """Test case for get_groups

        Retrieves a Specific Group  # noqa: E501
        """
        pass

    def test_get_wiki_group(self):
        """Test case for get_wiki_group

        Retrieves a Group Wiki Page  # noqa: E501
        """
        pass

    def test_list_groups(self):
        """Test case for list_groups

        Retrieves all Groups  # noqa: E501
        """
        pass

    def test_remove_groups(self):
        """Test case for remove_groups

        Deletes a Group  # noqa: E501
        """
        pass

    def test_update_groups(self):
        """Test case for update_groups

        Updates a Group  # noqa: E501
        """
        pass

    def test_update_groups_zones(self):
        """Test case for update_groups_zones

        Updates a Group's Zones  # noqa: E501
        """
        pass

    def test_update_wiki_group(self):
        """Test case for update_wiki_group

        Update a Group Wiki Page  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()