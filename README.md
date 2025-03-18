# Toledo API Client

Клиент для работы с API Toledo24.

## Установка

```
go get -u github.com/24example/toledoapi
```

## Доступные методы

> GetBrands - получение списка производителей
> GetProducts - получение списка продуктов
> GetProductPrices - получение списка цен

## Пример использования

```go
package main

import (
	"fmt"
	"log"

	"github.com/24example/toledoapi/api"
	"github.com/24example/toledoapi/models"
)

func main() {
	api := api.NewApi("apiKey")
	products, err := api.GetProducts(&models.ProductFilter{
		Filter: &models.Filter{
			Products: []string{"15533"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(products)
}
```
