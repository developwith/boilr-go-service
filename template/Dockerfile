FROM debian:latest

ENV APP_ROOT /opt/apps/go/{{AppName}}
ENV APP_USER {{AppName}}

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install \
    python-pip \
    wget \
    software-properties-common \
    openssh-client

# Setup user
RUN groupadd -r $APP_USER -g 433 && \
    useradd -u 431 -r -g $APP_USER -d /tmp -s /sbin/nologin -c "Docker image user" $APP_USER


ENV VERSION 1.5.3
ENV OS linux
ENV ARCH amd64

RUN wget --no-check-certificate https://storage.googleapis.com/golang/go$VERSION.$OS-$ARCH.tar.gz
RUN tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
RUN rm go$VERSION.$OS-$ARCH.tar.gz
ENV PATH $PATH:/usr/local/go/bin

WORKDIR $APP_ROOT
ENV GOPATH $APP_ROOT/
ENV GOBIN $APP_ROOT/bin
ADD src/ $APP_ROOT/src/
ADD vendor/ $APP_ROOT/vendor/
ADD config/ $APP_ROOT/config/

RUN go get github.com/constabulary/gb/...
RUN $APP_ROOT/bin/gb vendor update --all
RUN $APP_ROOT/bin/gb build

ENTRYPOINT $APP_ROOT/bin/{{AppName}}