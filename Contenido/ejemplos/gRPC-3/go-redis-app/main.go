package main

import (
    "context"
    "fmt"
    "math/rand"
    "os"
    "time"

    "github.com/redis/go-redis/v9"
)

func main() {
    ctx := context.Background()

    redisHost := os.Getenv("REDIS_HOST")
    redisPassword := os.Getenv("REDIS_PASSWORD")

    rdb := redis.NewClient(&redis.Options{
        Addr:     redisHost,
        Password: redisPassword,
        DB:       0,
    })

    nombre := "Sergio Ferrer"
    ciudad := "Guatemala"

    rand.Seed(time.Now().UnixNano())

    for {
        edad := rand.Intn(41) + 20 // 20-60 años
        timestamp := time.Now().Format("20060102_150405")
        key := fmt.Sprintf("registro:%s", timestamp)
        value := fmt.Sprintf("Nombre: %s | Ciudad: %s | Edad: %d", nombre, ciudad, edad)

        err := rdb.Set(ctx, key, value, 0).Err()
        if err != nil {
            fmt.Printf("Error al guardar en Redis: %v\n", err)
        } else {
            fmt.Printf("Registro guardado [%s]: %s\n", key, value)
        }

        time.Sleep(5 * time.Second)
    }
}
