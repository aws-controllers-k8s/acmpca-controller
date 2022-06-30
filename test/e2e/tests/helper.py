# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Helper functions for acmpca tests
"""

from typing import Union, Dict


class ACMPCAValidator:
    def __init__(self, acmpca_client):
        self.acmpca_client = acmpca_client

    def assert_certificate_authority(self, ca_id: str, exists=True):
        res_found = False
        try:
            aws_res = self.acmpca_client.describe_certificate_authority(CertificateAuthorityID=[ca_id])
            res_found = aws_res["Status"] is "PENDING_CERTIFICATE"
        except self.acmpca_client.exceptions.ClientError:
            pass
        assert res_found is exists
