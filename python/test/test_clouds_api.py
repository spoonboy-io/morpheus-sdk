"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.clouds_api import CloudsApi  # noqa: E501


class TestCloudsApi(unittest.TestCase):
    """CloudsApi unit test stubs"""

    def setUp(self):
        self.api = CloudsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_cloud_resource_pool(self):
        """Test case for add_cloud_resource_pool

        Creates a Specified Resource Pool for Specified Cloud  # noqa: E501
        """
        pass

    def test_add_clouds(self):
        """Test case for add_clouds

        Creates a Cloud  # noqa: E501
        """
        pass

    def test_get_cloud_datastores(self):
        """Test case for get_cloud_datastores

        Retrieves a Datastore for Specified Cloud  # noqa: E501
        """
        pass

    def test_get_cloud_folders(self):
        """Test case for get_cloud_folders

        Retrieves a Resource Folder for Specified Cloud  # noqa: E501
        """
        pass

    def test_get_cloud_resource_pools(self):
        """Test case for get_cloud_resource_pools

        Retrieves a Resource Pool for Specified Cloud  # noqa: E501
        """
        pass

    def test_get_cloud_types(self):
        """Test case for get_cloud_types

        Retrieves a Specific Cloud Type  # noqa: E501
        """
        pass

    def test_get_clouds(self):
        """Test case for get_clouds

        Retrieves a Specific Cloud  # noqa: E501
        """
        pass

    def test_get_wiki_cloud(self):
        """Test case for get_wiki_cloud

        Retrieves a Cloud Wiki Page  # noqa: E501
        """
        pass

    def test_list_cloud_datastores(self):
        """Test case for list_cloud_datastores

        Retrieves all Datastores for Specified Cloud  # noqa: E501
        """
        pass

    def test_list_cloud_folders(self):
        """Test case for list_cloud_folders

        Retrieves all resource folders for Specified Cloud  # noqa: E501
        """
        pass

    def test_list_cloud_resource_pools(self):
        """Test case for list_cloud_resource_pools

        Retrieves all Resource Pools for Specified Cloud  # noqa: E501
        """
        pass

    def test_list_cloud_security_groups(self):
        """Test case for list_cloud_security_groups

        Retrieves all Security Groups for a Cloud  # noqa: E501
        """
        pass

    def test_list_cloud_types(self):
        """Test case for list_cloud_types

        Retrieves all Cloud Types  # noqa: E501
        """
        pass

    def test_list_clouds(self):
        """Test case for list_clouds

        Retrieves all Clouds  # noqa: E501
        """
        pass

    def test_refresh_clouds(self):
        """Test case for refresh_clouds

        Refreshes a Cloud  # noqa: E501
        """
        pass

    def test_remove_cloud_resource_pools(self):
        """Test case for remove_cloud_resource_pools

        Deletes a Resource Pool for Specified Cloud  # noqa: E501
        """
        pass

    def test_remove_clouds(self):
        """Test case for remove_clouds

        Deletes a Cloud  # noqa: E501
        """
        pass

    def test_update_cloud_datastores(self):
        """Test case for update_cloud_datastores

        Updates a Specified Datastore for Specified Cloud  # noqa: E501
        """
        pass

    def test_update_cloud_folders(self):
        """Test case for update_cloud_folders

        Updates a Resource Folder for Specified Cloud  # noqa: E501
        """
        pass

    def test_update_cloud_logo(self):
        """Test case for update_cloud_logo

        Update Logo For Cloud  # noqa: E501
        """
        pass

    def test_update_cloud_resource_pool(self):
        """Test case for update_cloud_resource_pool

        Updates a Specified Resource Pool for Specified Cloud  # noqa: E501
        """
        pass

    def test_update_cloud_security_groups(self):
        """Test case for update_cloud_security_groups

        Sets Security Groups for a Cloud  # noqa: E501
        """
        pass

    def test_update_clouds(self):
        """Test case for update_clouds

        Updates a Cloud  # noqa: E501
        """
        pass

    def test_update_wiki_cloud(self):
        """Test case for update_wiki_cloud

        Update a Cloud Wiki Page  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
