FROM --platform=linux/amd64 amazonlinux:2

RUN yum install -y sqlite sqlite-devel golang

RUN mkdir /app

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o bootstrap ./functions/lambda-entrypoint
