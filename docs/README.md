### 本地文档

```sh
godoc -http=:6060

#查看
http://localhost:6060/pkg/github.com/researchlab
```

### 生成markdown

```sh
go install github.com/robertkrimen/godocdown/godocdown@latest
godocdown . > docs/v0.0.1.md
```
