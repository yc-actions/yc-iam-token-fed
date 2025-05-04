## GitHub Action issuing Yandex Cloud IAM Token

The action issues new Yandex Cloud IAM token and puts it in the output.

**Table of Contents**

<!-- toc -->

- [Usage](#usage)
- [Prerequisites](#prerequisites)
- [License Summary](#license-summary)

<!-- tocstop -->

## Usage

This action can be used in your workflow as follows:

```yaml
- name: Get Yandex Cloud IAM token
  id: get-iam-token
  uses: docker://ghcr.io/yc-actions/yc-iam-token-fed:1.0.0
  with:
    yc-sa-id: aje***
```

GitHub JWT token will be used to authenticate the action. It will be exchanged for Yandex Cloud IAM token using the
service account ID provided in the `yc-sa-id` input.

### Action Inputs

| Name       | Description        |
|------------|--------------------|
| `yc-sa-id` | Service Account ID |

### Action Outputs

| Name    | Description         |
|---------|---------------------|
| `token` | Generated IAM token |

## Prerequisites

To perform this action, service account is required.

## License Summary

This code is made available under the MIT license.