package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/appliedgocourses/bank"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	// Restore the bank data.
	err := bank.Load()
	if err != nil {
		log.Println(err)
		return
	}

	// Save the bank data.
	defer func() {
		err := bank.Save()
		if err != nil {
			log.Println(err)
		}
	}()

	// Perform the action.
	switch os.Args[1] {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bank create <name>")
			fmt.Println("Please provide the account name that is to be created.")
			return
		}
		name := os.Args[2]
		account := bank.NewAccount(name)
		fmt.Printf("New account created with name %q with balance %d\n", account.Name, account.Bal)
	case "list":
		fmt.Println(bank.ListAccounts())
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: bank update <name> <amount>")
			fmt.Println("Please provide the account name and the amount. Provide positive amount for deposit or negative amount for withdraw.")
			return
		}

		name := os.Args[2]
		amount, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Please provide a positive or negative integer for amount.")
			return
		}

		balance, err := update(name, amount)
		if err != nil {
			fmt.Println(err)
			return
		}

		if amount < 0 {
			fmt.Printf("Withdraw %d successful\n", -amount)
		} else {
			fmt.Printf("Deposit %d successful\n", amount)
		}
		fmt.Printf("Balance %q: %d\n", name, balance)

	case "transfer":
		if len(os.Args) < 5 {
			fmt.Println("Usage: bank transfer <name> <name> <amount>")
			fmt.Println("Please provide two account names (from and to) and the amount to transfer. Amount must be positive.")
			return
		}

		fromAccountName := os.Args[2]
		toAccountName := os.Args[3]

		amount, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println("Please provide transfer amount.")
			return
		}

		if amount < 0 {
			fmt.Println("Amount to transfer must be positive.")
			return
		}

		fromAccountBalance, toAccountBalance, err := transfer(fromAccountName, toAccountName, amount)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Transfer made from %q to %q for %d", fromAccountName, toAccountName, amount)
		fmt.Printf("Balance %q: %d\n", fromAccountName, fromAccountBalance)
		fmt.Printf("Balance %q: %d\n", toAccountName, toAccountBalance)
	case "history":
		if len(os.Args) < 3 {
			fmt.Println("Usage: bank history <name> ")
			fmt.Println("Please provide the account name.")
			return
		}

		name := os.Args[2]
		history, err := history(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(history)
	default:
		usage()
	}
}

func usage() {
	fmt.Println(`Usage:
------
bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
}

// update takes a name and an amount, deposits the amount if it
// is greater than zero, or withdraws it if it is less than zero,
// and returns the new balance and any error that occurred.
func update(name string, amount int) (int, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Account %q does not exist.\n", name))
	}

	// If the amount is less that 0 then withdraw
	if amount < 0 {
		withdraw := -amount // make withdraw amount positive
		balance, err := bank.Withdraw(account, withdraw)
		if err != nil {
			return 0, err
		}
		return balance, nil
	}

	// Deposit
	balance, err := bank.Deposit(account, amount)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

// transfer takes two names and an amount, transfers the amount from
// the account belonging to name #1 to the account belonging to name #2,
// and returns the new balances of both accounts and any error that occurred.
func transfer(name, name2 string, amount int) (int, int, error) {
	fromAccount, err := bank.GetAccount(name)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("From account %q does not exist.\n", name))
	}

	toAccount, err := bank.GetAccount(name2)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("To account %q does not exist.\n", name2))
	}

	if fromAccount.Bal < amount {
		return 0, 0, errors.New(fmt.Sprintf("Insufficient balance (%d)\n.", fromAccount.Bal))
	}

	fromAccountBalance, err := bank.Withdraw(fromAccount, amount)
	if err != nil {
		return 0, 0, err
	}

	toAccountBalance, err := bank.Deposit(toAccount, amount)
	if err != nil {
		return 0, 0, err
	}

	return fromAccountBalance, toAccountBalance, nil
}

// history takes an account name, retrieves the account, and calls bank.History()
// to get the history closure function. Then it calls the closure in a loop,
// formatting the return values and appending the result to the output string, until the boolean return parameter of the closure is `false`.
func history(name string) (string, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Account %q does not exist.\n", name))
	}

	var history []string
	next := bank.History(account)
	for {
		amount, balance, more := next()
		action := "Deposit"
		if amount < 0 {
			action = "Withdraw"
			amount = -amount
		}

		history = append(history, fmt.Sprintf("%s %d, Balance %d", action, amount, balance))

		if more == false {
			break
		}
	}

	return strings.Join(history, "\n"), nil
}
