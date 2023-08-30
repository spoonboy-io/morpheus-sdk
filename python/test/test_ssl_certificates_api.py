"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.ssl_certificates_api import SSLCertificatesApi  # noqa: E501


class TestSSLCertificatesApi(unittest.TestCase):
    """SSLCertificatesApi unit test stubs"""

    def setUp(self):
        self.api = SSLCertificatesApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_add_certificate(self):
        """Test case for add_certificate

        Create a Certificate  # noqa: E501
        """
        pass

    def test_delete_certificate(self):
        """Test case for delete_certificate

        Delete a Certificate  # noqa: E501
        """
        pass

    def test_get_certificate(self):
        """Test case for get_certificate

        Get a Specific Certificate  # noqa: E501
        """
        pass

    def test_list_certificates(self):
        """Test case for list_certificates

        Get All SSL Certificates  # noqa: E501
        """
        pass

    def test_update_certificate(self):
        """Test case for update_certificate

        Update a Certificate  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()