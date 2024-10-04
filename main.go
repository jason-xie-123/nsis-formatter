package main

import (
	"bufio"
	"fmt"
	packageVersion "nsis-formatter/version"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	rules = struct {
		beforeIndenters []string
		beforeDedenters []string
		afterIndenters  []string
		afterDedenters  []string
		// specialIndenters []string
	}{
		beforeIndenters: []string{
			"${Select}",
		},
		beforeDedenters: []string{
			"${EndSelect}", "${EndSwitch}", "${EndSwitch}", "SectionEnd", "FunctionEnd", "${AndIf}", "${EndIf}", "${OrIf}", "${ElseIf}", "${AndIfNot}",
			"${ElseIfNot}", "${OrIfNot}", "${ElseUnless}", "${Else}", "${Next}", "${Case}", "${Default}",
			"!endif", "!else", "!macroend", "PageExEnd", "SectionGroupEnd",
		},
		afterIndenters: []string{
			"!if", "!ifdef", "!ifmacrodef", "!ifmacrondef", "!ifndef", "!macro",
			"${Case2}", "${Case3}", "${Case4}", "${Case5}",
			"${Do}", "${DoUntil}", "${DoWhile}", "${For}", "${ForEach}",
			"${If}", "${IfNot}", "${MementoSection}", "${MementoUnselectedSection}",
			"${Unless}", "Function", "PageEx", "Section",
			"SectionGroup", "${While}", "${AndIf}", "${OrIf}", "${ElseIf}", "${AndIfNot}", "${ElseIfNot}",
			"${OrIfNot}", "${ElseUnless}", "${Else}",
			"${Case}", "${Switch}", "${Switch}", "${Default}",
			"!else",
		},
		afterDedenters: []string{
			"!macroend",
			"${EndWhile}", "${Loop}", "${LoopUntil}", "${LoopWhile}",
			"${MementoSectionEnd}", "${EndUnless}",
			"!elseif",
			"${AndUnless}",
			"${OrUnless}", "${CaseElse}",
		},
		// specialIndenters: []string{},
	}
)

const defaultIndentation = 2

type FormatterOptions struct {
	EndOfLines     string
	IndentSize     int
	TrimEmptyLines bool
	UseTabs        bool
}

func createFormatter(options FormatterOptions) func(scanner *bufio.Scanner) (string, error) {
	mergedOptions := FormatterOptions{
		EndOfLines:     detectPlatformEOL(),
		IndentSize:     defaultIndentation,
		TrimEmptyLines: true,
		UseTabs:        true,
	}

	// Override mergedOptions with provided options
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
		// lines := strings.Split(fileContents, mergedOptions.EndOfLines)

		// Flag to track consecutive empty lines
		previousLineEmpty := false

		for scanner.Scan() {
			line := scanner.Text() // 每次读取一行
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
			keyword := strings.TrimSpace(strings.Split(trimmedLine, " ")[0])

			if isGotoLine(trimmedLine) {
				if !previousLineEmpty {
					// Keep a single empty line if it's not a consecutive empty line
					formattedLines = append(formattedLines, "")
					previousLineEmpty = true
				}
				formattedLines = append(formattedLines, formatLineForGoto(trimmedLine, indentationLevel, mergedOptions))
				// Reset flag when we hit a non-empty line
				previousLineEmpty = false
				continue
			}

			currentIndentation := indentationLevel
			if counter := checkKeyPass(rules.beforeIndenters, keyword); counter > 0 {
				indentationLevel += counter
				currentIndentation = indentationLevel
				// formattedLines = append(formattedLines, formatLine(trimmedLine, indentationLevel, mergedOptions))
			}
			if counter := checkKeyPass(rules.beforeDedenters, keyword); counter > 0 {
				indentationLevel -= counter
				if indentationLevel < 0 {
					indentationLevel = 0
				}
				currentIndentation = indentationLevel
				// formattedLines = append(formattedLines, formatLine(trimmedLine, indentationLevel, mergedOptions))
			}
			if counter := checkKeyPass(rules.afterIndenters, keyword); counter > 0 {
				currentIndentation = indentationLevel
				// formattedLines = append(formattedLines, formatLine(trimmedLine, indentationLevel, mergedOptions))
				indentationLevel += counter
			}
			if counter := checkKeyPass(rules.afterDedenters, keyword); counter > 0 {
				currentIndentation = indentationLevel
				// formattedLines = append(formattedLines, formatLine(trimmedLine, indentationLevel, mergedOptions))
				indentationLevel -= counter
				if indentationLevel < 0 {
					indentationLevel = 0
				}
			}

			formattedLines = append(formattedLines, formatLine(trimmedLine, currentIndentation, mergedOptions))

			// Reset flag when we hit a non-empty line
			previousLineEmpty = false
		}

		formattedLines = append(formattedLines, "")

		return strings.Join(formattedLines, mergedOptions.EndOfLines), nil
	}
}

func checkKeyPass(ruleData []string, keyword string) int {
	counter := 0
	for _, rule := range ruleData {
		if strings.EqualFold(keyword, rule) {
			counter++
		}
	}

	return counter
}

func isGotoLine(line string) bool {
	re := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*:$`)
	return re.MatchString(line)
}

func formatLine(line string, level int, options FormatterOptions) string {
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

func formatLineForGoto(line string, level int, options FormatterOptions) string {
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
			options := FormatterOptions{
				IndentSize:     c.Int("indent-size"),
				TrimEmptyLines: c.Bool("trim"),
				UseTabs:        !c.Bool("use-spaces"),
			}

			if eol == "crlf" {
				options.EndOfLines = "\r\n"
			} else if eol == "lf" {
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
				fileHandler.Close()

				fmt.Printf("failed to format file: %s, err=%+v\n", file, err)
				os.Exit(2)
			}

			fileHandler.Close()

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
