FROM golang:1.9

WORKDIR /go/src/ct-budget-manager
COPY . .

RUN echo $GOPATH
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 80

CMD ["go-wrapper", "run"]
