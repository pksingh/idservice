
## Description

Unique ID generator Service.
This can be used in various ways; which provides unique id for each request.

## Start Service and Output
- **CODE:**
    - We need to define; go.mod file if NOT done yet. We do by following
    > `go mod init idservice`
    - Lets download any depedency.
    > `go mod tidy`

- **BUILD:**
    - Lets build our project by following command
    > `go build .`
    If go build success; this will generate a binary based on your OS/Platform

- **RUN:**
    - Lets run/execute the binary to bring the service up.
    > `idservice.exe` or `idservice` if you use *NIX os
    ```
    C:\Users\Home\Desktop\idservice>idserver.exe
    Server Starting on 8080
    ```

- **CHECK:**
    - Lets hit on browser or from cli using curl on /hello endpoint.
    > `curl http://localhost:8080/hello`
    > or
    > `curl -v http://localhost:8080/hello` to log or inspect details
    ```
    C:\Users\Home\Desktop\idservice>curl http://localhost:8080/hello
hi!
C:\Users\Home\Desktop\idservice>curl http://localhost:8080/hello -v
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hello HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.83.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Wed, 16 Mar 2022 16:32:51 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
<
hi!* Connection #0 to host localhost left intact

    ```

## License

MIT
