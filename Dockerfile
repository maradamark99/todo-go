FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go build -o /bin/app

FROM scratch
COPY --from=0 /bin/app /bin/app

CMD ["/bin/app"]

EXPOSE 3000