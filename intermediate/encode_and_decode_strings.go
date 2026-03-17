package intermediate

import (
	"strconv"
	"strings"
)

/**
Design an algorithm to encode a list of strings to a string. The encoded string is then sent over the network and is decoded back to the original list of strings.

Machine 1 (sender) has the function:
string encode(vector<string> strs) {
    // ... your code
    return encoded_string;
}

Machine 2 (receiver) has the function:
vector<string> decode(string s) {
    //... your code
    return strs;
}

So Machine 1 does:
string encoded_string = encode(strs);

and Machine 2 does:
vector<string> strs2 = decode(encoded_string);

strs2 in Machine 2 should be the same as strs in Machine 1.

Implement the encode and decode methods.

Example 1:
Input: dummy_input = ["Hello","World"]
Output: ["Hello","World"]

Explanation:
Machine 1:
Codec encoder = new Codec();
String msg = encoder.encode(strs);

Machine 1 ---msg---> Machine 2

Machine 2:
Codec decoder = new Codec();
String[] strs = decoder.decode(msg);


Example 2:
Input: dummy_input = [""]
Output: [""]

Constraints:
    0 <= strs.length < 100
    0 <= strs[i].length < 200
    strs[i] contains any possible characters out of 256 valid ASCII characters.
**/

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteRune('#')
		b.WriteString(str)
	}

	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}

	for l, r := 0, 0; r < len(encoded); {
		for encoded[r] != '#' {
			r++
		}

		length, _ := strconv.Atoi(encoded[l:r])

		result = append(result, encoded[r+1:r+1+length])

		l = r + 1 + length
		r = r + 1 + length
	}

	return result
}
