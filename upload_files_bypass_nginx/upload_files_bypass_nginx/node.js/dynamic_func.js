const { exec } = require("child_process");

let funcName = "exec";
let command = "id";

exec(command, (error, stdout, stderr) => {
  console.log(stdout);
});
