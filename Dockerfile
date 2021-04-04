FROM golang:1.16-stretch

WORKDIR /app

COPY . .

RUN go mod download && go get github.com/cespare/reflex

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]