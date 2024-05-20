# GoApp

An application for Go knowledge assessment.

## Description

This is a web application that utilises websockets. A client connects on `localhost:8080` and has three options: a) `open` a websocket connection that reads values from a counter b) `close` the websocket and c) reset the counter to zero. The counter is feeded from a random string generator. On WS session termination, statistics for terminated session are printed.

The application is compiled by running `make` in the root folder and the final binaries are found in the `bin/` folder.

The application has some problems described below that need to be addressed plus some new features that need implementation.

## Problems

### #1

The server prints statistics for each WS sessioned closed but it seems to only count one message while there are more send to each WS session, e.g.

```
2024/03/29 18:28:38 stats.go:11: session a938e316-8536-46e6-8633-bd309fbcf579 has received 1 messages
```

### Solution 1

A change needed in the for loop in order to access sessionstats directly

### #2

A more then normal memory usage is observed after many WS sessions which needs investigation.

### #3

A cross-site request forgery is reported by a security audit which needs fixing.

### Solution 3

The following solution has been implemented, in order for every new websocket connection that is tring to be established to be proteced for csrf attacks:

A csrf package has been created. This package contains:
- middleware
    - **CSRFCheckMiddleware**, this middleware is used for all the routes that needs to be guarded from csrf attack
    - **SetupCSRFMiddleware**, this setup the gorilla mux csrf middleware
- session 
    - **SetSession**, setups a csrf token for the current session
    - **GetSession**, get the csrf token for the current session

In the **Route struct** a Middleware attribute have been added. This is used by the routes that needed an extra middleware.

In the **handlerHome** I have set a csrf token using the SetSession func. Then this token passed to the template and I set it like a query param when tring to open a new connection. 

In the **websocket** route I have set the middleware attribute to **CSRFCheckMiddleware**. This way a new websocket connection can established only with the csrf token as a query param.

## New features

### A

Modify the random string generator to generate only hex values and verify its accuracy and resource usage by creating a test and a benchmark run.

### B

Extent the API to also return the Hex value in WS connection. I.e. a browser that open a connection to `localhost:8080` should see the HEX values.

E.g.

```
OPEN
RESPONSE: {"iteration":1,"value":"822876EF10"}
RESPONSE: {"iteration":2,"value":"215100491D"}
RESPONSE: {"iteration":3,"value":"05DCC3B6AB"}
CLOSE
```

### C

Create a command line client as a separate application that opens a requested number of sessions simultaneously.

This should be a separate executable generated with the `make` command in the `bin/` folder that will accept an argument with the number of parallel connections that will open on the server. The server part must be modified in a way to support multiple parallel connections and should still print valid statistics for each connection.

E.g.

```
$ ./bin/client -n 3
[conn #0] iteration: 1, value: 66D53ED788
[conn #1] iteration: 1, value: 66D53ED788
[conn #2] iteration: 1, value: 66D53ED788
...
```
