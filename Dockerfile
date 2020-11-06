FROM golang:1.15 as BUILD
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" /bin/
RUN GOARCH=wasm GOOS=js go build -o /bin/main.wasm main.go
RUN CGO_ENABLED=0 go build -o /bin/server server.go
FROM nginx
WORKDIR /app
COPY --from=BUILD /bin/main.wasm .
COPY --from=BUILD /bin/server .
COPY --from=BUILD /bin/wasm_exec.js .
COPY ./index.html .
COPY entrypoint.sh .
COPY default.conf /etc/nginx/conf.d/default.conf
RUN chmod 777 entrypoint.sh
ENTRYPOINT ./entrypoint.sh
