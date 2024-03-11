FROM golang:latest AS build

COPY . .
WORKDIR /go/app

RUN go mod download
RUN go build .

FROM alpine:latest

WORKDIR /app

RUN apk add ocrmypdf --no-cache

COPY --from=build /go/app/ocrTool /app/ocrTool

VOLUME [ "/data" ]
VOLUME [ "/usr/share/tessdata/" ]

CMD [ "./ocrTool" ]