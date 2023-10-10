package services

import (
	"fmt"
	"math/rand"
	"strings"
	// eliza "github.com/yosuke7040/go-todo-grpc/services/eliza"
)

type ElizaServiceInterface interface {
	Reply(input string) (string, bool)
	GetIntroResponses(name string) []string
}

type ElizaService struct{}

func NewElizaService() *ElizaService {
	return &ElizaService{}
}

func preprocess(input string) string {
	return strings.Trim(strings.ToLower(strings.TrimSpace(input)), `.!?'"`)
}

func randomElementFrom(list []string) string {
	return list[rand.Intn(len(list))]
}

// reflect flips a few words in an input fragment (such as "I" -> "you").
func reflect(fragment string) string {
	words := strings.Fields(fragment)
	for i, word := range words {
		if reflectedWord, ok := reflectedWords[word]; ok {
			words[i] = reflectedWord
		}
	}
	return strings.Join(words, " ")
}

func lookupResponse(input string) string {
	for re, responses := range requestInputRegexToResponseOptions {
		matches := re.FindStringSubmatch(input)
		if len(matches) < 1 {
			continue
		}
		response := randomElementFrom(responses)
		// If the response has an entry point, reflect the input phrase (so "I"
		// becomes "you").
		if !strings.Contains(response, "%s") {
			return response
		}
		if len(matches) > 1 {
			fragment := reflect(matches[1])
			response = fmt.Sprintf(response, fragment)
			return response
		}
	}
	return randomElementFrom(defaultResponses)
}

// Reply responds to a statement as a pyschotherapist might.
func (s *ElizaService) Reply(input string) (string, bool) {
	input = preprocess(input)
	if _, ok := goodbyeInputSet[input]; ok {
		return randomElementFrom(goodbyeResponses), true
	}
	return lookupResponse(input), false
}

// GetIntroResponses returns a collection of introductory responses tailored to the given name.
func (s *ElizaService) GetIntroResponses(name string) []string {
	intros := make([]string, 0, len(introResponses)+2)
	for _, n := range introResponses {
		intros = append(intros, fmt.Sprintf(n, name))
	}

	intros = append(intros, randomElementFrom(elizaFacts))
	intros = append(intros, "How are you feeling today?")
	return intros
}
