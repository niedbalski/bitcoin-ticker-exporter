FROM debian:stretch
MAINTAINER  Jorge Niedbalski <jnr@metaklass.org>
RUN apt-get update -yyq
RUN apt-get install -y ca-certificates

COPY ./bitcoin-ticker-exporter /opt/
RUN chmod 0755 /opt/bitcoin-ticker-exporter
ENTRYPOINT ["/opt/bitcoin-ticker-exporter"]
