FROM busybox:latest
LABEL maintainer="Jonathan Mainguy <jon@soh.re>"
ENV RELEASE=v0.0.1
RUN mkdir /opt
ADD run.sh /opt/
ADD static /opt/static
WORKDIR /opt
RUN wget -O /opt/ghreview_Linux_x86_64.tar.gz "https://github.com/Jmainguy/ghreview/releases/download/v0.0.1/ghreview_0.0.1_linux_386.tar.gz"
RUN chmod -R 777 /opt 
CMD ["/opt/run.sh"]
