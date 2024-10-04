package main

import (
	"bufio"
	"os"
	"testing"
)

func TestEmptyLines(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\n",
		IndentSize:     4,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/empty-lines.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/empty-lines.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/empty-lines.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestEmptyLines failed")
	}
}

func TestSpaceIndentation(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\n",
		IndentSize:     2,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/space-indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/space-indentation.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestSpaceIndentation failed")
	}
}

func TestSpaceIndentation2(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\n",
		IndentSize:     16,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/space-indentation-2.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/space-indentation-2.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestSpaceIndentation2 failed")
	}
}

func TestSpaceIndentation3(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\r\n",
		IndentSize:     16,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/space-indentation-3.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/space-indentation-3.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestSpaceIndentation3 failed")
	}
}

func TestTabIndentation(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\n",
		IndentSize:     2,
		TrimEmptyLines: true,
		UseTabs:        true,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/tab-indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/tab-indentation.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestTabIndentation failed")
	}
}

func TestTabIndentation2(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\n",
		IndentSize:     2,
		TrimEmptyLines: true,
		UseTabs:        true,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/indentation.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/tab-indentation-2.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/tab-indentation-2.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestTabIndentation2 failed")
	}
}

func TestCommonTest(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\r\n",
		IndentSize:     16,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/commontest.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/commontest.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/commontest.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestCommonTest failed")
	}
}

func TestCommonTest2(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\r\n",
		IndentSize:     16,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/commontest2.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/commontest2.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/commontest2.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestCommonTest2 failed")
	}
}

func TestCommonTest3(t *testing.T) {
	options := FormatterOptions{
		EndOfLines:     "\r\n",
		IndentSize:     16,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/commontest3.nsi")
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	formattedContent, err := format(scanner)
	if err != nil {
		t.Error(err)
	}

	contentExpected, err := os.ReadFile("./expected/commontest3.nsi")
	if err != nil {
		t.Error(err)
	}

	// err = os.WriteFile("./expected/commontest3.nsi", []byte(formattedContent), 0644)
	// if err != nil {
	// 	t.Error(err)
	// }

	if formattedContent != string(contentExpected) {
		t.Errorf("TestCommonTest3 failed")
	}
}

// go test -bench=BenchmarkIndentation -benchtime=5s
func BenchmarkIndentation(b *testing.B) {
	options := FormatterOptions{
		EndOfLines:     "\n",
		IndentSize:     2,
		TrimEmptyLines: true,
		UseTabs:        false,
	}
	format := createFormatter(options)

	file, err := os.Open("./fixtures/indentation.nsi")
	if err != nil {
		b.Error(err)
	}

	defer file.Close()

	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(file)
		_, err = format(scanner)
		if err != nil {
			b.Error(err)
		}
	}
}
