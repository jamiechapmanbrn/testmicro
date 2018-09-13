FROM alpine:3.2
ADD testmicro-srv /testmicro-srv
ENTRYPOINT [ "/testmicro-srv" ]
