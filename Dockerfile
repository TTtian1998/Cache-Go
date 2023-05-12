FROM ubuntu:latest
LABEL authors="tzx"

ENTRYPOINT ["top", "-b"]