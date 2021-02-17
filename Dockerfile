FROM golang:1.15
WORKDIR /app/src/
COPY . .
RUN cd ./ && go build
CMD ["/app/src/nix_blog"]