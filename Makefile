GOPATH = $(shell go env GOPATH)
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
#GOPARM  = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GOPARM = 
DOCKERFILE = ./docker/Dockerfile
APP = ./docker/app
APPPLUGINFOLDER = $(APP)/datacreator
CREATOR_DIR = $(dir ./datacreator/*/)

all: clean build buildplugin

clean:
	$(GOCLEAN) -r

	for dir in $(CREATOR_DIR); do \
        $(GOCLEAN) $$dir; \
    done

build:
	$(GOPARM) $(GOBUILD) main.go

buildplugin:
	$(GOPARM) $(GOBUILD) -buildmode=plugin -o ./datacreator/apiproviderlookup/apiproviderlookup.so ./datacreator/apiproviderlookup/apiproviderlookup.go
	$(GOPARM) $(GOBUILD) -buildmode=plugin -o ./datacreator/etlproviderlookup/etlproviderlookup.so ./datacreator/etlproviderlookup/etlproviderlookup.go
	$(GOPARM) $(GOBUILD) -buildmode=plugin -o ./datacreator/userlookup/userlookup.so ./datacreator/userlookup/userlookup.go

copy:
	mkdir -p $(APPPLUGINFOLDER)/apiproviderlookup
	mkdir -p $(APPPLUGINFOLDER)/etlproviderlookup
	mkdir -p $(APPPLUGINFOLDER)/userlookup
	cp ./main $(APP)/main
	cp ./datacreator.yaml $(APP)/datacreator.yaml

	cp ./datacreator/apiproviderlookup/apiproviderlookup.so $(APPPLUGINFOLDER)/apiproviderlookup/apiproviderlookup.so
	cp ./datacreator/etlproviderlookup/etlproviderlookup.so $(APPPLUGINFOLDER)/etlproviderlookup/etlproviderlookup.so
	cp ./datacreator/userlookup/userlookup.so $(APPPLUGINFOLDER)/userlookup/userlookup.so

dockerbuild: copy
	docker build --file $(DOCKERFILE) --build-arg APP=$(APP) -t registry.gitlab.com/iisidotnetgroup/ptxgrafanadatacreator .
	rm -r $(APP)

dockerpush:
	docker login registry.gitlab.com
	docker push registry.gitlab.com/iisidotnetgroup/ptxgrafanadatacreator

dockerrun:
	docker run --rm -it -p 8888:8888 registry.gitlab.com/iisidotnetgroup/ptxgrafanadatacreator

test:
	for dir in $(DIRECTORY); do \
        ls -l $$dir; \
    done
