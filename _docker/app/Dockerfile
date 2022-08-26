FROM golang:1.18.5 as develop

RUN apt-get update && apt-get install -y \
    build-essential \
    default-mysql-client \
 && apt-get clean

WORKDIR /app

COPY Makefile ./

RUN make install-tools

CMD ["air"]
