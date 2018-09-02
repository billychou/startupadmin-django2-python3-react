// 9月份，js
var person = new Object();
person.name = "Nicholas";
person.age = 29;
person.job = "Software";

person.sayName = function () {
    console.log("function sayName");
}

person.sayName()

Object.defineProperty(person, "name", {
    writable: false, 
    value: "Songchuan.zhou",
});

person.name = "Seconde";
console.log(person.name);

// 创建对象有三种模式

// 一、工厂模式
function createPerson(name, age, job) {
    var o = new Object();
    o.name = name;
    o.age = age;
    o.job = job;
    o.sayName = function () {
        console.log(this.name);
    }
    return o;
}
var person1 = createPerson('Nicholas', 29, "Software engineer");
var person2 = createPerson('Songchuan.zhou', 31, "Software engineer");


// 二、构造函数模式
function Person(name, age, job) {
    this.name = name;
    this.age = age;
    this.job = job;
    this.sayName = function () {
        console.log(this.name);
    };
}
var constructor_person = new Person("Nicholas", 29, "software engineer")
var constructor_person1 = new Person("Myname", 31, "CTO");
console.error(constructor_person.sayName == constructor_person1.sayName);
// 三、原型模式

function PrototypePerson() {

}

PrototypePerson.prototype.name = "Nich";
PrototypePerson.prototype.age = 29;
PrototypePerson.prototype.job = "CTO"
PrototypePerson.prototype.sayName = function () {
    console.error(this.name);
}

console.log("#####开始调试原型模式#####")
var prototypePerson1 = new PrototypePerson();
console.log(prototypePerson1.hasOwnProperty('name'))  // false
// hasOwnProperty 可以检测一个属性是存在于实例中还是存在于原型中，只有给定属性存在于对象实例中的时候，才会返回true
// in操作符只要对象能够访问到属性就返回true， hasOwnProperty只在属性存在于实例中才返回true

function hasPrototypeProperty(object, name) {
    // 函数判断属性在原型中
    return !object.hasOwnProperty(name) && (name in Object);
}

// keys返回数组
let keys = Object.keys(PrototypePerson.prototype);
console.log(keys);

console.log(typeof Array.prototype.sort);
console.log(typeof String.prototype.substring);

