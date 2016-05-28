FROM scratch

ADD ca-bundle.crt /etc/ssl/certs/
ADD rfc-bot /
ADD .env /

CMD ["/rfc-bot"]
