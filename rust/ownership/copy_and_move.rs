// 常用的具有 copy trait 的有：
// 1. 整形
// 2. 浮点型
// 3. 布尔值
// 4. 字符类型
// 5. 元组

fn main() {
    let a_char = '1';
    let a_tuple = (1, 2, 3);

    let a_string = String::from("hello world");
    let a_vec = vec![1, 2, 3];

    {
        let _b_char = a_char; // char copy
        let _b_tuple = a_tuple; // tuple copy

        let _b_string = a_string; // stirng move
        let _b_vec = a_vec; // vec move
    }

    println!("a_char = {:?}", a_char);
    println!("a_tuple = {:?}", a_tuple);

    // println!("a_vec = {:?}", a_vec);
    // println!("a_string = {:?}", a_string);
}

