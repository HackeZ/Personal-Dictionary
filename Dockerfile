FROM scratch
MAINTAINER HackerZ
RUN mkdir /PD
ADD main /PD
ADD conf /PD/conf
ADD views /PD/views
ADD static /PD/static
EXPOSE 8081
CMD ["/PD/main"]
