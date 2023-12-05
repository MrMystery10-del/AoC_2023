import java.io.File

val data: MutableList<String> = ArrayList()

fun main() {
    collectData()

    println("Solution of part 1 is: " + part1())
    println("Solution of part 2 is: " + part2())
}

fun part1(): Int {
    var result = 0

    for (line in data) {
        val luckyNumbers = line.substring(line.indexOf(":", 0, true) + 1, line.indexOf('|'))
        val numbers = line.substring(line.indexOf("|") + 1).plus("  ")

        var value = 0.5
        var digit = 0
        val luckyNumberList = mutableSetOf<Int>()
        for (char in luckyNumbers) {
            if (char.isDigit()) {
                digit = digit * 10 + char.digitToInt()
            } else {
                if (digit != 0)
                    luckyNumberList.add(digit)
                digit = 0
            }
        }
        for (char in numbers) {
            if (char.isDigit()) {
                digit = digit * 10 + char.digitToInt()
            } else {
                if (luckyNumberList.contains(digit)) {
                    value *= 2
                }
                digit = 0
            }
        }
        result += value.toInt()
    }
    return result
}

fun part2(): Int {
    var result = 0

    val cards = mutableListOf<Card>()
    for (index in data.indices) {
        val luckyNumbers = data[index].substring(data[index].indexOf(":", 0, true) + 1, data[index].indexOf('|'))
        val numbers = data[index].substring(data[index].indexOf("|") + 1).plus("  ")

        var matches = 0
        var digit = 0
        val luckyNumberList = mutableSetOf<Int>()
        for (char in luckyNumbers) {
            if (char.isDigit()) {
                digit = digit * 10 + char.digitToInt()
            } else {
                if (digit != 0)
                    luckyNumberList.add(digit)
                digit = 0
            }
        }
        for (char in numbers) {
            if (char.isDigit()) {
                digit = digit * 10 + char.digitToInt()
            } else {
                if (luckyNumberList.contains(digit)) {
                    matches++
                }
                digit = 0
            }
        }
        cards.add(Card(index, matches))
    }
    for (card in cards) {
        for (times in 1..card.count) {
            for (num in 1..card.matches) {
                cards.filter { c -> c.index == card.index + num }.forEach { c -> c.count++ }
            }
        }
    }
    for (card in cards) {
        result += card.count
    }
    return result
}

class Card(i: Int, m: Int) {
    val index = i
    val matches = m
    var count = 1
}

fun collectData() {
    File("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_4\\src\\main\\resources\\data.txt").bufferedReader()
        .useLines { lines ->
            lines.forEach { item -> data.add(item) }
        }
}