# Faiss-server

## Development

```bash
# update dependencies
bazel run //:gazelle
bazel run //:gazelle-update-repos
bazel run //:gazelle
```

```bash
# run faiss-server
bazel run //:faiss-server
```

```bash
# run test
bazel test //...
```
