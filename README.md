# faiss-server

`faiss-server` is gRPC/HTTP Server for ANN using [facebookresearch/faiss](https://github.com/facebookresearch/faiss).

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
