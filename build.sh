mkdir -p output/bin
cp script/bootstrap.sh script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh output/bootstrap.sh
go build -o ./output/bin/ai-assitant ./cmd/http