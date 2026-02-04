# Go `switch` with `fallthrough`

## 📌 Description
This project demonstrates the use of the **`fallthrough` keyword** in Go `switch` statements.

Unlike most languages, Go automatically breaks after each case.  
`fallthrough` forces execution to continue to the next case.

---

## 🚀 Technologies Used
- Golang
- fmt package

---

## 🔍 What Is `fallthrough`?
- Executes the **next case block unconditionally**
- Does **not evaluate the next case condition**
- Used only inside `switch`

⚠️ Use with care — it can reduce code clarity.

---

## 📄 Code Explanation

### 🔹 Example 1
```go
x := 3
switch x {
case 1:
	fmt.Println("Case 1")
	fallthrough
case 2:
	fmt.Println("Case 2")
}
