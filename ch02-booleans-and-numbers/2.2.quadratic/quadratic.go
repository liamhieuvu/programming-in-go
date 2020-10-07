// Copyright © 2010-12 Qtrac Ltd.
// 
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "math"
)

const (
    pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head>
<title>Quadratic Equation Solver</title>
<body><h3>Quadratic Equation Solver</h3>
<p>Solve equation of the form ax<sup>2</sup> + bx + c = 0</p>`
    form       = `<form action="/" method="POST">
<input type="text" name="a" size="10"> x<sup>2</sup> + 
<input type="text" name="b" size="10"> x + 
<input type="text" name="c" size="10"> = 0 &rarr;
<input type="submit" value="Calculate">
<br /><br />
</form>`
    pageBottom = `</body></html>`
    anError    = `<p class="error">%s</p>`
)

func main() {
    http.HandleFunc("/", homePage)
    if err := http.ListenAndServe(":9001", nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}

func homePage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm() // Must be called before writing response
    fmt.Fprint(writer, pageTop, form)
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        if numbers, message, ok := processRequest(request); ok {
            fmt.Fprint(writer, formatSolution(numbers))
        } else if message != "" {
            fmt.Fprintf(writer, anError, message)
        }
    }
    fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]complex128, string, bool) {
    var coef [3]float64
    emptyCnt := 3
    for i, key := range []string{"a", "b", "c"} {
        if slice, found := request.Form[key]; found && len(slice) > 0 {
            var err error
            if slice[0] == "" {
                slice[0] = "0"
            }
            if coef[i], err = strconv.ParseFloat(slice[0], 64); err != nil {
                return nil, "'" + slice[0] + "' is invalid", false
            }
            emptyCnt--
        }
    }
    if emptyCnt == 3 {
        return nil, "", false
    }

    a := coef[0]
    b := coef[1]
    c := coef[2]

    if a == 0 {
        if b == 0 {
            if c == 0 {
                return nil, "Countless solutions!", false
            } else {
                return []complex128{}, "", true
            }
        }
        return []complex128{complex(-c/b, 0)}, "", true
    }

    delta := b*b - 4*a*c
    if delta < 0 {
        sdelta := math.Sqrt(-delta)
        return []complex128{complex(-b/2/a, sdelta/2/a), complex(-b/2/a, -sdelta/2/a)}, "", true
    } else if delta == 0 {
        return []complex128{complex(-b/2/a, 0)}, "", true
    } else {
        sdelta := math.Sqrt(delta)
        return []complex128{complex((-b + sdelta)/2/a, 0), complex((-b - sdelta)/2/a, 0)}, "", true
    }
}

func formatSolution(numbers []complex128) string {
    if len(numbers) == 0 {
        return fmt.Sprintf("No solution!")
    } else if len(numbers) == 1 {
        return fmt.Sprintf("x=%v", numbers[0])
    } else if len(numbers) == 2 {
        return fmt.Sprintf("x<sub>1</sub>=%v; x<sub>2</sub>=%v", numbers[0], numbers[1])
    }
    return ""
}
