import java.io.File

val data: MutableList<String> = ArrayList()

fun main() {
    collectData()

    //println("Solution of part 1 is: " + part1())
    println("Solution of part 2 is: " + part2())
}

fun part1(): Int {
    var result = 0
    for (line in data) {
        val index1 = line.indexOfFirst { char -> char.isDigit() }
        val index2 = line.indexOfLast { char -> char.isDigit() }

        result += line[index1].digitToInt() * 10 + line[index2].digitToInt()
    }

    return result
}

fun part2(): Int {
    var result = 0
    for (line in data) {
        var value1 = -1
        var value2 = -1

        fun found(value: Int) {
            if (value1 == -1) {
                value1 = value
                value2 = value
            } else {
                value2 = value
            }
        }

        val words: Map<String, Int> = mapOf("one" to 1, "two" to 2, "three" to 3, "four" to 4, "five" to 5, "six" to 6, "seven" to 7, "eight" to 8, "nine" to 9)
        for (index in line.indices) {
            if (line[index].isDigit()) {
                found(line[index].digitToInt())
            } else {
                var possibleWords = words.keys.filter { it.startsWith(line[index]) }
                possibleWords = possibleWords.sortedBy { word -> word.length }

                if (possibleWords.isEmpty()) continue
                for (word in possibleWords) {
                    if (index + word.length > line.length) break

                    val possibleWord = line.substring(index, index + word.length)
                    if (possibleWord.contains(word)) {
                        found(words.get(word)!!)
                        break
                    }
                }
            }
        }

        result += value1 * 10 + value2;
    }

    return result
}

fun collectData() {
    File("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_1\\src\\main\\resources\\data.txt").bufferedReader()
        .useLines { lines ->
            lines.forEach { item -> data.add(item) }
        }
}