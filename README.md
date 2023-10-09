# MENAMPUNG PIKIRAN PAS BELAJAR UNTUK DITULIS KALAU DI INGET-INGET SUKA KE INGET DIA

### notes
```
    data := struct {
		User []User
	}{
		User: users,
	}
    err = tmpl.Execute(w, data)
	// jika seperti ini akses di html range User

    var users = []User{}
    err = tmpl.Execute(w, users)
	// jika langsung users range .
```