FROM alpine:3.7

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD cloud-controller-manager /bin/

ENTRYPOINT [ "/bin/cloud-controller-manager" ]