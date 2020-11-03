// if-else分支判断和其他语言类似。与很多语言不同的是，
// Rust 语言中的布尔判断条件不用小括号包住，每个判断条件后连着一个代码块。
// 如果if-else条件选择是一个表达式，所有分支都必须返回相同的类型。

fn main() {
    let n = 5;

    if n < 0 {
        print!("{} is negative", n);
    } else if n > 0 {
        print!("{} is positive", n);
    } else {
        print!("{} is zero", n);
    }

    let big_n = {
        if n < 10 && n > -10 {
            println!(", and is a small number, increase ten-fold");
            // 这条表达式返回一个`i32`类型
            10 * n
        } else {
            println!(", and is a big number, reduce by two");
            // 这条表达式也必须返回一个`i32`类型
            n / 2
        }
    };
    println!("{} -> {}", n, big_n);
}
