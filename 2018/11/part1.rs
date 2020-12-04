fn power_level(x: i32, y: i32, serial: i32) -> i32 {
    let rack_id = x + 10;
    let mut power: i32 = rack_id * y;
    power += serial;
    power *= rack_id;
    power %= 1000;
    power /= 100;
    power -= 5;

    return power;
}

fn main() {
    let serial = 6548;

    let mut top_score = 0;
    let mut top_x = 0;
    let mut top_y = 0;

    for i in 1..299 {
        for j in 1..299 {
            let mut total = 0;

            for x in i..i+3 {
                for y in j..j+3 {
                    total += power_level(x, y, serial);
                }
            }

            if total > top_score {
                top_score = total;
                top_x = i;
                top_y = j;
            }
        }
    }

    println!("{},{}", top_x, top_y);
}
