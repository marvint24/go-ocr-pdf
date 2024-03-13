FROM golang:latest AS build

COPY . .
WORKDIR /go/app

RUN go mod download
RUN go build .

FROM ubuntu:latest	

WORKDIR /app

RUN apt update
RUN apt install -y ocrmypdf

COPY --from=build /go/app/ocrTool /app/ocrTool

VOLUME [ "/data" ]
VOLUME [ "/usr/share/tesseract-ocr/4.00/tessdata/" ]

CMD [ "./ocrTool" ]