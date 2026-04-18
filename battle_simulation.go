package main

import "fmt"

func main() {
    pirateHealth := 100
    enemyHealth := 80

    fmt.Println("⚔️ Battle begins!")

    pirateHealth -= 20
    enemyHealth -= 35

    fmt.Printf("Pirate Health: %d\n", pirateHealth)
    fmt.Printf("Enemy Health: %d\n", enemyHealth)
}
