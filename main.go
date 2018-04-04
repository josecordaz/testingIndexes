// Copyright © 2018 NAME HERE josecordaz@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	con, err := sql.Open("mysql", "root:MyNewPass@tcp(localhost:3317)/ss_03112018")

	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	for _, file := range files {
		if strings.Contains(file.Name(), ".sql") {
			content, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File %s\n", file.Name())
			fmt.Println(string(content))

			times := 30
			var nanoTotal int64

			for i := 0; i < times; i++ {
				start := time.Now()
				_, err = con.Query(string(content))
				if err != nil {
					fmt.Println("Error en la consulta", err)
					fmt.Println(err)
					os.Exit(1)
				}
				// timesSum.Ad = (time.Since(start) * time.Millisecond)
				nanoTotal += time.Since(start).Nanoseconds()
				fmt.Println(time.Since(start))
			}

			fmt.Println("Query avg => ", time.Duration(nanoTotal/int64(times))*time.Nanosecond)
		}
	}
}
