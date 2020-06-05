## Usage

### Creating an account

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "time"

    accountapi "github.com/alexdreptu/form3-accountapi-client"
)

func main() {
    options := &accountapi.Options{
        Type:           "accounts",
        ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
        OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
        Attributes: []accountapi.Attribute{
            accountapi.WithAttrCountry("GB"), // you can use accountapi.CountryUnitedKingdom
            accountapi.WithAttrBIC("NWBKGB22"),
            accountapi.WithAttrBankID("400300"),
            accountapi.WithAttrBankIDCode("GBDSC"), // you can use accountapi.BankIDCodeUnitedKingdom
            accountapi.WithAttrAccountNumber("41426815"),
            accountapi.WithAttrBaseCurrency("GBP"), // you can use accountapi.CurrencyUnitedKingdom
            accountapi.WithAttrJointAccount(false),
            accountapi.WithAttrFirstName("Samantha"),
            accountapi.WithAttrAlternativeBankAccountNames(
                "Lola Andrews", "Mitchell Davis", "Francisco Andrews",
            ),
            accountapi.WithAttrAccountMatchingOptOut(true),
            accountapi.WithAttrCustomerID("5019343427"),
        },
    }

    account, err := accountapi.NewAccount(options)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    client := accountapi.NewClient(&http.Client{})
    createdAccount, err := client.CreateAccount(ctx, account)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    fmt.Printf("Created account '%s' on %s\n",
        createdAccount.Data.ID, createdAccount.Data.CreatedOn)
}
```

### Fetching an account

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "time"

    accountapi "github.com/alexdreptu/form3-accountapi-client"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    client := accountapi.NewClient(&http.Client{})
    fetchedAccount, err := client.FetchAccount(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    fmt.Printf("Account '%s' was created on %s\n",
        fetchedAccount.Data.ID, fetchedAccount.Data.CreatedOn)
}
```

### Listing accounts

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "time"

    accountapi "github.com/alexdreptu/form3-accountapi-client"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    client := accountapi.NewClient(&http.Client{})
    accountList, err := client.ListAccounts(ctx, 0, 5)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    for _, account := range accountList.Data {
        fmt.Printf("Account '%s' was created on %s\n",
            account.ID, account.CreatedOn)
    }
}
```

### Deleting an account

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "time"

    accountapi "github.com/alexdreptu/form3-accountapi-client"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    client := accountapi.NewClient(&http.Client{})
    err := client.DeleteAccount(ctx, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```
