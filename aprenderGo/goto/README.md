# Go `goto` Statement Example

## 📌 Description
This project demonstrates the use of the **`goto` statement** in Go.

`goto` transfers program control to a labeled statement within the same function.

⚠️ Its use is generally discouraged except in very specific scenarios.

---

## 🚀 Technologies Used
- Golang
- fmt package

---

## 🔀 What Does `goto` Do?
- Jumps execution to a labeled line of code
- The label must be in the same function
- Skips any code between the jump and the label

---

## 📄 Code Explanation

```go
fmt.Println("Antes de goto")
goto End
fmt.Println("Esto no se ejecutara")

End:
fmt.Println("Despues goto")
