// 创建一个函数，仅仅封装创建对象的代码，然后再返回新建的对象
// 理解 new对象的时候的方法
function Person(name, age, job) {
    let o = new Object();
    o.name = name;
    o.age = age;
    o.job = job;
    o.sayName = function () {
        console.log(this.name);
    }
    // 作为构造函数使用，不能使用箭头函数
    return o;
}

// 箭头函数表达式的语法比函数表达式更短，并且没有自己的this，arguments，super或 new.target。这些函数表达式更适用于那些本来需要匿名函数的地方，并且它们不能用作构造函数
// 不能使用构造函数

let friend = new Person("Nicholas", 29, "Software Engineer");
friend.sayName()

// 定义一个特殊的数组
function SpecialArray() {
    let array = []; 
    array.push.apply(array, arguments);
    return array;
}

let aArray = SpecialArray("apple", "bana", "cleptha");
console.log("AAA", aArray);
// ES6 
// 
for (let i=0; i<aArray.length; i++) {
    console.log("array", aArray[i]);
}

aArray.map( item=>{
    console.log("map", item);
})


var numbers = [1, 4, 2, 5]

// ...展开操作符
console.log(Math.max(...numbers))
console.log(Math.min(...numbers))

console.log(Math.max.apply(null, numbers))
console.log(Math.min.apply(null, numbers))
