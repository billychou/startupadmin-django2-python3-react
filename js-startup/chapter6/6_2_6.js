// 创建一个函数，仅仅封装创建对象的代码，然后再返回新建的对象

const Person = (name, age, job) => {
    let o = new Object();
    o.name = name;
    o.age = age;
    o.job = job;
    o.sayName = ()=> {
        console.log(this.name);
    }
    return o;
}

let friend = new Person("Nicholas", 29, "Software Engineer");
friend.sayName()