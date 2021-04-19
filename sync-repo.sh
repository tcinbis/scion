#!/usr/bin/env zsh
USER=$(id -un)
REMOTE_USER=$USER
REPO_PATH=$(dirname "$(readlink -f "$0")")
HOSTS="flowtele-ohio flowtele-singapore" # flowtele-ethz

if [ ! "$(command -v prsync)" ]; then
    echo "Please install parallel rsync"
    echo "For example with Arch: yay -S python-pssh"
    exit 1
fi

echo "Creating $REPO_PATH on hosts before rsync..."
pssh -H "$HOSTS" "mkdir -p $REPO_PATH"

echo "Syncing $REPO_PATH to $HOSTS"
prsync -av -H "$HOSTS" \
    -x "-L" \
    -x "--exclude '.git'" \
    -x "--exclude .bazel-cache" \
    -x "--include bazel-bin/go/flowtele/***" \
    -x "--include bazel-bin/go/examples/pingpong/***" \
    -x "--include */" \
    -x "--exclude bazel-bin/***" \
    -x "--exclude bazel-out/***" \
    -x "--exclude bazel-scion/***" \
    -x "--prune-empty-dirs" \
    "$REPO_PATH/" "$REPO_PATH/" \
