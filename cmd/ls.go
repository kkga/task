package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/kkga/togo/txt"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls [query...]",
	Short:   "List todos",
	Example: "togo ls\ntogo ls +myproject\ntogo ls myquery",
	Aliases: []string{"l, list"},
	Run: func(cmd *cobra.Command, args []string) {

		m, err := txt.TodoMap(TodoFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		todos := make(map[int]txt.Todo)

		for k, todo := range m {
			if len(args) > 0 {
				for _, arg := range args {
					_, exists := todos[k]
					matches := strings.Contains(todo.Subject, arg)
					if !exists && matches {
						todos[k] = todo
					}
				}
			} else {
				todos[k] = todo
			}
		}

		// iteration over map happens in random order, so we store the order
		// in a separate slice
		var keys []int
		for k := range todos {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			PrintTodo(k, todos[k])
		}

		fmt.Println("-----")
		fmt.Printf("%d/%d todos shown (%s)\n", len(todos), len(m), TodoFile)
	},
}

// PrintTodo colorizes and outputs the given Todo
func PrintTodo(key int, todo txt.Todo) {
	var result string

	if key == 0 {
		result += "-" + " "
	} else {
		result += fmt.Sprintf("%-2d", key)
	}

	if todo.Done {
		color := color.New(color.FgGreen, color.Bold).SprintFunc()
		result += color("[x]") + " "
	} else {
		color := color.New(color.Reset).SprintFunc()
		result += color("[ ]") + " "
	}

	if todo.Priority != "" {
		color := color.New(color.FgHiRed, color.Bold).SprintFunc()
		result += color("("+todo.Priority+")") + " "
	}

	if !todo.CompletionDate.IsZero() {
		color := color.New(color.Reset).SprintFunc()
		result += color(todo.CompletionDate.Format("2006-01-02")) + " "
	}

	if !todo.CreationDate.IsZero() {
		color := color.New(color.Reset).SprintFunc()
		result += color(todo.CreationDate.Format("2006-01-02")) + " "
	}

	if todo.Subject != "" {
		color := color.New(color.Reset).SprintFunc()
		result += color(todo.Subject)
	}

	if len(todo.Projects) > 0 {
		color := color.New(color.FgCyan).SprintFunc()
		for _, p := range todo.Projects {
			result = strings.ReplaceAll(result, p, color(p))
		}
	}

	if len(todo.Contexts) > 0 {
		color := color.New(color.FgMagenta).SprintFunc()
		for _, c := range todo.Contexts {
			result = strings.ReplaceAll(result, c, color(c))
		}
	}

	fmt.Println(result)
}

func init() {
	rootCmd.AddCommand(lsCmd)
	// lsCmd.Flags().BoolP("done", "d", false, "List done tasks from done.txt")
}
