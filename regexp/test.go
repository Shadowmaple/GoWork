package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	s := `<!DOCTYPE html>
					<td class="align-center ">成绩</td>
				</tr>
			</thead>
			<tbody>

					<tr>
						<td valign="middle">【 平时 】</td>
						<td valign="middle">40%&nbsp;</td>
						<td valign="middle">99&nbsp;</td>
					</tr>

					<tr>
						<td valign="middle">【 期末 】</td>
						<td valign="middle">60%&nbsp;</td>
						<td valign="middle">77.3&nbsp;</td>
					</tr>

					<tr>
						<td valign="middle">【 总评 】</td>
						<td valign="middle">&nbsp;</td>
						<td valign="middle">70.8&nbsp;</td>
					</tr></html>`
	//	rg, err := regexp.Compile(`<td valign="middle">[^%]*&nbsp;</td>`)
	rg, err := regexp.Compile(`<td valign="middle">([0-9|\\.]*)&nbsp;</td>`)
	// rg, err := regexp.Compile(`<td valign="middle">(.*)</td>`)
	if err != nil {
		fmt.Println(err)
		return
	}
	// t := rg.FindAllString(s, -1)
	// fmt.Println(t)

	// i := rg.FindAllStringIndex(s, -1)
	// fmt.Println(i)

	// valuesIndex := rg.FindAllStringIndex(s, -1)
	// for _, v := range valuesIndex {
	// 	fmt.Println(s[v[0]:v[1]])
	// }

	x := rg.FindAllStringSubmatch(s, 2)
	fmt.Println(x)

	if len(x) < 2 {
		return
	}

	u, err := strconv.ParseFloat(x[0][1], 32)
	f, err := strconv.ParseFloat(x[1][1], 32)
	fmt.Println(float32(u), " ", float32(f))
}
