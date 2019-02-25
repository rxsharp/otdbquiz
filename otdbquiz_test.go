package otdbquiz

import (
	"log"
)

func ExampleRaw() {
    config := DefaultClient("https://opentdb.com/api.php?amount=1")
	quiz, _ := Raw(config)
	log.Println(quiz)
    // Output: &{[{History multiple medium Where and when was the first cardboard box made for industrial use? England, 1817 [England, 1917 United States, 1917 United States, 1817 England, 1817]}]}
}

func ExampleStandard() {
    config := DefaultClient("https://opentdb.com/api.php?amount=1")
	quiz, _ := Standard(config)
	log.Println(quiz)
    // Output: &{[{Mythology multiple hard Talos, the mythical giant bronze man, was the protector of which island? Crete [Sardinia Sicily Cyprus]}]}
}