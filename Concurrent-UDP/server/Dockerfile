FROM ubuntu:18.04

LABEL Developer="Lakshya Nanjangud <lnanjangud@vmware.com>"

LABEL Maintainer="Solution Engineering Team - TKG(MAPBU)"

WORKDIR /usr/src/app

EXPOSE 10001/udp

COPY udpserver .

RUN chmod 777 udpserver

CMD ./udpserver
