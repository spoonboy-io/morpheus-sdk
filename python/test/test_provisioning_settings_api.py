"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.provisioning_settings_api import ProvisioningSettingsApi  # noqa: E501


class TestProvisioningSettingsApi(unittest.TestCase):
    """ProvisioningSettingsApi unit test stubs"""

    def setUp(self):
        self.api = ProvisioningSettingsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_list_provisioning_settings(self):
        """Test case for list_provisioning_settings

        Get Provisioning Settings  # noqa: E501
        """
        pass

    def test_update_provisioning_settings(self):
        """Test case for update_provisioning_settings

        Update Provisioning Settings  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
