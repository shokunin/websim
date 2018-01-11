# Note you cannot run golang binaries on Alpine directly
FROM            debian:buster-slim

MAINTAINER      chris.mague@shokunin.co

COPY            websim /websim

WORKDIR		/
ENV		GIN_MODE=release

EXPOSE          8080

ENTRYPOINT      [ "/websim" ]
