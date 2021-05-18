#!/usr/bin/env zsh

mkdir docker-build-results
docker build -t scion-build --build-arg USER_ID="$(id -u)" --build-arg GROUP_ID="$(id -g)" .
./scion.sh bazel_remote
docker run -it --rm --network bazel_remote_default -v "$(pwd)"/docker-build-results:/build-res scion-build /bin/bash -c "make flowtele && cp -R bazel-bin/* /build-res && exit"
