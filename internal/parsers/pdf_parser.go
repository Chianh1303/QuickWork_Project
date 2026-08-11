package parsers

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ParsedCVData struct {
	FullText     string
	Projects     []string
	Skills       []string
	Universities []string
}

type PDFParser interface {
	GetCVFileBytesAndInfo(cvURL string) (fileBytes []byte, filename string, size int64, exists bool)
	ParsePDFTextContent(pdfBytes []byte) ParsedCVData
	RemoveDuplicates(elements []string) []string
}

type pdfParser struct{}

func NewPDFParser() PDFParser {
	return &pdfParser{}
}

func (p *pdfParser) GetCVFileBytesAndInfo(cvURL string) ([]byte, string, int64, bool) {
	if strings.TrimSpace(cvURL) == "" {
		return nil, "", 0, false
	}
	const marker = "/uploads/cvs/"
	parts := strings.SplitN(cvURL, marker, 2)
	if len(parts) != 2 {
		return nil, "", 0, false
	}
	filename := filepath.Base(strings.TrimSpace(parts[1]))
	if filename == "." || filename == string(filepath.Separator) || filename == "" {
		return nil, "", 0, false
	}
	filePath := filepath.Join("uploads", "cvs", filename)
	info, err := os.Stat(filePath)
	if err != nil || info.IsDir() {
		return nil, filename, 0, false
	}
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, filename, info.Size(), false
	}
	return fileBytes, filename, info.Size(), true
}

func (p *pdfParser) ParsePDFTextContent(pdfBytes []byte) ParsedCVData {
	if len(pdfBytes) == 0 {
		return ParsedCVData{}
	}
	content := string(pdfBytes)
	re := regexp.MustCompile(`\(([^\(\)\\\r\n]{2,500})\)`)
	matches := re.FindAllStringSubmatch(content, -1)

	var textLines []string
	var detectedProjects []string
	var detectedSkills []string
	var detectedUniv []string

	techKeywords := []string{"java", "spring boot", "spring mvc", "mysql", "javascript", "thymeleaf", "html", "css", "vue", "react", "node", "python", "backend", "frontend", "fullstack", "go", "golang", "postgresql", "mongodb", "docker", "git", "github", "rest api", "api"}

	for _, match := range matches {
		if len(match) <= 1 {
			continue
		}
		txt := strings.TrimSpace(match[1])
		if len(txt) < 3 || strings.HasPrefix(txt, "http") || strings.Contains(txt, "Adobe") || strings.Contains(txt, "Identity") || strings.Contains(txt, "Font") || strings.Contains(txt, "xml") || txt == "EDUCATION" || txt == "SKILLS" || txt == "WORK EXPERIENCE" || txt == "SOFT SKILLS" {
			continue
		}
		textLines = append(textLines, txt)
		lower := strings.ToLower(txt)

		for _, keyword := range techKeywords {
			if strings.Contains(lower, keyword) {
				if len(txt) <= 60 && !strings.Contains(lower, "student") && !strings.Contains(lower, "experience") && !strings.Contains(lower, "majoring") && !strings.Contains(lower, "building") {
					detectedSkills = append(detectedSkills, txt)
				}
				break
			}
		}

		isProject := strings.Contains(lower, "website") || strings.Contains(lower, "platform") || strings.Contains(lower, "app") || strings.Contains(lower, "system") || strings.Contains(lower, "project")
		if isProject && !strings.Contains(lower, "majoring") {
			if len(txt) <= 100 && !strings.Contains(lower, "experience") {
				detectedProjects = append(detectedProjects, txt)
			}
		}

		if strings.Contains(lower, "thanhdo") || strings.Contains(lower, "university") || strings.Contains(lower, "information technology") || strings.Contains(lower, "bachelor") {
			if len(txt) <= 80 {
				detectedUniv = append(detectedUniv, txt)
			}
		}
	}

	fullText := strings.Join(textLines, "\n")
	return ParsedCVData{
		FullText:     fullText,
		Projects:     p.RemoveDuplicates(detectedProjects),
		Skills:       p.RemoveDuplicates(detectedSkills),
		Universities: p.RemoveDuplicates(detectedUniv),
	}
}

func (p *pdfParser) RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for _, element := range elements {
		clean := strings.TrimSpace(element)
		if clean == "" || encountered[clean] {
			continue
		}
		encountered[clean] = true
		result = append(result, clean)
	}
	return result
}
