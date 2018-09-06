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
    return o;
}

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
for (let i=0; i<aArray.length; i++) {
    console.log("array", aArray[i]);
}
