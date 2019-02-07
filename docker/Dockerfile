FROM golang



# create a working directory
WORKDIR /go/src/app


# add source code
ADD Log  Log

EXPOSE 9090

# run main.go
CMD ["go", "run", "Log/src/logging.go"]
