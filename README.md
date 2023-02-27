<h1 align="center">
  AuthX
</h1>

<br />

Add sophisticated authentication and authorization to your applications using AuthX. Free version of Auth0 for Iranian users.
Providing a quick IDaaS.

## Quick road

An index page for creating an application.
Each application should be for a user and provides
one API key and Address. Each application can store a limited number
of users. Each application will be deleted after 10 days.

## Endpoints

There are two endpoints.

Add a new user to system:

```shell
curl -X PUT www.authx.com/api/user
```

```json
{
  "creds": {
    "name": "",
    "fellan": ""
  }
}
```

Check a user:

```shell
curl -X GET www.authx.com/api/user?token=
```