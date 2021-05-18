FROM ubuntu:bionic

ARG USER_ID
ARG GROUP_ID

RUN apt-get -qq update && apt-get -qq upgrade -y && \
    DEBIAN_FRONTEND=noninteractive apt-get -qq install -y sudo \
                                                          git \
                                                          build-essential \
                                                          apt-utils \
                                                          software-properties-common \
                                                          apt-transport-https \
                                                          curl \
                                                          gnupg > /dev/null && rm -rf /var/lib/apt/lists/*
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get -qq update > /dev/null && \
    DEBIAN_FRONTEND=noninteractive apt-get -qq install -y golang-1.14 > /dev/null && \
    rm -rf /var/lib/apt/lists/*
RUN curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg && \
    mv bazel.gpg /etc/apt/trusted.gpg.d/ && \
    echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list && \
    apt-get -qq update > /dev/null && DEBIAN_FRONTEND=noninteractive apt-get -qq install -y bazel-3.6.0 > /dev/null && \
    rm -rf /var/lib/apt/lists/* && \
    ln -s /usr/bin/bazel-3.6.0 /usr/bin/bazel

RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
RUN groupadd -g $GROUP_ID scion && \
    groupadd docker && \
    useradd -u $USER_ID -g $GROUP_ID -ms /bin/bash scion && \
    usermod -a -G sudo scion && \
    usermod -a -G docker scion

USER scion
RUN mkdir -p /home/scion/scion && git clone -b tcinbis/flowtele https://github.com/tcinbis/scion.git /home/scion/scion
COPY .bazelrc-docker /home/scion/scion/.bazelrc
WORKDIR /home/scion/scion
RUN sudo apt-get update -qq > /dev/null && ./env/deps > /dev/null
