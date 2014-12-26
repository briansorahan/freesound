PROGS = fsnd

LIBS = freesound

.PHONY: all clean test freesound

all: $(LIBS) $(PROGS)

fsnd: fsnd.go
	go build $^

freesound:
	cd freesound && go install -a

clean:
	rm -rf $(PROGS) *~

test:
	cd freesound && go test
