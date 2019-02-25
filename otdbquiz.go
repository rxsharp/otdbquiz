// Package otdbquiz A quiz generator with questions from opentdb.com
package otdbquiz

import (
	"encoding/json"
	_"log" //testing
	"math/rand"
	"net/http"
	"time"
)

const defaultTimeout = time.Second * 10

var defaultClient = http.Client{Timeout: defaultTimeout}

// Config is the obj for each quiz request
type Config struct {
	*http.Client
	URL string
}

// Response is the object returned from a request to the quiz api
type Response struct {
	Result []Result `json:"results"`
}

// Result is the individual quiz items of the original Response object
type Result struct {
	Category        string   `json:"category"`
	Type            string   `json:"type"`
	Difficulty      string   `json:"difficulty"`
	Question        string   `json:"question"`
	CorrectAnswer   string   `json:"correct_answer"`
	IncorrectAnswer []string `json:"incorrect_answers"`
}

func init() {
	// Set randomization seed
	rand.Seed(time.Now().UnixNano())
}

// SetConfig for custom configuration
func SetConfig(client *http.Client, url string) (config *Config) {
	return &Config{
		client,
		url,
	}
}

// DefaultClient returns a Config object with default client and specified url
func DefaultClient(url string) (config *Config) {
	return &Config{
		&defaultClient,
		url,
	}
}

// Raw makes a client request with Config fields
func Raw(config *Config) (raw *Response, err error) {
	raw = new(Response)
	request, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return nil, err
	}

	response, err := config.Client.Do(request)
	err = json.NewDecoder(response.Body).Decode(raw)
	return
}

// Standard creates a quiz with random multiple choices and ordered true/false questions
func Standard(config *Config) (standardized *Response, err error) {
	quiz, err := Raw(config)
	if err != nil {
		return nil, err
	}

	quiz.standardize()
	return quiz, nil
}

// setChoices will randomize muliple choices or order true/fase questions
func (quiz Response) standardize() (err error) {
	for i, item := range quiz.Result {
		quiz.Result[i].CombineChoices()
		if len(item.IncorrectAnswer) > 1 {
			// randomize multiple choice
			quiz.Result[i].ShuffleChoices()
		} else {
			// order true/false
			quiz.Result[i].OrderBoolean()
		}

	}
	return nil
}

// CombineChoices appends correct and incorrect answers to Result's []Choice field
func (result *Result) CombineChoices() []string {
	result.IncorrectAnswer = append(result.IncorrectAnswer, result.CorrectAnswer)
	return result.IncorrectAnswer
}

// ShuffleChoices will randomize items in choices in Result struct
func (result *Result) ShuffleChoices() []string {
	//Shuffle choices
	for l := len(result.IncorrectAnswer) - 1; l > 0; l-- {
		r := rand.Intn(l + 1)
		result.IncorrectAnswer[l], result.IncorrectAnswer[r] = result.IncorrectAnswer[r], result.IncorrectAnswer[l]
	}
	return result.IncorrectAnswer
}

// OrderBoolean orders the choices for true/false questions
func (result *Result) OrderBoolean() []string {
	if result.IncorrectAnswer[0] == "False" {
		// Switch answers
		result.IncorrectAnswer[0], result.IncorrectAnswer[1] = result.IncorrectAnswer[1], result.IncorrectAnswer[0]
	}

	return result.IncorrectAnswer
}
