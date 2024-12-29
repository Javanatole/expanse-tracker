# Expense Tracker CLI Example

This document demonstrates how to use the Expense Tracker CLI. The commands below show how to add, list, delete, and summarize expenses.

## Adding Expenses

```bash
$ go run main.go add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ go run main.go add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)
```

## Listing Expenses
```bash
$ go run main.go list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10
```

## Viewing Expense Summary

```bash
$ go run main.go summary
# Total expenses: $30
```

## Deleting an Expense

```bash
$ go run main.go delete --id 2
# Expense deleted successfully
```

## Viewing Updated Summary

```bash
$ go run main.go summary
# Total expenses: $20
```

## Viewing Monthly Expense Summary

```bash
$ go run main.go summary --month 8
# Total expenses for month 8: $20
```