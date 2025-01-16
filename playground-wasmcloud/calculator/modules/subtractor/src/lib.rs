wit_bindgen::generate!({ generate_all });

use exports::yanacalculator::subtractor::subtractor::Guest;

struct Subtractor;

impl Guest for Subtractor {
    fn subtract(arg1: u32, arg2: u32) -> u32 {
      return arg1 - arg2;
    }
}

export!(Subtractor);