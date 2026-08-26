package engine

import (
	"testing"
)

func TestMatchingEngine_GoVsGolang(t *testing.T) {
	eng := NewMatchingEngine()
	s1 := eng.NormalizeSkill("Go")
	s2 := eng.NormalizeSkill("Golang")
	s3 := eng.NormalizeSkill("go programming")

	if s1 != "go" || s2 != "go" || s3 != "go" {
		t.Errorf("Expected all to normalize to 'go', got: %s, %s, %s", s1, s2, s3)
	}
}

func TestMatchingEngine_GitVsGitHubNotMatched(t *testing.T) {
	eng := NewMatchingEngine()
	sGit := eng.NormalizeSkill("Git")
	sGitHub := eng.NormalizeSkill("GitHub")
	sGitLab := eng.NormalizeSkill("GitLab")

	if sGit != "git" {
		t.Errorf("Expected 'git', got '%s'", sGit)
	}
	if sGitHub != "github" {
		t.Errorf("Expected 'github', got '%s'", sGitHub)
	}
	if sGitLab != "gitlab" {
		t.Errorf("Expected 'gitlab', got '%s'", sGitLab)
	}

	// Student has "git", Job requires "github"
	studentSkills := []string{"git"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"github"},
		CoreRequirements:  []string{"github"},
	}
	res := eng.Match(studentSkills, reqSet)
	if res.MatchScore != 0 {
		t.Errorf("Expected MatchScore 0 when student has git but job requires github, got %d", res.MatchScore)
	}
}

func TestMatchingEngine_DeduplicateSkills(t *testing.T) {
	eng := NewMatchingEngine()
	skills := []string{"Go", "golang", "go", "JAVA", "Java", "java"}
	normalized := eng.NormalizeSkills(skills)

	if len(normalized) != 2 {
		t.Errorf("Expected 2 unique skills ('go', 'java'), got %d: %v", len(normalized), normalized)
	}
}

func TestMatchingEngine_CaseInsensitive(t *testing.T) {
	eng := NewMatchingEngine()
	s1 := eng.NormalizeSkill("REACT")
	s2 := eng.NormalizeSkill("ReactJS")
	s3 := eng.NormalizeSkill("react")

	if s1 != "react" || s2 != "react" || s3 != "react" {
		t.Errorf("Expected 'react', got %s, %s, %s", s1, s2, s3)
	}
}

func TestMatchingEngine_ZeroRequirements(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{"go", "react"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{},
		CoreRequirements:  []string{},
	}
	res := eng.Match(studentSkills, reqSet)
	if res.MatchScore != 50 {
		t.Errorf("Expected score 50 for zero requirements, got %d", res.MatchScore)
	}
}

func TestMatchingEngine_ZeroStudentSkills(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"go", "mysql"},
		CoreRequirements:  []string{"go"},
	}
	res := eng.Match(studentSkills, reqSet)
	if res.MatchScore != 0 {
		t.Errorf("Expected score 0 for zero student skills, got %d", res.MatchScore)
	}
}

func TestMatchingEngine_FullMatch(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{"go", "mysql", "docker"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"go", "mysql"},
		CoreRequirements:  []string{"go"},
	}
	res := eng.Match(studentSkills, reqSet)
	if res.MatchScore != 100 {
		t.Errorf("Expected score 100 for full match, got %d", res.MatchScore)
	}
}

func TestMatchingEngine_PartialMatch(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{"go"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"go", "mysql"},
		CoreRequirements:  []string{},
	}
	res := eng.Match(studentSkills, reqSet)
	if res.MatchScore <= 0 || res.MatchScore >= 100 {
		t.Errorf("Expected score between 0 and 100 for partial match, got %d", res.MatchScore)
	}
}

func TestMatchingEngine_MissingSkills(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{"go"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"go", "mysql", "docker"},
		CoreRequirements:  []string{"go"},
	}
	res := eng.Match(studentSkills, reqSet)
	if len(res.MissingSkills) != 2 {
		t.Errorf("Expected 2 missing skills ('mysql', 'docker'), got %d", len(res.MissingSkills))
	}
}

func TestMatchingEngine_CoreRequirements(t *testing.T) {
	eng := NewMatchingEngine()
	// Student A has the core skill "go"
	resA := eng.Match([]string{"go"}, JobRequirementSet{
		TotalRequirements: []string{"go", "mysql"},
		CoreRequirements:  []string{"go"},
	})

	// Student B has the non-core skill "mysql"
	resB := eng.Match([]string{"mysql"}, JobRequirementSet{
		TotalRequirements: []string{"go", "mysql"},
		CoreRequirements:  []string{"go"},
	})

	if resA.MatchScore <= resB.MatchScore {
		t.Errorf("Expected Student with Core Skill (score %d) to rank higher than non-core (score %d)", resA.MatchScore, resB.MatchScore)
	}
}

func TestMatchingEngine_CoreRequirementsZeroNoDivisionByZero(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{"go"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"go", "mysql"},
		CoreRequirements:  []string{}, // 0 core requirements
	}
	res := eng.Match(studentSkills, reqSet)
	// overallScore = 1/2 = 0.5 -> 50%
	if res.MatchScore != 50 {
		t.Errorf("Expected score 50 when core requirements = 0, got %d", res.MatchScore)
	}
}

func TestMatchingEngine_ScoreClamped(t *testing.T) {
	eng := NewMatchingEngine()
	studentSkills := []string{"go", "mysql", "docker", "kubernetes", "react"}
	reqSet := JobRequirementSet{
		TotalRequirements: []string{"go"},
		CoreRequirements:  []string{"go"},
	}
	res := eng.Match(studentSkills, reqSet)
	if res.MatchScore < 0 || res.MatchScore > 100 {
		t.Errorf("Expected score to be clamped [0, 100], got %d", res.MatchScore)
	}
}
