"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.plugin import Plugin
from openapi_client.model.upload_plugin200_response_all_of import UploadPlugin200ResponseAllOf
globals()['Model200Success'] = Model200Success
globals()['Plugin'] = Plugin
globals()['UploadPlugin200ResponseAllOf'] = UploadPlugin200ResponseAllOf
from openapi_client.model.upload_plugin200_response import UploadPlugin200Response


class TestUploadPlugin200Response(unittest.TestCase):
    """UploadPlugin200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testUploadPlugin200Response(self):
        """Test UploadPlugin200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = UploadPlugin200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
