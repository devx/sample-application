FROM scratch
MAINTAINER Ryan Richard <ryan@kumoru.io>

ADD templates templates
COPY kumoru-sample-app /

EXPOSE 8080
ENTRYPOINT [ "/kumoru-sample-app" ]
