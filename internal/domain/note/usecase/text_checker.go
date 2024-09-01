package usecase

import "fmt"

type TextValidatorUseCase struct {
	textChecker TextChecker
}

type TextChecker interface {
	CheckText(inputText string) ([]string, error)
}

func (t *TextValidatorUseCase) ValidateText(inputText string) (string, error) {
	verifiedWords, err := t.textChecker.CheckText(inputText)
	if err != nil {
		return "", fmt.Errorf("usecase - TextValidatorUseCase - ValidateText: %w", err)
	}

	var result string
	for i, word := range verifiedWords {
		result += word
		if i < len(verifiedWords)-1 {
			result += " "
		}
	}

	return result, nil
}

func NewTextValidator(textChecker TextChecker) *TextValidatorUseCase {
	return &TextValidatorUseCase{textChecker: textChecker}
}
