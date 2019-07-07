FROM golang:1.12 as build

WORKDIR /tmp
# use modules
RUN go get -d -v github.com/filex/echo-json && mv /go/src/github.com/filex/echo-json .
RUN cd echo-json && go build -ldflags "-s -w" -o /echo-json

FROM scratch
COPY --from=build /echo-json /echo-json
RUN ["/echo-json", "works", "great!"]
ENTRYPOINT ["/echo-json"]

# This image is not really useful for direct execution. However, you can use it
# to see how echo-json works:
#
# $ docker run --rm -it filex/echo-json foo bar baz:int 123
# {"baz:int":"123","foo":"bar"}
#
# The actual purpose of this image is being available as a COPY --from target
# if you want to use echo-json in your Docker image:
#
# FROM filex/echo-json as echo-json
#
# FROM alpine # … your image starts here
# …
# COPY --from=echo-json /echo-json /usr/local/bin/
