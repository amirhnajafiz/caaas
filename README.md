# AuthX

![](https://img.shields.io/github/languages/top/amirhnajafiz/authx?style=for-the-badge)
![](https://img.shields.io/github/go-mod/go-version/amirhnajafiz/authx?color=003377&label=Go%20Version&style=for-the-badge)
![](https://img.shields.io/github/languages/code-size/amirhnajafiz/authx?style=for-the-badge)
![](https://img.shields.io/github/v/release/amirhnajafiz/authx?color=990011&style=for-the-badge)

Add sophisticated authentication and authorization to your applications using AuthX.
Free version of Auth0 for Iranian users. Providing a quick IDaaS for system clients.

AuthX is a IDaaS (Identity as a Service) which can help you to
authenticate your users by setting any credentials that you desire.
For using this service, register with your student number and password
to get an
<span style="background: #b9b9b9; color: #ff4747; padding: 0 5px;">API KEY</span>
in order to use the service APIs.
With this API you can authenticate your clients for free.

## Config

Create a ```config.yml``` file in order to set the configs of ```AuthX```
service.

```yaml
auth:
  private_key: ""
  expire_time: 100 # seconds
logger:
  level: debug
http:
  port: 8000
storage:
  host: ""
  port: 3306
  user: "root"
  pass: ""
  database: "authx"
```

## Docker

Build and start ```AuthX``` using ```Docker```:

```shell
docker build . -f build/package/Dockerfile -t authx:v0.1.0
docker run -d \
  -v type=bind,source=$(pwd)/config.yml,dest=/config.yml \
  authx:v0.1.0
```
