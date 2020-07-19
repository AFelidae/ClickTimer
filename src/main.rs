mod hack_util;
use std::time::Instant;

fn main() {
    let mut running = false;
    let mut iteration: u16 = 0;
    let mut time = Instant::now();
    loop {
        if hack_util::pressed(0x14) {
            if !running {
                time = Instant::now();
                iteration += 1;
                running = true
            }
        } else {
            if running {
                running = false;
                println!("Burst {} lasted {}", iteration, time.elapsed().as_millis());
            }
        }
    }
}
