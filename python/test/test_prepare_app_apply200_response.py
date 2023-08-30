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
from openapi_client.model.app_prepare_apply import AppPrepareApply
from openapi_client.model.app_prepare_apply_data import AppPrepareApplyData
from openapi_client.model.model200_success import Model200Success
globals()['AppPrepareApply'] = AppPrepareApply
globals()['AppPrepareApplyData'] = AppPrepareApplyData
globals()['Model200Success'] = Model200Success
from openapi_client.model.prepare_app_apply200_response import PrepareAppApply200Response


class TestPrepareAppApply200Response(unittest.TestCase):
    """PrepareAppApply200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testPrepareAppApply200Response(self):
        """Test PrepareAppApply200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = PrepareAppApply200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
