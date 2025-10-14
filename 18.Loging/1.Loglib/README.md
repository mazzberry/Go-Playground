```Go
func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("something went wrong with log file", err)
	}
	log.SetOutput(file)
}
```
<br>
we've just set up a log file 
<br>

