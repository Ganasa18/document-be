FROM golang:1.22.0-alpine AS gobuilder
LABEL stage=builder-document-be-CHANGECOMMIT
WORKDIR /go/src/github.com/Ganasa18/document-be
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /document-be ./main.go