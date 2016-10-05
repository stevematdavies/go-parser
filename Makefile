BINARY=gomake

# These are the values you want to pass to the build
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

# Set up the -ldflags option for go, interpolate the values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL:	${BINARY}

${BINARY}:
		go build ${LDFLAGS} -o ${BINARY} ./..

# installs nescessary binaries
install:
	go install ${LDFLAGS} -o ${BINARY} ./..

# cleans out projects cleans binaries
.PHONY: clean install

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi