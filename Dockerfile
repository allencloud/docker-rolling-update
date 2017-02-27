FROM daocloud.io/golang:1.8

EXPOSE 3000 4000

HEALTHCHECK --interval=1s --timeout=1s --retries=5 CMD curl --connect-timeout 1 –f http://127.0.0.1:4000/ || exit 1

COPY . /go

RUN go build -o webserver

CMD ["./webserver"]
