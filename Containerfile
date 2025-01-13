FROM alpine
COPY serv /
ENTRYPOINT ["/serv"]
