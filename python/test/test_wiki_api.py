"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.wiki_api import WikiApi  # noqa: E501


class TestWikiApi(unittest.TestCase):
    """WikiApi unit test stubs"""

    def setUp(self):
        self.api = WikiApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_wiki(self):
        """Test case for add_wiki

        Create a Wiki Page  # noqa: E501
        """
        pass

    def test_get_wiki(self):
        """Test case for get_wiki

        Retrieves a specific Wiki page  # noqa: E501
        """
        pass

    def test_get_wiki_app(self):
        """Test case for get_wiki_app

        Retrieves an App Wiki Page  # noqa: E501
        """
        pass

    def test_get_wiki_categories(self):
        """Test case for get_wiki_categories

        Retrieves all Wiki categories associated with the account  # noqa: E501
        """
        pass

    def test_get_wiki_cloud(self):
        """Test case for get_wiki_cloud

        Retrieves a Cloud Wiki Page  # noqa: E501
        """
        pass

    def test_get_wiki_cluster(self):
        """Test case for get_wiki_cluster

        Retrieves a Cluster Wiki Page  # noqa: E501
        """
        pass

    def test_get_wiki_group(self):
        """Test case for get_wiki_group

        Retrieves a Group Wiki Page  # noqa: E501
        """
        pass

    def test_get_wiki_instance(self):
        """Test case for get_wiki_instance

        Retrieves an Instance Wiki Page  # noqa: E501
        """
        pass

    def test_get_wiki_server(self):
        """Test case for get_wiki_server

        Retrieves a Server Wiki Page  # noqa: E501
        """
        pass

    def test_list_wiki(self):
        """Test case for list_wiki

        Retrieves wiki pages associated with the account.  # noqa: E501
        """
        pass

    def test_remove_wiki(self):
        """Test case for remove_wiki

        Deletes a Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki(self):
        """Test case for update_wiki

        Updates a Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki_app(self):
        """Test case for update_wiki_app

        Update an App Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki_cloud(self):
        """Test case for update_wiki_cloud

        Update a Cloud Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki_cluster(self):
        """Test case for update_wiki_cluster

        Update a Cluster Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki_group(self):
        """Test case for update_wiki_group

        Update a Group Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki_instance(self):
        """Test case for update_wiki_instance

        Update an Instance Wiki Page  # noqa: E501
        """
        pass

    def test_update_wiki_server(self):
        """Test case for update_wiki_server

        Update a Server Wiki Page  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
