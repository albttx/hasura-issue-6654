# [Hasura request_id error](https://github.com/hasura/graphql-engine/issues/6654)

This error environment to re-create the issue.

## Expanation

This issue is that on an action called, the header `x-request-id` isn't passed to the API.

#### NOTE:

- The `forward_client_headers: true` is set, see [actions.yml](./hasura/metadata/actions.yml)

## Start the environment

```bash
$ docker-compose up -d
Creating hasura-issue-6654_api_1 ... done
Creating hasura-issue-6654_psql_1 ... done
Creating hasura-issue-6654_hasura_1 ... done

$ cd hasura/
hasura/ $ hasura migrate apply
INFO migrations applied

hasura/ $ hasura metadata apply
INFO metadata applied

hasura/ $ hasura seeds apply
INFO Seeds planted
```

## Reproduce the error

First in a terminal, watch the logs using the commands

```bash
$ docker-compose logs -f api hasura
...
```

Then you can either open the console on [http://localhost:8080/](http://localhost:8080/) and run the mutation

```graphql
mutation {
  login(username: "foo", password: "pass") {
    access_token
  }
}
```

or you can use this `curl` command

```bash
$ curl -v 'http://localhost:8080/v1/graphql' --compressed \
    -H 'content-type: application/json' \
    -H 'x-hasura-admin-secret: admin' \
    --data-raw '{"query":"mutation {\n  login(username: \"foo\", password: \"pass\") {\n    access_token\n  }\n}","variables":null}'
```

The logs in docker-compose are

```bash
$ docker-compose logs -f api hasura
...
api_1     | POST /login HTTP/1.1
api_1     | Host: api:3000
api_1     | Accept-Encoding: gzip
api_1     | Content-Length: 118
api_1     | Content-Type: application/json
api_1     | User-Agent: hasura-graphql-engine/v1.3.3
api_1     | X-Forwarded-Host: localhost:8080
api_1     | X-Forwarded-User-Agent: curl/7.76.0
api_1     | X-Hasura-Spanid: 1283379409906698631
api_1     | X-Hasura-Traceid: 15574078171665254689
api_1     |
api_1     |
hasura_1  | {"type":"http-log","timestamp":"2021-04-19T14:31:30.265+0000","level":"info","detail":{"operation":{"query_execution_time":1.691667e-3,"user_vars":{"x-hasura-role":"admin"},"request_id":"1d64dece-0f59-4ccc-a5a3-125a31f002ed","response_size":57,"request_read_time":4.961e-6},"http_info":{"status":200,"http_version":"HTTP/1.1","url":"/v1/graphql","ip":"172.22.0.1","method":"POST","content_encoding":"gzip"}}}
```

As you can see, the header `X-Request-Id` is missing.

Now let's send a `X-request-Id` from the request

```bash
curl -v 'http://localhost:8080/v1/graphql' --compressed \                                                                                            ✔
    -H 'content-type: application/json' \
    -H 'x-request-id: toto' \
    -H 'x-hasura-admin-secret: admin' \
    --data-raw '{"query":"mutation {\n  login(username: \"foo\", password: \"pass\") {\n    access_token\n  }\n}","variables":null}'

# Logs
$ docker-compose logs -f api hasura
api_1     | POST /login HTTP/1.1
api_1     | Host: api:3000
api_1     | Accept-Encoding: gzip
api_1     | Content-Length: 118
api_1     | Content-Type: application/json
api_1     | User-Agent: hasura-graphql-engine/v1.3.3
api_1     | X-Forwarded-Host: localhost:8080
api_1     | X-Forwarded-User-Agent: curl/7.76.0
api_1     | X-Hasura-Spanid: 787684514260826683
api_1     | X-Hasura-Traceid: 12899392442796024351
api_1     | X-Request-Id: toto         # <- The x-request-id header is present !
api_1     |
api_1     |
hasura_1  | {"type":"http-log","timestamp":"2021-04-19T14:33:09.479+0000","level":"info","detail":{"operation":{"query_execution_time":2.069069e-3,"user_vars":{"x-hasura-role":"admin"},"request_id":"toto","response_size":57,"request_read_time":4.466e-6},"http_info":{"status":200,"http_version":"HTTP/1.1","url":"/v1/graphql","ip":"172.22.0.1","method":"POST","content_encoding":"gzip"}}}
# request_id is `toto` in the json logs.
```
