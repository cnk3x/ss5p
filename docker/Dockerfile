FROM alpine:3.8

ADD ss5p entrypoint.sh /

RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT [ "/entrypoint.sh" ]