FROM centos:latest
LABEL maintainer="Jonathan Mainguy <jon@soh.re>"
ENV RELEASE=0.0.5
ADD run.sh /opt/
WORKDIR /opt
RUN yum install -y wget; wget -O /opt/ghreview_Linux_x86_64.tar.gz "https://github.com/Jmainguy/ghreview/releases/download/v${RELEASE}/ghreview_${RELEASE}_linux_386.tar.gz"; yum clean all
RUN chmod -R 777 /opt 
CMD ["/opt/run.sh"]
