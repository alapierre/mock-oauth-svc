FROM scratch

MAINTAINER Adrian Lapierre <al@soft-project.pl>

EXPOSE 9005

ADD mock-oauth-svr /mock-oauth-svr

#RUN addgroup -g 666 -S app && adduser -u 666 -S -G app app
#USER app

CMD ["/mock-oauth-svr"]
