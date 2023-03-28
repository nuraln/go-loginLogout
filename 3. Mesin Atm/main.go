package main

import (
    "bufio"
    "fmt"
    "os"
)

type Account struct {
    AccountNumber string
    Pin           string
    Balance       float64
}

var accounts = []Account{
    {"123456", "1234", 50000},
    {"654321", "4321", 100000},
}

func main() {
    var accountNumber string
    var pin string

    fmt.Println("Welcome to My Bank ATM")
    fmt.Println("Please enter your account number: ")
    fmt.Scanln(&accountNumber)

    fmt.Println("Please enter your PIN: ")
    fmt.Scanln(&pin)

    // Authenticate the user
    account, authenticated := authenticate(accountNumber, pin)
    if !authenticated {
        fmt.Println("Invalid account number or PIN")
        os.Exit(1)
    }

    // Display the account menu
    for {
        fmt.Println("Welcome, ", account.AccountNumber)
        fmt.Println("1. Withdraw")
        fmt.Println("2. Deposit")
        fmt.Println("3. Balance Inquiry")
        fmt.Println("4. Quit")

        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("Enter option number: ")
        scanner.Scan()
        option := scanner.Text()

        switch option {
        case "1":
            fmt.Print("Enter withdrawal amount: ")
            scanner.Scan()
            amount := scanner.Text()
            withdraw(account, amount)
        case "2":
            fmt.Print("Enter deposit amount: ")
            scanner.Scan()
            amount := scanner.Text()
            deposit(account, amount)
        case "3":
            balanceInquiry(account)
        case "4":
            fmt.Println("Thank you for using My Bank ATM!")
            os.Exit(0)
        default:
            fmt.Println("Invalid option number")
        }
    }
}

func authenticate(accountNumber string, pin string) (Account, bool) {
    for _, account := range accounts {
        if account.AccountNumber == accountNumber && account.Pin == pin {
            return account, true
        }
    }
    return Account{}, false
}

func withdraw(account Account, amount string) {
    var withdrawalAmount float64
    _, err := fmt.Sscanf(amount, "%f", &withdrawalAmount)
    if err != nil {
        fmt.Println("Invalid withdrawal amount")
        return
    }

    if account.Balance < withdrawalAmount {
        fmt.Println("Insufficient funds")
        return
    }

    account.Balance -= withdrawalAmount
    fmt.Printf("Withdrawal of $%.2f successful. Remaining balance: $%.2f\n", withdrawalAmount, account.Balance)
}

func deposit(account Account, amount string) {
    var depositAmount float64
    _, err := fmt.Sscanf(amount, "%f", &depositAmount)
    if err != nil {
        fmt.Println("Invalid deposit amount")
        return
    }

    account.Balance += depositAmount
    fmt.Printf("Deposit of $%.2f successful. New balance: $%.2f\n", depositAmount, account.Balance)
}

func balanceInquiry(account Account) {
    fmt.Printf("Your current balance is: $%.2f\n", account.Balance)
}
