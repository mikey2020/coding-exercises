// read file input 
package main
import (
    "fmt"
    "os"
    "log"
    "bufio"
    "strconv"
)

func main() {
    var s []int
    var count int
    f, err := os.Open("day1_input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        i, _ := strconv.Atoi(scanner.Text())
        s = append(s,i)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for i := 0; i < len(s); i++ {
        if i == 0 {
            continue
        }
        if s[i] - s[i-1] > 0 {
            count += 1
        }
	}
    fmt.Println(count)
}

