var person = new Object();
person.name = "Nicholas";
person.age = 29;
person.job = "Software";

person.sayName = function () {
    console.log("function sayName");
}

person.sayName()

Object.defineProperty(person, "name", {
    writable: true, 
    value: "Songchuan.zhou",
});

person.name = "Seconde";
console.log(person.name);

person.sayAge = function () {
    console.log(this.age);
}

person.sayAge()

var person1 = {
    name: 'Chris',
    greeting: function() {
        console.log('welcome', this.name);
    }
}

var person2 = {
    name: 'Brian',
    greeting: function () {
        console.log('welcome', this.name);
    }
}

function createNewPerson(name) {
    var obj = {};
    obj.name = name;
    obj.greeting = function () {
        console.log()
    }
    return obj;
}