var minimatch = require("minimatch")
var os = require("os")
var chalk = require("chalk")

console.log("app running")

minimatch("bar.foo", "*.foo") // true!
minimatch("bar.foo", "*.bar") // false!
minimatch("bar.foo", "*.+(bar|foo)", { debug: true }) // true, and noisy!
