# timber
 
A flexable logging package

```Go
var output bytes.Buffer

jack := timber.NewJack(
	timber.WithWriter(&output),
)

jack.Debug("this will debug")
jack.Error("this will error")
jack.Alert("this will alert")

fmt.Println(output.String())
// Output:
// DEBUG: this will debug
// ERROR: this will error
// ALERT: this will alert
```