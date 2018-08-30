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