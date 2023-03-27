package main

import (
    "fmt"
)

func main() {
    var username, password string
    isLoggedIn := false

    for {
        if isLoggedIn {
            fmt.Println("Anda sudah login")
            fmt.Println("Ketik 'logout' untuk keluar")
        } else {
            fmt.Println("Silakan login")
        }

        fmt.Print("Username: ")
        fmt.Scanln(&username)
        fmt.Scanln(&username)

        if username == "admin" {
            fmt.Print("Password: ")
            fmt.Scanln(&password)

            if password == "password123" {
                isLoggedIn = true
                fmt.Println("Login berhasil!")
            } else {
                fmt.Println("Password salah. Silakan coba lagi")
            }
        } else if username == "logout" && isLoggedIn {
            isLoggedIn = false
            fmt.Println("Logout berhasil!")
        } else {
            fmt.Println("Username salah. Silakan coba lagi")
        }

        fmt.Println()
    }
}
