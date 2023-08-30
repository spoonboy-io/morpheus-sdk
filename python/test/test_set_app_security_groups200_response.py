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
from openapi_client.model.app_security_groups import AppSecurityGroups
from openapi_client.model.get_app_security_groups200_response_all_of import GetAppSecurityGroups200ResponseAllOf
from openapi_client.model.model200_success import Model200Success
globals()['AppSecurityGroups'] = AppSecurityGroups
globals()['GetAppSecurityGroups200ResponseAllOf'] = GetAppSecurityGroups200ResponseAllOf
globals()['Model200Success'] = Model200Success
from openapi_client.model.set_app_security_groups200_response import SetAppSecurityGroups200Response


class TestSetAppSecurityGroups200Response(unittest.TestCase):
    """SetAppSecurityGroups200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSetAppSecurityGroups200Response(self):
        """Test SetAppSecurityGroups200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SetAppSecurityGroups200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()