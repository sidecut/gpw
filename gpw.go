package main

import (
	"errors"
	"fmt"
	"strings"
)

/* GPW - Generate pronounceable passwords
   This program uses statistics on the frequency of three-letter sequences
   in English to generate passwords.  The statistics are
   generated from your dictionary by the program load_trigram.

   See www.multicians.org/thvv/gpw.html for history and info.
   Tom Van Vleck

   THVV 06/01/94 Coded
   THVV 04/14/96 converted to Java
   THVV 07/30/97 fixed for Netscape 4.0
   THVV 11/27/09 ported to Javascript
*/

/* Permission is hereby granted, free of charge, to any person obtaining
   a copy of this software and associated documentation files (the
   "Software"), to deal in the Software without restriction, including
   without limitation the rights to use, copy, modify, merge, publish,
   distribute, sublicense, and/or sell copies of the Software, and to
   permit persons to whom the Software is furnished to do so, subject to
   the following conditions:

   The above copyright notice and this permission notice shall be included
   in all copies or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
   EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
   MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
   IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
   CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
   TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
   SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/**
 * var pw = GPW.pronounceable(10);
 */

func pronounceable(pwl int) (output string, err error) {
	var pwds []string
	for i := 0; i < 3; i++ {
		output, err = pronounceableInner(pwl)
		if err != nil {
			return
		}
		pwds = append(pwds, output)
		if len(output) == pwl {
			return
		}
	}

	err = errors.New(fmt.Sprintf("Could not generate password of sufficient length %v: %v", pwl, pwds))
	return
}

func pronounceableInner(pwl int) (output string, err error) {
	// Pick a random starting point.
	pik, err := randFloat64() // random number [0,1]
	if err != nil {
		return
	}
	ranno := int(pik * 125729.0)
	sum := 0
	for c1 := 0; c1 < 26; c1++ {
		for c2 := 0; c2 < 26; c2++ {
			for c3 := 0; c3 < 26; c3++ {
				sum += _trigram[c1][c2][c3]
				if sum > ranno {
					output += string(_alphabet[c1])
					output += string(_alphabet[c2])
					output += string(_alphabet[c3])
					c1 = 26 // Found start. Break all 3 loops.
					c2 = 26
					c3 = 26
				} // if sum
			} // for c3
		} // for c2
	} // for c1
	// Now do a random walk.
	nchar := 3
	for nchar < pwl {
		c1 := strings.IndexRune(_alphabet, rune(output[nchar-2]))
		c2 := strings.IndexRune(_alphabet, rune(output[nchar-1]))
		sum = 0
		for c3 := 0; c3 < 26; c3++ {
			sum += _trigram[c1][c2][c3]
		}
		if sum == 0 {
			//alert("sum was 0, outut="+output);
			break // exit while loop
		}
		//pik = ran.nextDouble();
		pik, err = randFloat64()
		if err != nil {
			return
		}
		ranno := pik * float64(sum)
		sum = 0
		for c3 := 0; c3 < 26; c3++ {
			sum += _trigram[c1][c2][c3]
			if sum > int(ranno) {
				output += string(_alphabet[c3])
				c3 = 26 // break for loop
			} // if sum
		} // for c3
		nchar++
	} // while nchar

	return
} // pronounceable
