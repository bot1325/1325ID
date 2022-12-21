---
title: "Seafile"
description: "Integrating Seafile with the Authelia OpenID Connect Provider."
lead: ""
date: 2022-06-15T17:51:47+10:00
draft: false
images: []
menu:
  integration:
    parent: "openid-connect"
weight: 620
toc: true
community: true
---

## Tested Versions

* [Authelia]
  * [v4.36.9](https://github.com/authelia/authelia/releases/tag/v4.36.9)
* [Seafile] Server
  * [9.0.9](https://manual.seafile.com/changelog/server-changelog/#909-2022-09-22)

## Before You Begin

### Common Notes

1. You are *__required__* to utilize a unique client id for every client.
2. The client id on this page is merely an example and you can theoretically use any alphanumeric string.
3. You *__should not__* use the client secret in this example, We *__strongly recommend__* reading the
   [Generating Client Secrets] guide instead.

[Generating Client Secrets]: ../specific-information.md#generating-client-secrets

### Assumptions

This example makes the following assumptions:

* __Application Root URL:__ `https://seafile.example.com`
* __Authelia Root URL:__ `https://auth.example.com`
* __Client ID:__ `seafile`
* __Client Secret:__ `seafile_client_secret`

## Configuration

### Application

To configure [Seafile] to utilize Authelia as an [OpenID Connect] Provider:

1. [Seafile] may require some dependencies such as `requests_oauthlib` to be manually installed.
   See the [Seafile] documentation in the [see also](#see-also) section for more information.

2. Edit your [Seafile] `seahub_settings.py` configuration file and add configure the following:

```python
ENABLE_OAUTH = True
OAUTH_ENABLE_INSECURE_TRANSPORT = False
OAUTH_CLIENT_ID = "seafile"
OAUTH_CLIENT_SECRET = "seafile_client_secret"
OAUTH_REDIRECT_URL = 'https://seafile.example.com/oauth/callback/'
OAUTH_PROVIDER_DOMAIN = 'auth.example.com'
OAUTH_AUTHORIZATION_URL = 'https://auth.example.com/api/oidc/authorization'
OAUTH_TOKEN_URL = 'https://auth.example.com/api/oidc/token'
OAUTH_USER_INFO_URL = 'https://auth.example.com/api/oidc/userinfo'
OAUTH_SCOPE = [
    "openid",
    "profile",
    "email",
]
OAUTH_ATTRIBUTE_MAP = {
    "email": (True, "email"),
    "name": (False, "name"),
    "id": (False, "not used"),
}
```

### Authelia

The following YAML configuration is an example __Authelia__
[client configuration](../../../configuration/identity-providers/open-id-connect.md#clients) for use with [Seafile]
which will operate with the above example:

```yaml
- id: seafile
  description: Seafile
  secret: '$plaintext$seafile_client_secret'
  public: false
  authorization_policy: two_factor
  redirect_uris:
    - https://seafile.example.com/oauth/callback/
  scopes:
    - openid
    - profile
    - email
  userinfo_signing_algorithm: none
```

## See Also

* [Seafile OAuth Authentication Documentation](https://manual.seafile.com/deploy/oauth/)

[Authelia]: https://www.authelia.com
[Seafile]: https://www.seafile.com/
[OpenID Connect]: ../../openid-connect/introduction.md
