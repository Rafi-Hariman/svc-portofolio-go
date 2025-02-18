FROM golang:1.21.0-bullseye

RUN useradd -rm -d /home/bullseye -s /bin/bash -g root -G sudo -u 1001 bullseye

USER bullseye

ADD main /

EXPOSE 80

CMD ["/main"]