# Faiss-server

## Development

```bash
# update dependencies and lint
bazel run //:gazelle
bazel run //:gazelle-update-repos
bazel run //:gazelle
bazel run //:buildifier
```

```bash
# run faiss-server
bazel run //:faiss-server
```

```bash
# run test
bazel test //...
```
