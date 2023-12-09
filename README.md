# AoC_2023
My solutions for the AoC of 2023! Languages im going to use here are randomly choosed by the [TogetherJava](https://discord.gg/together-java-272761734820003841) community!

## Puzzle list by language
- Puzzle 1-4 | Kotlin, choosed by [Zabuzard](https://github.com/Zabuzard)
- Puzzle 5-8 | Go, choosed by [Zabuzard](https://github.com/Zabuzard)
- Puzzle 9-12 | Rust, choosed by [Zabuzard](https://github.com/Zabuzard)
- Puzzle 13-16 | C#, choosed by [Zabuzard](https://github.com/Zabuzard)
- Puzzle 17-20 | C, choosed by [legoaggelos](https://github.com/legoaggelos)
- Puzzle 21-25 | Unknown

## Puzzle opinion
### Puzzle 1
I found Kotlin easy to learn and use, even though it was my first time. I could implement my ideas quickly with it.
The second part was tricky, though. They didn't tell us that twone should count as 21, not just 2. That messed up the regex I had before.

### Puzzle 2
I found Day 2 much simpler than Day 1. I had to read the instructions several times to figure out what the puzzle wanted from me, but I got the right answer on my first attempt. For Part 2, I only needed to modify one line of code and it worked. I used an object-oriented approach, which might not be very efficient in Kotlin, but I don't mind as long as it solves the problem.

### Puzzle 3
I spent some time on part 1 because I didn't consider the case when the number is at the end of the line. My initial solution was to add the part-number to the list only if the next symbol was not a digit, but that didn't work well. Part 2 was much easier and I finished it in about 2 minutes. I just had to change a few lines of code.

### Puzzle 4
I missed the deadline for this solution by a day. I was too distracted by an exam that I had to take. I'm not happy with my solution and I feel mad because I haven't started day 5 puzzle yet. I also want to learn Go and use it for day 5-8 puzzles, but I don't know anything about it. It's a tough situation, but I'm determined to catch up eventually :)

### Puzzle 5
I spent a lot of time on this puzzle... It wasn't too hard and I solved part 1 easily, but part 2 was a different story. I had a solution that worked on the example input, but it failed on the main input. The problem was that it gave me the right answer plus one, which was obviously wrong. I don't know why that happened, but I fixed it by subtracting one from the result. That took me about five hours of staring at my screen and wondering what was wrong with my brute-force approach.

### Puzzle 6
I breezed through the whole thing in no time, just 10 minutes for both parts. It felt like a piece of cake, maybe it should have been the first day's challenge.

### Puzzle 7
This puzzle took me a while to finish. It wasn't too difficult, but it required a lot of patience. I also made a silly mistake and assumed that a joker in part 2 was stronger than 2, not weaker. That confused me for a bit and wasted some time.

### Puzzle 8 
For the first part, I used brute force and it worked pretty fast, but then I got stuck on part 2. I tried brute force again, but after 20 minutes of running, I realized it would take forever. Then I noticed that it always took the same number of steps to reach ..Z. So I used my brain and figured out that I could just multiply all the values to get a result that was divisible by all the results to ..Z. But that was wrong because it wanted the fastest way to reach ..Z on all paths. Then I remembered that I could simplify the result, but since I didn't know how to code that, I just googled a lcm function implementation lol.

### Puzzle 9
Today is the first day of my four-day Rust puzzles. Fortunately, I already know Rust, so I don’t have to learn a new language in a hurry like I did with Kotlin and Go for previous puzzles. Today’s puzzle was easy but enjoyable, and I really liked it. Rust is always fun to code with. For part 2, I literally only had to change two lines: from vector.last() to vector.first(), and the return statement of the recursive inner function.
