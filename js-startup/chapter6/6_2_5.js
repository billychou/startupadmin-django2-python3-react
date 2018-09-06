// 动态原型模式
// Digest:
// 动态原型模式把所有信息都封装在了构造函数中，而通过在构造函数中初始化原型，而通过在构造函数中初始化原型，
// 保持了同事使用构造函数和原型的优点

function Person(name, age, job) {
    this.name = name;
    this.age = age;
    this.job = job;
    // 在sayName方法不存在的情况下，会将它添加到原型中
    if (typeof this.sayName != "function") {
        Person.prototype.sayName = () => {
            console.log(this.name);
        }
    }
}

var friend = new Person("QPP", 27, "Software Engineer");
friend.sayName()


// 下面使用 class 关键词
class PersonClass {
    // 构造函数
    constructor(name, age, job) {
        this.name = name;
        this.age = age;
        this.job = job;
    }

    sayName() {
        console.log(this.name);
        console.log(this.age);
        console.log(this.job);
    }
}

var personClass = new PersonClass("songchuan.zhou", 31, "Software Engineer");
personClass.sayName();