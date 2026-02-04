# Go `defer` Statement (LIFO Behavior)

## 📌 Description
This project demonstrates the use of the **`defer` keyword in Go**, which delays the execution of a function until the surrounding function returns.

`defer` is commonly used for:
- Resource cleanup
- Closing files
- Releasing memory
- Logging

---

## 🚀 Technologies Used
- Golang
- fmt package

---

## ⏳ What Does `defer` Do?
- `defer` postpones function execution
- Deferred functions run **after `main()` finishes**
- Deferred calls follow **LIFO order** (Last In, First Out)

> 🧠 First in → Last out

---

## 📄 Code Explanation

### 🔹 Simple `defer`
```go
defer fmt.Println("Esto esta impreso ultimo")
fmt.Println("Esto esta impreso primero")
