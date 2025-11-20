```go
// method receiver 
//	(p Person) →  value receiver
//	(p *Person) → pointer receiver

func (p Person) seyHello() string {
	return fmt.Sprintf("%s %s - %s", p.lastname, p.name, p.age)
}
```