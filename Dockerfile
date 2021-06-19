FROM golang AS buildStage

WORKDIR /go/src/github.com/power-sentinel
COPY . .
ENV GIT_TERMINAL_PROMPT=1
RUN go get -insecure golang.org/x/sys/unix
RUN CGO_ENABLED=0  go get -insecure  -v .../.

RUN  CGO_ENABLED=0 go build

FROM alpine

RUN apk --update --no-cache add openssh ca-certificates && update-ca-certificates 2>/dev/null || true

WORKDIR /app
COPY --from=buildStage /go/src/github.com/power-sentinel .

EXPOSE 6669
ENTRYPOINT ["/app/power-sentinel"]
