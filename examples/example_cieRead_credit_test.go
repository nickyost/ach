// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/moov-io/ach"
)

func Example_cieReadCredit() {
	f, err := os.Open(filepath.Join("testdata", "cie-credit.ach"))
	if err != nil {
		log.Fatalln(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		log.Fatalf("reading file: %v\n", err)
	}
	// Validate the ACH file
	if achFile.Validate(); err != nil {
		log.Fatalf("validating file: %v\n", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if err := achFile.Create(); err != nil {
		log.Fatalf("creating file: %v\n", err)
	}

	fmt.Printf("Total Amount Debit: %s\n", strconv.Itoa(achFile.Control.TotalDebitEntryDollarAmountInFile))
	fmt.Printf("Total Amount Credit: %s\n", strconv.Itoa(achFile.Control.TotalCreditEntryDollarAmountInFile))
	fmt.Printf("SEC Code: %s\n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Addenda05: %s\n", achFile.Batches[0].GetEntries()[0].Addenda05[0].String())

	// Output:
	// Total Amount Debit: 0
	// Total Amount Credit: 100000000
	// SEC Code: CIE
	// Addenda05: 705Credit Store Account                                                            00010000001
}
