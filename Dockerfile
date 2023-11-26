FROM golang:buster

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/hello-world

ENV SCOPE=${SCOPE}
ENV TELEGRAM_TOKEN=${TELEGRAM_TOKEN}
ENV BUILDER_READ_RANGE=${BUILDER_READ_RANGE}
ENV BUILDER_SPREADSHEET_ID=${BUILDER_SPREADSHEET_ID}
ENV REQUESTS_SPREADSHEET_ID=${REQUESTS_SPREADSHEET_ID}
ENV REQUESTS_READ_RANGE=${REQUESTS_READ_RANGE}
ENV GOOGLE_CLOUD_CONFIG=${GOOGLE_CLOUD_CONFIG}

RUN echo $SCOPE
RUN echo $TELEGRAM_TOKEN
RUN echo $BUILDER_READ_RANGE
RUN echo $BUILDER_SPREADSHEET_ID
RUN echo $REQUESTS_SPREADSHEET_ID
RUN echo $REQUESTS_READ_RANGE
RUN echo $GOOGLE_CLOUD_CONFIG

RUN env

EXPOSE 8080
CMD ["/usr/local/bin/hello-world"]