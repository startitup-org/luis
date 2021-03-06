package luis_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/kkdai/luis"
)

var API_KEY string
var APPID string

const ()

func init() {
	API_KEY = os.Getenv("SUB_KEY")
	APPID = os.Getenv("APP_ID")
	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export SUB_KEY=12234 && export APP_ID=5678`")
	}
}

func getLuis(t *testing.T) *Luis {
	e := NewLuis(API_KEY, APPID)
	if e == nil {
		t.Error("Cannot connect to server")
	}
	return e
}

func TestIntentList(t *testing.T) {
	if API_KEY == "" {
		return
	}

	e := getLuis(t)

	res, err := e.IntentList()

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
	result := NewIntentListResponse(res)
	fmt.Println("Luis Intent Ret", result)
}

func TestActionChannels(t *testing.T) {
	if API_KEY == "" {
		return
	}
	e := getLuis(t)
	res, err := e.ActionChannels()

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
}

func TestPredict(t *testing.T) {
	if API_KEY == "" {
		return
	}
	e := getLuis(t)
	res, err := e.Predict("test string")

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
	fmt.Println("Get the best predict result:", GetBestScoreIntent(NewPredictResponse(res)))
}

func TestTrain(t *testing.T) {
	if API_KEY == "" {
		return
	}
	e := getLuis(t)
	res, err := e.Train()

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
}

func TestExample(t *testing.T) {
	if API_KEY == "" {
		return
	}
	e := getLuis(t)
	ex := ExampleJson{ExampleText: "test", SelectedIntentName: "test2"}
	res, err := e.AddLabel(ex)

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
}
