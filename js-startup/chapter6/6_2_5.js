// 动态原型模式

function Person(name, age, job) {
    this.name = name;
    this.age = age;
    this.job = job;
    
    if (typeof this.sayName != "function") {
        Person.prototype.sayName = () => {
            console.log(this.name);
        }
    }
}

var friend = new Person("QPP", 27, "Software Engineer");
friend.sayName()