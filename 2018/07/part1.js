"use strict";
exports.__esModule = true;
var fs = require("fs");
var data = fs.readFileSync("input", "utf-8");
var actionRe = /^Step (\w) must be finished before step (\w) can begin.$/;
var Node = /** @class */ (function () {
    function Node() {
        this.deps = {};
        this.done = false;
    }
    Node.prototype.add = function (dep) {
        this.deps[dep] = true;
    };
    Node.prototype.complete = function (dep) {
        this.deps[dep] = false;
    };
    Node.prototype.count = function () {
        var _this = this;
        return Object.keys(this.deps).filter(function (k) { return _this.deps[k]; }).length;
    };
    Node.prototype.is_ready = function () {
        return this.count() == 0;
    };
    return Node;
}());
;
var nodes = {};
for (var _i = 0, _a = data.trim().split("\n"); _i < _a.length; _i++) {
    var line = _a[_i];
    var match = line.match(actionRe);
    var dep = match[1];
    var src = match[2];
    if (!(src in nodes)) {
        nodes[src] = new Node();
    }
    if (!(dep in nodes)) {
        nodes[dep] = new Node();
    }
    nodes[src].deps[dep] = true;
}
var keys = Object.keys(nodes).sort();
while (true) {
    // Find start point
    var start = "";
    for (var _b = 0, keys_1 = keys; _b < keys_1.length; _b++) {
        var key = keys_1[_b];
        if (!nodes[key].done && nodes[key].is_ready()) {
            start = key;
            break;
        }
    }
    if (start == "") {
        break;
    }
    process.stdout.write(start);
    // Complete it's children
    nodes[start].done = true;
    for (var _c = 0, keys_2 = keys; _c < keys_2.length; _c++) {
        var key = keys_2[_c];
        nodes[key].complete(start);
    }
}
console.log();
