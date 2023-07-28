FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -v -o /app/go_hospitalnative

EXPOSE 8071

ENTRYPOINT [ "/app/go_hospitalnative" ]
CMD [ "serve" ]