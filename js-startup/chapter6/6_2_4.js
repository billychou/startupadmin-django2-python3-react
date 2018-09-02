// 组合使用构造函数模式和原型模式

function Person(name, age, job) {
    this.name = name;
    this.age = age;
    this.job = job;
}

Person.prototype = {
    constructor: Person,
    sayName: ()=>{
        console.log(this.name);
    }
}

let person1 = new Person("Songchuan.zhou", 31, "Software Engineer");
let person2 = new Person("Greg", 243, "Doctor");

console.log(person1.name)
console.log(person2.name)
console.log(person1.sayName === person2.sayName)

// 软件工程方法： 实例属性在构造函数中定义，所有实例共享的属性constructor和方法在原型中定义