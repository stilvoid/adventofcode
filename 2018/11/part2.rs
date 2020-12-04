fn power_level(x: usize, y: usize, serial: u32) -> i32 {
    let rack_id: i32 = x as i32 + 10;
    let mut power: i32 = rack_id * y as i32;
    power += serial as i32;
    power *= rack_id;
    power %= 1000;
    power /= 100;
    power -= 5;

    return power;
}

fn main() {
    let serial = 6548;

    let mut top_score = 0;
    let mut top_size = 0;
    let mut top_x = 0;
    let mut top_y = 0;

    // Build the grid first
    let mut grid = [[0i32; 301]; 301];

    for x in 1usize..301 {
        for y in 1usize..301 {
            let p = power_level(x, y, serial);
            grid[x][y] = p + grid[x-1][y] + grid[x][y-1] - grid[x-1][y-1];
        }
    }

    // Now find the winner
    for size in 1..301 {
        for x in size..301 {
            for y in size..301 {
                let total = grid[x][y] - grid[x][y - size] - grid[x - size][y] + grid[x - size][y - size];

                if total > top_score {
                    top_score = total;
                    top_x = x - size + 1;
                    top_y = y - size + 1;
                    top_size = size;
                }
            }
        }
    }

    println!("{},{},{}", top_x, top_y, top_size);
}
