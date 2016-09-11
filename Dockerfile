FROM ubuntu
MAINTAINER Richard Turek

RUN apt-get update

RUN mkdir /var/dockerapp
RUN echo "Docker image built: $(date)" >> /var/dockerapp/build-datetime

ADD httpsvr /var/dockerapp/

CMD ["/var/dockerapp/httpsvr", "8080"]
