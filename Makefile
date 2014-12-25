PROGS = freesnd

LIBS = freesound

.PHONY: freesound

all: $(PROGS) $(LIBS)

freesnd: freesnd.go
	go build freesnd.go

freesound:
	cd freesound && go install -a

clean:
	rm -rf $(PROGS) *~
