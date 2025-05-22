CHAFAVERSION = 1.16
PREFIX       = $(CURDIR)/build
LIBDIR       = libs

.PHONY: all clean

all: $(LIBDIR)/libchafa-linux-amd64.so $(LIBDIR)/libchafa-darwin-amd64.dylib

$(LIBDIR)/libchafa-linux-amd64.so:
	mkdir -p build/chafa-linux && cd build/chafa-linux && \
	git clone --branch $(CHAFAVERSION) --depth 1 https://github.com/hpjansson/chafa.git . && \
	./autogen.sh --without-tools && make
	cp build/chafa-linux/chafa/.libs/libchafa.so $(CURDIR)/$(LIBDIR)/libchafa-linux-amd64.so

$(LIBDIR)/libchafa-darwin-amd64.dylib:
	mkdir -p build/chafa-darwin && cd build/chafa-darwin && \
	git clone --branch $(CHAFAVERSION) --depth 1 https://github.com/hpjansson/chafa.git . && \
	./autogen.sh && make
	cp build/chafa-darwin/chafa/.libs/libchafa.dylib $(CURDIR)/$(LIBDIR)/libchafa-darwin-amd64.dylib

clean:
	rm -rf build $(LIBDIR)/*.so $(LIBDIR)/*.dylib

