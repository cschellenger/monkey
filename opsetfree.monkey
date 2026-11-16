let counter = fn() {
    let x = 0;

    fn() {
        x = x + 1;
        puts(x);
    };
};

let next = counter();

next();
next();
next();
