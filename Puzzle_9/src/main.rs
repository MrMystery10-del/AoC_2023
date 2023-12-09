use std::io::{BufRead, BufReader};

fn main() {
    let data = get_data("C:\\Users\\Mr.Mystery 1.0\\Desktop\\GitHub\\AoC_2023\\Puzzle_9\\resources\\data.txt");

    println!("Result of part 1 is: {}", part1(&data));
    println!("Result of part 2 is: {}", part2(&data));
}

fn part1(data: &Vec<Vec<i32>>) -> i32 {
    let mut result = 0;

    fn recursive(vector: &Vec<i32>) -> i32 {
        let (difference_vector, recursive_end) = new_difference_vector(vector);

        if recursive_end {
            return *vector.last().unwrap();
        }

        return recursive(&difference_vector) + vector.last().unwrap();
    }

    for line_vector in data {
        result += recursive(line_vector);
    }

    result
}

fn part2(data: &Vec<Vec<i32>>) -> i32 {
    let mut result = 0;

    fn recursive(vector: &Vec<i32>) -> i32 {
        let (difference_vector, recursive_end) = new_difference_vector(vector);

        if recursive_end {
            return *vector.first().unwrap();
        }

        return vector.first().unwrap() - recursive(&difference_vector);
    }

    for line_vector in data {
        result += recursive(line_vector);
    }

    result
}

fn new_difference_vector(vector: &Vec<i32>) -> (Vec<i32>, bool) {
    let mut difference_vector: Vec<i32> = Vec::new();

    let mut index = 0;
    let mut recursive_end = true;
    while index < vector.len() - 1 {
        let difference = vector[index + 1] - vector[index];

        if difference != 0 {
            recursive_end = false
        }

        difference_vector.push(difference);

        index += 1;
    }
    return (difference_vector, recursive_end);
}

pub fn get_data(path: &str) -> Vec<Vec<i32>> {
    let file = std::fs::OpenOptions::new()
        .read(true)
        .open(path)
        .expect("Failed to open file");
    let buf_reader = BufReader::new(file);

    let mut lines = Vec::new();
    for line in buf_reader.lines() {
        match line {
            Ok(line) => {
                let mut inner_vec = Vec::new();
                for s in line.split(" ") {
                    let num = s.parse::<i32>().unwrap();
                    inner_vec.push(num);
                }
                lines.push(inner_vec);
            }
            Err(e) => println!("Error reading line: {}", e),
        }
    }
    lines
}