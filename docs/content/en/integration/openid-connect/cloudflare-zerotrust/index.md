---
title: "Cloudflare Zero Trust"
description: "Integrating Cloudflare Zero Trust with the Authelia OpenID Connect Provider."
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
  * [v4.35.6](https://github.com/authelia/authelia/releases/tag/v4.35.6)

## Before You Begin

### Common Notes

1. You are *__required__* to utilize a unique client id for every client.
2. The client id on this page is merely an example and you can theoretically use any alphanumeric string.
3. You *__should not__* use the client secret in this example, We *__strongly recommend__* reading the
   [Generating Client Secrets] guide instead.

[Generating Client Secrets]: ../specific-information.md#generating-client-secrets

### Assumptions

This example makes the following assumptions:

* __Cloudflare Team Name:__ `example-team`
* __Authelia Root URL:__ `https://auth.example.com`
* __Client ID:__ `cloudflare`
* __Client Secret:__ `cloudflare_client_secret`

*__Important Note:__ [Cloudflare Zero Trust] does not properly URL encode the secret per [RFC6749 Appendix B] at the
time this article was last modified (noted at the bottom). This means you'll either have to use only alphanumeric
characters for the secret or URL encode the secret yourself.*

[RFC6749 Appendix B]: https://www.rfc-editor.org/rfc/rfc6749#appendix-B

## Configuration

### Application

*__Important Note:__ It is a requirement that the Authelia URL's can be requested by Cloudflare's servers. This usually
means that the URL's are accessible to foreign clients on the internet. There may be a way to configure this without
accessibility to foreign clients on the internet on Cloudflare's end but this is beyond the scope of this document.*

To configure [Cloudflare Zero Trust] to utilize Authelia as an [OpenID Connect] Provider:

1. Visit the [Cloudflare Zero Trust Dashboard](https://dash.teams.cloudflare.com)
2. Visit `Settings`
3. Visit `Authentication`
4. Under `Login nethods` select `Add new`
5. Select `OpenID Connect`
6. Set the following values:
   1. Name: `Authelia`
   2. App ID: `cloudflare`
   3. Client Secret: `cloudflare_client_secret`
   4. Auth URL: `https://auth.example.com/api/oidc/authorization`
   5. Token URL: `https://auth.example.com/api/oidc/token`
   6. Certificate URL: `https://auth.example.com/jwks.json`
   7. Enable `Proof Key for Code Exchange (PKCE)`
   8. Add the following OIDC Claims: `preferred_username`, `mail`
7. Click Save

### Authelia

The following YAML configuration is an example __Authelia__
[client configuration](../../../configuration/identity-providers/open-id-connect.md#clients) for use with [Cloudflare]
which will operate with the above example:

```yaml
- id: cloudflare
  description: Cloudflare ZeroTrust
  secret: '$plaintext$cloudflare_client_secret'
  public: false
  authorization_policy: two_factor
  redirect_uris:
    - https://example-team.cloudflareaccess.com/cdn-cgi/access/callback
  scopes:
    - openid
    - profile
    - email
  userinfo_signing_algorithm: none
```

## See Also

* [Cloudflare Zero Trust Generic OIDC Documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/generic-oidc/)

[Authelia]: https://www.authelia.com
[Cloudflare]: https://www.cloudflare.com/
[Cloudflare Zero Trust]: https://www.cloudflare.com/products/zero-trust/
[OpenID Connect]: ../../openid-connect/introduction.md
