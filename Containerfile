FROM scratch
COPY serv /
ENTRYPOINT ["/serv"]
