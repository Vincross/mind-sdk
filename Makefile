.PHONY : xcompile cli

build : xcompile cli

xcompile : 
	cd xcompile && make

cli :
	cd cli && make

clean :
	cd xcompile && make clean
	cd cli && make clean
