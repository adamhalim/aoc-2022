use std::fs::File;
use std::io::{BufRead, BufReader};
use std::collections::HashSet;
// use std::collections::Vector;

fn main() {
    let filename = "input.txt";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    let mut sum_one: u32 = 0;
    let mut sum_two: u32 = 0;
    let mut line_count: u32 = 0;
    let mut lines: [String; 6] = Default::default();
    for (_, line) in reader.lines().enumerate() {
        let line = line.unwrap();
        let line_copy = line.clone();

        sum_one += part_one(line);
       
        lines[line_count as usize] = line_copy;
        sum_two += part_two(&mut lines, &mut line_count);
    }
    println!("Sum part one: {}.", sum_one);
    println!("Sum part two: {}.", sum_one);
}

fn part_one(line: String) -> u32 {
    let first_compartment = &line[..line.chars().count() / 2];
    let second_compartment = &line[line.chars().count() / 2 ..];

    let first_set: HashSet<char> = first_compartment.chars().collect();
    let second_set: HashSet<char> = second_compartment.chars().collect();

    let intersection = first_set.intersection(&second_set);

    for item in intersection {
        let value = *item as u32;
        return priority(value);
    }
    return 0;
}


fn part_two(lines: &mut[String], line_count: &mut u32) -> u32 {
    if *line_count != 5  {
        *line_count += 1;
        return 0;
    }

    let mut first_prio : u32 = 0;
    for c in lines[0].chars()  {
        if lines[1].contains(c) && lines[2].contains(c) {
            first_prio = priority(c as u32);
        }
    }
    
    let mut second_prio : u32 = 0;
    for c in lines[3].chars() {
        if lines[4].contains(c) && lines[5].contains(c) {
            second_prio = priority(c as u32);
        }
    }

    *line_count = 0;
    return first_prio + second_prio;
}

fn priority(mut value: u32) -> u32 {
    if value > 90 {
        value -= 96;
    } else {
        value -= 64 - 26;
    }
    return value;
}