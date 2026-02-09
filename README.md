## GitHub Action issuing Yandex Cloud IAM Token

The action issues new Yandex Cloud IAM token and puts it in the output.

**Table of Contents**

<!-- toc -->

- [GitHub Action issuing Yandex Cloud IAM Token](#github-action-issuing-yandex-cloud-iam-token)
- [Usage](#usage)
  - [Action Inputs](#action-inputs)
  - [Action Outputs](#action-outputs)
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

To perform this action, Service account and Workload identity federation is required.
For setting up Workload identity federation please refer to the [official YC tutorial](https://yandex.cloud/en/docs/tutorials/security/wlif-github-integration). Lockbox-related steps could be ommited.


Kindly note that your workflow definition must include `id-token` permission in configuration root to use this action.
Otherwise, error `failed to get ID token: missing ACTIONS_ID_TOKEN_REQUEST_URL in environment` will be thrown.

```yaml
name: Release

on:
  push:
    branches:
      - master

permissions:
  id-token: write # This is required for requesting Github OIDC token used for authentocation in YC.

jobs:
  ...
```

For more details refer to the [official Github documentation](https://docs.github.com/en/actions/reference/security/oidc#required-permission)
or to the [blog post](https://nikolaymatrosov.ru/2025-05-04-Authorizing-in-GitHub-Actions-via-Workload-Identities/?utm_source=github&utm_medium=yc-actions&utm_campaign=yc-iam-token-fed) describing the action (in Russian).

## License Summary

This code is made available under the MIT license.