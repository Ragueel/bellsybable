FROM go-1.21.6:alpine


COPY go.mod go.mod go.sum go.sum
RUN go get

RUN go build -o /bin/web ./cmd/web/main.go
