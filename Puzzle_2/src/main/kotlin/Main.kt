import java.io.File
import kotlin.math.max

val data: MutableList<String> = ArrayList()

fun main() {
    collectData()

    println("Solution of part 1 is: " + part1())
    println("Solution of part 2 is: " + part2())
}

fun part1(): Int {
    var result = 0

    for (line in data) {
        val game: Game = Game()
        game.create(line)
        if (game.isValid()) {
            result += game.gameId
        }
    }

    return result
}

fun part2(): Int {
    var result = 0

    for (line in data) {
        val game: Game = Game()
        game.create(line)

        result += game.green * game.blue * game.red
    }

    return result
}

class Game() {
    var gameId = 0;
    var blue = 0
    var red = 0
    var green = 0

    fun create(information: String) {
        for (index in information.indices) {
            if (information[index].isDigit()) {
                gameId *= 10
                gameId += information[index].digitToInt()
            } else if (information[index] == ':') {
                collectSets(information.substring(index + 1).replace(" ", "").split(";"))
                break
            }
        }
    }

    fun isValid(): Boolean {
        if (blue > 14) return false
        if (green > 13) return false
        return red <= 12
    }

    private fun collectSets(information: List<String>) {
        for (set in information) {
            var value = 0
            for (c in set) {
                if (c.isDigit()) {
                    value = value * 10 + c.digitToInt()
                } else {
                    when (c) {
                        'r' -> red = max(red, value)
                        'b' -> blue = max(blue, value)
                        'g' -> green = max(green, value)
                    }
                    value = 0
                }
            }
        }
    }
}

fun collectData() {
    File("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_2\\src\\main\\resources\\data.txt").bufferedReader()
        .useLines { lines ->
            lines.forEach { item -> data.add(item) }
        }
}