"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.deploys_api import DeploysApi  # noqa: E501


class TestDeploysApi(unittest.TestCase):
    """DeploysApi unit test stubs"""

    def setUp(self):
        self.api = DeploysApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_instance_deploy(self):
        """Test case for add_instance_deploy

        Deploy to an Instance  # noqa: E501
        """
        pass

    def test_deletedeploy(self):
        """Test case for deletedeploy

        Delete a Deploy  # noqa: E501
        """
        pass

    def test_get_instance_deploys(self):
        """Test case for get_instance_deploys

        Get all Deploys for an Instance  # noqa: E501
        """
        pass

    def test_list_deploys(self):
        """Test case for list_deploys

        Get all Deploys  # noqa: E501
        """
        pass

    def test_run_deploy(self):
        """Test case for run_deploy

        Run a Deploy  # noqa: E501
        """
        pass

    def test_update_deploy(self):
        """Test case for update_deploy

        Update a Deploy  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
