package main

import (
	"fmt"
	"regexp"
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
						<td valign="middle">缓考&nbsp;</td>
					</tr>

					<tr>
						<td valign="middle">【 总评 】</td>
						<td valign="middle">0%&nbsp;</td>
						<td valign="middle">70.8&nbsp;</td>
					</tr></html>`
	rg, err := regexp.Compile(`<td valign="middle">(.*)&nbsp;</td>`)
	// rg, err := regexp.Compile(`<td valign="middle">([\d|\\.|%]*)&nbsp;</td>`)
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

	x := rg.FindAllStringSubmatch(s, -1)
	fmt.Println(x)
}
