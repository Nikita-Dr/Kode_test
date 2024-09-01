package yadex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TextChecker struct {
	url string
}

type CheckerWord struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func (t *TextChecker) CheckText(inputText string) ([]string, error) {
	var convertedText string
	words := strings.Fields(inputText)
	for i, word := range words {
		convertedText = convertedText + word
		if i < len(words)-1 {
			convertedText += "+"
		}
	}

	urlWithText := t.url + convertedText

	res, err := http.Get(urlWithText)
	if err != nil {
		return nil, fmt.Errorf("usecase - NoteUseCase - GetNotes: %w", err)
	}
	data, err := io.ReadAll(res.Body)

	checkerWords := []CheckerWord{}
	//TODO обработка этой ошибки
	err = json.Unmarshal(data, &checkerWords)

	var verifiedWords []string
	for _, w := range checkerWords {
		verifiedWords = append(verifiedWords, w.S[0])
	}

	return verifiedWords, err
}

func NewTextChecker() *TextChecker {
	url := "https://speller.yandex.net/services/spellservice.json/checkText?text="
	return &TextChecker{url: url}
}
