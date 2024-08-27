const { exec } = require("child_process");

let command = "exec('id', (error, stdout, stderr) => console.log(stdout))";
eval(command);
