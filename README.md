# wasm-calculator
Simple Web Assembly Calculator written in Golang.<br>
For build and run locally<br>
> cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" . <br>
> GOARCH=wasm GOOS=js go build -o main.wasm main.go <br>
> goexec 'http.ListenAndServe(\`:9090\`, http.FileServer(http.Dir(\`.\`)))' <br>
## PoC
