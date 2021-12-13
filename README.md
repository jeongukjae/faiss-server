# faiss-server

`faiss-server` is a ANN server using [facebookresearch/faiss](https://github.com/facebookresearch/faiss).

* Supports gRPC/HTTP
* Prometheus handler for monitoring (`/metrics`)

## Usage

* Tag List: <https://ghcr.io/jeongukjae/faiss-server>
* API Specification: [service proto file (protos/faiss_server/service.proto)](protos/faiss_server/service.proto)

```bash
$ docker pull ghcr.io/jeongukjae/faiss-server
$ docker run --rm -it ghcr.io/jeongukjae/faiss-server --help
Usage of /faiss-server:
  -alsologtostderr
        log to standard error as well as files
  -faiss_index string
        path of prebuilt faiss index
  -grpc_server_endpoint string
        gRPC server endpoint (default "0.0.0.0:8000")
  -http_server_endpoint string
        http server endpoint (default "0.0.0.0:8001")
  -log_backtrace_at value
        when logging hits line file:N, emit a stack trace
  -log_dir string
        If non-empty, write log files in this directory
  -logtostderr
        log to standard error instead of files
  -stderrthreshold value
        logs at or above this threshold go to stderr
  -v value
        log level for V logs
  -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging
  -with_reload_method
        enable reload method
$ docker run --rm -it \
    -v `pwd`/testdata/random-index.faiss:/random-index.faiss:ro \
    ghcr.io/jeongukjae/faiss-server \
    -faiss_index /random-index.faiss -logtostderr
I1213 02:27:32.106571       1 server.go:108] Loading faiss index from /random-index.faiss
I1213 02:27:32.132563       1 server.go:70] Serve grpc server at 0.0.0.0:8000
I1213 02:27:32.132648       1 server.go:83] Register grpc gateway server at 0.0.0.0:8000
I1213 02:27:32.133806       1 server.go:98] Serve http server at 0.0.0.0:8001
```

## Build faiss-server

### Docker image

I recommend you to build this repo with docker.

```bash
docker build -t faiss-server .
```

### Binary file

But you can also build from source.

```bash
bazel build //:faiss-server
```

## Dev notes

### Resolve dependencies and run buildifier

```bash
# update dependencies and lint
bazel run //:gazelle
bazel run //:gazelle-update-repos
bazel run //:gazelle
bazel run //:buildifier
```

### TODO

- [ ] Load from GCS/S3
