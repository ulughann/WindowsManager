package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"
	"syscall"
)

func levenshteinDistance(str1 string, str2 string) int {
	n, m := len(str1)+1, len(str2)+1
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		d[i][0] = i
	}
	for j := 0; j < m; j++ {
		d[0][j] = j
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if str1[i-1] == str2[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+1)
			}
		}
	}

	return d[n-1][m-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func similaritySort(a []string, text string, similarityMetric func(string, string) int) {
	sort.Slice(a, func(i, j int) bool {
		return similarityMetric(a[i], text) < similarityMetric(a[j], text) // Sort in ascending order of distance
	})
}

func (a *App) Search(text string) string {
	cmd := exec.Command("powershell", "-Command", "choco find "+text, "--id-only", "--limit-output", "--not-broken")

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Chocolatey search script:", err)
		fmt.Println("Output:", string(output))
		return "this resulted in an error, truly saddening."
	}

	results := strings.Split(string(output), "\n")
	similaritySort(results, text, levenshteinDistance)

	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	sortedOutput := strings.Join(results, "\n")

	return sortedOutput
}
