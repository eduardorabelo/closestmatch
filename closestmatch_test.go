package closestmatch

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/schollz/closestmatch/test"
)

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Open(test.WordsToTest, []int{3})
	}
}

func BenchmarkSplitOne(b *testing.B) {
	cm := Open(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.splitWord(searchWord)
	}
}

func BenchmarkClosestOne(b *testing.B) {
	cm := Open(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func BenchmarkClosest3(b *testing.B) {
	cm := Open(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.ClosestN(searchWord, 3)
	}
}

func BenchmarkClosest30(b *testing.B) {
	cm := Open(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.ClosestN(searchWord, 30)
	}
}

func BenchmarkLargeFile(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := Open(wordsToTest, []int{3})
	searchWord := "island of a thod mirrors"
	fmt.Println(cm.Closest(searchWord))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func ExampleMatching() {
	cm := Open(test.WordsToTest, []int{1, 2, 3})
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'don quixote by miguel de cervantes saavedra'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'charles dickens' matched 'hard times by charles dickens'
	// 'william shakespeare' matched 'the tragedy of romeo and juliet by william shakespeare'
	// 'war by hg wells' matched 'the war of the worlds by h. g. wells'
}

func ExampleMatchingN() {
	cm := Open(test.WordsToTest, []int{1, 2, 3})
	fmt.Println(cm.ClosestN("war by hg wells", 3))
	// Output:
	// [the war of the worlds by h. g. wells the time machine by h. g. wells the iliad by homer]
}

func TestAccuray(t *testing.T) {
	cm := Open(test.WordsToTest, []int{3})
	fmt.Println(cm.Accuracy())
	// Output:
	// [the war of the worlds by h. g. wells the time machine by h. g. wells the iliad by homer]
}
