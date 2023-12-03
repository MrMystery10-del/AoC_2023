import java.awt.Point
import java.io.File

val data: MutableList<String> = ArrayList()

fun main() {
    collectData()

    println("Solution of part 1 is: " + part1())
    println("Solution of part 2 is: " + part2())
}

fun part1(): Int {
    var result = 0
    val pair = findPartsAndSymbols()

    for (part in pair.first) {
        for (symbol in pair.second) {
            if (part.edgePositions.contains(symbol.point)) {
                result += part.number
                break
            }
        }
    }
    return result
}

fun part2(): Int {
    var result = 0
    val pair = findPartsAndSymbols()

    for (symbol in pair.second.filter { s -> s.char == '*' }) {
        var value = 1
        var count = 0
        for (part in pair.first) {
            if (part.edgePositions.contains(symbol.point)) {
                value *= part.number
                count++;
            }
        }

        if (count == 2) {
            result += value
        }
    }
    return result
}

fun findPartsAndSymbols(): Pair<List<PartNumber>, List<Symbol>> {
    val parts = mutableListOf<PartNumber>()
    val symbolPositions = mutableListOf<Symbol>()

    for (yPos in 0 until data.size) {
        val line = data[yPos]

        var value = 0
        val partNumberEdges = mutableListOf<Point>()
        for (xPos in line.indices) {
            if (line[xPos].isDigit()) {
                value = value * 10 + line[xPos].digitToInt()
                for (xOff in -1 until 2) {
                    for (yOff in -1 until 2) {
                        partNumberEdges.add(Point(xPos + xOff, yPos + yOff))
                    }
                }
            } else {
                if (partNumberEdges.isNotEmpty()) {
                    parts.add(PartNumber(value, partNumberEdges.distinct()))
                    partNumberEdges.clear()
                    value = 0
                }
                if (line[xPos] != '.')
                    symbolPositions.add(Symbol(line[xPos], Point(xPos, yPos)))
            }
        }
        if (partNumberEdges.isNotEmpty()) {
            parts.add(PartNumber(value, partNumberEdges.distinct()))
            partNumberEdges.clear()
            value = 0
        }
    }
    return Pair(parts, symbolPositions)
}

class PartNumber(n: Int, e: List<Point>) {
    val number = n
    val edgePositions = e
}

class Symbol(c: Char, p: Point) {
    val char = c
    val point = p
}

fun collectData() {
    File("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_3\\src\\main\\resources\\data.txt").bufferedReader()
        .useLines { lines ->
            lines.forEach { item -> data.add(item) }
        }
}