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

```Go
 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
```
<br>
<p>with setFlag we could specify the needed information and get it in log file as log. like date, time or file name that log happened in to it </p>
<br>