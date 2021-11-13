FROM golang:1.17.3-buster
RUN apt install curl
RUN mkdir /app
ADD . /app
WORKDIR /app
#ENV TA_CONFIG_FILE /app/config.yml
RUN ls -la
#RUN go build -tags 'traceable_filter' -o auth .
RUN go build -o auth .
#RUN curl -sSL https://raw.githubusercontent.com/Traceableai/goagent/main/filter/traceable/copy-library.sh | bash -s -- .
CMD ["/app/auth"]
