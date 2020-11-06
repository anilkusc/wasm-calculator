FROM golang:1.15 as BUILD
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" /bin/
RUN GOARCH=wasm GOOS=js go build -o /bin/main.wasm
FROM nginx
RUN rm -fr /usr/share/nginx/html/*
COPY --from=BUILD /bin/main.wasm /usr/share/nginx/html/
COPY --from=BUILD /bin/wasm_exec.js /usr/share/nginx/html/
COPY ./index.html /usr/share/nginx/html/
