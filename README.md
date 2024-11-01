# Go Socks5 Proxy
Socks5 Proxy Server in Golang

# Examples

- Run docker container using default container port 1080 and expose it to world using host port 1080, with auth creds

    ```docker run -d --name socks5 -p 1080:1080 -e PROXY_USER=<PROXY_USER> -e PROXY_PASSWORD=<PROXY_PASSWORD>  aiql/go-socks5-proxy```

    - Leave `PROXY_USER` and `PROXY_PASSWORD` empty for skip authentication options while running socks5 server, see example below

- Run docker container using specifit container port and expose it to host port 1090, without auth creds

    ```docker run -d --name socks5 -p 1090:9090 -e PROXY_PORT=9090 aiql/go-socks5-proxy```

# List of supported config parameters

|ENV variable|Type|Default|Description|
|------------|----|-------|-----------|
|PROXY_USER|String|EMPTY|Set proxy user (also required existed PROXY_PASS)|
|PROXY_PASSWORD|String|EMPTY|Set proxy password for auth, used with PROXY_USER|
|PROXY_PORT|String|1080|Set listen port for application inside docker container|
|ALLOWED_DEST_FQDN|String|EMPTY|Allowed destination address regular expression pattern. Default allows all.|
|ALLOWED_IPS|String|Empty|Set allowed IP's that can connect to proxy, separator `,`|


# Build your own image:
`docker-compose -f docker-compose.build.yml up -d`\
Just don't forget to set parameters in the `.env` file.

# Test running service

Assuming that you are using container on 1080 host docker port

## Without authentication

```curl --socks5 <docker host ip>:1080  https://ifcfg.co``` - result must show docker host ip (for bridged network)

or

```docker run --rm curlimages/curl:7.65.3 -s --socks5 <docker host ip>:1080 https://ifcfg.co```

## With authentication

```curl --socks5 <docker host ip>:1080 -U <PROXY_USER>:<PROXY_PASSWORD> http://ifcfg.co```

or

```docker run --rm curlimages/curl:7.65.3 -s --socks5 <PROXY_USER>:<PROXY_PASSWORD>@<docker host ip>:1080 http://ifcfg.co```

# Reference
This repository was originally cloned from [socks5-server](https://github.com/serjs/socks5-server) due to its long period of inactivity. It is a small yet elegant repo, and we aim to use it as a pilot project to transform it into a comprehensive library with thorough documentation and automation. We will also strive to optimize and enhance it using AI.
