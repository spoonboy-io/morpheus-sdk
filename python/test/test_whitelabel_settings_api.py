"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.whitelabel_settings_api import WhitelabelSettingsApi  # noqa: E501


class TestWhitelabelSettingsApi(unittest.TestCase):
    """WhitelabelSettingsApi unit test stubs"""

    def setUp(self):
        self.api = WhitelabelSettingsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_whitelabel_image(self):
        """Test case for get_whitelabel_image

        Download Image  # noqa: E501
        """
        pass

    def test_list_whitelabel_settings(self):
        """Test case for list_whitelabel_settings

        Get Whitelabel Settings  # noqa: E501
        """
        pass

    def test_remove_whitelabel_image(self):
        """Test case for remove_whitelabel_image

        Reset Image  # noqa: E501
        """
        pass

    def test_update_whitelabel_images(self):
        """Test case for update_whitelabel_images

        Update Images  # noqa: E501
        """
        pass

    def test_update_whitelabel_settings(self):
        """Test case for update_whitelabel_settings

        Update Whitelabel Settings  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
