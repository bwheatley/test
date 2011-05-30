include $(GOROOT)/src/Make.inc

TARG=gocurses
CGOFILES=\
        gocurses.go\

#GOFILES=\
#        main.go\

CGO_LDFLAGS=-lncurses

CLEANFILES+= $(TARG) main 
CGO_OFILES+= 
OBJECTS = main.go talkmore.go

include $(GOROOT)/src/Make.pkg

#%: install %.go main.o
#	$(GC) $*.go
#	$(LD) -o $@ $*.$O


#%: install %.go
#	$(GC) $*.go
#	$(LD) -o $@ $*.$O


main: install main.go talkmore.go
	$(GC) $(OBJECTS)
	$(LD) -o $(TARG) main.$O

#all: install main.go talkmore.go
#	$(GC) $(OBJECTS)
#	$(LD) -o $(TARG) main.$O

%: install $(OBJECTS)
	$(GC) $(OBJECTS)
	$(LD) -o $(TARG) main.$O

