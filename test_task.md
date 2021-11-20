1) Реализовать CRUD сервер для Блога, используя фреймворк Fiber(https://github.com/gofiber/fiber) и шаблон чистой архитектуры (https://github.com/evrone/go-clean-template). В системе реализовать имитацию интеграции с внешним REST сервером

2) Что выведется в приведенном ниже коде? Почему?

```
package main

import "fmt"

func a() {
    x := []int{}
    x = append(x, 0)
    x = append(x, 1)
    x = append(x, 2)
    y := append(x, 3)
    z := append(x, 4)
    fmt.Println(y, z)
}

func main() {
    a()
}
```
