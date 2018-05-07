FROM ubuntu:16.04

ADD conf /conf

ADD static /static

ADD index.html /

ADD testclient /

EXPOSE 8090

CMD ["./testClient"]
