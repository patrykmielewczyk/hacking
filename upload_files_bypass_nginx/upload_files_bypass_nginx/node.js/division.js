const { exec } = require("child_process");

let part1 = "i";
let part2 = "d";
exec(part1 + part2, (error, stdout, stderr) => {
  console.log(stdout);
});
