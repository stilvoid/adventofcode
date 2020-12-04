import fs = require("fs");

let data = fs.readFileSync("input", "utf-8");

const actionRe = /^Step (\w) must be finished before step (\w) can begin.$/;

class Node {
    deps:any = {};
    done:boolean = false;

    add(dep: string) {
        this.deps[dep] = true;
    }

    complete(dep: string) {
        this.deps[dep] = false;
    }

    count(): number {
        return Object.keys(this.deps).filter((k) => this.deps[k]).length;
    }

    is_ready(): boolean {
        return this.count() == 0;
    }
};

let nodes:any = {};

for(var line of data.trim().split("\n")) {
    let match = line.match(actionRe);

    let dep = match![1];
    let src = match![2];

    if(!(src in nodes)) {
        nodes[src] = new Node();
    }

    if(!(dep in nodes)) {
        nodes[dep] = new Node();
    }

    nodes[src].deps[dep] = true;
}

let keys = Object.keys(nodes).sort();

while(true) {
    // Find start point
    let start:string = "";

    for(let key of keys) {
        if(!nodes[key].done && nodes[key].is_ready()) {
            start = key;
            break;
        }
    }

    if(start == "") {
        break;
    }

    process.stdout.write(start);

    // Complete it's children
    nodes[start].done = true;
    for(let key of keys) {
        nodes[key].complete(start);
    }
}

console.log();
