# This how we want to name the binary output
BINARY=gomake

# Builds the project
build:
	go build -o ${BINARY}

# Installs our project: copies binaries
install:
	go install

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
