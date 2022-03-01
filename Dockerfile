FROM golang:latest

ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p "awesomeProject"

WORKDIR /awesomeProject
COPY . .

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN cd /awesomeProject && go build ./main.go

EXPOSE 9004
CMD [ "./main" ]
