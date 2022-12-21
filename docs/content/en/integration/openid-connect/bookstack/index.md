---
title: "BookStack"
description: "Integrating BookStack with the Authelia OpenID Connect Provider."
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
  * [v4.35.5](https://github.com/authelia/authelia/releases/tag/v4.35.5)
* [BookStack]
  * 20.10

## Before You Begin

### Common Notes

1. You are *__required__* to utilize a unique client id for every client.
2. The client id on this page is merely an example and you can theoretically use any alphanumeric string.
3. You *__should not__* use the client secret in this example, We *__strongly recommend__* reading the
   [Generating Client Secrets] guide instead.

[Generating Client Secrets]: ../specific-information.md#generating-client-secrets

### Assumptions

This example makes the following assumptions:

* __Application Root URL:__ `https://bookstack.example.com`
* __Authelia Root URL:__ `https://auth.example.com`
* __Client ID:__ `bookstack`
* __Client Secret:__ `bookstack_client_secret`

*__Important Note:__ [BookStack] does not properly URL encode the secret per [RFC6749 Appendix B] at the time this
article was last modified (noted at the bottom). This means you'll either have to use only alphanumeric characters for
the secret or URL encode the secret yourself.*

[RFC6749 Appendix B]: https://www.rfc-editor.org/rfc/rfc6749#appendix-B

## Configuration

### Application

To configure [BookStack] to utilize Authelia as an [OpenID Connect] Provider:

1. Edit your .env file
2. Set the following values:
   1. AUTH_METHOD: `oidc`
   2. OIDC_NAME: `Authelia`
   3. OIDC_DISPLAY_NAME_CLAIMS: `name`
   4. OIDC_CLIENT_ID: `bookstack`
   5. OIDC_CLIENT_SECRET: `bookstack_client_secret`
   6. OIDC_ISSUER: `https://auth.example.com`
   7. OIDC_ISSUER_DISCOVER: `true`

### Authelia

The following YAML configuration is an example __Authelia__
[client configuration](../../../configuration/identity-providers/open-id-connect.md#clients) for use with [BookStack]
which will operate with the above example:

```yaml
- id: bookstack
  description: BookStack
  secret: '$plaintext$bookstack_client_secret'
  public: false
  authorization_policy: two_factor
  redirect_uris:
    - https://bookstack.example.com/oidc/callback
  scopes:
    - openid
    - profile
    - email
  userinfo_signing_algorithm: none
```

## See Also

* [BookStack OpenID Connect Documentation](https://www.bookstackapp.com/docs/admin/oidc-auth/)

[Authelia]: https://www.authelia.com
[BookStack]: https://www.bookstackapp.com/
[OpenID Connect]: ../../openid-connect/introduction.md
