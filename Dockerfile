FROM qnib/uplain-golang

WORKDIR /usr/local/src/github.com/qnib/go-test/
COPY vendor/vendor.json ./vendor/vendor.json
COPY main.go ./main.go
RUN govendor fetch +m \
 && govendor build 

FROM scratch

COPY --from=0 /usr/local/src/github.com/qnib/go-test/go-test  /
ENTRYPOINT ["/go-test"]
