package main

import ( 
	"flag"
	"github.com/umang345/todo-app"
	"os"
	"fmt"
	"io"
	"strings"
	"bufio"
	"errors"
)

const (
	todoFile = ".todos.json"
)

func main() {
	
	add := flag.Bool("add" , false, "add a new todo")
	complete := flag.Int("complete", 0, "mark todo as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list" , false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile);err!=nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task,err := getInput(os.Stdin, flag.Args()...)
		if err!=nil {
			fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
		}

		todos.Add(task)

		err1 := todos.Store(todoFile)
		if err1!= nil {
            fmt.Fprintln(os.Stderr, err1.Error())
            os.Exit(1)
        }
        
	case *complete > 0:
		err := todos.Complete(*complete)
		if err!= nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
		err1 := todos.Store(todoFile)
		if err1!= nil {
            fmt.Fprintln(os.Stderr, err1.Error())
            os.Exit(1)
        }
	case *del > 0:
		err := todos.Delete(*del)
		if err!= nil {
            fmt.Fprintln(os.Stderr, err.Error())
            os.Exit(1)
        }
		err1 := todos.Store(todoFile)
		if err1!= nil {
            fmt.Fprintln(os.Stderr, err1.Error())
            os.Exit(1)
        }
	case *list:
		todos.Print()
	default: 
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) >0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "",errors.New("empty todo is not allowed")
	} 

	return text, nil
}