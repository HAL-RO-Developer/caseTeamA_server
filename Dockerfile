FROM golang:latest

WORKDIR /go/src/github.com/HAL-RO-Developer/caseTeamA_server


ADD ./ ./

RUN cp config.yml.template config.yml

RUN ls -la

EXPOSE 8000

ENTRYPOINT ["go","run","main.go"]

