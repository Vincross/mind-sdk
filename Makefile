.PHONY : xcompile cli

build : xcompile cli

install : xcompile cli
	cd cli && make install 

push : xcompile cli
	cd xcompile && make push
	cd cli && make push

xcompile : 
	cd xcompile && make

cli :
	cd cli && make

clean :
	cd xcompile && make clean
	cd cli && make clean
