FROM centos:latest
MAINTAINER Marcello Paris <marcello.paris@gmail.com>

RUN mkdir /source

RUN yum install -y \
    epel-release

RUN yum install -y \
    net-tools pciutils htop wget unzip git telnet which tar lzip

RUN yum install -y \
    make autoconf automake libtool gcc clang cmake3

RUN yum install -y \
    flex bison

RUN yum install -y \
    python3 python3-libs python3-devel python3-pip

RUN python3 -m pip install --upgrade pip setuptools

# building z3 from sources
RUN cd /source && \
    wget https://github.com/Z3Prover/z3/archive/z3-4.8.7.tar.gz && \
    tar xvf z3-4.8.7.tar.gz && \
    cd z3-z3-4.8.7 && \
    python3 scripts/mk_make.py && \
    cd build && \
    make -j 8 && make install && \
    rm -rf /source/z3-z3-4.8.7

# building gmp from sources
#RUN cd /source && \
#    wget https://gmplib.org/download/gmp/gmp-6.2.0.tar.lz && \
#    lzip -d gmp-6.2.0.tar.lz && \
#    tar xvf gmp-6.2.0.tar && \
#    cd gmp-6.2.0 && \
#    ./configure && make -j 8 && make install && ldconfig && \
#    rm -rf /source/gmp-6.2.0

# building pbc from sources
#RUN cd /source && \
#    wget https://crypto.stanford.edu/pbc/files/pbc-0.5.14.tar.gz && \
#    tar xvf pbc-0.5.14.tar.gz && \
#    cd pbc-0.5.14 && \
#    ./configure && make -j 8 && make install && \
#    echo "/usr/local/lib" > /etc/ld.so.conf.d/pbc.conf && ldconfig -v && \
#    rm -rf /source/pbc-0.5.14

RUN wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz

RUN yum install -y \
    redis

RUN yum install -y \
    leveldb leveldb-devel

#RUN cd /source && \
#    git clone https://github.com/miracl/MIRACL

ENV PATH="${PATH}:/usr/local/go/bin"

RUN mkdir /mosaic
ADD ./abe /mosaic/abe
ADD ./lib/mosaic.go /mosaic/lib
ADD ./examples /mosaic/examples
ADD ./service /mosaic/service
ADD go.mod /mosaic
ADD go.sum /mosaic

ADD ./docker/docker-entrypoint.sh /mosaic

WORKDIR /mosaic

ENTRYPOINT ["/mosaic/docker-entrypoint.sh"]

CMD ["/bin/bash"]
