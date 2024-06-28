# CAAAS

Centeral Authentication and Authorization service. CAAAS is a simple and fast tool for identifying your users and authorizing them. It supports global groups and roles resources in order to authorize your users.

By using `Golang` and `Echo`, it is designed to be fast and light. CAAAS uses `postgresQL` as it database, where it stores three main tables named `users`, `user-groups`, and `user-roles`.

## Structure

- For each user, there is database record that has `Username` and `Password`
- For each user-group, there is a set of `Username` and `GroupName`
- For each user-role, there is a set of `Username` and `Role`

## APIs

CAAAS provides two types of service. An `api` and a `gateway`. APIs are used to handle
administrator operations as follow:

- Add a user
- Remove a user
- Update a user password
- Add user to a group
- Remove user from a group
- Remove a group
- Add a role to a user
- Remove a role from a user
- Remove a role
- Fetch user and users data

Gateways are very fast http services that are designed to handle the following methods:

- Login
- Validate
- Groups
- Roles

You can read more about its usages in [DOCS.md](./docs/README.md).
