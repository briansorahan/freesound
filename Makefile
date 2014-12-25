PROGS = freesnd

all: $(PROGS)

freesnd: freesnd.go
	go build freesnd.go

clean:
	rm -rf $(PROGS) *~
