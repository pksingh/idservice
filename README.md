
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
    C:\Users\Home\Desktop\idservice>idservice.exe
    Server Starting on 8080
    ```

- **CHECK:**
    - Lets hit on browser or from cli using curl on /hello endpoint.
    > `curl http://localhost:8080/hello`
    or
    > `curl -v http://localhost:8080/hello` to log or inspect details
    ```
    C:\Users\Home\Desktop\idservice>curl http://localhost:8080/hello
    hi!
    C:\Users\Home\Desktop\idservice>curl -v http://localhost:8080/hello
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
## APIs
- */idgen*: 
    - using curl: 
        `curl -X GET http://localhost:8080/idgen`

    - output:
        `{"uid": "47393816367333377"}` 

- */parseid*:
    - using curl:
        `curl http://localhost:8080/parseid?uid=47393816367333377`

    - output:
        ```
        {"id": "47393816367333377", "time": 22599132713, "nodeId": 0, "sequence": 1}
        ```

- */idmeta*:
    - using curl:
        `curl -X GET "http://localhost:8080/idmeta"`

    - output:
        ```
        {"start_time": "2022-01-01 00:00:00 +0000 UTC", "node_id": 0, "time_bits": 42, "node_bits": 5, "count_bits": 16}
        ```

## Testing
- *UnitTests*:
    - executing : 
        ```
        go test -timeout 30s -run ^(TestGetHealthHttpt|TestGetHealth|TestGetIdgen|TestGetIdmeta|TestGetIdparsed|TestGetIdparsedError)$ github.com/pksingh/idservice
        ```

    - result : 
        ```
        === RUN   TestGetHealthHttpt
        --- PASS: TestGetHealthHttpt (0.00s)
        === RUN   TestGetHealth
        --- PASS: TestGetHealth (0.00s)
        === RUN   TestGetIdgen
        --- PASS: TestGetIdgen (0.01s)
        === RUN   TestGetIdmeta
        --- PASS: TestGetIdmeta (0.00s)
        === RUN   TestGetIdparsed
        --- PASS: TestGetIdparsed (0.00s)
        === RUN   TestGetIdparsedError
        --- PASS: TestGetIdparsedError (0.00s)
        PASS
        ok      github.com/pksingh/idservice    0.724s

        > Test run finished at 5/26/2022, 6:52:04 PM <
        ```
    - executing : 
        ```
        go test -timeout 30s -run ^(TestGetHealthError|TestGetIdgenError)$ github.com/pksingh/idservice
        ```

    - result : 
        ```
        === RUN   TestGetHealthError
        --- PASS: TestGetHealthError (0.00s)
        === RUN   TestGetIdgenError
        --- PASS: TestGetIdgenError (0.00s)
        PASS
        ok      github.com/pksingh/idservice    0.713s

        > Test run finished at 5/26/2022, 6:57:51 PM <
        ```
    - executing : 
        ```
        go test -timeout 30s -run ^(TestNextIds|TestNextIdPanics|TestSetNode|TestParseId)$ github.com/pksingh/idservice/snowid
        ````

    - result : 
        ```
        === RUN   TestNextIds
        --- PASS: TestNextIds (0.00s)
        === RUN   TestNextIdPanics
        === RUN   TestNextIdPanics/max_time_exceeded
        --- PASS: TestNextIdPanics (0.00s)
            --- PASS: TestNextIdPanics/max_time_exceeded (0.00s)
        === RUN   TestSetNode
        === RUN   TestSetNode/invalid_node_id
        === RUN   TestSetNode/invalid_timestamp_bits
        === RUN   TestSetNode/invalid_node_bits
        === RUN   TestSetNode/invalid_sequence_bits
        === RUN   TestSetNode/max_time_exceeded
        --- PASS: TestSetNode (0.00s)
            --- PASS: TestSetNode/invalid_node_id (0.00s)
            --- PASS: TestSetNode/invalid_timestamp_bits (0.00s)
            --- PASS: TestSetNode/invalid_node_bits (0.00s)
            --- PASS: TestSetNode/invalid_sequence_bits (0.00s)
            --- PASS: TestSetNode/max_time_exceeded (0.00s)
        === RUN   TestParseId
        === RUN   TestParseId/parse_snowid
        --- PASS: TestParseId (0.00s)
            --- PASS: TestParseId/parse_snowid (0.00s)
        PASS
        ok      github.com/pksingh/idservice/snowid     0.630s

        > Test run finished at 5/26/2022, 6:47:41 PM <
        ```

- *Benchmark*:
    - executing : 
        ```
        go test -benchmem -run=^$ -bench ^BenchmarkSnowid$ github.com/pksingh/idservice/snowid
        ```

    - result : 
        ```
        goos: windows
        goarch: amd64
        pkg: github.com/pksingh/idservice/snowid
        cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx
        BenchmarkSnowid
        BenchmarkSnowid-8
         8655067               128.8 ns/op             0 B/op          0 allocs/op
        PASS
        ok      github.com/pksingh/idservice/snowid     1.882s

        > Test run finished at 5/26/2022, 6:44:43 PM <
        ```

## License

MIT
