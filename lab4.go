// พัฒนาโดย นายโชคชัย แจ่มน้อย
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//สร้างตัวแปร
var (
	process   []string
	allocate  []int
	need      []int
	max       []int
	available []int
)

func defaultx() {
	process = make([]string, 10)
	allocate = make([]int, 30)
	need = make([]int, 30)
	max = make([]int, 30)
	available = make([]int, 3)

	//ให้  available เริ่มต้นที่ A=10 B=10 C=10
	for i := range available {
		available[i] = 10
	}
}

//แสดงตาราง
func showTable() {
	fmt.Printf("\n-----------------------------------------------\n")
	fmt.Printf(" Process |Allocate|  Need |  Max  | Available ")
	fmt.Printf("\n         | A B C  | A B C | A B C | ")
	fmt.Printf("\n-----------------------------------------------\n")
	if process[0] == "" {
		fmt.Printf("    -    | - - -  | - - - | - - - | %d %d %d\n", available[0], available[1], available[2])
	} else {
		for i := range process {
			//ถ้า process ยังไม่ถูกสร้างให้ข้าม
			if process[i] == "" {
				continue
			} else {
				if i == 0 {
					fmt.Printf("    %s   | %d %d %d  | %d %d %d | %d %d %d | %d %d %d\n", process[i], allocate[0], allocate[1], allocate[2], need[0], need[1], need[2], max[0], max[1], max[2], available[0], available[1], available[2])
				} else {
					fmt.Printf("    %s   | %d %d %d  | %d %d %d | %d %d %d |\n", process[i], allocate[0+(3*i)], allocate[1+(3*i)], allocate[2+(3*i)], need[0+(3*i)], need[1+(3*i)], need[2+(3*i)], max[0+(3*i)], max[1+(3*i)], max[2+(3*i)])
				}
			}
		}
	}
	fmt.Printf("\n")
	fmt.Printf("\nCommand>")
}

//รับค่าจาก keyboard
func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

//ลบ Process ที่ได้ทรัพยากรครบตามที่ต้องการแล้ว และคืนค่าทรัพยากรที่ยืมมา
func terminate(index int) {
	//คืนค่าทรัพยากรที่ยืมมา
	available[0] += allocate[0+(index*3)]
	available[1] += allocate[1+(index*3)]
	available[2] += allocate[2+(index*3)]

	//ลบ Process ที่ได้ทรัพยากรครบตามที่ต้องการแล้ว
	for i := range process {
		if process[i] == "" {
			break
		} else if process[i] != process[index] {
			continue
		} else {
			process[i] = process[i+1]
			need[0+(i*3)] = need[0+(i*3)+3]
			need[1+(i*3)] = need[1+(i*3)+3]
			need[2+(i*3)] = need[2+(i*3)+3]
			max[0+(i*3)] = max[0+(i*3)+3]
			max[1+(i*3)] = max[1+(i*3)+3]
			max[2+(i*3)] = max[2+(i*3)+3]
			allocate[0+(i*3)] = allocate[0+(i*3)+3]
			allocate[1+(i*3)] = allocate[1+(i*3)+3]
			allocate[2+(i*3)] = allocate[2+(i*3)+3]
			index = i + 1
		}
	}
}

//อัพเดทค่าในตาราง
func update() {
	for i := range process {
		if process[i] == "" {
			continue
		} else {
			need[0+(i*3)] = max[0+(i*3)] - allocate[0+(i*3)]
			need[1+(i*3)] = max[1+(i*3)] - allocate[1+(i*3)]
			need[2+(i*3)] = max[2+(i*3)] - allocate[2+(i*3)]
			if (need[0+(i*3)] == 0) && (need[1+(i*3)] == 0) && (need[2+(i*3)] == 0) {
				terminate(i)
			}
		}
	}
}

//สร้าง Process ใหม่
func newProcess(p string, max1, max2, max3 int) {
	for i := range process {
		if process[i] == "" {
			process[i] = p
			max[0+(i*3)] = max1
			max[1+(i*3)] = max2
			max[2+(i*3)] = max3
			update()
			break
		}
	}

}

//ฟังก์ชันเมื่อมีการร้องขอทรัพยากร
func req(p string, a, b, c int) {
	//req p6 3 3 3
	if (available[0]-a > 0) && (available[1]-b > 0) && (available[2]-c > 0) {
		test1 := available[0] - a //4-3 = 1
		test2 := available[1] - b //1
		test3 := available[2] - c //1
		safe := false
		fmt.Printf("\ntest1\t|\ttest2\t|\ttest3\t|\tsafe\n")
		fmt.Printf("%d\t|\t%d\t|\t%d\t|\t%t\n", test1, test2, test3, safe)

		for i := range process {
			if process[i] == "" {
				continue
			} else if process[i] != p {
				if (test1 >= need[0+(i*3)]) && (test2 >= need[1+(i*3)]) && (test3 >= need[2+(i*3)]) {
					safe = true
					break
				}
			} else { //p6 (6-3 = 3)
				if (test1 >= (need[0+(i*3)] - a)) && (test2 >= (need[1+(i*3)] - b)) && (test3 >= (need[2+(i*3)] - c)) {
					safe = true
					break
				}
			}
		}

		for i := range process {
			if process[i] == p {
				if (a <= need[0+(i*3)]) && (b <= need[1+(i*3)]) && (c <= need[2+(i*3)]) && safe == true {
					allocate[0+(i*3)] += a
					allocate[1+(i*3)] += b
					allocate[2+(i*3)] += c
					available[0] -= a
					available[1] -= b
					available[2] -= c
					fmt.Printf("\n*******************Safe!*******************\n")
					safe = false
					// safeStatus = false
				} else {
					fmt.Printf("\n*******************Not Safe!*******************\n")
				}
			} else {
				continue
			}
		}
		update()
	} else if (available[0]-a == 0) && (available[1]-b == 0) && (available[2]-c == 0) {
		test1 := available[0] - a //4-3 = 1
		test2 := available[1] - b //1
		test3 := available[2] - c //1
		safe := false
		fmt.Printf("\ntest1\t|\ttest2\t|\ttest3\t|\tsafe\n")
		fmt.Printf("%d\t|\t%d\t|\t%d\t|\t%t\n", test1, test2, test3, safe)

		for i := range process {
			if process[i] == "" {
				continue
			} else if process[i] != p {
				if (test1 >= need[0+(i*3)]) && (test2 >= need[1+(i*3)]) && (test3 >= need[2+(i*3)]) {
					safe = true
					break
				}
			} else { //p6 (6-3 = 3)
				if (test1 >= (need[0+(i*3)] - a)) && (test2 >= (need[1+(i*3)] - b)) && (test3 >= (need[2+(i*3)] - c)) {
					safe = true
					break
				}
			}
		}

		for i := range process {
			if process[i] == p {
				if (available[0]-need[0+(i*3)] == 0) && (available[1]-need[1+(i*3)] == 0) && (available[2]-need[2+(i*3)] == 0) && safe == true {
					allocate[0+(i*3)] += a
					allocate[1+(i*3)] += b
					allocate[2+(i*3)] += c
					available[0] -= a
					available[1] -= b
					available[2] -= c
					fmt.Printf("\n*******************Safe!*******************\n")
					safe = false
					// safeStatus = false
				} else {
					fmt.Printf("\n*******************Not Safe!*******************\n")
				}
			} else {
				continue
			}
		}
		update()
	} else {
		fmt.Printf("\n*******************Not Safe!*******************\n")
	}

}

func main() {
	defaultx()
	for {
		showTable()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			max1, _ := strconv.Atoi(commandx[2])
			max2, _ := strconv.Atoi(commandx[3])
			max3, _ := strconv.Atoi(commandx[4])
			newProcess(commandx[1], max1, max2, max3)
		case "req":
			a, _ := strconv.Atoi(commandx[2])
			b, _ := strconv.Atoi(commandx[3])
			c, _ := strconv.Atoi(commandx[4])
			req(commandx[1], a, b, c)
		}

	}
}
