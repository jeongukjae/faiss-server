# BUILDER ==========
FROM golang:1.17-buster as BUILDER

# Install Bazel
RUN apt update && \
    apt install -y apt-transport-https curl gnupg patch && \
    curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg && \
    mv bazel.gpg /etc/apt/trusted.gpg.d/ && \
    echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" \
        | tee /etc/apt/sources.list.d/bazel.list && \
    apt update && apt install -y bazel && \
    bazel version

WORKDIR /app
COPY . .

# Build faiss-server
RUN apt install -y libopenblas-dev libomp-7-dev libstdc++6 && \
    bazel build //:faiss-server --nokeep_state_after_build

# Prepare libraries for distroless image
RUN cd /tmp && \
    apt-get download \
        $(apt-cache depends \
            --recurse --no-recommends --no-suggests --no-conflicts \
            --no-breaks --no-replaces --no-enhances \
            libomp5-7 libopenblas-base libstdc++6 \
                | grep "^\w" | sort -u) && \
    for deb in *.deb; do dpkg --extract $deb /dpkg || exit 10; done && \
    cd /dpkg && mkdir /dpkg-so-files && \
    for so in $(find . -name "*.so*" -not -name "*.py" -not -name "*.conf.d"); do cp $so /dpkg-so-files; done && \
    ls -alh /dpkg-so-files

# OUTPUT ==========
FROM gcr.io/distroless/base-debian10

COPY --from=BUILDER /dpkg-so-files /usr/lib
COPY --from=BUILDER /app/bazel-bin/faiss-server_/faiss-server /faiss-server

EXPOSE 8000 8001
WORKDIR /
ENTRYPOINT ["/faiss-server"]
