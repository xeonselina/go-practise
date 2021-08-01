1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

   
   
   答：不应该 warp，直接返回  如果返回值是 `[]rows,error`,那么我会返回  `[], nil`. 因为如果返回一个 warp 过的 ErrNoRows, warp 的过程你就已经用了 `if err.Is(sql.ErrNoRows)` 做了一些处理，然后 service 层又要这样判断一次，违反了错误只处理一次的理念，所以不如干脆不返回错误
    代码例如：

```go
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil && err != sql.ErrNoRows {
  return nil, err
}
return name, err
```