package main

import (
	"fmt"

	"github.com/Kryszak/aoc2023/common"
	day1 "github.com/Kryszak/aoc2023/day_01"
	day2 "github.com/Kryszak/aoc2023/day_02"
	day3 "github.com/Kryszak/aoc2023/day_03"
	day4 "github.com/Kryszak/aoc2023/day_04"
	day5 "github.com/Kryszak/aoc2023/day_05"
	day6 "github.com/Kryszak/aoc2023/day_06"
	day7 "github.com/Kryszak/aoc2023/day_07"
	day8 "github.com/Kryszak/aoc2023/day_08"
	day9 "github.com/Kryszak/aoc2023/day_09"
	day10 "github.com/Kryszak/aoc2023/day_10"
	day11 "github.com/Kryszak/aoc2023/day_11"
	day12 "github.com/Kryszak/aoc2023/day_12"
	day13 "github.com/Kryszak/aoc2023/day_13"
)

func main() {
	fmt.Println("*** Day 01 ***")
	common.TimeMethodCall("day_01/input1.txt", day1.Part1)
	common.TimeMethodCall("day_01/input2.txt", day1.Part2)

	fmt.Println("*** Day 02 ***")
	common.TimeMethodCall("day_02/input.txt", day2.Part1)
	common.TimeMethodCall("day_02/input.txt", day2.Part2)

	fmt.Println("*** Day 03 ***")
	common.TimeMethodCall("day_03/input.txt", day3.Part1)
	common.TimeMethodCall("day_03/input.txt", day3.Part2)

	fmt.Println("*** Day 04 ***")
	common.TimeMethodCall("day_04/input.txt", day4.Part1)
	common.TimeMethodCall("day_04/input.txt", day4.Part2)

	fmt.Println("*** Day 05 ***")
	common.TimeMethodCall("day_05/input.txt", day5.Part1)
	common.TimeMethodCall("day_05/input.txt", day5.Part2)

	fmt.Println("*** Day 06 ***")
	common.TimeMethodCall("day_06/input.txt", day6.Part1)
	common.TimeMethodCall("day_06/input.txt", day6.Part2)

	fmt.Println("*** Day 07 ***")
	common.TimeMethodCall("day_07/input.txt", day7.Part1)
	common.TimeMethodCall("day_07/input.txt", day7.Part2)

	fmt.Println("*** Day 08 ***")
	common.TimeMethodCall("day_08/input1.txt", day8.Part1)
	common.TimeMethodCall("day_08/input2.txt", day8.Part2)

	fmt.Println("*** Day 09 ***")
	common.TimeMethodCall("day_09/input.txt", day9.Part1)
	common.TimeMethodCall("day_09/input.txt", day9.Part2)

	fmt.Println("*** Day 10 ***")
	common.TimeMethodCall("day_10/input1.txt", day10.Part1)
	common.TimeMethodCall("day_10/input2.txt", day10.Part2)

	fmt.Println("*** Day 11 ***")
	common.TimeMethodCall("day_11/input.txt", day11.Part1)
	common.TimeMethodCall("day_11/input.txt", day11.Part2)

	fmt.Println("*** Day 12 ***")
	common.TimeMethodCall("day_12/input.txt", day12.Part1)
	common.TimeMethodCall("day_12/input.txt", day12.Part2)

	fmt.Println("*** Day 13 ***")
	common.TimeMethodCall("day_13/input.txt", day13.Part1)
	common.TimeMethodCall("day_13/input.txt", day13.Part2)
}
