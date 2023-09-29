# Snowflake Proxy

## Proxy

The proxy is an haproxy instance that will route traffic to snowflake

| Variable | Description |
| -------- | ----------- |
| `SNOWFLAKE_ACCOUNTID` | The account id of the snowflake account |
| `SNOWFLAKE_REGION` | The region of the snowflake account |

## Test Client

The test client is a simple golang script that will connect to the proxy and run a query.  It configures the connection to use the local proxy.  It prints each row

| Flag | Env Variable | Description |
| -------- | -------- | ----------- |
| `snowflake_accountid` | `SNOWFLAKE_ACCOUNTID` | The account id of the snowflake account |
| `snowflake_region` | `SNOWFLAKE_REGION` | The region of the snowflake account |
| `snowflake_user`| `SNOWFLAKE_USER` | The user to connect as |
| `snowflake_password` | `SNOWFLAKE_PASSWORD` | The password to connect with |
| `proxy_host` | `127.0.0.1` | The host of the proxy |
| `proxy_port` | `8080` | The port of the proxy |

## Running

To run this we need to build the docker image and then run it.

```bash
docker build -t snowflake-proxy .
```

Now that it is built we will need to run it

```bash
docker run -it -p 8080:8080 -e SNOWFLAKE_ACCOUNTID=12345 -e SNOWFLAKE_REGION=us-east-1 snowflake-proxy
```

Replace the account id and region with your own.  You can also change the port if you want to run it on a different port.

Now that the proxy is running we can run the test client.  Replace the user and password with your own.

```bash
go run main.go -snowflake_accountid 12345 -snowflake_region us-east-1 -snowflake_user test -snowflake_password test -proxy_host 127.0.0.1 -proxy_port 8080
```

This should print something like the following if it connected properly

```text
[INFO] 2023/09/28 21:46:27 main.go:45: ONE
[INFO] 2023/09/28 21:46:27 main.go:46: ---
[INFO] 2023/09/28 21:46:27 main.go:54: 1
```

That is it!  You are now connected to snowflake through the proxy.
