# OTBB Quiz package
A quiz generator with questions from opentdb.com



## Examples
### Raw Quiz Results
```
// To obtain the raw response from opentdb, use the Raw function
config := otdbquiz.DefaultClient("https://opentdb.com/api.php?amount=1")

quiz, err := otdbquiz.Raw(config)

```

### Standard Quiz Results
```
// The Standard function takes the raw response and formats the response with random multiple choices and ordered true/false questions
url := "https://opentdb.com/api.php?amount=1"
config := otdbquiz.DefaultClient(url)

quiz, err := otdbquiz.Standard(config)

```


## Full documentation
 - `godoc -http=:6060`