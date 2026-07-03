package main

import (
	"bufio"
	"fmt"
	packageVersion "nsis-formatter/internal/version"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	rules = struct {
		beforeIndenters          []string
		beforeDedenters          []string
		afterIndenters           []string
		afterDedenters           []string
		standardSyntaxVocabulary []string
		syntaxValidations        []syntaxValidation
	}{
		beforeIndenters: []string{},
		beforeDedenters: []string{
			"${EndSelect}", "${EndSelect}", "${EndSwitch}", "${EndSwitch}", "SectionEnd", "FunctionEnd", "${AndIf}", "${EndIf}", "${OrIf}", "${ElseIf}", "${AndIfNot}", "${ElseIfNot}", "${OrIfNot}", "${ElseUnless}", "${Else}", "${Next}", "${Case}", "${Default}", "${CaseElse}", "!endif", "!else", "!macroend", "PageExEnd", "SectionGroupEnd", "${AndUnless}", "${OrUnless}", "${EndWhile}", "${Loop}", "${LoopWhile}", "${LoopUntil}", "${EndUnless}", "!elseif",
		},
		afterIndenters: []string{
			"!if", "!ifdef", "!ifmacrodef", "!ifmacrondef", "!ifndef", "!macro", "${Do}", "${DoUntil}", "${DoWhile}", "${For}", "${ForEach}", "${If}", "${IfNot}", "${Unless}", "Function", "PageEx", "Section", "SectionGroup", "${While}", "${AndIf}", "${OrIf}", "${ElseIf}", "${AndIfNot}", "${ElseIfNot}", "${OrIfNot}", "${ElseUnless}", "${Else}", "${Case}", "${Switch}", "${Switch}", "${Default}", "!else", "${Select}", "${Select}", "${CaseElse}", "${AndUnless}", "${OrUnless}", "!elseif",
		},
		afterDedenters:           []string{},
		standardSyntaxVocabulary: []string{},
		syntaxValidations: []syntaxValidation{
			{
				"Goto", 2,
			},
			{
				"LangString", 4,
			},
			{
				"Unicode", 2,
			},
			{
				"!define", 3,
			},
			{
				"VIProductVersion", 2,
			},
			{
				"VIAddVersionKey", 3,
			},
			{
				"!include", 2,
			},
			{
				"!insertmacro", 2,
			},
			{
				"LicenseLangString", 3,
			},
			{
				"Call", 2,
			},
			{
				"SetOutPath", 2,
			},
			{
				"SetOverwrite", 2,
			},
			{
				"File", 2,
			},
			{
				"CreateDirectory", 2,
			},
			{
				"CreateShortCut", 2,
			},
			{
				"DetailPrint", 2,
			},
		},
	}
)

const defaultIndentation = 2

type formatterOptions struct {
	EndOfLines     string
	IndentSize     int
	TrimEmptyLines bool
	UseTabs        bool
}

type syntaxValidation struct {
	keyword               string
	defaultParameterCount int
}

func createFormatter(options formatterOptions) func(scanner *bufio.Scanner) (string, error) {
	mergedOptions := formatterOptions{
		EndOfLines:     detectPlatformEOL(),
		IndentSize:     defaultIndentation,
		TrimEmptyLines: true,
		UseTabs:        true,
	}

	if options.IndentSize > 0 {
		mergedOptions.IndentSize = options.IndentSize
	}
	if options.EndOfLines != "" {
		mergedOptions.EndOfLines = options.EndOfLines
	}
	mergedOptions.UseTabs = options.UseTabs
	mergedOptions.TrimEmptyLines = options.TrimEmptyLines

	// Define formatting function
	return func(scanner *bufio.Scanner) (string, error) {
		indentationLevel := 0
		var formattedLines []string

		// Flag to track consecutive empty lines
		previousLineEmpty := false

		for scanner.Scan() {
			line := scanner.Text() // read one line at a time
			trimmedLine := strings.TrimSpace(line)

			// If current line is empty
			if trimmedLine == "" {
				if !previousLineEmpty {
					// Keep a single empty line if it's not a consecutive empty line
					formattedLines = append(formattedLines, "")
				}
				previousLineEmpty = true
				continue
			}

			// Process the line as per the indentation rules
			trimmedLineDatas := strings.Split(trimmedLine, " ")
			if len(trimmedLineDatas[0]) > 1 && strings.HasPrefix(trimmedLineDatas[0], ";") && !strings.HasPrefix(trimmedLineDatas[0], "; ") {
				trimmedLineDatas[0] = "; " + trimmedLineDatas[0][1:]
			}
			keyword := strings.TrimSpace(trimmedLineDatas[0])

			if isGotoLine(trimmedLine) {
				if !previousLineEmpty {
					// Keep a single empty line if it's not a consecutive empty line
					formattedLines = append(formattedLines, "")
				}
				formattedLines = append(formattedLines, formatLineForGoto(trimmedLine, indentationLevel, mergedOptions))
				// Reset flag when we hit a non-empty line
				previousLineEmpty = false
				continue
			}

			currentIndentation := indentationLevel
			if counter, formattedKeyWord := checkKeyPass(rules.beforeIndenters, keyword); counter > 0 {
				trimmedLineDatas[0] = formattedKeyWord
				indentationLevel += counter
				currentIndentation = indentationLevel
			}
			if counter, formattedKeyWord := checkKeyPass(rules.beforeDedenters, keyword); counter > 0 {
				trimmedLineDatas[0] = formattedKeyWord
				indentationLevel -= counter
				if indentationLevel < 0 {
					indentationLevel = 0
				}
				currentIndentation = indentationLevel
			}
			if counter, formattedKeyWord := checkKeyPass(rules.afterIndenters, keyword); counter > 0 {
				trimmedLineDatas[0] = formattedKeyWord
				currentIndentation = indentationLevel
				indentationLevel += counter
			}
			if counter, formattedKeyWord := checkKeyPass(rules.afterDedenters, keyword); counter > 0 {
				trimmedLineDatas[0] = formattedKeyWord
				currentIndentation = indentationLevel
				indentationLevel -= counter
				if indentationLevel < 0 {
					indentationLevel = 0
				}
			}
			if counter, formattedKeyWord := checkKeyPass(rules.standardSyntaxVocabulary, keyword); counter > 0 {
				trimmedLineDatas[0] = formattedKeyWord
			}

			formattedLines = append(formattedLines, formatLine(lineSyntaxValidation(trimmedLineDatas), currentIndentation, mergedOptions))

			// Reset flag when we hit a non-empty line
			previousLineEmpty = false
		}

		formattedLines = append(formattedLines, "")

		return strings.Join(formattedLines, mergedOptions.EndOfLines), nil
	}
}

func lineSyntaxValidation(trimmedLineDatas []string) string {
	correctedLineData := make([]string, 0, len(trimmedLineDatas))
	isNeedSyntaxValidation := false
	parameterCount := 0
	for _, rule := range rules.syntaxValidations {
		if strings.EqualFold(trimmedLineDatas[0], rule.keyword) {
			isNeedSyntaxValidation = true
			parameterCount = rule.defaultParameterCount
			trimmedLineDatas[0] = rule.keyword
			break
		}
	}

	if isNeedSyntaxValidation {
		currentIndex := 0
		for _, data := range trimmedLineDatas {
			if currentIndex < parameterCount {
				if len(data) > 0 {
					correctedLineData = append(correctedLineData, data)
					currentIndex++
				}
			} else {
				correctedLineData = append(correctedLineData, data)
			}
		}
	} else {
		correctedLineData = append(correctedLineData, trimmedLineDatas...)
	}

	return strings.Join(correctedLineData, " ")
}

func checkKeyPass(ruleData []string, keyword string) (int, string) {
	counter := 0
	currentRule := ""
	for _, rule := range ruleData {
		if strings.EqualFold(keyword, rule) {
			currentRule = rule
			counter++
		}
	}

	return counter, currentRule
}

func isGotoLine(line string) bool {
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*:$`)
	return re.MatchString(line)
}

func formatLine(line string, level int, options formatterOptions) string {
	if len(line) == 0 {
		return ""
	}

	if level < 0 {
		level = 0
	}

	indent := strings.Repeat("\t", options.IndentSize*level)
	if !options.UseTabs {
		indent = strings.Repeat(" ", options.IndentSize*level)
	}
	return fmt.Sprintf("%s%s", indent, strings.TrimSpace(line))
}

func formatLineForGoto(line string, level int, options formatterOptions) string {
	if len(line) == 0 {
		return ""
	}

	var indent string
	if options.UseTabs {
		indent = strings.Repeat("\t", options.IndentSize*level)
	} else if options.IndentSize >= 4 && level > 0 {
		indent = strings.Repeat(" ", options.IndentSize*level-2)
	} else {
		indent = strings.Repeat(" ", options.IndentSize*level)
	}

	return fmt.Sprintf("%s%s", indent, strings.TrimSpace(line))
}

func detectPlatformEOL() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func main() {
	app := &cli.App{
		Name:    "nsis-formatter",
		Usage:   "CLI tool to format NSIS scripts",
		Version: packageVersion.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "eol",
				Usage: "control how line-breaks are represented (crlf, lf)",
			},
			&cli.IntFlag{
				Name:  "indent-size",
				Usage: "number of units per indentation level",
				Value: 2,
			},
			&cli.BoolFlag{
				Name:  "use-spaces",
				Usage: "indent with spaces instead of tabs",
			},
			&cli.BoolFlag{
				Name:  "trim",
				Usage: "trim empty lines",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "write",
				Usage: "edit files in-place",
			},
			&cli.BoolFlag{
				Name:  "quiet",
				Usage: "suppress output",
			},
		},
		Action: func(c *cli.Context) error {
			eol := c.String("eol")
			if len(eol) > 0 && eol != "crlf" && eol != "lf" {
				return fmt.Errorf("invalid value for --eol: %s. Valid options are 'crlf' or 'lf'", eol)
			}

			file := c.Args().First()
			options := formatterOptions{
				IndentSize:     c.Int("indent-size"),
				TrimEmptyLines: c.Bool("trim"),
				UseTabs:        !c.Bool("use-spaces"),
			}

			switch eol {
			case "crlf":
				options.EndOfLines = "\r\n"
			case "lf":
				options.EndOfLines = "\n"
			}

			format := createFormatter(options)
			fmt.Println("Processing file:", file)

			fileHandler, err := os.Open(file)
			if err != nil {
				fmt.Printf("failed to open file: %s, err=%+v\n", file, err)
				os.Exit(1)
			}

			scanner := bufio.NewScanner(fileHandler)
			formattedContent, err := format(scanner)
			if err != nil {
				_ = fileHandler.Close()

				fmt.Printf("failed to format file: %s, err=%+v\n", file, err)
				os.Exit(2)
			}

			if err := fileHandler.Close(); err != nil {
				fmt.Printf("failed to close file: %s, err=%+v\n", file, err)
				os.Exit(2)
			}

			if c.Bool("write") {
				err = os.WriteFile(file, []byte(formattedContent), 0644)
				if err != nil {
					fmt.Printf("failed to rewrite file: %s, err=%+v\n", file, err)
					os.Exit(2)
				}
				if !c.Bool("quiet") {
					fmt.Println(formattedContent)
				}
			} else {
				if !c.Bool("quiet") {
					fmt.Println(formattedContent)
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
