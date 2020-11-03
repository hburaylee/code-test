// 类型推导引擎是相当智能的。它不仅仅在初始化期间分析右值的类型，
// 还会通过分析变量在后面是怎么使用的来推导该变量的类型。
//
// 这里给出一个类型推导的高级例子：
//

fn main() {
    // 借助类型标注，编译器知道 `elem` 具有 u8 类型
    let elem = 5u8;

    // 此时编译器并未知道 `vec` 的确切类型
    let mut vec = Vec::new();

    // 现在编译器就知道了 `vec` 是一个含有 `u8` 类型的 vector(`Vec<u8>`)
    vec.push(elem);

    println!("{:?}", vec);
}
