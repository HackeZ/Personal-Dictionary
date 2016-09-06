FROM golang
MAINTAINER HackerZ
RUN go get github.com/astaxie/beego && go get github.com/beego/bee 
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/go-sql-driver/mysql && go get github.com/lib/pq && go get github.com/mattn/go-sqlite3
RUN go get github.com/russross/blackfriday
RUN mkdir /PD
ADD main /PD
ADD conf /PD/conf
ADD views /PD/views
ADD static /PD/static
EXPOSE 8081
CMD ["bee","run","/PD"]
