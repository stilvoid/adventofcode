use std::fs::File;
use std::io::prelude::*;

fn main() -> std::io::Result<()> {
    let mut file = File::open("input")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    let mut total = 0;
    let mut n: i32;

    for line in contents.trim().lines() {
        n = line[1..].parse().expect("Uh oh, not a number");

        total += match &line[0..1] {
            "+" => Ok(n),
            "-" => Ok(-n),
            _ => Err("Malformed input"),
        }.expect("Error")
    }



    println!("{}", total);

    Ok(())
}
