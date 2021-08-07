# 行空

天马行空, 以梦为马.

## 常用命令

```bash
swag init --output api --generalInfo ./cmd/server/main.go
```

## 风格

- 变量命名尽量简短, 反正有类型提示, 名字越短越好
- struct 方法的本体参数都命名为 g, 即 `func (g Type) funcName`
- struct 里如果包含其他 struct, 即嵌套的话, 都使用指针
- struct 方法的本体参数的类型, 要统一, 那么都用值, 要么都用指针
