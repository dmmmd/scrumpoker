# First, build a binary for Linux manually
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./docker/server/scrumpoker .
# TODO use multistep build to avoid this manual labor
FROM scratch
ADD scrumpoker /
CMD ["/scrumpoker"]
