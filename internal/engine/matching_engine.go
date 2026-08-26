package engine

import (
	"math"
	"strings"
)

type JobRequirementSet struct {
	TotalRequirements []string
	CoreRequirements  []string
}

type MatchResult struct {
	MatchScore     int
	MatchingSkills []string
	MissingSkills  []string
}

type MatchingEngine interface {
	NormalizeSkill(skill string) string
	NormalizeSkills(skills []string) []string
	ExtractJobRequirements(title, description, category string) JobRequirementSet
	Match(studentSkills []string, reqSet JobRequirementSet) MatchResult
}

type matchingEngine struct {
	synonyms map[string]string
}

func NewMatchingEngine() MatchingEngine {
	synonyms := map[string]string{
		"go":             "go",
		"golang":         "go",
		"go programming": "go",

		"js":         "javascript",
		"javascript": "javascript",
		"ecmascript": "javascript",

		"ts":         "typescript",
		"typescript": "typescript",

		"java": "java",

		"py":      "python",
		"python":  "python",
		"python3": "python",

		"cpp": "c++",
		"c++": "c++",

		"c#":     "c#",
		"csharp": "c#",

		"php":    "php",
		"ruby":   "ruby",
		"swift":  "swift",
		"kotlin": "kotlin",

		"html":  "html",
		"html5": "html",

		"css":  "css",
		"css3": "css",

		"react":    "react",
		"reactjs":  "react",
		"react.js": "react",

		"vue":    "vue",
		"vuejs":  "vue",
		"vue.js": "vue",

		"angular":   "angular",
		"angularjs": "angular",

		"node":    "node.js",
		"nodejs":  "node.js",
		"node.js": "node.js",

		"express":   "express",
		"expressjs": "express",

		"spring":      "spring",
		"spring boot": "spring",
		"spring mvc":  "spring",

		"mysql":      "mysql",
		"postgre":    "postgresql",
		"postgres":   "postgresql",
		"postgresql": "postgresql",
		"mongodb":    "mongodb",
		"mongo":      "mongodb",
		"redis":      "redis",

		"docker":     "docker",
		"k8s":        "kubernetes",
		"kubernetes": "kubernetes",

		// Conservative taxonomy: git, github, gitlab remain distinct
		"git":    "git",
		"github": "github",
		"gitlab": "gitlab",

		"rest":        "rest api",
		"restful":     "rest api",
		"rest api":    "rest api",
		"restful api": "rest api",
		"graphql":     "graphql",

		"tailwind":    "tailwind",
		"tailwindcss": "tailwind",
		"bootstrap":   "bootstrap",

		"figma":     "figma",
		"photoshop": "photoshop",
		"excel":     "excel",
		"word":      "word",
		"marketing": "marketing",
		"seo":       "seo",
		"sales":     "sales",
	}

	return &matchingEngine{synonyms: synonyms}
}

func (e *matchingEngine) NormalizeSkill(raw string) string {
	clean := strings.ToLower(strings.TrimSpace(raw))
	if clean == "" {
		return ""
	}
	if canonical, found := e.synonyms[clean]; found {
		return canonical
	}
	return clean
}

func (e *matchingEngine) NormalizeSkills(skills []string) []string {
	encountered := make(map[string]bool)
	var result []string
	for _, skill := range skills {
		norm := e.NormalizeSkill(skill)
		if norm == "" {
			continue
		}
		if !encountered[norm] {
			encountered[norm] = true
			result = append(result, norm)
		}
	}
	return result
}

func (e *matchingEngine) ExtractJobRequirements(title, description, category string) JobRequirementSet {
	titleLower := strings.ToLower(title)
	descLower := strings.ToLower(description)
	catLower := strings.ToLower(category)

	coreMap := make(map[string]bool)
	totalMap := make(map[string]bool)

	for rawKey, canonical := range e.synonyms {
		// Match in Title or Category -> Core Requirement
		if strings.Contains(titleLower, rawKey) || strings.Contains(catLower, rawKey) {
			coreMap[canonical] = true
			totalMap[canonical] = true
		} else if strings.Contains(descLower, rawKey) {
			totalMap[canonical] = true
		}
	}

	var coreReqs []string
	for k := range coreMap {
		coreReqs = append(coreReqs, k)
	}

	var totalReqs []string
	for k := range totalMap {
		totalReqs = append(totalReqs, k)
	}

	return JobRequirementSet{
		TotalRequirements: totalReqs,
		CoreRequirements:  coreReqs,
	}
}

func (e *matchingEngine) Match(studentSkills []string, reqSet JobRequirementSet) MatchResult {
	normStudentSkills := e.NormalizeSkills(studentSkills)
	studentMap := make(map[string]bool)
	for _, s := range normStudentSkills {
		studentMap[s] = true
	}

	totalReqs := e.NormalizeSkills(reqSet.TotalRequirements)
	coreReqs := e.NormalizeSkills(reqSet.CoreRequirements)

	// Edge Case: 0 requirements -> score 50
	if len(totalReqs) == 0 {
		return MatchResult{
			MatchScore:     50,
			MatchingSkills: []string{},
			MissingSkills:  []string{},
		}
	}

	var matchingSkills []string
	var missingSkills []string
	matchingMap := make(map[string]bool)

	for _, req := range totalReqs {
		if studentMap[req] {
			matchingSkills = append(matchingSkills, req)
			matchingMap[req] = true
		} else {
			missingSkills = append(missingSkills, req)
		}
	}

	// Edge Case: 0 matching skills -> score 0
	if len(matchingSkills) == 0 {
		return MatchResult{
			MatchScore:     0,
			MatchingSkills: []string{},
			MissingSkills:  missingSkills,
		}
	}

	// Full match -> score 100
	if len(matchingSkills) == len(totalReqs) {
		return MatchResult{
			MatchScore:     100,
			MatchingSkills: matchingSkills,
			MissingSkills:  []string{},
		}
	}

	overallScore := float64(len(matchingSkills)) / float64(len(totalReqs))

	var finalScore float64
	if len(coreReqs) > 0 {
		var matchingCoreCount int
		for _, core := range coreReqs {
			if studentMap[core] {
				matchingCoreCount++
			}
		}
		coreScore := float64(matchingCoreCount) / float64(len(coreReqs))
		finalScore = (0.70*overallScore + 0.30*coreScore) * 100
	} else {
		// Adjustment 2: totalCoreRequirements == 0 -> no division by zero
		finalScore = overallScore * 100
	}

	scoreInt := int(math.Round(finalScore))
	if scoreInt < 0 {
		scoreInt = 0
	}
	if scoreInt > 100 {
		scoreInt = 100
	}

	return MatchResult{
		MatchScore:     scoreInt,
		MatchingSkills: matchingSkills,
		MissingSkills:  missingSkills,
	}
}
