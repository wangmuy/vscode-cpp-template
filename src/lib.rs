pub trait Solution {
    fn resolve(&self) -> i32;
}

pub struct MySolution {}
impl Solution for MySolution {
    fn resolve(&self) -> i32 {
        let result = self.my_solution();
        println!("{:?}", result);
        0
    }
}

impl MySolution {
    pub fn my_solution(&self) -> i32 {
        42
    }
}
